package task

import (
	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox/request/filter"
)

// DocumentTextDataExtractionTask represents a document text data extraction task
type DocumentTextDataExtractionTask struct {
	documentTask
	Result documentTextDataExtractionTaskResult `json:"result"`
}

// DocumentTextDataExtractionTaskBuilder builds a DocumentTextDataExtractionTask
type DocumentTextDataExtractionTaskBuilder struct {
	documentTaskBuilder
	documentFields map[string]string
}

type documentTextDataExtractionTaskResult struct {
	DocumentFields map[string]string `json:"document_fields,omitempty"`
}

// NewDocumentTextDataExtractionTaskBuilder creates a new DocumentTextDataExtractionTaskBuilder
func NewDocumentTextDataExtractionTaskBuilder() *DocumentTextDataExtractionTaskBuilder {
	return &DocumentTextDataExtractionTaskBuilder{}
}

// WithDocumentFilter adds a document filter to the task
func (b *DocumentTextDataExtractionTaskBuilder) WithDocumentFilter(filter filter.DocumentFilter) *DocumentTextDataExtractionTaskBuilder {
	b.documentTaskBuilder.withDocumentFilter(filter)
	return b
}

// WithDocumentField adds a document field to the task
func (b *DocumentTextDataExtractionTaskBuilder) WithDocumentField(key string, value string) *DocumentTextDataExtractionTaskBuilder {
	if b.documentFields == nil {
		b.documentFields = make(map[string]string)
	}
	b.documentFields[key] = value
	return b
}

// Build creates a new DocumentTextDataExtractionTask
func (b *DocumentTextDataExtractionTaskBuilder) Build() (DocumentTextDataExtractionTask, error) {
	documentTextDataExtractionTask := DocumentTextDataExtractionTask{
		documentTask: b.documentTaskBuilder.build(),
	}

	documentTextDataExtractionTask.Result = documentTextDataExtractionTaskResult{
		DocumentFields: b.documentFields,
	}

	return documentTextDataExtractionTask, nil
}
