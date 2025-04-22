// layer-3-policy.cue

// TODO
// Create schema for a policy document
// Policy docs should reference common and unique risk definitions, so we need schema for that as well

#Policy: {
    // Metadata useful for evaluation and automation
    metadata?: {
        version: string
        author: string
        last_modified: string
        sci_version: string
        remarks: string
    }

    // Unique identifier for this policy
    id: string

    // Human-readable title of the policy
    title: string

    // Short description of this policy’s intent or purpose
    description: string

    // Optional reference to a parent policy this one inherits from or refines
    parent_policy_id?: string

    // Policy classification level (e.g., mandatory, recommended)
    classification: oneOf("mandatory", "recommended")

    // Reference to one or more Layer 2 control catalogs
    control_catalogs: [...#CatalogReference]

}

#CatalogReference: {

    id: string

    version: string

    // List of IDs to applicability values defined in the catalog
    applicability: [...string]

    modify?: [#ControlModification]

    // Reason for including this catalog in the policy
    objective?: string
}

#ControlModification: {
    id: string

    // The modified applicability level of this control, using IDs defined in the catalog
    // An empty list means the control should be omitted entirely
    applicability: [...string]

    // Justification for modifying this control
    rationale?: string
}

#Reference: {
    title: string
    url?: string
    description?: string
}
