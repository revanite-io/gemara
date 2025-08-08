package layer4

import (
	"encoding/json"
	"reflect"
	"runtime"
)

// AssessmentMethod describes the method used to assess the layer 2 control requirement referenced by requirementID.
type AssessmentMethod struct {
	// Id is the unique identifier of the assessment method being executed
	Id string `json:"id" yaml:"id"`

	// Name is the name of the method used to assess the requirement.
	Name string `json:"name" yaml:"name"`

	// Description is a detailed explanation of the method.
	Description string `json:"description" yaml:"description"`

	// Run is a boolean indicating whether the method was run or not. When run is true, result is expected to be present.
	Run bool `json:"run" yaml:"run"`

	// Remediation guide is a URL to remediation guidance associated with the control's assessment requirement and this specific assessment method.
	RemediationGuide URL `json:"remediation-guide,omitempty" yaml:"remediation-guide,omitempty"`

	// URL to documentation that describes how the assessment method evaluates the control requirement.
	Documentation URL `json:"documentation,omitempty" yaml:"documentation,omitempty"`

	// Result is the status or outcome of an assessed method present. This field is present when Run is true.
	Result *Result `json:"result,omitempty" yaml:"result,omitempty"`
	// Executor is a function type that inspects the provided targetData and returns a Result with a message.
	// The message may be an error string or other descriptive text.
	Executor MethodExecutor
}

// URL describes a specific subset of URLs of interest to the framework
type URL string

// MethodExecutor is a function type that inspects the provided payload and returns the result of the assessment.
// The payload is the data/evidence that the assessment will be run against.
type MethodExecutor func(payload interface{}, c map[string]*Change) (Result, string)

// RunMethod executes the assessment method using the provided payload and changes.
// It returns the result of the assessment and any error encountered during execution.
// The payload is the data/evidence that the assessment will be run against.
func (a *AssessmentMethod) RunMethod(payload interface{}, changes map[string]*Change) (Result, string) {
	result, message := a.Executor(payload, changes)
	a.Result = &result
	a.Run = true
	return result, message
}

func (e MethodExecutor) String() string {
	// Get the function pointer correctly
	fn := runtime.FuncForPC(reflect.ValueOf(e).Pointer())
	if fn == nil {
		return "<unknown function>"
	}
	return fn.Name()
}

func (e MethodExecutor) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.String())
}

func (e MethodExecutor) MarshalYAML() (interface{}, error) {
	return e.String(), nil
}
