package task

import (
	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox/filter"
)

type documentTask struct {
	DocumentFilter filter.DocumentFilter `json:"document_filter"`
}

type documentTaskBuilder struct {
	documentFilter filter.DocumentFilter
	err            error
}

func (b *documentTaskBuilder) withDocumentFilter(filter filter.DocumentFilter) {
	b.documentFilter = filter
}

func (b *documentTaskBuilder) build() (documentTask, error) {
	documentTask := documentTask{}

	documentTask.DocumentFilter = b.documentFilter

	return documentTask, b.err
}
