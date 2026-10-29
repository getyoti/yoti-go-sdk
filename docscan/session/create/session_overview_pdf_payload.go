// Package create provides the request/payload types used to build a Yoti
// Doc Scan (IDV) session and related resources.
package create

// SessionOverviewPdfPayload specifies which sections to include when creating
// a session overview PDF. Any option left unset (nil) is omitted from the
// request, leaving the decision to the backend rather than the SDK.
type SessionOverviewPdfPayload struct {
	IncludeImages         *bool `json:"include_images,omitempty"`
	IncludeExtractedText  *bool `json:"include_extracted_text,omitempty"`
	IncludeDeviceTracking *bool `json:"include_device_tracking,omitempty"`
}

// NewSessionOverviewPdfPayload creates an empty payload. Use the With* methods
// to set only the options you want to specify.
func NewSessionOverviewPdfPayload() *SessionOverviewPdfPayload {
	return &SessionOverviewPdfPayload{}
}

// WithIncludeImages sets whether images should be included in the PDF.
func (p *SessionOverviewPdfPayload) WithIncludeImages(value bool) *SessionOverviewPdfPayload {
	p.IncludeImages = &value
	return p
}

// WithIncludeExtractedText sets whether extracted text should be included in the PDF.
func (p *SessionOverviewPdfPayload) WithIncludeExtractedText(value bool) *SessionOverviewPdfPayload {
	p.IncludeExtractedText = &value
	return p
}

// WithIncludeDeviceTracking sets whether device tracking information should be included in the PDF.
func (p *SessionOverviewPdfPayload) WithIncludeDeviceTracking(value bool) *SessionOverviewPdfPayload {
	p.IncludeDeviceTracking = &value
	return p
}
