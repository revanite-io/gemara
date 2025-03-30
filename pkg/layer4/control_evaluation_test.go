package layer4

import "testing"

var controlEvaluationTestData = []struct {
	testName          string
	control           *ControlEvaluation
	failBeforePass    bool
	expectedResult    Result
	expectedCorrupted bool
}{
	{
		testName:          "ControlEvaluation with no Assessments",
		expectedResult:    NeedsReview,
		expectedCorrupted: false,
		control: &ControlEvaluation{
			Assessments: []*Assessment{},
		},
	},
	{
		testName:          "ControlEvaluation with one passing Assessment",
		expectedResult:    Passed,
		expectedCorrupted: false,
		control: &ControlEvaluation{
			Assessments: []*Assessment{passingAssessmentPtr()},
		},
	},
	{
		testName:          "ControlEvaluation with one failing Assessment",
		expectedResult:    Failed,
		expectedCorrupted: false,
		control: &ControlEvaluation{
			Assessments: []*Assessment{failingAssessmentPtr()},
		},
	},
	{
		testName:          "ControlEvaluation with one NeedsReview Assessment",
		expectedResult:    NeedsReview,
		expectedCorrupted: false,
		control: &ControlEvaluation{
			Assessments: []*Assessment{needsReviewAssessmentPtr()},
		},
	},
	{
		testName:          "ControlEvaluation with one Unknown Assessment",
		expectedResult:    Unknown,
		expectedCorrupted: false,
		control: &ControlEvaluation{
			Assessments: []*Assessment{unknownAssessmentPtr()},
		},
	},
	{
		testName:          "ControlEvaluation with first NeedsReview and then Unknown Assessment",
		expectedResult:    Unknown,
		expectedCorrupted: false,
		control: &ControlEvaluation{
			Assessments: []*Assessment{
				needsReviewAssessmentPtr(),
				unknownAssessmentPtr(),
			},
		},
	},
	{
		testName:          "ControlEvaluation with first Unknown and then NeedsReview Assessment",
		expectedResult:    Unknown,
		expectedCorrupted: false,
		control: &ControlEvaluation{
			Assessments: []*Assessment{
				unknownAssessmentPtr(),
				needsReviewAssessmentPtr(),
			},
		},
	},
	{
		testName:          "ControlEvaluation with first Failed and then NeedsReview Assessment",
		expectedResult:    Failed,
		expectedCorrupted: false,
		control: &ControlEvaluation{
			Assessments: []*Assessment{
				failingAssessmentPtr(),
				needsReviewAssessmentPtr(),
			},
		},
	},
	{
		testName:          "ControlEvaluation with first Failing and then Passing Assessment",
		expectedResult:    Failed,
		failBeforePass:    true,
		expectedCorrupted: false,
		control: &ControlEvaluation{
			Assessments: []*Assessment{
				failingAssessmentPtr(),
				passingAssessmentPtr(),
			},
		},
	},
}

// TestEvaluate runs a series of tests on the ControlEvaluation.Evaluate method
func TestEvaluate(t *testing.T) {
	for _, test := range controlEvaluationTestData {
		t.Run(test.testName, func(t *testing.T) {
			c := test.control // copy the control to avoid duplication in the next test
			c.Evaluate(nil, testingApplicability, true)

			if c.Result != test.expectedResult {
				t.Errorf("Expected Result to be %v, but it was %v", test.expectedResult, c.Result)
			}

			if c.Corrupted_State != test.expectedCorrupted {
				t.Errorf("Expected Corrupted_State to be %v, but it was %v", test.expectedCorrupted, c.Corrupted_State)
			}
		})
		t.Run(test.testName+"no-changes", func(t *testing.T) {
			c := test.control // copy the control to avoid duplication in the next test
			c.Evaluate(nil, testingApplicability, false)

			for _, assessment := range c.Assessments {
				if assessment.Changes != nil {
					for _, change := range assessment.Changes {
						if change.Applied {
							t.Errorf("Expected no changes to be applied, but they were")
							return
						}
					}
				}
			}

			if c.Result != test.expectedResult {
				t.Errorf("Expected Result to be %v, but it was %v", test.expectedResult, c.Result)
			}

			if c.Corrupted_State != test.expectedCorrupted {
				t.Errorf("Expected Corrupted_State to be %v, but it was %v", test.expectedCorrupted, c.Corrupted_State)
			}
		})
	}
}

func TestAddAssesment(t *testing.T) {

	controlEvaluationTestData[0].control.AddAssessment("test", "test", []string{}, []AssessmentStep{})

	if controlEvaluationTestData[0].control.Result != Failed {
		t.Errorf("Expected Result to be Failed, but it was %v", controlEvaluationTestData[0].control.Result)
	}

	if controlEvaluationTestData[0].control.Message != "expected all Assessment fields to have a value, but got: requirementId=len(4), description=len=(4), applicability=len(0), steps=len(0)" {
		t.Errorf("Expected error message to be 'expected all Assessment fields to have a value, but got: requirementId=len(4), description=len=(4), applicability=len(0), steps=len(0)', but instead it was '%v'", controlEvaluationTestData[0].control.Message)
	}

}
