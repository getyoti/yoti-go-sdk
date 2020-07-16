package check

import (
	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox/request/check/report"
)

const (
	zoom = "ZOOM"
)

// ZoomLivenessCheckBuilder builds a "ZOOM" LivenessCheck
type ZoomLivenessCheckBuilder struct {
	livenessCheckBuilder
}

// NewZoomLivenessCheckBuilder creates a new ZoomLivenessCheckBuilder
func NewZoomLivenessCheckBuilder() *ZoomLivenessCheckBuilder {
	return &ZoomLivenessCheckBuilder{}
}

// WithRecommendation sets the recommendation on the check
func (b *ZoomLivenessCheckBuilder) WithRecommendation(recommendation report.Recommendation) *ZoomLivenessCheckBuilder {
	b.livenessCheckBuilder.withRecommendation(recommendation)
	return b
}

// WithBreakdown adds a breakdown item to the check
func (b *ZoomLivenessCheckBuilder) WithBreakdown(breakdown report.Breakdown) *ZoomLivenessCheckBuilder {
	b.livenessCheckBuilder.withBreakdown(breakdown)
	return b
}

// Build creates a new LivenessCheck
func (b *ZoomLivenessCheckBuilder) Build() (LivenessCheck, error) {
	livenessCheck, err := b.livenessCheckBuilder.withLivenessType(zoom).build()
	if err != nil {
		return livenessCheck, err
	}

	return livenessCheck, nil
}
