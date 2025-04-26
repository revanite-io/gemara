// layer-3-policy.cue
package layer3

import "time"

// TODO
// Create schema for a policy document
// Policy docs should reference common and unique risk definitions, so we need schema for that as well

#Policy: {
    // Metadata useful for evaluation and automation
    metadata?: {
        // version if the version of the policy itself
        version: string
        // owner is the organizational unit responsible for the policy. This could be a department, team, or person.
        owner: string @go(Owner)
        // last_modified is the date and time when the policy was last modified
        last_modified: time.Time @go(LastModified)
        // the version of the Simplified Compliance Infrastructure model/schema used to create this policy
        sci_version: string @go(SCIVersion)
        // useful commentary or notes about the policy
        remarks: string @go(Remarks)
    }

    // Unique identifier for this policy
    id: string @go(ID)

    // Human-readable title of the policy
    title: string @go(Title)

    // Short description of this policy’s intent or purpose
    description: string @go(Description)

    // Optional reference to a parent policy this one inherits from or refines
    parent_policy_id?: string @go(ParentPolicyID)

    // Policy classification level (e.g., mandatory, recommended)
    classification: "mandatory" | "recommended"  @go(Classification)

    // Reference to one or more Layer 2 control catalogs
    control_catalogs: [#CatalogReference, ...#CatalogReference] @go(ControlCatalogs)

}

#CatalogReference: {
    // Unique identifier for the catalog
    id: string @go(ID)
    // version of the catalog
    version: string @go(Version)

    // List of IDs of Layer 2 Catalog.Metadata.ApplicabilityCategories
    applicability: [...string] @Go(ApplicabilityCategories)

    modify?: [#ControlModification] @go(Modify)

    // Reason for including this catalog in the policy
    objective?: string @go(Objective)
}

#ControlModification: {
    id: string

    // The modified applicability level of this control, using IDs defined in the catalog
    // An empty list means the control should be omitted entirely
    applicability: [...string] @go(Applicability)

    // Justification for modifying this control
    rationale?: string @go(Rationale)
}

#Reference: {
    title: string @go(Title)
    url?: string @go(URL)
    description?: string @go(Description)
}
