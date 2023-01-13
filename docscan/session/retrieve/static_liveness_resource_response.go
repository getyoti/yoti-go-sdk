package retrieve

// StaticLivenessResourceResponse represents a Static Liveness resource for a given session
type StaticLivenessResourceResponse struct {
	*LivenessResourceResponse
	Image *Image `json:"image"`
}

type Image struct {
	Media *MediaResponse `json:"media"`
}
