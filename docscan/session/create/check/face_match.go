package check

import (
	"github.com/getyoti/yoti-go-sdk/v3/docscan/constants"
)

// RequestedFaceMatchCheck requests creation of a FaceMatch Check
type RequestedFaceMatchCheck struct {
	*RequestedCheck
	Config *RequestedFaceMatchConfig `json:"config"`
}

// NewRequestedFaceMatchCheck creates a new Document Authenticity Check
func NewRequestedFaceMatchCheck(config *RequestedFaceMatchConfig) *RequestedFaceMatchCheck {
	return &RequestedFaceMatchCheck{
		RequestedCheck: &RequestedCheck{
			Type: constants.IDDocumentFaceMatch,
		},
		Config: config,
	}
}

// RequestedFaceMatchConfig is the configuration applied when creating a FaceMatch Check
type RequestedFaceMatchConfig struct {
	*RequestedCheckConfig
	ManualCheck string `json:"manual_check"`
}

// NewRequestedFaceMatchCheckBuilder creates a new RequestedFaceMatchCheckBuilder
func NewRequestedFaceMatchCheckBuilder() *RequestedFaceMatchCheckBuilder {
	return &RequestedFaceMatchCheckBuilder{}
}

// RequestedFaceMatchCheckBuilder builds a RequestedFaceMatchCheck
type RequestedFaceMatchCheckBuilder struct {
	manualCheck string
}

// WithManualCheckAlways sets the value of manual check to "ALWAYS"
func (builder *RequestedFaceMatchCheckBuilder) WithManualCheckAlways() *RequestedFaceMatchCheckBuilder {
	builder.manualCheck = constants.Always
	return builder
}

// WithManualCheckFallback sets the value of manual check to "FALLBACK"
func (builder *RequestedFaceMatchCheckBuilder) WithManualCheckFallback() *RequestedFaceMatchCheckBuilder {
	builder.manualCheck = constants.Fallback
	return builder
}

// WithManualCheckNever sets the value of manual check to "NEVER"
func (builder *RequestedFaceMatchCheckBuilder) WithManualCheckNever() *RequestedFaceMatchCheckBuilder {
	builder.manualCheck = constants.Never
	return builder
}

// Build builds the RequestedFaceMatchCheck
func (builder *RequestedFaceMatchCheckBuilder) Build() (*RequestedFaceMatchCheck, error) {
	return NewRequestedFaceMatchCheck(
		&RequestedFaceMatchConfig{
			ManualCheck: builder.manualCheck,
		}), nil
}
