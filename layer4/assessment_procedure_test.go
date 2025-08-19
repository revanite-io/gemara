package layer4

import (
	"testing"
)

func getProcedures() []struct {
	testName           string
	procedure          AssessmentProcedure
	numberOfSteps      int
	numberOfStepsToRun int
	expectedResult     Result
} {
	return []struct {
		testName           string
		procedure          AssessmentProcedure
		numberOfSteps      int
		numberOfStepsToRun int
		expectedResult     Result
	}{
		{
			testName:       "Assessment with no steps",
			procedure:      AssessmentProcedure{},
			expectedResult: Unknown,
		},
		{
			testName:           "Procedure with one step",
			procedure:          passingProcedure,
			numberOfSteps:      1,
			numberOfStepsToRun: 1,
			expectedResult:     Passed,
		},
		{
			testName:           "Procedure and two steps",
			procedure:          failingProcedure,
			numberOfSteps:      2,
			numberOfStepsToRun: 1,
			expectedResult:     Failed,
		},
		{
			testName:           "Procedure three steps",
			procedure:          needsReviewProcedure,
			numberOfSteps:      3,
			numberOfStepsToRun: 3,
			expectedResult:     NeedsReview,
		},
	}
}

// TestRun ensures that Run executes all steps, halting if any Procedure does not return Passed
func TestRunProcedure(t *testing.T) {
	for _, data := range getProcedures() {
		t.Run(data.testName, func(t *testing.T) {
			a := data.procedure // copy the procedure to prevent duplicate executions in the next test
			result := a.RunProcedure(nil, nil)
			if result != a.Result {
				t.Errorf("expected match between Run return value (%s) and assessment Result value (%s)", result, data.expectedResult)
			}
			if a.StepsExecuted != data.numberOfStepsToRun {
				t.Errorf("expected to run %d tests, got %d", data.numberOfStepsToRun, a.StepsExecuted)
			}
		})
	}
}

// TestNewStep ensures that NewStep queues a new step in the Assessment
func TestAddStep(t *testing.T) {
	for _, test := range getProcedures() {
		t.Run(test.testName, func(t *testing.T) {
			if len(test.procedure.Steps) != test.numberOfSteps {
				t.Errorf("Bad test data: expected to start with %d, got %d", test.numberOfSteps, len(test.procedure.Steps))
			}
			test.procedure.AddStep(passingAssessmentStep)
			if len(test.procedure.Steps) != test.numberOfSteps+1 {
				t.Errorf("expected %d, got %d", test.numberOfSteps, len(test.procedure.Steps))
			}
		})
	}
}

// TestRunStep ensures that runStep runs the step and updates the Assessment Procedure
func TestRunStep(t *testing.T) {
	stepsTestData := []struct {
		testName string
		step     AssessmentStep
		result   Result
	}{
		{
			testName: "Failing step",
			step:     failingAssessmentStep,
			result:   Failed,
		},
		{
			testName: "Passing step",
			step:     passingAssessmentStep,
			result:   Passed,
		},
		{
			testName: "Needs review step",
			step:     needsReviewAssessmentStep,
			result:   NeedsReview,
		},
		{
			testName: "Unknown step",
			step:     unknownAssessmentStep,
			result:   Unknown,
		},
	}
	for _, test := range stepsTestData {
		t.Run(test.testName, func(t *testing.T) {
			anyOldProcedure := AssessmentProcedure{}

			result := anyOldProcedure.runStep(nil, nil, test.step)
			if result != test.result {
				t.Errorf("expected %s, got %s", test.result, result)
			}
			if anyOldProcedure.Result != test.result {
				t.Errorf("expected %s, got %s", test.result, anyOldProcedure.Result)
			}
		})
	}
}
