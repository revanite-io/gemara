package oscalexporter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalizeControl(t *testing.T) {
	tests := []struct {
		name        string
		inputString string
		isSubPart   bool
		wantString  string
	}{
		{
			name:        "NormalizedControl",
			inputString: "air-det-1",
			isSubPart:   false,
			wantString:  "air-det-1",
		},
		{
			name:        "CapitalizedControl",
			inputString: "AIR-DET-1",
			isSubPart:   false,
			wantString:  "air-det-1",
		},
		{
			name:        "InvalidInput",
			inputString: "AU-6(9)",
			isSubPart:   false,
			wantString:  "au-6.9",
		},
		{
			name:        "Subpart",
			inputString: "AU-6(9)",
			isSubPart:   true,
			wantString:  "9",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotString := NormalizeControl(tt.inputString, tt.isSubPart)
			assert.Equal(t, tt.wantString, gotString)
		})
	}
}
