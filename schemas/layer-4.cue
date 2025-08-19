package schemas

#Layer4: {
	evaluations: [#ControlEvaluation, ...#ControlEvaluation]
}

// ControlEvaluation provides all assessment results for a single control
#ControlEvaluation: {
	// Name is the name of the control being evaluated
	name: string
	// ControlID uniquely identifies the control
	"control-id": string
	// Result communicates whether the evaluation has been run, and if so, the outcome(s)
	result: #Result
	// Message describes the result of the evaluation
	message: string
	// CorruptedState is true if the control evaluation was interrupted and changes were not reverted
	"corrupted-state"?: bool
	// RemediationGuide provides a URL with guidance on how to remediate systems that fail control evaluation
	"remediation-guide"?: string
	// Assessments represents the collection of results from evaluation of each control requirement
	assessments: [...#Assessment]
}

// Assessment provides all assessment results from evaluation of a single control requirement
#Assessment: {
	// RequirementID uniquely identifies the requirement being tested
	"requirement-id": string @go(RequirementId)
	// Applicability provides identifier strings to determine when this assessment is applicable
	applicability: [...string]
	// Description provides a detailed explanation of the assessment
	description: string
	// Result communicates whether the assessment has been run, and if so, the outcome(s)
	result: #Result
	// Message describes the result of the assessment
	message: string
	// Procedures defines the assessment procedures associated with the assessment
	procedures: [...#AssessmentProcedure]
	// RunDuration is the time it took to run the assessment
	"run-duration"?: string @go(RunDuration)
	// Value is the object that was returned during the assessment
	value?: _
	// Changes describes changes that were made during the assessment
	changes?: [string]: #Change
}

// AssessmentProcedure describes a testing procedure for evaluating a Layer 2 control requirement.
#AssessmentProcedure: {
	// Id uniquely identifies the assessment procedure being executed
	id: string
	// Name provides a summary of the procedure
	name: string
	// Description provides a detailed explanation of the procedure
	description: string
	// Method describe the high-level method used to determine the results of the procedure
	method: #ProcedureMethod
	// Run is a boolean indicating whether the procedure was run or not. When run is true, result is expected to be present
	run: bool
	// RemediationGuide provides a URL with remediation guidance associated with the control's assessment requirement and this specific assessment procedure
	"remediation-guide"?: #URL @go(RemediationGuide)
	// Documentation provides a URL to documentation that describes how the assessment procedure evaluates the control requirement
	documentation?: #URL
	// Steps provides the address for the assessment steps executed
	"steps"?: [...string]
}

// Additional constraints on Assessment Procedure.
#AssessmentProcedure: {
	run: false
	// Message describes the result of the procedure
	message?: string
	// Result communicates the outcome(s) of the procedure
	result?: ("Not Run" | *null) @go(Result,optional=nillable)
} | {
	run:     true
	message!: string
	result!: #ResultWhenRun
}

// Result describes valid assessment outcomes before and after execution.
#Result: #ResultWhenRun | "Not Run"

// Result describes the outcome(s) of an assessment procedure when it is executed.
#ResultWhenRun: "Passed" | "Failed" | "Needs Review" | "Not Applicable" | "Unknown"

// ProcedureMethod describes method options that can be used to determine the results
#ProcedureMethod: "Test" | "Observation"

// Change is a struct that contains the data and functions associated with a single change to a target resource.
#Change: {
	// TargetName is the name or ID of the resource or configuration that is to be changed
	"target-name": string @go(TargetName)
	// Description is a human-readable description of the change
	description: string
	// TargetObject is supplemental data describing the object that was changed
	"target-object"?: _ @go(TargetObject)
	// Applied is true if the change was successfully applied at least once
	applied?: bool
	// Reverted is true if the change was successfully reverted and not applied again
	reverted?: bool
	// Error is used if any error occurred during the change
	error?: string
	// Allowed may be disabled to prevent the change from being applied
	allowed?: bool
}

// URL describes a specific subset of URLs of interest to the framework
#URL: =~"^https?://[^\\s]+$"
