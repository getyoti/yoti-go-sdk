package check

import (
	"encoding/json"

	"github.com/getyoti/yoti-go-sdk/v3/docscan/constants"
)

// RequestedFaceMatchCheck requests creation of a FaceMatch Check
type RequestedFaceMatchCheck struct {
	config RequestedFaceMatchConfig
}

// Type is the type of the Requested Check
func (c RequestedFaceMatchCheck) Type() string {
	return constants.IDDocumentFaceMatch
}

// Config is the configuration of the Requested Check
func (c RequestedFaceMatchCheck) Config() RequestedCheckConfig {
	return c.config
}

func (c RequestedFaceMatchCheck) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type   string               `json:"type"`
		Config RequestedCheckConfig `json:"config,omitempty"`
	}{
		Type:   c.Type(),
		Config: c.Config(),
	})
}

// RequestedFaceMatchConfig is the configuration applied when creating a FaceMatch Check
type RequestedFaceMatchConfig struct {
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
func (b *RequestedFaceMatchCheckBuilder) WithManualCheckAlways() *RequestedFaceMatchCheckBuilder {
	b.manualCheck = constants.Always
	return b
}

// WithManualCheckFallback sets the value of manual check to "FALLBACK"
func (b *RequestedFaceMatchCheckBuilder) WithManualCheckFallback() *RequestedFaceMatchCheckBuilder {
	b.manualCheck = constants.Fallback
	return b
}

// WithManualCheckNever sets the value of manual check to "NEVER"
func (b *RequestedFaceMatchCheckBuilder) WithManualCheckNever() *RequestedFaceMatchCheckBuilder {
	b.manualCheck = constants.Never
	return b
}

// Build builds the RequestedFaceMatchCheck
func (b *RequestedFaceMatchCheckBuilder) Build() (*RequestedFaceMatchCheck, error) {
	config := RequestedFaceMatchConfig{
		ManualCheck: b.manualCheck,
	}

	return &RequestedFaceMatchCheck{
		config: config,
	}, nil
}
