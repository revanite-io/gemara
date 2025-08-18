package layer2

import (
	"testing"

	"github.com/defenseunicorns/go-oscal/src/pkg/validation"
	"github.com/stretchr/testify/assert"
)

var TestCases = []struct {
	name             string
	catalog          *Catalog
	controlFamilyIDs map[string]string
	version          string
	controlHREF      string
	catalogUUID      string
	namespace        string
	wantErr          bool
	expectedTitle    string
}{
	{
		name: "Valid catalog with single control family",
		catalog: &Catalog{
			Metadata: Metadata{
				Id:    "test-catalog",
				Title: "Test Catalog",
			},
			ControlFamilies: []ControlFamily{
				{
					Id:          "AC",
					Title:       "access-control",
					Description: "Controls for access management",
					Controls: []Control{
						{
							Id:    "AC-01",
							Title: "Access Control Policy",
							AssessmentRequirements: []AssessmentRequirement{
								{
									Id:   "AC-01.1",
									Text: "Develop and document access control policy",
								},
							},
						},
					},
				},
			},
		},
		controlFamilyIDs: map[string]string{
			"AC": "AC",
		},
		version:       "devel",
		controlHREF:   "https://baseline.openssf.org/versions/%s#%s",
		catalogUUID:   "8c222a23-fc7e-4ad8-b6dd-289014f07a9f",
		namespace:     "http://baseline.openssf.org/ns/oscal",
		wantErr:       false,
		expectedTitle: "Test Catalog",
	},
	{
		name: "Valid catalog with multiple control families",
		catalog: &Catalog{
			Metadata: Metadata{
				Id:    "test-catalog-multi",
				Title: "Test Catalog Multiple",
			},
			ControlFamilies: []ControlFamily{
				{
					Id:          "AC",
					Title:       "access-control",
					Description: "Controls for access management",
					Controls: []Control{
						{
							Id:    "AC-01",
							Title: "Access Control Policy",
							AssessmentRequirements: []AssessmentRequirement{
								{
									Id:   "AC-01.1",
									Text: "Develop and document access control policy",
								},
							},
						},
					},
				},
				{
					Id:          "BR",
					Title:       "business-requirements",
					Description: "Controls for business requirements",
					Controls: []Control{
						{
							Id:    "BR-01",
							Title: "Business Requirements Policy",
							AssessmentRequirements: []AssessmentRequirement{
								{
									Id:   "BR-01.1",
									Text: "Define business requirements",
								},
							},
						},
					},
				},
			},
		},
		controlFamilyIDs: map[string]string{
			"AC": "AC",
			"BR": "BR",
		},
		version:       "devel",
		controlHREF:   "https://baseline.openssf.org/versions/%s#%s",
		catalogUUID:   "8c222a23-fc7e-4ad8-b6dd-289014f07a9f",
		namespace:     "http://baseline.openssf.org/ns/oscal",
		wantErr:       false,
		expectedTitle: "Test Catalog Multiple",
	},
}

func Test_toOSCAL(t *testing.T) {
	for _, tt := range TestCases {
		t.Run(tt.name, func(t *testing.T) {
			oscalCatalog, err := tt.catalog.ToOSCAL(
				tt.controlFamilyIDs,
				tt.version,
				tt.controlHREF,
				tt.catalogUUID,
				tt.namespace,
			)

			if (err == nil) == tt.wantErr {
				t.Errorf("ToOSCAL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}
			// Wrap oscal catalog
			// Create the proper OSCAL document structure
			oscalDocument := map[string]interface{}{
				"catalog": oscalCatalog,
			}

			// Create validation for the OSCAL catalog
			validator, err := validation.NewValidator(oscalDocument)
			if err != nil {
				t.Errorf("Failed to create validator: %v", err)
				return
			}
			// Validate the OSCAL document
			err = validator.Validate()
			if err != nil {
				t.Errorf("OSCAL validation failed: %v", err)
				return
			}
			// Compare each field
			assert.Equal(t, tt.catalogUUID, oscalCatalog.UUID)
			assert.Equal(t, tt.expectedTitle, oscalCatalog.Metadata.Title)
			assert.Equal(t, tt.version, oscalCatalog.Metadata.Version)
			assert.Equal(t, len(tt.catalog.ControlFamilies), len(*oscalCatalog.Groups))

			// Compare each control family
			for i, family := range tt.catalog.ControlFamilies {
				groups := (*oscalCatalog.Groups)
				group := groups[i]
				assert.Equal(t, group.ID, tt.controlFamilyIDs[family.Id])
			}
		})
	}
}
