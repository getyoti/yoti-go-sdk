package check

import (
	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox/request/check/report"
	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox/request/filter"
)

// DocumentFaceMatchCheck represents a face match check
type DocumentFaceMatchCheck struct {
	documentCheck
}

// DocumentFaceMatchCheckBuilder builds a DocumentFaceMatchCheck
type DocumentFaceMatchCheckBuilder struct {
	documentCheckBuilder
}

// NewDocumentFaceMatchCheckBuilder creates a new DocumentFaceMatchCheckBuilder
func NewDocumentFaceMatchCheckBuilder() *DocumentFaceMatchCheckBuilder {
	return &DocumentFaceMatchCheckBuilder{}
}

// WithRecommendation sets the recommendation on the check
func (b *DocumentFaceMatchCheckBuilder) WithRecommendation(recommendation report.Recommendation) *DocumentFaceMatchCheckBuilder {
	b.documentCheckBuilder.withRecommendation(recommendation)
	return b
}

// WithBreakdown adds a breakdown item to the check
func (b *DocumentFaceMatchCheckBuilder) WithBreakdown(breakdown report.Breakdown) *DocumentFaceMatchCheckBuilder {
	b.documentCheckBuilder.withBreakdown(breakdown)
	return b
}

// WithDocumentFilter adds a document filter to the check
func (b *DocumentFaceMatchCheckBuilder) WithDocumentFilter(filter filter.DocumentFilter) *DocumentFaceMatchCheckBuilder {
	b.documentCheckBuilder.withDocumentFilter(filter)
	return b
}

// Build creates a new DocumentFaceMatchCheck
func (b *DocumentFaceMatchCheckBuilder) Build() (DocumentFaceMatchCheck, error) {
	return DocumentFaceMatchCheck{
		documentCheck: b.documentCheckBuilder.build(),
	}, nil
}
