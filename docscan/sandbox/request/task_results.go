package request

import (
	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox/request/task"
)

type TaskResults struct {
	DocumentTextDataExtractionTasks []task.DocumentTextDataExtractionTask `json:"ID_DOCUMENT_TEXT_DATA_EXTRACTION"`
}

type taskResultsBuilder struct {
	documentTextDataExtractionTasks []task.DocumentTextDataExtractionTask
}

func NewTaskResultsBuilder() *taskResultsBuilder {
	return &taskResultsBuilder{
		documentTextDataExtractionTasks: []task.DocumentTextDataExtractionTask{},
	}
}

func (b *taskResultsBuilder) WithDocumentTextDataExtractionTask(documentTextDataExtractionTasks task.DocumentTextDataExtractionTask) *taskResultsBuilder {
	b.documentTextDataExtractionTasks = append(b.documentTextDataExtractionTasks, documentTextDataExtractionTasks)
	return b
}

func (b *taskResultsBuilder) Build() (TaskResults, error) {
	return TaskResults{
		DocumentTextDataExtractionTasks: b.documentTextDataExtractionTasks,
	}, nil
}