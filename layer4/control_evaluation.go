package layer4

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

// ControlEvaluation is a struct that contains all assessment results, organized by name.
type ControlEvaluation struct {
	// Name is the name of the control being evaluated
	Name string `yaml:"name"`
	// ControlID is the unique identifier for the control being evaluated
	ControlID string `yaml:"control-id"`
	// Result is the overall result of the control evaluation
	Result Result `yaml:"result"`
	// Message is the human-readable result of the final assessment to run in this evaluation
	Message string `yaml:"message"`
	// CorruptedState is true if the control evaluation was interrupted and changes were not reverted
	CorruptedState bool `yaml:"corrupted-state"`
	// RemediationGuide is the URL to the documentation for this evaluation
	RemediationGuide string `yaml:"remediation-guide"`
	// Assessments is a slice of pointers to Assessment objects to establish idempotency
	Assessments []*Assessment `yaml:"assessments"`
}

// AddAssessment creates a new Assessment object and adds it to the ControlEvaluation.
func (c *ControlEvaluation) AddAssessment(requirementId string, description string, applicability []string, procedures []*AssessmentProcedure) (assessment *Assessment) {
	assessment, err := NewAssessment(requirementId, description, applicability, procedures)
	if err != nil {
		c.Result = Failed
		c.Message = err.Error()
	}
	c.Assessments = append(c.Assessments, assessment)
	return
}

// Evaluate runs each test procedure in each assessment, updating the relevant fields on the control evaluation.
// It will halt an assessment if a procedure step returns a failed result. The targetData is the data that the assessment will be run against.
// The userApplicability is a slice of strings that determine when the assessment is applicable. The changesAllowed
// determines whether the assessment is allowed to execute its changes.
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

// Cleanup reverts all changes made by the ControlEvaluation.
func (c *ControlEvaluation) Cleanup() {
	for _, assessment := range c.Assessments {
		corrupted := assessment.RevertChanges()
		if corrupted {
			c.CorruptedState = true
		}
	}
}

// closeHandler creates a 'listener' on a new goroutine which will notify the program if it receives an interrupt from the operating system.
// If an interrupt is received, this will attempt to revert any changes made by the terminated ControlEvaluation.
func (c *ControlEvaluation) closeHandler() {
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-channel
		log.Print("\n*****\nUnexpected termination. Attempting to revert changes made by the active ControlEvaluation. Do not interrupt this process.\n*****\n")
		c.Cleanup()
		os.Exit(0)
	}()
}
