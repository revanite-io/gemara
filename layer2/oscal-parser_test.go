package layer2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var oscalTests = []struct {
	name             string
	catalog          *Catalog
	controlFamilyIDs map[string]string
	versionOSPS      string
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
					Title:       "Access Control",
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
		versionOSPS:   "devel",
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
					Title:       "Access Control",
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
					Title:       "Business Requirements",
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
		versionOSPS:   "devel",
		controlHREF:   "https://baseline.openssf.org/versions/%s#%s",
		catalogUUID:   "8c222a23-fc7e-4ad8-b6dd-289014f07a9f",
		namespace:     "http://baseline.openssf.org/ns/oscal",
		wantErr:       false,
		expectedTitle: "Test Catalog Multiple",
	},
	{
		name: "Empty catalog",
		catalog: &Catalog{
			Metadata: Metadata{
				Id:    "empty-catalog",
				Title: "Empty Catalog",
			},
			ControlFamilies: []ControlFamily{},
		},
		controlFamilyIDs: map[string]string{},
		versionOSPS:      "devel",
		controlHREF:      "https://baseline.openssf.org/versions/%s#%s",
		catalogUUID:      "8c222a23-fc7e-4ad8-b6dd-289014f07a9f",
		namespace:        "http://baseline.openssf.org/ns/oscal",
		wantErr:          false,
		expectedTitle:    "Empty Catalog",
	},
}

func Test_ToOSCAL(t *testing.T) {
	for _, tt := range oscalTests {
		t.Run(tt.name, func(t *testing.T) {
			oscalCatalog, err := tt.catalog.ToOSCAL(
				tt.controlFamilyIDs,
				tt.versionOSPS,
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

			assert.Equal(t, tt.catalogUUID, oscalCatalog.UUID)
			assert.Equal(t, tt.expectedTitle, oscalCatalog.Metadata.Title)
			assert.Equal(t, tt.versionOSPS, oscalCatalog.Metadata.Version)
			assert.Equal(t, len(tt.catalog.ControlFamilies), len(*oscalCatalog.Groups))

			for i, family := range tt.catalog.ControlFamilies {
				groups := (*oscalCatalog.Groups)
				group := groups[i]
				assert.Equal(t, group.ID, tt.controlFamilyIDs[family.Id])
			}
		})
	}
}
