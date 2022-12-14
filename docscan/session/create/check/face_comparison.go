package check

import (
	"encoding/json"

	"github.com/getyoti/yoti-go-sdk/v3/docscan/constants"
)

// RequestedFaceMatchCheck requests creation of a FaceMatch Check
type RequestedFaceComparisonCheck struct {
	config RequestedFaceComparisonConfig
}

// Type is the type of the Requested Check
func (c *RequestedFaceComparisonCheck) Type() string {
	return constants.FaceComparison
}

// Config is the configuration of the Requested Check
func (c *RequestedFaceComparisonCheck) Config() RequestedCheckConfig {
	return c.config
}

// MarshalJSON returns the JSON encoding
func (c *RequestedFaceComparisonCheck) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type   string               `json:"type"`
		Config RequestedCheckConfig `json:"config,omitempty"`
	}{
		Type:   c.Type(),
		Config: c.Config(),
	})
}

// RequestedFaceMatchConfig is the configuration applied when creating a FaceMatch Check
type RequestedFaceComparisonConfig struct {
	ManualCheck string `json:"manual_check,omitempty"`
}

// NewRequestedFaceMatchCheckBuilder creates a new RequestedFaceMatchCheckBuilder
func NewRequestedFaceComparisonCheckBuilder() *RequestedFaceComparisonCheckBuilder {
	return &RequestedFaceComparisonCheckBuilder{}
}

// RequestedFaceMatchCheckBuilder builds a RequestedFaceMatchCheck
type RequestedFaceComparisonCheckBuilder struct {
	manualCheck string
}

// WithManualCheckAlways sets the value of manual check to "ALWAYS"
func (b *RequestedFaceComparisonCheckBuilder) WithManualCheckAlways() *RequestedFaceComparisonCheckBuilder {
	b.manualCheck = constants.Always
	return b
}

// WithManualCheckFallback sets the value of manual check to "FALLBACK"
func (b *RequestedFaceComparisonCheckBuilder) WithManualCheckFallback() *RequestedFaceComparisonCheckBuilder {
	b.manualCheck = constants.Fallback
	return b
}

// WithManualCheckNever sets the value of manual check to "NEVER"
func (b *RequestedFaceComparisonCheckBuilder) WithManualCheckNever() *RequestedFaceComparisonCheckBuilder {
	b.manualCheck = constants.Never
	return b
}

// Build builds the RequestedFaceMatchCheck
func (b *RequestedFaceComparisonCheckBuilder) Build() (*RequestedFaceComparisonCheck, error) {
	config := RequestedFaceComparisonConfig{
		ManualCheck: b.manualCheck,
	}

	return &RequestedFaceComparisonCheck{
		config: config,
	}, nil
}
