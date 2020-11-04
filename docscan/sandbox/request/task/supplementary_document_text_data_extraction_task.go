package task

import (
	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox/request/filter"
)

// SupplementaryDocumentTextDataExtractionTask represents a document text data extraction task
type SupplementaryDocumentTextDataExtractionTask struct {
	*documentTask
	Result supplementaryDocumentTextDataExtractionTaskResult `json:"result"`
}

// SupplementaryDocumentTextDataExtractionTaskBuilder builds a SupplementaryDocumentTextDataExtractionTask
type SupplementaryDocumentTextDataExtractionTaskBuilder struct {
	documentTaskBuilder
	documentFields  map[string]interface{}
	detectedCountry string
	recommendation  *TextDataExtractionRecommendation
}

type supplementaryDocumentTextDataExtractionTaskResult struct {
	DocumentFields  map[string]interface{}            `json:"document_fields,omitempty"`
	DetectedCountry string                            `json:"detected_country,omitempty"`
	Recommendation  *TextDataExtractionRecommendation `json:"recommendation,omitempty"`
}

// NewSupplementaryDocumentTextDataExtractionTaskBuilder creates a new SupplementaryDocumentTextDataExtractionTaskBuilder
func NewSupplementaryDocumentTextDataExtractionTaskBuilder() *SupplementaryDocumentTextDataExtractionTaskBuilder {
	return &SupplementaryDocumentTextDataExtractionTaskBuilder{}
}

// WithDocumentFilter adds a document filter to the task
func (b *SupplementaryDocumentTextDataExtractionTaskBuilder) WithDocumentFilter(filter *filter.DocumentFilter) *SupplementaryDocumentTextDataExtractionTaskBuilder {
	b.documentTaskBuilder.withDocumentFilter(filter)
	return b
}

// WithDocumentField adds a document field to the task
func (b *SupplementaryDocumentTextDataExtractionTaskBuilder) WithDocumentField(key string, value interface{}) *SupplementaryDocumentTextDataExtractionTaskBuilder {
	if b.documentFields == nil {
		b.documentFields = make(map[string]interface{})
	}
	b.documentFields[key] = value
	return b
}

// WithDocumentFields sets document fields
func (b *SupplementaryDocumentTextDataExtractionTaskBuilder) WithDocumentFields(documentFields map[string]interface{}) *SupplementaryDocumentTextDataExtractionTaskBuilder {
	b.documentFields = documentFields
	return b
}

// WithDetectedCountry sets the detected country
func (b *SupplementaryDocumentTextDataExtractionTaskBuilder) WithDetectedCountry(detectedCountry string) *SupplementaryDocumentTextDataExtractionTaskBuilder {
	b.detectedCountry = detectedCountry
	return b
}

// WithRecommendation sets the recommendation
func (b *SupplementaryDocumentTextDataExtractionTaskBuilder) WithRecommendation(recommendation *TextDataExtractionRecommendation) *SupplementaryDocumentTextDataExtractionTaskBuilder {
	b.recommendation = recommendation
	return b
}

// Build creates a new SupplementaryDocumentTextDataExtractionTask
func (b *SupplementaryDocumentTextDataExtractionTaskBuilder) Build() (*SupplementaryDocumentTextDataExtractionTask, error) {
	return &SupplementaryDocumentTextDataExtractionTask{
		documentTask: b.documentTaskBuilder.build(),
		Result: supplementaryDocumentTextDataExtractionTaskResult{
			DocumentFields:  b.documentFields,
			DetectedCountry: b.detectedCountry,
			Recommendation:  b.recommendation,
		},
	}, nil
}
