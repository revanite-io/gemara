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
	passingAssessmentMethod = AssessmentMethod{Executor: func(interface{}, map[string]*Change) (Result, string) {
		return Passed, ""
	}}
	failingAssessmentMethod = AssessmentMethod{Executor: func(interface{}, map[string]*Change) (Result, string) {
		return Failed, ""
	}}
	needsReviewAssessmentMethod = AssessmentMethod{Executor: func(interface{}, map[string]*Change) (Result, string) {
		return NeedsReview, ""
	}}
	unknownAssessmentMethod = AssessmentMethod{Executor: func(interface{}, map[string]*Change) (Result, string) {
		return Unknown, ""
	}}
)

func pendingChangePtr() *Change {
	c := pendingChange()
	return &c
}
func pendingChange() Change {
	return Change{
		TargetName:  "pendingChange",
		Description: "description placeholder",
		applyFunc:   goodApplyFunc,
		revertFunc:  goodRevertFunc,
	}
}
func appliedRevertedChange() Change {
	return Change{
		TargetName:  "appliedRevertedChange",
		Description: "description placeholder",
		applyFunc:   goodApplyFunc,
		revertFunc:  goodRevertFunc,
		Applied:     true,
		Reverted:    true,
	}
}
func appliedNotRevertedChange() Change {
	return Change{
		TargetName:  "appliedNotRevertedChange",
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
		TargetName:  "badApplyChange",
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
		TargetName:  "badRevertChange",
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
		TargetName:  "goodRevertedChange",
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
		TargetName:  "goodNotRevertedChange",
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
		TargetName:  "noApplyChange",
		Description: "description placeholder",
		revertFunc:  goodRevertFunc,
	}
}
func noRevertChange() Change {
	return Change{
		TargetName:  "noRevertChange",
		Description: "description placeholder",
		applyFunc:   goodApplyFunc,
	}
}
func disallowedChange() Change {
	return Change{
		TargetName:  "disallowedChange",
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
		RequirementId: "failingAssessment()",
		Description:   "failing assessment",
		Methods: []*AssessmentMethod{
			&failingAssessmentMethod,
			&passingAssessmentMethod,
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
		RequirementId: "passingAssessment()",
		Description:   "passing assessment",
		Methods: []*AssessmentMethod{
			&passingAssessmentMethod,
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
		RequirementId: "needsReviewAssessment()",
		Description:   "needs review assessment",
		Methods: []*AssessmentMethod{
			&passingAssessmentMethod,
			&needsReviewAssessmentMethod,
			&passingAssessmentMethod,
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
		RequirementId: "unknownAssessment()",
		Description:   "unknown assessment",
		Methods: []*AssessmentMethod{
			&passingAssessmentMethod,
			&unknownAssessmentMethod,
			&passingAssessmentMethod,
		},
		Applicability: testingApplicability,
	}
}

func badRevertPassingAssessment() Assessment {
	return Assessment{
		RequirementId: "badRevertPassingAssessment()",
		Description:    "bad revert passing assessment",
		Changes: map[string]*Change{
			"badRevertChange": badRevertChangePtr(),
		},
		Methods: []*AssessmentMethod{
			&passingAssessmentMethod,
			&passingAssessmentMethod,
			&passingAssessmentMethod,
			&passingAssessmentMethod,
		},
		Applicability: testingApplicability,
	}
}
