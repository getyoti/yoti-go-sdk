// face_capture_resource_response.go
package retrieve

// FaceCaptureResourceResponse models the response for face capture resource.
type FaceCaptureResourceResponse struct {
	ID     string `json:"id"`
	Frames int    `json:"frames"`
}

// GetID returns the ID of the face capture resource.
func (r *FaceCaptureResourceResponse) GetID() string {
	return r.ID
}

// GetFrames returns the number of image frames required.
func (r *FaceCaptureResourceResponse) GetFrames() int {
	return r.Frames
}
