package check

import (
	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox/check/report"
)

const (
	zoom = "ZOOM"
)

type ZoomLivenessCheck struct {
	LivenessCheck
}

type zoomLivenessCheckBuilder struct {
	livenessCheckBuilder
}

func NewZoomLivenessCheckBuilder() *zoomLivenessCheckBuilder {
	return &zoomLivenessCheckBuilder{}
}

func (b *zoomLivenessCheckBuilder) WithRecommendation(recommendation report.Recommendation) *zoomLivenessCheckBuilder {
	b.livenessCheckBuilder.withRecommendation(recommendation)
	return b
}

func (b *zoomLivenessCheckBuilder) WithBreakdown(breakdown report.Breakdown) *zoomLivenessCheckBuilder {
	b.livenessCheckBuilder.withBreakdown(breakdown)
	return b
}

func (b *zoomLivenessCheckBuilder) Build() (LivenessCheck, error) {
	livenessCheck, err := b.livenessCheckBuilder.withLivenessType(zoom).build()
	if err != nil {
		return livenessCheck, err
	}

	return livenessCheck, nil
}
