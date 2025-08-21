package layer4

import (
	"testing"
)

func getAssessments() []struct {
	testName       string
	assessment     Assessment
	expectedResult Result
} {
	return []struct {
		testName       string
		assessment     Assessment
		expectedResult Result
	}{
		{
			testName:   "Assessment with no steps",
			assessment: Assessment{},
		},
		{
			testName:       "Passing assessment",
			assessment:     passingAssessment(),
			expectedResult: Passed,
		},
		{
			testName:       "Failing assessment",
			assessment:     failingAssessment(),
			expectedResult: Failed,
		},
		{
			testName:       "Assessment needs review",
			assessment:     needsReviewAssessment(),
			expectedResult: NeedsReview,
		},
		{
			testName:       "Bad change revert",
			assessment:     badRevertPassingAssessment(),
			expectedResult: Passed,
		},
	}
}

func TestRun(t *testing.T) {
	for _, data := range getAssessments() {
		t.Run(data.testName+"-no-changes", func(t *testing.T) {
			a := data.assessment // copy the assessment to prevent duplicate executions in the next test
			result := a.Run(nil, false)
			if result != a.Result {
				t.Errorf("expected match between Run return value (%s) and assessment Result value (%s)", result, data.expectedResult)
			}
			for _, change := range data.assessment.Changes {
				if change.Allowed {
					t.Errorf("expected all changes to be disallowed, but found an allowed change")
					return
				}
				if change.Applied || change.Reverted {
					t.Errorf("expected no changes to be applied, but found applied=%t, reverted=%t", change.Applied, change.Reverted)
					return
				}
			}
		})
	}
}

// TestNewChange ensures that NewChange creates a new Change object and adds it to the Assessment
func TestNewChange(t *testing.T) {
	anyOldAssessment := Assessment{}
	testName := "Add-a-new-change"
	t.Run(testName, func(t *testing.T) {
		if len(anyOldAssessment.Changes) != 0 {
			t.Errorf("Expected empty assessment object to start with 0 Change objects, got %d", len(anyOldAssessment.Changes))
		}
		change := anyOldAssessment.NewChange(testName, "targetName", "description", nil, goodApplyFunc, goodRevertFunc)
		if len(anyOldAssessment.Changes) != 1 {
			t.Errorf("Expected assessment object to have 1 Change object, got %d", len(anyOldAssessment.Changes))
		}
		if change == nil {
			t.Error("expected a change object to be returned by NewChange, got nil")
		}
		if change != anyOldAssessment.Changes[testName] {
			t.Errorf("Found different change object in assessment object than the one returned by NewChange")
		}

	})
}

// TestRevertChanges ensures that RevertChanges attempts to revert all changes in the Assessment
func TestRevertChanges(t *testing.T) {
	revertChangesTestData := []struct {
		testName   string
		assessment Assessment
		corrupted  bool
	}{
		{
			testName:   "No changes",
			assessment: Assessment{},
			corrupted:  false,
		},
		{
			testName:   "Change already applied and reverted",
			assessment: Assessment{Changes: map[string]*Change{"test": goodRevertedChangePtr()}},
			corrupted:  false,
		},
		{
			testName:   "Change without apply function",
			assessment: Assessment{Changes: map[string]*Change{"test": noApplyChangePtr()}},
			corrupted:  true,
		},
		{
			testName:   "Change with error from apply function",
			assessment: Assessment{Changes: map[string]*Change{"test": badApplyChangePtr()}},
			corrupted:  true,
		},
		{
			testName:   "Change with error from revert function",
			assessment: Assessment{Changes: map[string]*Change{"test": badRevertChangePtr()}},
			corrupted:  true,
		},
		{
			testName:   "Change previously applied and needs reverted",
			assessment: Assessment{Changes: map[string]*Change{"test": goodNotRevertedChangePtr()}},
			corrupted:  false,
		},
		{
			testName:   "Two changes already applied, with one already reverted",
			assessment: passingAssessment(),
			corrupted:  false,
		},
	}
	for _, data := range revertChangesTestData {
		t.Run(data.testName, func(t *testing.T) {
			for _, change := range data.assessment.Changes {
				if !change.Allowed {
					return
				}
				change.Apply("target_name", "target_object", "change_input")
			}
			corrupted := data.assessment.RevertChanges()
			if corrupted != data.corrupted {
				t.Errorf("expected corruption to be %t, got %t", data.corrupted, corrupted)
			}
		})
	}
}

func TestNewAssessment(t *testing.T) {
	newAssessmentsTestData := []struct {
		testName      string
		requirementId string
		description   string
		applicability []string
		procedures    []*AssessmentProcedure
		expectedError bool
	}{
		{
			testName:      "Empty requirementId",
			requirementId: "",
			description:   "test",
			applicability: []string{"test"},
			procedures:    []*AssessmentProcedure{&passingProcedure},
			expectedError: true,
		},
		{
			testName:      "Empty description",
			requirementId: "test",
			description:   "",
			applicability: []string{"test"},
			procedures:    []*AssessmentProcedure{&passingProcedure},
			expectedError: true,
		},
		{
			testName:      "Empty applicability",
			requirementId: "test",
			description:   "test",
			applicability: []string{},
			procedures:    []*AssessmentProcedure{&passingProcedure},
			expectedError: true,
		},
		{
			testName:      "Empty steps",
			requirementId: "test",
			description:   "test",
			applicability: []string{"test"},
			procedures:    []*AssessmentProcedure{},
			expectedError: true,
		},
		{
			testName:      "Good data",
			requirementId: "test",
			description:   "test",
			applicability: []string{"test"},
			procedures:    []*AssessmentProcedure{&passingProcedure},
			expectedError: false,
		},
	}
	for _, data := range newAssessmentsTestData {
		t.Run(data.testName, func(t *testing.T) {
			assessment, err := NewAssessment(data.requirementId, data.description, data.applicability, data.procedures)
			if data.expectedError && err == nil {
				t.Error("expected error, got nil")
			}
			if !data.expectedError && err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			if assessment == nil && !data.expectedError {
				t.Error("expected assessment object, got nil")
			}
		})
	}
}
