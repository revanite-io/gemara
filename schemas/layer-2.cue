package schemas

#Layer2: {
    metadata?: #Metadata

    "control-families"?: [...#ControlFamily]
    threats?: [...#Threat]
    capabilities?: [...#Capability]

    "shared-controls"?: [...#Mapping]
    "shared-threats"?: [...#Mapping]
    "shared-capabilities"?: [...#Mapping]
}

// Resuable types //

#Metadata: {
    id: string
    title: string
    description: string
    version?: string
    "last-modified"?: string
    "applicability-categories"?: [...#Category]
    "mapping-references"?: [...#MappingReference]
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
    "assessment-requirements": [...#AssessmentRequirement]

    "guideline-mappings"?: [...#Mapping]
    "threat-mappings"?: [...#Mapping]
}

#Threat: {
    id: string
    title: string
    description: string
    capabilities: [...#Mapping]

    "external-mappings"?: [...#Mapping]
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
    "reference-id": string
    identifiers: [...string]
}

#AssessmentRequirement: {
    id: string
    text: string
    applicability: [...string]

    recommendation?: string
}
