package request

import (
	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox/request/task"
)

// TaskResults represents task results
type TaskResults struct {
	DocumentTextDataExtractionTasks []task.DocumentTextDataExtractionTask `json:"ID_DOCUMENT_TEXT_DATA_EXTRACTION"`
}

// TaskResultsBuilder builds TaskResults
type TaskResultsBuilder struct {
	documentTextDataExtractionTasks []task.DocumentTextDataExtractionTask
}

// NewTaskResultsBuilder creates a new TaskResultsBuilder
func NewTaskResultsBuilder() *TaskResultsBuilder {
	return &TaskResultsBuilder{
		documentTextDataExtractionTasks: []task.DocumentTextDataExtractionTask{},
	}
}

// WithDocumentTextDataExtractionTask adds a document text data extraction task
func (b *TaskResultsBuilder) WithDocumentTextDataExtractionTask(documentTextDataExtractionTasks task.DocumentTextDataExtractionTask) *TaskResultsBuilder {
	b.documentTextDataExtractionTasks = append(b.documentTextDataExtractionTasks, documentTextDataExtractionTasks)
	return b
}

// Build creates TaskResults
func (b *TaskResultsBuilder) Build() (TaskResults, error) {
	return TaskResults{
		DocumentTextDataExtractionTasks: b.documentTextDataExtractionTasks,
	}, nil
}
