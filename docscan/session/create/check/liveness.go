package check

import (
	"encoding/json"

	"github.com/getyoti/yoti-go-sdk/v3/docscan/constants"
)

// RequestedLivenessCheck requests creation of a Liveness Check
type RequestedLivenessCheck struct {
	config RequestedLivenessConfig
}

// Type is the type of the Requested Check
func (c *RequestedLivenessCheck) Type() string {
	return constants.Liveness
}

// Config is the configuration of the Requested Check
func (c *RequestedLivenessCheck) Config() RequestedCheckConfig {
	return RequestedCheckConfig(
		c.config,
	)
}

// MarshalJSON returns the JSON encoding
func (c *RequestedLivenessCheck) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type   string               `json:"type"`
		Config RequestedCheckConfig `json:"config,omitempty"`
	}{
		Type:   c.Type(),
		Config: c.Config(),
	})
}

// RequestedLivenessConfig is the configuration applied when creating a Liveness Check
type RequestedLivenessConfig struct {
	MaxRetries   int    `json:"max_retries,omitempty"`
	LivenessType string `json:"liveness_type,omitempty"`
	ManualCheck  string `json:"manual_check,omitempty"`
}

// NewRequestedLivenessCheckBuilder creates a new RequestedLivenessCheckBuilder
func NewRequestedLivenessCheckBuilder() *RequestedLivenessCheckBuilder {
	return &RequestedLivenessCheckBuilder{}
}

// RequestedLivenessCheckBuilder builds a RequestedLivenessCheck
type RequestedLivenessCheckBuilder struct {
	livenessType string
	maxRetries   int
	manualCheck  string
}

// ForZoomLiveness sets the liveness type to "ZOOM"
func (b *RequestedLivenessCheckBuilder) ForZoomLiveness() *RequestedLivenessCheckBuilder {
	return b.ForLivenessType(zoom)
}

// ForStaticLiveness sets the liveness type to "STATIC"
func (b *RequestedLivenessCheckBuilder) ForStaticLiveness() *RequestedLivenessCheckBuilder {
	return b.ForLivenessType(static)
}

// ForLivenessType sets the liveness type on the builder
func (b *RequestedLivenessCheckBuilder) ForLivenessType(livenessType string) *RequestedLivenessCheckBuilder {
	b.livenessType = livenessType
	return b
}

// WithMaxRetries sets the maximum number of retries allowed for liveness check on the builder
func (b *RequestedLivenessCheckBuilder) WithMaxRetries(maxRetries int) *RequestedLivenessCheckBuilder {
	b.maxRetries = maxRetries
	return b
}

// Build builds the RequestedLivenessCheck
func (b *RequestedLivenessCheckBuilder) Build() (*RequestedLivenessCheck, error) {
	config := RequestedLivenessConfig{
		MaxRetries:   b.maxRetries,
		LivenessType: b.livenessType,
		ManualCheck:  "NEVER",
	}

	return &RequestedLivenessCheck{
		config: config,
	}, nil
}
