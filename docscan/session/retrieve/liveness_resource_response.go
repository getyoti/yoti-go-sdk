package retrieve

// LivenessResourceResponse represents a Liveness resource for a given session
type LivenessResourceResponse struct {
	*ResourceResponse
	LivenessType string `json:"liveness_type"`
}
