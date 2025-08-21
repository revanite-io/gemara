package schemas

#Layer4: {
	evaluations: [...#ControlEvaluation]
}

#ControlEvaluation: {
	name: string
	controlID: string
	result: #Result
	message: string
	corruptedState: bool
	assessments: [...#Assessment]
}

#Assessment: {
	requirementId: string
	applicability: [...string]
	description: string
	result: #Result
	message: string
	steps: [...#AssessmentStep]
	stepsExecuted?: int
	runDuration?: string
	value?: _
	changes?: { [string]: #Change }
	recommendation?: string
}

// AssessmentStep is a function type that inspects the provided targetData and returns a Result with a message.
// The message may be an error string or other descriptive text.
#AssessmentStep: string

#Change: {
	targetName: string
	description: string
	targetObject?: _
	applied?: bool
	reverted?: bool
	error?: string
	allowed?: bool
}

#Result: "Not Run" | "Passed" | "Failed" | "Needs Review" | "Not Applicable" | "Unknown"
