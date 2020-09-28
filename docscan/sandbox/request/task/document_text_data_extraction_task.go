package task

import (
	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox/request/filter"
)

// DocumentTextDataExtractionTask represents a document text data extraction task
type DocumentTextDataExtractionTask struct {
	*documentTask
	Result documentTextDataExtractionTaskResult `json:"result"`
}

// DocumentTextDataExtractionTaskBuilder builds a DocumentTextDataExtractionTask
type DocumentTextDataExtractionTaskBuilder struct {
	documentTaskBuilder
	documentFields map[string]interface{}
}

type documentTextDataExtractionTaskResult struct {
	DocumentFields map[string]interface{} `json:"document_fields,omitempty"`
}

// NewDocumentTextDataExtractionTaskBuilder creates a new DocumentTextDataExtractionTaskBuilder
func NewDocumentTextDataExtractionTaskBuilder() *DocumentTextDataExtractionTaskBuilder {
	return &DocumentTextDataExtractionTaskBuilder{}
}

// WithDocumentFilter adds a document filter to the task
func (b *DocumentTextDataExtractionTaskBuilder) WithDocumentFilter(filter *filter.DocumentFilter) *DocumentTextDataExtractionTaskBuilder {
	b.documentTaskBuilder.withDocumentFilter(filter)
	return b
}

// WithDocumentField adds a document field to the task
func (b *DocumentTextDataExtractionTaskBuilder) WithDocumentField(key string, value interface{}) *DocumentTextDataExtractionTaskBuilder {
	if b.documentFields == nil {
		b.documentFields = make(map[string]interface{})
	}
	b.documentFields[key] = value
	return b
}

// WithDocumentFields sets document fields
func (b *DocumentTextDataExtractionTaskBuilder) WithDocumentFields(documentFields map[string]interface{}) *DocumentTextDataExtractionTaskBuilder {
	b.documentFields = documentFields
	return b
}

// Build creates a new DocumentTextDataExtractionTask
func (b *DocumentTextDataExtractionTaskBuilder) Build() (*DocumentTextDataExtractionTask, error) {
	return &DocumentTextDataExtractionTask{
		documentTask: b.documentTaskBuilder.build(),
		Result: documentTextDataExtractionTaskResult{
			DocumentFields: b.documentFields,
		},
	}, nil
}
