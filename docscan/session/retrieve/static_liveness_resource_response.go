package retrieve

// StaticLivenessResourceResponse represents a Static Liveness resource for a given session
type StaticLivenessResourceResponse struct {
	*LivenessResourceResponse
	FaceMap *FaceMapResponse `json:"facemap"`
	Frames  []*FrameResponse `json:"frames"`
}
