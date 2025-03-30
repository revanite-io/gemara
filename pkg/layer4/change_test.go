package layer4

import "testing"

func changesTestData() []struct {
	testName string
	change   Change
} {
	return []struct {
		testName string
		change   Change
	}{
		{
			testName: "Change not yet applied",
			change:   pendingChange(),
		},
		{
			testName: "Change already applied and not yet reverted",
			change:   appliedNotRevertedChange(),
		},
		{
			testName: "Change already applied and reverted",
			change:   appliedRevertedChange(),
		},
		{
			testName: "No revert function specified",
			change:   noRevertChange(),
		},
		{
			testName: "No apply function specified",
			change:   noApplyChange(),
		},
		{
			testName: "Neither function specified",
			change:   Change{},
		},
		{
			testName: "Change is not allowed to execute",
			change:   disallowedChange(),
		},
	}
}

func TestApply(t *testing.T) {
	for _, test := range changesTestData() {
		t.Run(test.testName, func(t *testing.T) {

			test.change.Apply("should", "not", "run")
			if test.change.Applied && test.change.Error != nil {
				t.Errorf("Expected no error, but got: %v", test.change.Error)
			}

			test.change.Allow()
			test.change.Apply("target_name", "target_object", "change_input")

			if test.change.applyFunc == nil && test.change.Error == nil {
				t.Errorf("Expected error to be set due to nil applyFunc, but it was not")
			}
			if test.change.revertFunc == nil && test.change.Error == nil {
				t.Errorf("Expected error to be set due to nil revertFunc, but it was not")
			}
			if test.change.applyFunc != nil && test.change.revertFunc != nil {
				if test.change.Allowed && !test.change.Applied {
					t.Errorf("Expected change to be applied, but it was not")
				} else if !test.change.Allowed && test.change.Applied {
					t.Errorf("Expected change to not be applied, but it was")
				}

				test.change.Revert("revert_change_input")

				if test.change.Allowed && !test.change.Applied {
					t.Errorf("Reverting should not erase the record that the change was applied")
				}
			}
		})
	}
}

func TestAllow(t *testing.T) {
	// TODO: Make this a table test
	t.Run("simple test of Allow setter", func(t *testing.T) {
		change := NewChange("pendingChange", "description placeholder", nil, goodApplyFunc, goodRevertFunc)
		if change.Applied {
			return // not applicable
		}
		if change.Allowed {
			t.Errorf("Expected change to not be allowed by default, but it was")
		}
		change.Allow()
		if !change.Allowed {
			t.Errorf("Expected change to be allowed, but it was not")
		}
		change.Apply("target_name", "target_object", "change_input")
		if !change.Applied {
			t.Errorf("Expected change to be applied, but it was not")
		}
	})
}

func TestRevert(t *testing.T) {
	for _, test := range changesTestData() {
		t.Run(test.testName, func(t *testing.T) {

			test.change.Revert("revert_change_input")

			if test.change.applyFunc == nil && test.change.Error == nil {
				t.Errorf("Expected error to be set due to nil applyFunc, but it was not")
			}
			if test.change.revertFunc == nil && test.change.Error == nil {
				t.Errorf("Expected error to be set due to nil revertFunc, but it was not")
			}
			if test.change.applyFunc != nil && test.change.revertFunc != nil {
				if test.change.Applied && !test.change.Reverted {
					t.Errorf("Expected change to be reverted, but it was not")
				}
				if !test.change.Applied && test.change.Reverted {
					t.Errorf("Reverting should not be recorded if a change was not applied to revert")
				}

				test.change.Apply("should", "not", "run")
				if test.change.Applied && test.change.Error != nil {
					t.Errorf("Expected no error, but got: %v", test.change.Error)
				}

				test.change.Allow()
				test.change.Apply("target_name", "target_object", "change_input")

				if test.change.Reverted {
					t.Errorf("Applying further times should mark the change as not reverted")
				}
			}

		})
	}
}
