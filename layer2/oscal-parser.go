package layer2

import (
	"fmt"
	"strings"
	"time"

	oscal "github.com/defenseunicorns/go-oscal/src/types/oscal-1-1-3"
)

// ToOSCAL converts a Catalog to OSCAL Catalog format.
func (c *Catalog) ToOSCAL(controlFamilyIDs map[string]string,
	versionOSPS, controlHREF, catalogUUID, namespace string) (oscal.Catalog, error) {

	now := time.Now()
	oscalCatalog := oscal.Catalog{
		UUID:   catalogUUID,
		Groups: nil,
		Metadata: oscal.Metadata{
			LastModified: now,
			Links: &[]oscal.Link{
				{
					Href: fmt.Sprintf(controlHREF, versionOSPS, ""),
					Rel:  "canonical",
				},
			},
			OscalVersion: "1.1.3",
			Published:    &now,
			Title:        c.Metadata.Title,
			Version:      versionOSPS,
		},
	}

	catalogGroups := []oscal.Group{}

	for _, family := range c.ControlFamilies {
		group := oscal.Group{
			Class:    "OSPS",
			Controls: nil,
			ID:       controlFamilyIDs[family.Id],
			Title:    family.Description,
		}

		controls := []oscal.Control{}
		for _, control := range family.Controls {
			parts := []oscal.Part{}
			for _, ar := range control.AssessmentRequirements {
				parts = append(parts, oscal.Part{
					Class: control.Id,
					ID:    ar.Id,
					Name:  ar.Id,
					Ns:    "",
					Parts: &[]oscal.Part{
						{
							ID:    ar.Id + ".R",
							Name:  "recommendation",
							Ns:    namespace,
							Prose: ar.Recommendation,
							Links: &[]oscal.Link{
								{
									Href: fmt.Sprintf(controlHREF, versionOSPS, ar.Id),
									Rel:  "canonical",
								},
							},
						},
					},
					Prose: ar.Text,
					Title: "",
				})
			}

			newCtl := oscal.Control{
				Class: family.Title,
				ID:    control.Id,
				Links: &[]oscal.Link{
					{
						Href: fmt.Sprintf(controlHREF, versionOSPS, strings.ToLower(control.Id)),
						Rel:  "canonical",
					},
				},
				Parts: &parts,
				Title: strings.TrimSpace(control.Title),
			}
			controls = append(controls, newCtl)
		}

		group.Controls = &controls
		catalogGroups = append(catalogGroups, group)
	}
	oscalCatalog.Groups = &catalogGroups

	return oscalCatalog, nil
}
