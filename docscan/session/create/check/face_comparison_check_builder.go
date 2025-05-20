package check

import "github.com/getyoti/yoti-go-sdk/v3/docscan/constants"

// RequestedFaceComparisonCheckBuilder builds a RequestedFaceComparisonCheck.
type RequestedFaceComparisonCheckBuilder struct {
	manualCheck string
}

// NewRequestedFaceComparisonCheckBuilder creates a new builder.
func NewRequestedFaceComparisonCheckBuilder() *RequestedFaceComparisonCheckBuilder {
	return &RequestedFaceComparisonCheckBuilder{}
}

// WithManualCheckNever sets the manual check mode to NEVER.
func (b *RequestedFaceComparisonCheckBuilder) WithManualCheckNever() *RequestedFaceComparisonCheckBuilder {
	b.manualCheck = constants.Never
	return b
}

func (b *RequestedFaceComparisonCheckBuilder) WithManualCheckAlways() *RequestedFaceComparisonCheckBuilder {
	b.manualCheck = constants.Always
	return b
}

func (b *RequestedFaceComparisonCheckBuilder) WithManualCheckFallback() *RequestedFaceComparisonCheckBuilder {
	b.manualCheck = constants.Fallback
	return b
}

// Build constructs the RequestedFaceComparisonCheck.
func (b *RequestedFaceComparisonCheckBuilder) Build() (*RequestedFaceComparisonCheck, error) {
	config := RequestedFaceComparisonConfig{
		ManualCheck: b.manualCheck,
	}
	return &RequestedFaceComparisonCheck{
		config: config,
	}, nil
}
