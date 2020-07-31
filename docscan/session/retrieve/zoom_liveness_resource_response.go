package retrieve

// ZoomLivenessResourceResponse represents a Zoom Liveness resource for a given session
type ZoomLivenessResourceResponse struct {
	LivenessResourceResponse
	FaceMap FaceMapResponse `json:"facemap"`
	Frames  []FrameResponse `json:"frames"`
}
