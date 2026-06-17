package check

import (
	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox/request/check/report"
	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox/request/filter"
)

// DocumentAuthenticityCheck Represents a document authenticity check
type DocumentAuthenticityCheck struct {
	*documentCheck
}

// DocumentAuthenticityCheckBuilder builds a DocumentAuthenticityCheck
type DocumentAuthenticityCheckBuilder struct {
	documentCheckBuilder
}

// NewDocumentAuthenticityCheckBuilder creates a new DocumentAuthenticityCheckBuilder
func NewDocumentAuthenticityCheckBuilder() *DocumentAuthenticityCheckBuilder {
	return &DocumentAuthenticityCheckBuilder{}
}

// WithRecommendation sets the recommendation on the check
func (b *DocumentAuthenticityCheckBuilder) WithRecommendation(recommendation *report.Recommendation) *DocumentAuthenticityCheckBuilder {
	b.documentCheckBuilder.withRecommendation(recommendation)
	return b
}

// WithBreakdown adds a breakdown item to the check
func (b *DocumentAuthenticityCheckBuilder) WithBreakdown(breakdown *report.Breakdown) *DocumentAuthenticityCheckBuilder {
	b.documentCheckBuilder.withBreakdown(breakdown)
	return b
}

// WithDocumentFilter adds a document filter to the check
func (b *DocumentAuthenticityCheckBuilder) WithDocumentFilter(filter *filter.DocumentFilter) *DocumentAuthenticityCheckBuilder {
	b.documentCheckBuilder.withDocumentFilter(filter)
	return b
}

// WithHandledCheckLimit sets the number of times this check report configuration will be used
// by the sandbox before moving to the next configured response.
func (b *DocumentAuthenticityCheckBuilder) WithHandledCheckLimit(handledCheckLimit int) *DocumentAuthenticityCheckBuilder {
	b.withHandledCheckLimit(handledCheckLimit)
	return b
}

// Build creates a new DocumentAuthenticityCheck
func (b *DocumentAuthenticityCheckBuilder) Build() (*DocumentAuthenticityCheck, error) {
	return &DocumentAuthenticityCheck{
		documentCheck: b.documentCheckBuilder.build(),
	}, nil
}
