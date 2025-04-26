package layer2

type Catalog struct {
	// All are optional, as multiple may be combined into a complete Layer2 object
	Metadata           Metadata        `yaml:"metadata"`
	ControlFamilies    []ControlFamily `yaml:"control-families"`
	Threats            []Threat        `yaml:"threats"`
	Capabilities       []Capability    `yaml:"capabilities"`
	SharedControls     []Mapping       `yaml:"shared-controls"`
	SharedThreats      []Mapping       `yaml:"shared-threats"`
	SharedCapabilities []Mapping       `yaml:"shared-capabilities"`
}

type Metadata struct {
	Id                      string             `yaml:"id"`
	Title                   string             `yaml:"title"`
	Description             string             `yaml:"description"`
	Version                 string             `yaml:"version"`
	LastModified            string             `yaml:"last-modified"`
	ApplicabilityCategories []Category         `yaml:"applicability-categories"`
	MappingReferences       []MappingReference `yaml:"mapping-references"`
}

type Category struct {
	Id          string `yaml:"id"`
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
}

type ControlFamily struct {
	Title       string    `yaml:"title"`
	Description string    `yaml:"description"`
	Controls    []Control `yaml:"controls"`
}

type Control struct {
	Id           string        `yaml:"id"`
	Title        string        `yaml:"title"`
	Objective    string        `yaml:"objective"`
	Requirements []Requirement `yaml:"requirements"`

	// optional
	ThreatMappings    []Mapping `yaml:"threat-mappings"`
	GuidelineMappings []Mapping `yaml:"guideline-mappings"`
}

type Threat struct {
	Id           string    `yaml:"id"`
	Title        string    `yaml:"title"`
	Description  string    `yaml:"description"`
	Capabilities []Mapping `yaml:"capabilities"`

	// optional
	Mappings []Mapping `yaml:"mappings"`
}

type Capability struct {
	Id          string `yaml:"id"`
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
}

type MappingReference struct {
	Id      string `yaml:"id"`
	Title   string `yaml:"title"`
	Version string `yaml:"version"`

	// optional
	URL string `yaml:"url"`
}

type Mapping struct {
	ReferenceId string   `yaml:"reference-id"`
	Identifiers []string `yaml:"identifiers"`
}

type Requirement struct {
	Id            string   `yaml:"id"`
	Text          string   `yaml:"text"`
	Applicability []string `yaml:"applicability"`

	// optional
	Recommendation string `yaml:"recommendation"`
}
