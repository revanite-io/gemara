package layer4

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"time"
)

// Assessment is a struct that contains the results of a single step within a ControlEvaluation.
type Assessment struct {
	// RequirementID is the unique identifier for the requirement being tested
	RequirementId string `yaml:"requirement-id"`
	// Applicability is a slice of identifier strings to determine when this test is applicable
	Applicability []string `yaml:"applicability"`
	// Description is a human-readable description of the test
	Description string `yaml:"description"`
	// Result is true if the test passed
	Result Result `yaml:"result"`
	// Message is the human-readable result of the test
	Message string `yaml:"message"`
	// Steps is a slice of steps that were executed during the test
	Steps []AssessmentStep `yaml:"steps"`
	// StepsExecuted is the number of steps that were executed during the test
	StepsExecuted int `yaml:"steps-executed,omitempty"`
	// RunDuration is the time it took to run the test
	RunDuration string `yaml:"run-duration,omitempty"`
	// Value is the object that was returned during the test
	Value interface{} `yaml:"value,omitempty"`
	// Changes is a slice of changes that were made during the test
	Changes map[string]*Change `yaml:"changes,omitempty"`
	// Recommendation should optionally present the corresponding text from a layer 2 assessment requirement
	Recommendation string `yaml:"recommendation,omitempty"`
}

// AssessmentStep is a function type that inspects the provided targetData and returns a Result with a message.
// The message may be an error string or other descriptive text.
type AssessmentStep func(payload interface{}, c map[string]*Change) (Result, string)

func (as AssessmentStep) String() string {
	// Get the function pointer correctly
	fn := runtime.FuncForPC(reflect.ValueOf(as).Pointer())
	if fn == nil {
		return "<unknown function>"
	}
	return fn.Name()
}

func (as AssessmentStep) MarshalJSON() ([]byte, error) {
	return json.Marshal(as.String())
}

func (as AssessmentStep) MarshalYAML() (interface{}, error) {
	return as.String(), nil
}

// NewAssessment creates a new Assessment object and returns a pointer to it.
func NewAssessment(requirementId string, description string, applicability []string, steps []AssessmentStep) (*Assessment, error) {
	a := &Assessment{
		RequirementId: requirementId,
		Description:   description,
		Applicability: applicability,
		Result:        NotRun,
		Steps:         steps,
	}
	err := a.precheck()
	return a, err
}

// AddStep queues a new step in the Assessment
func (a *Assessment) AddStep(step AssessmentStep) {
	a.Steps = append(a.Steps, step)
}

func (a *Assessment) runStep(targetData interface{}, step AssessmentStep) Result {
	a.StepsExecuted++
	result, message := step(targetData, a.Changes)
	a.Result = UpdateAggregateResult(a.Result, result)
	a.Message = message
	return result
}

// Run will execute all steps, halting if any step does not return layer4.Passed.
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
	for _, step := range a.Steps {
		if a.runStep(targetData, step) == Failed {
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
	if a.RequirementId == "" || a.Description == "" || a.Applicability == nil || a.Steps == nil || len(a.Applicability) == 0 || len(a.Steps) == 0 {
		message := fmt.Sprintf(
			"expected all Assessment fields to have a value, but got: requirementId=len(%v), description=len=(%v), applicability=len(%v), steps=len(%v)",
			len(a.RequirementId), len(a.Description), len(a.Applicability), len(a.Steps),
		)
		a.Result = Unknown
		a.Message = message
		return errors.New(message)
	}

	return nil
}
