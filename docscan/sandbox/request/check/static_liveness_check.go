package check

import (
	"github.com/getyoti/yoti-go-sdk/v3/docscan/constants"
	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox/request/check/report"
)

// StaticLivenessCheckBuilder builds a "ZOOM" LivenessCheck
type StaticLivenessCheckBuilder struct {
	livenessCheckBuilder
}

// NewStaticLivenessCheckBuilder creates a new StaticLivenessCheckBuilder
func NewStaticLivenessCheckBuilder() *StaticLivenessCheckBuilder {
	return &StaticLivenessCheckBuilder{}
}

// WithRecommendation sets the recommendation on the check
func (b *StaticLivenessCheckBuilder) WithRecommendation(recommendation *report.Recommendation) *StaticLivenessCheckBuilder {
	b.livenessCheckBuilder.withRecommendation(recommendation)
	return b
}

// WithBreakdown adds a breakdown item to the check
func (b *StaticLivenessCheckBuilder) WithBreakdown(breakdown *report.Breakdown) *StaticLivenessCheckBuilder {
	b.livenessCheckBuilder.withBreakdown(breakdown)
	return b
}

// Build creates a new LivenessCheck
func (b *StaticLivenessCheckBuilder) Build() (*LivenessCheck, error) {
	livenessCheck := b.livenessCheckBuilder.
		withLivenessType(constants.Zoom).
		build()

	return livenessCheck, nil
}
