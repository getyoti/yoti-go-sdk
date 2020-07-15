package task

import (
	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox/request/filter"
)

type DocumentTextDataExtractionTask struct {
	documentTask
	Result documentTextDataExtractionTaskResult `json:"result"`
}

type documentTextDataExtractionTaskBuilder struct {
	documentTaskBuilder
	documentFields map[string]string
}

type documentTextDataExtractionTaskResult struct {
	DocumentFields map[string]string `json:"document_fields,omitempty"`
}

func NewDocumentTextDataExtractionTaskBuilder() *documentTextDataExtractionTaskBuilder {
	return &documentTextDataExtractionTaskBuilder{}
}

func (b *documentTextDataExtractionTaskBuilder) WithDocumentFilter(filter filter.DocumentFilter) *documentTextDataExtractionTaskBuilder {
	b.documentTaskBuilder.withDocumentFilter(filter)
	return b
}

func (b *documentTextDataExtractionTaskBuilder) WithDocumentField(key string, value string) *documentTextDataExtractionTaskBuilder {
	if b.documentFields == nil {
		b.documentFields = make(map[string]string)
	}
	b.documentFields[key] = value
	return b
}

func (b *documentTextDataExtractionTaskBuilder) Build() (DocumentTextDataExtractionTask, error) {
	documentTextDataExtractionTask := DocumentTextDataExtractionTask{}

	documentTask, err := b.documentTaskBuilder.build()
	if err != nil {
		return documentTextDataExtractionTask, err
	}

	documentTextDataExtractionTask.documentTask = documentTask
	documentTextDataExtractionTask.Result = documentTextDataExtractionTaskResult{
		DocumentFields: b.documentFields,
	}

	return documentTextDataExtractionTask, nil
}
