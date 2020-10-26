package request

import (
	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox/request/task"
)

// TaskResults represents task results
type TaskResults struct {
	DocumentTextDataExtractionTasks              []*task.DocumentTextDataExtractionTask              `json:"ID_DOCUMENT_TEXT_DATA_EXTRACTION"`
	SupplementaryDocumentTextDataExtractionTasks []*task.SupplementaryDocumentTextDataExtractionTask `json:"SUPPLEMENTARY_DOCUMENT_TEXT_DATA_EXTRACTION"`
}

// TaskResultsBuilder builds TaskResults
type TaskResultsBuilder struct {
	documentTextDataExtractionTasks              []*task.DocumentTextDataExtractionTask
	supplementaryDocumentTextDataExtractionTasks []*task.SupplementaryDocumentTextDataExtractionTask
}

// NewTaskResultsBuilder creates a new TaskResultsBuilder
func NewTaskResultsBuilder() *TaskResultsBuilder {
	return &TaskResultsBuilder{
		documentTextDataExtractionTasks:              []*task.DocumentTextDataExtractionTask{},
		supplementaryDocumentTextDataExtractionTasks: []*task.SupplementaryDocumentTextDataExtractionTask{},
	}
}

// WithDocumentTextDataExtractionTask adds a supplementary document text data extraction task
func (b *TaskResultsBuilder) WithDocumentTextDataExtractionTask(documentTextDataExtractionTask *task.DocumentTextDataExtractionTask) *TaskResultsBuilder {
	b.documentTextDataExtractionTasks = append(b.documentTextDataExtractionTasks, documentTextDataExtractionTask)
	return b
}

// WithSupplementaryDocumentTextDataExtractionTask adds a supplementary document text data extraction task
func (b *TaskResultsBuilder) WithSupplementaryDocumentTextDataExtractionTask(supplementaryTextDataExtractionTask *task.SupplementaryDocumentTextDataExtractionTask) *TaskResultsBuilder {
	b.supplementaryDocumentTextDataExtractionTasks = append(b.supplementaryDocumentTextDataExtractionTasks, supplementaryTextDataExtractionTask)
	return b
}

// Build creates TaskResults
func (b *TaskResultsBuilder) Build() (TaskResults, error) {
	return TaskResults{
		DocumentTextDataExtractionTasks:              b.documentTextDataExtractionTasks,
		SupplementaryDocumentTextDataExtractionTasks: b.supplementaryDocumentTextDataExtractionTasks,
	}, nil
}
