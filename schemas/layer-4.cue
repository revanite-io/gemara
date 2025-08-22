package schemas

#EvaluationResults: {
	"evaluation-set": [#ControlEvaluation, ...#ControlEvaluation] @go(EvaluationSet)
	...
}

#ControlEvaluation: {
	name:              string
	"control-id":      string @go(ControlId)
	run:               bool
	"corrupted-state": bool @go(CorruptedState)
	assessments: [...#Assessment]
} & (
	{
		run:     false
		result?: #Result
	} | {
		run:     true
		result!: #Result
	})

#Assessment: {
	"requirement-id": string @go(RequirementId)
	applicability: [...string]
	description: string
	run:         bool
	steps: [...#AssessmentStep]
	"steps-executed"?: int    @go(StepsExecuted)
	"run-duration"?:   string @go(RunDuration)
	value?:            _
	changes?: {[string]: #Change}
	recommendation?: string
} & (
	{
		run:     false
		result?: #Result
	} | {
		run:     true
		result!: #Result
	})

#AssessmentStep: string

#Change: {
	"target-name":    string @go(TargetName)
	description:      string
	"target-object"?: _ @go(TargetObject)
	applied?:         bool
	reverted?:        bool
	error?:           string
}

#Result: {
	state:   #State
	message: string
}

#State: "Passed" | "Failed" | "Needs Review" | "Not Applicable" | "Unknown"
