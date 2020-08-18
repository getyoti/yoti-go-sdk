package retrieve

import "time"

// MediaResponse represents a media resource
type MediaResponse struct {
	ID          string    `json:"id"`
	Type        string    `json:"type"`
	Created     time.Time `json:"created"`
	LastUpdated time.Time `json:"last_updated"`
}
