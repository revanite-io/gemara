package schemas

@go(layer4)

#EvaluationResults: {
	"evaluation-set": [#ControlEvaluation, ...#ControlEvaluation] @go(EvaluationSet)
	...
}

#ControlEvaluation: {
	name:              string
	"control-id":      string @go(ControlId)
	result:            #Result
	message:           string
	"corrupted-state": bool @go(CorruptedState)
	assessments: [...#Assessment]
}

#Assessment: {
	"requirement-id": string @go(RequirementId)
	applicability: [...string]
	description: string
	result:      #Result
	message:     string
	steps: [...#AssessmentStep]
	"steps-executed"?: int    @go(StepsExecuted)
	"run-duration"?:   string @go(RunDuration)
	value?:            _
	changes?: [string]: #Change @go(,optional=nillable)
	recommendation?: string
}

#AssessmentStep: string @go(-)

#Change: {
	"target-name":    string @go(TargetName)
	description:      string
	"target-object"?: _ @go(TargetObject)
	applied?:         bool
	reverted?:        bool
	error?:           string
}

#Result: "Not Run" | "Passed" | "Failed" | "Needs Review" | "Not Applicable" | "Unknown" @go(-)
