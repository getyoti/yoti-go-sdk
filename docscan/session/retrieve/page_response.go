package retrieve

// PageResponse represents information about an uploaded document Page
type PageResponse struct {
	CaptureMethod string          `json:"capture_method"`
	Media         []MediaResponse `json:"media"`
}
