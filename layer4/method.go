package layer4

import (
	"encoding/json"
)

// Method is an enum representing the method used to determine the assessment procedure result.
// This is designed to restrict the possible method values to a set of known types.
type Method int

const (
	UnknownMethod Method = iota
	// TestMethod represents an automated testing assessment method
	TestMethod
	// ObservationMethod represents an assessment method that requirement
	// inspection done by a human
	ObservationMethod
)

// methodToString maps Method values to their string representations.
var methodToString = map[Method]string{
	ObservationMethod: "Observation",
	TestMethod:        "Test",
	UnknownMethod:     "Unknown",
}

func (m Method) String() string {
	return methodToString[m]
}

// MarshalYAML ensures that Method is serialized as a string in YAML
func (m Method) MarshalYAML() (interface{}, error) {
	return m.String(), nil
}

// MarshalJSON ensures that Method is serialized as a string in JSON
func (m Method) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.String())
}
