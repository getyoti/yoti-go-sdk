package retrieve

import (
	"time"
)

// TaskResponse represents the attributes of a task, for any given session
type TaskResponse struct {
	ID              string                   `json:"id"`
	Type            string                   `json:"type"`
	State           string                   `json:"state"`
	Created         *time.Time               `json:"created"`
	LastUpdated     *time.Time               `json:"last_updated"`
	GeneratedChecks []GeneratedCheckResponse `json:"generated_checks"`
	GeneratedMedia  []GeneratedMedia         `json:"generated_media"`
}
