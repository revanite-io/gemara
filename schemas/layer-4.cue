package layer4
// Top level schema //

#EvaluationPlan: {
}

"evaluation-plans": [...#EvaluationPlan]

// Types

#ControlEvaluation: {
    name: string
    "control-id": string
    result: #Result
    message: string
    "documentation-url"?: =~"^https?://[^\\s]+$"
    "corrupted-state"?: bool
    "assessment-results"?: [...#AssessmentResult]
}

#AssessmentResult: {
    result: #Result
    name: string
    description: string
    message: string
    "function-address": string
    change?: #Change
    value?: _
}

#Result: "Passed" | "Failed" | "Needs Review"

#Change: {
    "target-name": string
    applied: bool
    reverted: bool
    error?: string
    "target-object"?: _
}