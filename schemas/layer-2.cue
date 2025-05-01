package schemas
@go(layer2)

#Layer2: {
    metadata?: #Metadata

    "control-families"?: [...#ControlFamily] @go(ControlFamilies)
    threats?: [...#Threat] @go(Threats)
    capabilities?: [...#Capability] @go(Capabilities)

    "shared-controls"?: [...#Mapping] @go(SharedControls)
    "shared-threats"?: [...#Mapping] @go(SharedThreats)
    "shared-capabilities"?: [...#Mapping] @go(SharedCapabilities)
}

// Resuable types //

#Metadata: {
    id: string
    title: string
    description: string
    version?: string
    "last-modified"?: string @go(LastModified)
    "applicability-categories"?: [...#Category] @go(ApplicabilityCategories)
    "mapping-references"?: [...#MappingReference] @go(MappingReferences)
}

#Category: {
    id: string
    title: string
    description: string
}

#ControlFamily: {
    title: string
    description: string
    controls: [...#Control]
}

#Control: {
    id: string
    title: string
    objective: string
    "assessment-requirements": [...#AssessmentRequirement] @go(AssessmentRequirements)

    "guideline-mappings"?: [...#Mapping] @go(GuidelineMappings)
    "threat-mappings"?: [...#Mapping] @go(ThreatMappings)
}

#Threat: {
    id: string
    title: string
    description: string
    capabilities: [...#Mapping]

    "external-mappings"?: [...#Mapping] @go(ExternalMappings)
}

#Capability: {
    id: string
    title: string
    description: string
}

#MappingReference: {
    id: string
    title: string
    version: string
    description?: string
    url?: =~"^https?://[^\\s]+$"
}

#Mapping: {
    "reference-id": string @go(ReferenceId)
    identifiers: [...string]
}

#AssessmentRequirement: {
    id: string
    text: string
    applicability: [...string]

    recommendation?: string
}
