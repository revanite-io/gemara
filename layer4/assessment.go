package layer4

import (
	"errors"
	"fmt"
	"time"
)

// Assessment is a struct that contains the results of a single method within a ControlEvaluation.
type Assessment struct {
	// RequirementID is the unique identifier for the requirement being tested
	RequirementId string `yaml:"requirement-id"`
	// Applicability is a slice of identifier strings to determine when this test is applicable
	Applicability []string `yaml:"applicability"`
	// Description is a human-readable description of the test
	Description string `yaml:"description"`
	// Result is the overall result of the assessment
	Result Result `yaml:"result"`
	// Message is the human-readable result of the test
	Message string `yaml:"message"`
	// Methods is a slice of assessment methods that were executed during the test
	Methods []*AssessmentMethod `yaml:"methods"`
	// MethodsExecuted is the number of assessment methods that were executed during the test
	MethodsExecuted int `yaml:"methods-executed,omitempty"`
	// RunDuration is the time it took to run the test
	RunDuration string `yaml:"run-duration,omitempty"`
	// Value is the object that was returned during the test
	Value interface{} `yaml:"value,omitempty"`
	// Changes is a map of changes that were made during the test
	Changes map[string]*Change `yaml:"changes,omitempty"`
}

// NewAssessment creates a new Assessment object and returns a pointer to it.
func NewAssessment(requirementId string, description string, applicability []string, methods []*AssessmentMethod) (*Assessment, error) {
	a := &Assessment{
		RequirementId: requirementId,
		Description:   description,
		Applicability: applicability,
		Result:        NotRun,
		Methods:       methods,
	}
	err := a.precheck()
	return a, err
}

// AddMethod queues a new method in the Assessment
func (a *Assessment) AddMethod(method AssessmentMethod) {
	a.Methods = append(a.Methods, &method)
}

func (a *Assessment) runMethod(targetData interface{}, method *AssessmentMethod) Result {
	a.MethodsExecuted++
	result, message := method.RunMethod(targetData, a.Changes)
	a.Result = UpdateAggregateResult(a.Result, result)
	a.Message = message
	return result
}

// Run will execute all steps, halting if any method does not return layer4.Passed.
func (a *Assessment) Run(targetData interface{}, changesAllowed bool) Result {
	if a.Result != NotRun {
		return a.Result
	}

	startTime := time.Now()
	err := a.precheck()
	if err != nil {
		a.Result = Unknown
		return a.Result
	}
	for _, change := range a.Changes {
		if changesAllowed {
			change.Allow()
		}
	}
	for _, method := range a.Methods {
		if a.runMethod(targetData, method) == Failed {
			return Failed
		}
	}
	a.RunDuration = time.Since(startTime).String()
	return a.Result
}

// NewChange creates a new Change object and adds it to the Assessment.
func (a *Assessment) NewChange(
	changeName,
	targetName,
	description string,
	targetObject interface{},
	applyFunc ApplyFunc,
	revertFunc RevertFunc,
) *Change {
	change := NewChange(targetName, description, targetObject, applyFunc, revertFunc)
	if a.Changes == nil {
		a.Changes = make(map[string]*Change)
	}
	a.Changes[changeName] = &change
	return &change
}

// RevertChanges reverts all changes made by the assessment.
// It will not revert changes that have not been applied.
func (a *Assessment) RevertChanges() (corrupted bool) {
	for _, change := range a.Changes {
		if !corrupted && (change.Applied || change.Error != nil) {
			if !change.Reverted {
				change.Revert(nil)
			}
			if change.Error != nil || !change.Reverted {
				corrupted = true // do not break loop here; continue attempting to revert all changes
			}
		}
	}
	return
}

// precheck verifies that the assessment has all the required fields.
// It returns an error if the assessment is not valid.
func (a *Assessment) precheck() error {
	if a.RequirementId == "" || a.Description == "" || a.Applicability == nil || a.Methods == nil || len(a.Applicability) == 0 || len(a.Methods) == 0 {
		message := fmt.Sprintf(
			"expected all Assessment fields to have a value, but got: requirementId=len(%v), description=len=(%v), applicability=len(%v), methods=len(%v)",
			len(a.RequirementId), len(a.Description), len(a.Applicability), len(a.Methods),
		)
		a.Result = Unknown
		a.Message = message
		return errors.New(message)
	}

	return nil
}
