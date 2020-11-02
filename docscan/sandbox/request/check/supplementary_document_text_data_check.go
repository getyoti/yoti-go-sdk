package check

import (
	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox/request/check/report"
	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox/request/filter"
)

// SupplementaryDocumentTextDataCheck represents a supplementary document text data check
type SupplementaryDocumentTextDataCheck struct {
	Result SupplementaryDocumentTextDataCheckResult `json:"result"`
	*documentCheck
}

// SupplementaryDocumentTextDataCheckBuilder builds a SupplementaryDocumentTextDataCheck
type SupplementaryDocumentTextDataCheckBuilder struct {
	documentCheckBuilder
	documentFields map[string]interface{}
}

// SupplementaryDocumentTextDataCheckResult represents a document text data check result
type SupplementaryDocumentTextDataCheckResult struct {
	checkResult
	DocumentFields map[string]interface{} `json:"document_fields,omitempty"`
}

// NewSupplementaryDocumentTextDataCheckBuilder builds a new SupplementaryDocumentTextDataCheckResult
func NewSupplementaryDocumentTextDataCheckBuilder() *SupplementaryDocumentTextDataCheckBuilder {
	return &SupplementaryDocumentTextDataCheckBuilder{}
}

// WithRecommendation sets the recommendation on the check
func (b *SupplementaryDocumentTextDataCheckBuilder) WithRecommendation(recommendation *report.Recommendation) *SupplementaryDocumentTextDataCheckBuilder {
	b.documentCheckBuilder.withRecommendation(recommendation)
	return b
}

// WithBreakdown adds a breakdown item to the check
func (b *SupplementaryDocumentTextDataCheckBuilder) WithBreakdown(breakdown *report.Breakdown) *SupplementaryDocumentTextDataCheckBuilder {
	b.documentCheckBuilder.withBreakdown(breakdown)
	return b
}

// WithDocumentFilter adds a document filter to the check
func (b *SupplementaryDocumentTextDataCheckBuilder) WithDocumentFilter(filter *filter.DocumentFilter) *SupplementaryDocumentTextDataCheckBuilder {
	b.documentCheckBuilder.withDocumentFilter(filter)
	return b
}

// WithDocumentField adds a document field to the text data check
func (b *SupplementaryDocumentTextDataCheckBuilder) WithDocumentField(key string, value interface{}) *SupplementaryDocumentTextDataCheckBuilder {
	if b.documentFields == nil {
		b.documentFields = make(map[string]interface{})
	}
	b.documentFields[key] = value
	return b
}

// WithDocumentFields sets document fields
func (b *SupplementaryDocumentTextDataCheckBuilder) WithDocumentFields(documentFields map[string]interface{}) *SupplementaryDocumentTextDataCheckBuilder {
	b.documentFields = documentFields
	return b
}

// Build creates a new SupplementaryDocumentTextDataCheck
func (b *SupplementaryDocumentTextDataCheckBuilder) Build() (*SupplementaryDocumentTextDataCheck, error) {
	docCheck := b.documentCheckBuilder.build()

	return &SupplementaryDocumentTextDataCheck{
		documentCheck: docCheck,
		Result: SupplementaryDocumentTextDataCheckResult{
			checkResult:    docCheck.Result,
			DocumentFields: b.documentFields,
		},
	}, nil
}
