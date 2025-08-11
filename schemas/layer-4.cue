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

// Assessment is a struct that contains the results of a single step within a ControlEvaluation.
#Assessment: {
	// RequirementID is the unique identifier for the requirement being tested
	"requirement-id": string @go(RequirementId)
	// Applicability is a slice of identifier strings to determine when this test is applicable
	applicability: [...string]
	// Description is a human-readable description of the test
	description: string
	// Result is the overall result of the assessment
	result: #Result
	// Message is the human-readable result of the test
	message: string
	// Methods is a slice of assessment methods that were executed during the test
	methods: [...#AssessmentMethod]
	// MethodsExecuted is the number of assessment methods that were executed during the test
	"methods-executed"?: int @go(MethodsExecuted)
	// RunDuration is the time it took to run the test
	"run-duration"?: string @go(RunDuration)
	// Value is the object that was returned during the test
	value?: _
	// Changes is a map of changes that were made during the test
	changes?: [string]: #Change
}

// AssessmentMethod describes a specific procedure for evaluating a Layer 2 control requirement.
#AssessmentMethod: {
	// Id is the unique identifier of the assessment method being executed.
	id: string
	// Name is the human-readable name of the method.
	name: string
	// Description is a detailed explanation of the method.
	description: string
	// Run is a boolean indicating whether the method was run or not. When run is true, result is expected to be present.
	run: bool
	// Remediation guide is a URL to remediation guidance associated with the control's assessment requirement and this specific assessment method.
	"remediation-guide"?: #URL @go(RemediationGuide)
	// URL to documentation that describes how the assessment method evaluates the control requirement.
	documentation?: #URL
	 // Executor is a string identifier for the address or location for the specific assessment function to be used.
	executor?: string
}

// Additional constraints on Assessment Method.
#AssessmentMethod: {
	run: false
	result?: ("Not Run" | *null) @go(Result,optional=nillable)
} | {
	run:     true
	result!: #ResultWhenRun
}

// Result is valid assessment outcomes before and after execution.
#Result: #ResultWhenRun | "Not Run"

// Result is the outcome of an assessment method when it is executed.
#ResultWhenRun: "Passed" | "Failed" | "Needs Review" | "Not Applicable" | "Unknown"

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
