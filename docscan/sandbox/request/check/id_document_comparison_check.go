package check

import (
	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox/request/check/report"
	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox/request/filter"
)

// IDDocumentComparisonCheck Represents a document authenticity check
type IDDocumentComparisonCheck struct {
	*check
	SecondaryDocumentFilter *filter.DocumentFilter `json:"secondary_document_filter,omitempty"`
}

// IDDocumentComparisonCheckBuilder builds a IDDocumentComparisonCheck
type IDDocumentComparisonCheckBuilder struct {
	checkBuilder
	secondaryDocumentFilter *filter.DocumentFilter
}

// NewIDDocumentComparisonCheckBuilder creates a new IDDocumentComparisonCheckBuilder
func NewIDDocumentComparisonCheckBuilder() *IDDocumentComparisonCheckBuilder {
	return &IDDocumentComparisonCheckBuilder{}
}

// WithRecommendation sets the recommendation on the check
func (b *IDDocumentComparisonCheckBuilder) WithRecommendation(recommendation *report.Recommendation) *IDDocumentComparisonCheckBuilder {
	b.checkBuilder.withRecommendation(recommendation)
	return b
}

// WithBreakdown adds a breakdown item to the check
func (b *IDDocumentComparisonCheckBuilder) WithBreakdown(breakdown *report.Breakdown) *IDDocumentComparisonCheckBuilder {
	b.checkBuilder.withBreakdown(breakdown)
	return b
}

// WithSecondaryDocumentFilter adds a secondary document filter to the check
func (b *IDDocumentComparisonCheckBuilder) WithSecondaryDocumentFilter(filter *filter.DocumentFilter) *IDDocumentComparisonCheckBuilder {
	b.secondaryDocumentFilter = filter
	return b
}

// Build creates a new IDDocumentComparisonCheck
func (b *IDDocumentComparisonCheckBuilder) Build() (*IDDocumentComparisonCheck, error) {
	return &IDDocumentComparisonCheck{
		check:                   b.checkBuilder.build(),
		SecondaryDocumentFilter: b.secondaryDocumentFilter,
	}, nil
}
