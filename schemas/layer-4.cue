package layer4

import "time"

// Top level schema //

// Evaluation is a collection of assessments of the framework controls and their requirements.
#Evaluation: {
    // name is a descriptive identifier for the evaluation
    name: string @go(Name)
    // ID of the Layer 2 Catalog being evaluated in this evaluation
    catalog_id: string @go(CatalogID)
    // final outcome of the evaluation
    result: #Result @go(Result)
    // timestamp of when the evaluation execution began. If the field is not provided, the evaluation has not been executed yet.
    start_time?: time.Time @go(StartTime)
    // timestamp of when the evaluation execution ended. If the field is not provided, the evaluation has not been executed yet.
    end_time?: time.Time @go(EndTime)
    // will be true when the evaluation execution changed the evaluated service and could not successfully revert
    corrupted_state: bool @go(CorruptedState) 
    // one or more evaluations of the framework controls
    control_evaluations: [#ControlEvaluation, ...#ControlEvaluation] @go(ControlEvaluations)
}


// Types

#ControlEvaluation: {
    "control-id": string
    result: #Result
    message: string
    "documentation-url"?: =~"^https?://[^\\s]+$"
    "corrupted-state"?: bool
    "assessment-results"?: [...#AssessmentResult]
}

#AssessmentResult: {
    result: #Result
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