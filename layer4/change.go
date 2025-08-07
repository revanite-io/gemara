package layer4

import (
	"fmt"
)

// Prepared function to apply the change
type ApplyFunc func(interface{}) (interface{}, error)

// Prepared function to revert the change after it has been applied
type RevertFunc func(interface{}) error

// Change is a struct that contains the data and functions associated with a single change to a target resource.
type Change struct {
	// TargetName is the name or ID of the resource or configuration that is to be changed
	TargetName string `yaml:"target-name"`
	// Description is a human-readable description of the change
	Description string `yaml:"description"`
	// applyFunc is the function that will be executed to make the change
	applyFunc ApplyFunc
	// revertFunc is the function that will be executed to undo the change
	revertFunc RevertFunc
	// TargetObject is supplemental data describing the object that was changed
	TargetObject interface{} `yaml:"target-object,omitempty"`
	// Applied is true if the change was successfully applied at least once
	Applied bool `yaml:"applied,omitempty"`
	// Reverted is true if the change was successfully reverted and not applied again
	Reverted bool `yaml:"reverted,omitempty"`
	// Error is used if any error occurred during the change
	Error error `yaml:"error,omitempty"`
	// Allowed may be disabled to prevent the change from being applied
	Allowed bool `yaml:"allowed,omitempty"`
}

// Allow marks the change as allowed to be applied.
func (c *Change) Allow() {
	c.Allowed = true
}

// Apply the prepared function for the change. It will not apply the change if it has already been applied and not reverted.
// It will also not apply the change if it is not allowed.
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
	c.TargetName = targetName
	c.TargetObject = targetObject
	changeOutput, err = c.applyFunc(changeInput)
	if err != nil {
		return false, changeOutput
	}
	c.Applied = true
	c.Reverted = false
	return true, changeOutput
}

// Revert the change by executing the revert function. It will not revert the change if it has not been applied.
func (c *Change) Revert(data interface{}) {
	err := c.precheck()
	if err != nil {
		c.Error = err
		return
	}
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

// precheck verifies that the applyFunc and revertFunc are defined for the change.
// It returns an error if the change is not valid.
func (c *Change) precheck() error {
	if c.applyFunc == nil || c.revertFunc == nil {
		return fmt.Errorf("applyFunc and revertFunc must be defined for a change, but got applyFunc: %v, revertFunc: %v",
			c.applyFunc != nil, c.revertFunc != nil)
	}
	if c.TargetName == "" || c.Description == "" {
		return fmt.Errorf("change must have a TargetName and Description defined, but got TargetName: %v, Description: %v",
			c.TargetName, c.Description)
	}
	if c.Error != nil {
		return fmt.Errorf("change has a previous error and can no longer be applied: %s", c.Error.Error())
	}
	return nil
}

// NewChange creates a new Change object.
func NewChange(targetName string, description string, targetObject interface{}, applyFunc ApplyFunc, revertFunc RevertFunc) Change {
	return Change{
		TargetName:   targetName,
		TargetObject: targetObject,
		Description:  description,
		applyFunc:    applyFunc,
		revertFunc:   revertFunc,
	}
}
