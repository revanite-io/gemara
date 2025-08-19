package layer2

import (
	"fmt"
	"strings"
	"time"

	oscal "github.com/defenseunicorns/go-oscal/src/types/oscal-1-1-3"
)

// ToOSCAL converts a Catalog to OSCAL Catalog format.
// Parameters:
//   - controlFamilyIDs: Maps internal family IDs to OSCAL group IDs (e.g., "AC" -> "AC" for Access Control)
//   - version: Catalog version string (e.g., "1.0.0", "devel")
//   - controlHREF: URL template for control links, uses fmt.Sprintf(controlHREF, version, controlID)
//   - catalogUUID: Unique identifier for the OSCAL catalog
//   - namespace: XML namespace for OSCAL elements (e.g., "http://baseline.openssf.org/ns/oscal")
//
// TODO: Consider using go-oscal's UUID generation for future OSCAL elements:
// - uuid.NewUUID() for random UUIDs in production
// - uuid.NewUUIDWithSource() for deterministic UUIDs in testing
func (c *Catalog) ToOSCAL(controlFamilyIDs map[string]string,
	version, controlHREF, catalogUUID, namespace string) (oscal.Catalog, error) {
	now := time.Now()
	lastModified := now
	if c.Metadata.LastModified != "" {
		if parsedTime, err := time.Parse(time.RFC3339, c.Metadata.LastModified); err == nil {
			lastModified = parsedTime
		}
	}
	oscalCatalog := oscal.Catalog{
		UUID:   catalogUUID,
		Groups: nil,
		Metadata: oscal.Metadata{
			LastModified: lastModified,
			Links: &[]oscal.Link{
				{
					Href: fmt.Sprintf(controlHREF, version, ""),
					Rel:  "canonical",
				},
			},
			OscalVersion: "1.1.3",
			Published:    &now,
			Title:        c.Metadata.Title,
			Version:      version,
		},
	}

	catalogGroups := []oscal.Group{}

	for _, family := range c.ControlFamilies {
		group := oscal.Group{
			Class:    "family",
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
									Href: fmt.Sprintf(controlHREF, version, ar.Id),
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
						Href: fmt.Sprintf(controlHREF, version, strings.ToLower(control.Id)),
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
