package retrieve

import "time"

// CheckResponse represents the attributes of a check, for any given session
type CheckResponse struct {
	ID             string           `json:"id"`
	Type           string           `json:"type"`
	State          string           `json:"state"`
	ResourcesUsed  []string         `json:"resources_used"`
	GeneratedMedia []GeneratedMedia `json:"generated_media"`
	Report         ReportResponse   `json:"report"`
	Created        *time.Time       `json:"created"`
	LastUpdated    *time.Time       `json:"last_updated"`
}
