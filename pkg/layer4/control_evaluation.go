package layer4

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

// ControlEvaluation is a struct that contains all assessment results, organized by name
type ControlEvaluation struct {
	Name             string        `yaml:"name"`              // Name is the name of the control being evaluated
	ControlId        string        `yaml:"control_id"`        // ControlId is the unique identifier for the control being evaluated
	Result           Result        `yaml:"result"`            // Result is the overall result of the control evaluation
	Message          string        `yaml:"message"`           // Message is the human-readable result of the final assessment to run in this evaluation
	CorruptedState   bool          `yaml:"corrupted_state"`   // CorruptedState is true if the control evaluation was interrupted and changes were not reverted
	RemediationGuide string        `yaml:"remediation_guide"` // Remediation_Guide is the URL to the documentation for this evaluation
	Assessments      []*Assessment `yaml:"assessments"`       // Assessments is a map of pointers to Assessment objects to establish idempotency
}

func (c *ControlEvaluation) AddAssessment(requirementId string, description string, applicability []string, steps []AssessmentStep) (assessment *Assessment) {
	assessment, err := NewAssessment(requirementId, description, applicability, steps)
	if err != nil {
		c.Result = Failed
		c.Message = err.Error()
	}
	c.Assessments = append(c.Assessments, assessment)
	return
}

// Evaluate runs each step in each assessment, updating the relevant fields on the control evaluation.
// It will halt if a step returns a failed result.
// `targetData` is the data that the assessment will be run against.
// `userApplicability` is a slice of strings that determine when the assessment is applicable.
// `changesAllowed` determines whether the assessment is allowed to execute its changes.
func (c *ControlEvaluation) Evaluate(targetData interface{}, userApplicability []string, changesAllowed bool) {
	if len(c.Assessments) == 0 {
		c.Result = NeedsReview
		return
	}
	c.closeHandler()
	for _, assessment := range c.Assessments {
		var applicable bool
		for _, aa := range assessment.Applicability {
			for _, ua := range userApplicability {
				if aa == ua {
					applicable = true
					break
				}
			}
		}
		if applicable {
			result := assessment.Run(targetData, changesAllowed)
			c.Result = UpdateAggregateResult(c.Result, result)
			c.Message = assessment.Message
			if c.Result == Failed {
				break
			}
		}
	}
	c.Cleanup()
}

func (c *ControlEvaluation) Cleanup() {
	for _, assessment := range c.Assessments {
		corrupted := assessment.RevertChanges()
		if corrupted {
			c.CorruptedState = true
		}
	}
}

// CloseHandler creates a 'listener' on a new goroutine which will notify the
// program if it receives an interrupt from the operating system.
// If an interrupt is received, this will attempt to revert any changes
// made by the terminated ControlEvaluation.
func (c *ControlEvaluation) closeHandler() {
	// Ref: https://golangcode.com/handle-ctrl-c-exit-in-terminal/
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-channel
		log.Print("\n*****\nUnexpected termination. Attempting to revert changes made by the active ControlEvaluation. Do not interrupt this process.\n*****\n")
		c.Cleanup()
		os.Exit(0)
	}()
}
