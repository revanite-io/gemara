package schemas

@go(layer2)

#Catalog: {
	metadata?: #Metadata

	"control-families"?: [...#ControlFamily] @go(ControlFamilies)
	threats?: [...#Threat] @go(Threats)
	capabilities?: [...#Capability] @go(Capabilities)

	"imported-controls"?: [...#Mapping] @go(ImportedControls)
	"imported-threats"?: [...#Mapping] @go(ImportedThreats)
	"imported-capabilities"?: [...#Mapping] @go(ImportedCapabilities)
}

// Resuable types //
#Metadata: {
	id:               string
	title:            string
	description:      string
	version?:         string
	"last-modified"?: string @go(LastModified) @yaml("last-modified,omitempty")
	"applicability-categories"?: [...#Category] @go(ApplicabilityCategories) @yaml("applicability-categories,omitempty")
	"mapping-references"?: [...#MappingReference] @go(MappingReferences) @yaml("mapping-references,omitempty")
}

#Category: {
	id:          string
	title:       string
	description: string
}

#ControlFamily: {
	id:          string
	title:       string
	description: string
	controls: [...#Control]
}

#Control: {
	id:        string
	title:     string
	objective: string
	"assessment-requirements": [...#AssessmentRequirement] @go(AssessmentRequirements)
	"guideline-mappings"?: [...#Mapping] @go(GuidelineMappings)
	"threat-mappings"?: [...#Mapping] @go(ThreatMappings)
}

#Threat: {
	id:          string
	title:       string
	description: string
	capabilities: [...#Mapping]

	"external-mappings"?: [...#Mapping] @go(ExternalMappings)
}

#Capability: {
	id:          string
	title:       string
	description: string
}

#MappingReference: {
	id:           string
	title:        string
	version:      string
	description?: string
	url?:         =~"^https?://[^\\s]+$"
}

#Mapping: {
	"reference-id": string @go(ReferenceId)
	entries: [...#MappingEntry]
	remarks?: string
}

#MappingEntry: {
	"reference-id": string @go(ReferenceId)
	strength:       int & >=1 & <=10
	remarks?:       string
}

#AssessmentRequirement: {
	id:   string
	text: string
	applicability: [...string]

	recommendation?: string
}
