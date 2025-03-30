package layer4

// This file is for reusable test data to help seed ideas and reduce duplication.

import "errors"

var (
	// Generic applicability for testing
	testingApplicability = []string{"test-applicability"}

	// Functions
	goodApplyFunc = func(interface{}) (interface{}, error) {
		return nil, nil
	}
	goodRevertFunc = func(interface{}) error {
		return nil
	}
	badApplyFunc = func(interface{}) (interface{}, error) {
		return nil, errors.New("error")
	}
	badRevertFunc = func(interface{}) error {
		return errors.New("error")
	}

	// Assessment Results
	passingAssessmentStep = func(interface{}, map[string]*Change) (Result, string) {
		return Passed, ""
	}
	failingAssessmentStep = func(interface{}, map[string]*Change) (Result, string) {
		return Failed, ""
	}
	needsReviewAssessmentStep = func(interface{}, map[string]*Change) (Result, string) {
		return NeedsReview, ""
	}
	unknownAssessmentStep = func(interface{}, map[string]*Change) (Result, string) {
		return Unknown, ""
	}
)

func pendingChangePtr() *Change {
	c := pendingChange()
	return &c
}
func pendingChange() Change {
	return Change{
		Target_Name: "pendingChange",
		Description: "description placeholder",
		applyFunc:   goodApplyFunc,
		revertFunc:  goodRevertFunc,
	}
}
func appliedRevertedChange() Change {
	return Change{
		Target_Name: "appliedRevertedChange",
		Description: "description placeholder",
		applyFunc:   goodApplyFunc,
		revertFunc:  goodRevertFunc,
		Applied:     true,
		Reverted:    true,
	}
}
func appliedNotRevertedChange() Change {
	return Change{
		Target_Name: "appliedNotRevertedChange",
		Description: "description placeholder",
		applyFunc:   goodApplyFunc,
		revertFunc:  goodRevertFunc,
		Applied:     true,
	}
}
func badApplyChangePtr() *Change {
	c := badApplyChange()
	return &c
}
func badApplyChange() Change {
	return Change{
		Target_Name: "badApplyChange",
		Description: "description placeholder",
		applyFunc:   badApplyFunc,
		revertFunc:  goodRevertFunc,
	}
}
func badRevertChangePtr() *Change {
	c := badRevertChange()
	return &c
}
func badRevertChange() Change {
	return Change{
		Target_Name: "badRevertChange",
		Description: "description placeholder",
		applyFunc:   goodApplyFunc,
		revertFunc:  badRevertFunc,
	}
}
func goodRevertedChangePtr() *Change {
	c := goodRevertedChange()
	return &c
}
func goodRevertedChange() Change {
	return Change{
		Target_Name: "goodRevertedChange",
		Description: "description placeholder",
		applyFunc:   goodApplyFunc,
		revertFunc:  goodRevertFunc,
		Reverted:    true,
	}
}
func goodNotRevertedChangePtr() *Change {
	c := goodNotRevertedChange()
	return &c
}
func goodNotRevertedChange() Change {
	return Change{
		Target_Name: "goodNotRevertedChange",
		Description: "description placeholder",
		applyFunc:   goodApplyFunc,
		revertFunc:  goodRevertFunc,
		Applied:     true,
	}
}
func noApplyChangePtr() *Change {
	c := noApplyChange()
	return &c
}
func noApplyChange() Change {
	return Change{
		Target_Name: "noApplyChange",
		Description: "description placeholder",
		revertFunc:  goodRevertFunc,
	}
}
func noRevertChange() Change {
	return Change{
		Target_Name: "noRevertChange",
		Description: "description placeholder",
		applyFunc:   goodApplyFunc,
	}
}
func disallowedChange() Change {
	return Change{
		Target_Name: "disallowedChange",
		Description: "description placeholder",
		Allowed:     false,
		applyFunc:   goodApplyFunc,
		revertFunc:  goodRevertFunc,
	}
}

func failingAssessmentPtr() *Assessment {
	a := failingAssessment()
	return &a
}

func failingAssessment() Assessment {
	return Assessment{
		Requirement_Id: "failingAssessment()",
		Description:    "failing assessment",
		Steps: []AssessmentStep{
			failingAssessmentStep,
			passingAssessmentStep,
		},
		Applicability: testingApplicability,
	}
}
func passingAssessmentPtr() *Assessment {
	a := passingAssessment()
	return &a
}

func passingAssessment() Assessment {
	return Assessment{
		Requirement_Id: "passingAssessment()",
		Description:    "passing assessment",
		Steps: []AssessmentStep{
			passingAssessmentStep,
		},
		Applicability: testingApplicability,
		Changes: map[string]*Change{
			"pendingChange": pendingChangePtr(),
		},
	}
}
func needsReviewAssessmentPtr() *Assessment {
	a := needsReviewAssessment()
	return &a
}

func needsReviewAssessment() Assessment {
	return Assessment{
		Requirement_Id: "needsReviewAssessment()",
		Description:    "needs review assessment",
		Steps: []AssessmentStep{
			passingAssessmentStep,
			needsReviewAssessmentStep,
			passingAssessmentStep,
		},
		Applicability: testingApplicability,
	}
}
func unknownAssessmentPtr() *Assessment {
	a := unknownAssessment()
	return &a
}

func unknownAssessment() Assessment {
	return Assessment{
		Requirement_Id: "unknownAssessment()",
		Description:    "unknown assessment",
		Steps: []AssessmentStep{
			passingAssessmentStep,
			unknownAssessmentStep,
			passingAssessmentStep,
		},
		Applicability: testingApplicability,
	}
}

func badRevertPassingAssessment() Assessment {
	return Assessment{
		Requirement_Id: "badRevertPassingAssessment()",
		Description:    "bad revert passing assessment",
		Changes: map[string]*Change{
			"badRevertChange": badRevertChangePtr(),
		},
		Steps: []AssessmentStep{
			passingAssessmentStep,
			passingAssessmentStep,
			passingAssessmentStep,
			passingAssessmentStep,
		},
		Applicability: testingApplicability,
	}
}
