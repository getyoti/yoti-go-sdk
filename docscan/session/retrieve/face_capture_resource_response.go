package retrieve

import (
	"encoding/json"
)

// FaceCaptureResourceResponse represents a face capture resource response for a given session
type FaceCaptureResourceResponse struct {
	ID     string            `json:"id"`
	Source Source            `json:"source"`
	Image  *Image            `json:"image"`
	Tasks  []json.RawMessage `json:"tasks"`
}
