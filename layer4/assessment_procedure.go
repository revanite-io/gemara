package layer4

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"runtime"
)

// AssessmentProcedure is a specific outline defining how to assess a Layer 2 control requirement.
type AssessmentProcedure struct {
	// Id is the unique identifier of the assessment procedure being executed
	Id string `json:"id" yaml:"id"`
	// Name is the human-readable name of the procedure.
	Name string `json:"name" yaml:"name"`
	// Description is a detailed explanation of the procedure.
	Description string `json:"description" yaml:"description"`
	// Method describe the high-level method used to determine the results of the procedure
	Method Method `yaml:"method"`
	// Remediation guide is a URL to remediation guidance associated with the control's assessment requirement and this specific assessment procedure.
	RemediationGuide string `json:"remediation-guide,omitempty" yaml:"remediation-guide,omitempty"`
	// URL to documentation that describes how the assessment procedure evaluates the control requirement.
	Documentation string `json:"documentation,omitempty" yaml:"documentation,omitempty"`
	// Run is a boolean indicating whether the procedure was run or not. When run is true, result is expected to be present.
	Run bool `json:"run" yaml:"run"`
	// Message is the human-readable result of the procedure
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
	// Result is the outcome of the assessment procedure.
	// This field must be present if Run is true.
	Result Result `json:"result,omitempty" yaml:"result,omitempty"`
	// StepsExecuted is the number of steps that were executed during the assessment execution
	StepsExecuted int `json:"-" yaml:"-"`
	// Steps define logical steps to inspect the provided targetData and returns a Result with a message.
	// The message may be an error string or other descriptive text.
	Steps []AssessmentStep
}

// NewProcedure creates a new AssessmentProcedure object and returns a pointer to it.
func NewProcedure(id, name, description string, steps []AssessmentStep) (*AssessmentProcedure, error) {
	a := &AssessmentProcedure{
		Id:          id,
		Name:        name,
		Description: description,
		Result:      NotRun,
		Steps:       steps,
	}
	err := a.precheck()
	return a, err
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

// AddStep queues a new step in the Assessment Procedure
func (a *AssessmentProcedure) AddStep(step AssessmentStep) {
	a.Steps = append(a.Steps, step)
}

// RunProcedure executes all assessment steps, halting if any assessment does not return layer4.Passed.
func (a *AssessmentProcedure) RunProcedure(targetData interface{}, changes map[string]*Change) Result {
	a.Run = true

	err := a.precheck()
	if err != nil {
		a.Result = Unknown
		return a.Result
	}
	for _, steps := range a.Steps {
		if a.runStep(targetData, changes, steps) == Failed {
			return Failed
		}
	}
	return a.Result
}

func (a *AssessmentProcedure) runStep(targetData interface{}, changes map[string]*Change, step AssessmentStep) Result {
	a.StepsExecuted++
	result, message := step(targetData, changes)
	a.Result = UpdateAggregateResult(a.Result, result)
	a.Message = message
	return result
}

// precheck verifies that the assessment procedure has step fields.
// It returns an error if the assessment is not valid.
func (a *AssessmentProcedure) precheck() error {
	if len(a.Steps) == 0 {
		message := fmt.Sprintf(
			"expected all Assessment Procedure fields steps=len(%v)",
			len(a.Steps),
		)
		return errors.New(message)
	}
	return nil
}
