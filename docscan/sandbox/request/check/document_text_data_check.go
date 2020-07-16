package check

import (
	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox/request/check/report"
	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox/request/filter"
)

// DocumentTextDataCheck Represents a document text data check
type DocumentTextDataCheck struct {
	Result DocumentTextDataCheckResult `json:"result"`
	documentCheck
}

// DocumentTextDataCheckBuilder builds a DocumentTextDataCheck
type DocumentTextDataCheckBuilder struct {
	documentCheckBuilder
	documentFields map[string]string
}

// DocumentTextDataCheckResult represent a document text data check result
type DocumentTextDataCheckResult struct {
	checkResult
	DocumentFields map[string]string `json:"document_fields"`
}

// NewDocumentTextDataCheckBuilder builds a new DocumentTextDataCheckResult
func NewDocumentTextDataCheckBuilder() *DocumentTextDataCheckBuilder {
	return &DocumentTextDataCheckBuilder{}
}

// WithRecommendation sets the recommendation on the check
func (b *DocumentTextDataCheckBuilder) WithRecommendation(recommendation report.Recommendation) *DocumentTextDataCheckBuilder {
	b.documentCheckBuilder.withRecommendation(recommendation)
	return b
}

// WithBreakdown adds a breakdown item to the check
func (b *DocumentTextDataCheckBuilder) WithBreakdown(breakdown report.Breakdown) *DocumentTextDataCheckBuilder {
	b.documentCheckBuilder.withBreakdown(breakdown)
	return b
}

// WithDocumentFilter adds a document filter to the check
func (b *DocumentTextDataCheckBuilder) WithDocumentFilter(filter filter.DocumentFilter) *DocumentTextDataCheckBuilder {
	b.documentCheckBuilder.withDocumentFilter(filter)
	return b
}

// WithDocumentField adds a document field to the text data check
func (b *DocumentTextDataCheckBuilder) WithDocumentField(key string, value string) *DocumentTextDataCheckBuilder {
	if b.documentFields == nil {
		b.documentFields = make(map[string]string)
	}
	b.documentFields[key] = value
	return b
}

// Build creates a new DocumentTextDataCheck
func (b *DocumentTextDataCheckBuilder) Build() (DocumentTextDataCheck, error) {
	documentTextDataCheck := DocumentTextDataCheck{}

	documentCheck, err := b.documentCheckBuilder.build()
	if err != nil {
		return documentTextDataCheck, err
	}

	documentTextDataCheck.documentCheck = documentCheck
	documentTextDataCheck.Result = DocumentTextDataCheckResult{
		checkResult:    documentCheck.Result,
		DocumentFields: b.documentFields,
	}

	return documentTextDataCheck, nil
}
