package check

import (
	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox/request/check/report"
	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox/request/filter"
)

// DocumentTextDataCheck Represents a document text data check
type DocumentTextDataCheck struct {
	Result DocumentTextDataCheckResult `json:"result"`
	*documentCheck
}

// DocumentTextDataCheckBuilder builds a DocumentTextDataCheck
type DocumentTextDataCheckBuilder struct {
	documentCheckBuilder
	documentFields map[string]interface{}
}

// DocumentTextDataCheckResult represent a document text data check result
type DocumentTextDataCheckResult struct {
	checkResult
	DocumentFields map[string]interface{} `json:"document_fields,omitempty"`
}

// NewDocumentTextDataCheckBuilder builds a new DocumentTextDataCheckResult
func NewDocumentTextDataCheckBuilder() *DocumentTextDataCheckBuilder {
	return &DocumentTextDataCheckBuilder{}
}

// WithRecommendation sets the recommendation on the check
func (b *DocumentTextDataCheckBuilder) WithRecommendation(recommendation *report.Recommendation) *DocumentTextDataCheckBuilder {
	b.documentCheckBuilder.withRecommendation(recommendation)
	return b
}

// WithBreakdown adds a breakdown item to the check
func (b *DocumentTextDataCheckBuilder) WithBreakdown(breakdown *report.Breakdown) *DocumentTextDataCheckBuilder {
	b.documentCheckBuilder.withBreakdown(breakdown)
	return b
}

// WithDocumentFilter adds a document filter to the check
func (b *DocumentTextDataCheckBuilder) WithDocumentFilter(filter *filter.DocumentFilter) *DocumentTextDataCheckBuilder {
	b.documentCheckBuilder.withDocumentFilter(filter)
	return b
}

// WithDocumentField adds a document field to the text data check
func (b *DocumentTextDataCheckBuilder) WithDocumentField(key string, value interface{}) *DocumentTextDataCheckBuilder {
	if b.documentFields == nil {
		b.documentFields = make(map[string]interface{})
	}
	b.documentFields[key] = value
	return b
}

// WithDocumentFields sets document fields
func (b *DocumentTextDataCheckBuilder) WithDocumentFields(documentFields map[string]interface{}) *DocumentTextDataCheckBuilder {
	b.documentFields = documentFields
	return b
}

// Build creates a new DocumentTextDataCheck
func (b *DocumentTextDataCheckBuilder) Build() (*DocumentTextDataCheck, error) {
	documentCheck := b.documentCheckBuilder.build()

	return &DocumentTextDataCheck{
		documentCheck: documentCheck,
		Result: DocumentTextDataCheckResult{
			checkResult:    documentCheck.Result,
			DocumentFields: b.documentFields,
		},
	}, nil
}
