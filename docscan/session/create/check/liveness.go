package check

import (
	"github.com/getyoti/yoti-go-sdk/v3/docscan/constants"
)

// RequestedLivenessCheck requests creation of a Liveness Check
type RequestedLivenessCheck struct {
	RequestedCheck                          //TODO: pointer in embedded type?
	Config         *RequestedLivenessConfig `json:"config"`
}

// NewRequestedLivenessCheck creates a new Document Authenticity Check
func NewRequestedLivenessCheck(config *RequestedLivenessConfig) *RequestedLivenessCheck {
	return &RequestedLivenessCheck{
		RequestedCheck: RequestedCheck{
			Type: constants.Liveness,
		},
		Config: config,
	}
}

// RequestedLivenessConfig is the configuration applied when creating a Liveness Check
type RequestedLivenessConfig struct {
	RequestedCheckConfig
	MaxRetries   *int   `json:"max_retries"`
	LivenessType string `json:"liveness_type"`
}

// NewRequestedLivenessCheckBuilder creates a new RequestedLivenessCheckBuilder
func NewRequestedLivenessCheckBuilder() *RequestedLivenessCheckBuilder {
	return &RequestedLivenessCheckBuilder{}
}

// RequestedLivenessCheckBuilder builds a RequestedLivenessCheck
type RequestedLivenessCheckBuilder struct {
	livenessType string
	maxRetries   *int
	//TODO: do we want this with pointer so we can set to nil?
}

// ForZoomLiveness sets the liveness type to "ZOOM"
func (builder *RequestedLivenessCheckBuilder) ForZoomLiveness() *RequestedLivenessCheckBuilder {
	return builder.ForLivenessType(constants.Zoom)
}

// ForLivenessType sets the liveness type on the builder
func (builder *RequestedLivenessCheckBuilder) ForLivenessType(livenessType string) *RequestedLivenessCheckBuilder {
	builder.livenessType = livenessType
	return builder
}

// WithMaxRetries sets the maximum number of retries allowed for liveness check on the builder
func (builder *RequestedLivenessCheckBuilder) WithMaxRetries(maxRetries int) *RequestedLivenessCheckBuilder {
	builder.maxRetries = &maxRetries
	return builder
}

// Build builds the RequestedLivenessCheck
func (builder *RequestedLivenessCheckBuilder) Build() (*RequestedLivenessCheck, error) {

	config := &RequestedLivenessConfig{
		RequestedCheckConfig: RequestedCheckConfig{},
		MaxRetries:           builder.maxRetries,
		LivenessType:         builder.livenessType,
	}

	return NewRequestedLivenessCheck(config), nil
}
