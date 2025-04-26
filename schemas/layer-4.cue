package layer4

import "time"

// Top level schema //

"evaluation-suite": {
    name: string // descriptive identifier for the execution that produced this output
    catalog_id: string // id for the layer2 control catalog that this suite is evaluating against
    start_time: time.Time // timestamp for when this evaluation started
    end_time: time.Time // timestamp for when this evaluation completed
    result: #Result // final outcome of the evaluation
    corrupted_state: bool // whether the evaluated service has been changed without successful reversion

    control_evaluations: [...#ControlEvaluation]
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