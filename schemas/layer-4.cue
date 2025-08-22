package schemas

import "time"

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
	"steps-executed"?: int @go(StepsExecuted)
	"start":           #Datetime
	"end"?:            #Datetime
	value?:            _
	changes?: {[string]: #Change}
	recommendation?: string
}

#AssessmentStep: string

#Change: {
	"target-name":    string @go(TargetName)
	description:      string
	"target-object"?: _ @go(TargetObject)
	applied?:         bool
	reverted?:        bool
	error?:           string
}

#Result: "Not Run" | "Passed" | "Failed" | "Needs Review" | "Not Applicable" | "Unknown"

#Datetime: time.Format("2006-01-02T15:04:05Z07:00") @go(Datetime,format="date-time")
