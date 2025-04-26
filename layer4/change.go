package layer4

import (
	"fmt"
)

type ApplyFunc func(interface{}) (interface{}, error)
type RevertFunc func(interface{}) error

// Change is a struct that contains the data and functions associated with a single change to a target resource.
type Change struct {
	Target_Name string     // Required. TargetName is the name or ID of the resource or configuration that is to be changed
	Description string     // Required. Description is a human-readable description of the change
	applyFunc   ApplyFunc  // Required. applyFunc is the function that will be executed to make the change
	revertFunc  RevertFunc // Required. revertFunc is the function that will be executed to undo the change

	Target_Object interface{} // TargetObject is supplemental data describing the object that was changed
	Applied       bool        // Applied is true if the change was successfully applied at least once
	Reverted      bool        // Reverted is true if the change was successfully reverted and not applied again
	Error         error       // Error is used if any error occurred during the change
	Allowed       bool        // Allowed may be disabled to prevent the change from being applied
}

func (c *Change) Allow() {
	c.Allowed = true
}

// Apply executes the Apply function for the change
func (c *Change) Apply(targetName string, targetObject interface{}, changeInput interface{}) (applied bool, changeOutput interface{}) {
	if !c.Allowed {
		return
	}
	err := c.precheck()
	if err != nil {
		c.Error = err
		return
	}
	// Do nothing if the change has already been applied and not reverted
	if c.Applied && !c.Reverted {
		return true, nil
	}
	c.Target_Name = targetName
	c.Target_Object = targetObject
	changeOutput, err = c.applyFunc(changeInput)
	if err != nil {
		return false, changeOutput
	}
	c.Applied = true
	c.Reverted = false
	return true, changeOutput
}

// Revert executes the Revert function for the change
func (c *Change) Revert(data interface{}) {
	err := c.precheck()
	if err != nil {
		c.Error = err
		return
	}
	// Do nothing if the change has not been applied
	if !c.Applied {
		return
	}
	err = c.revertFunc(data)
	if err != nil {
		c.Error = err
		return
	}
	c.Reverted = true
}

// precheck verifies that the applyFunc and revertFunc are defined for the change
func (c *Change) precheck() error {
	if c.applyFunc == nil || c.revertFunc == nil {
		return fmt.Errorf("applyFunc and revertFunc must be defined for a change, but got applyFunc: %v, revertFunc: %v",
			c.applyFunc != nil, c.revertFunc != nil)
	}
	if c.Target_Name == "" || c.Description == "" {
		return fmt.Errorf("change must have a TargetName and Description defined, but got TargetName: %v, Description: %v",
			c.Target_Name, c.Description)
	}
	if c.Error != nil {
		return fmt.Errorf("change has a previous error and can no longer be applied: %s", c.Error.Error())
	}
	return nil
}

// NewChange creates a new Change object and adds it to the Assessment
func NewChange(targetName string, description string, targetObject interface{}, applyFunc ApplyFunc, revertFunc RevertFunc) Change {
	return Change{
		Target_Name:   targetName,
		Target_Object: targetObject,
		Description:   description,
		applyFunc:     applyFunc,
		revertFunc:    revertFunc,
	}
}
