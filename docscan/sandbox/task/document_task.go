package task

import (
	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox"
)

type documentTask struct {
	DocumentFilter sandbox.DocumentFilter `json:"document_filter"`
}

type documentTaskBuilder struct {
	documentFilter sandbox.DocumentFilter
	err            error
}

func (b *documentTaskBuilder) withDocumentFilter(filter sandbox.DocumentFilter) {
	b.documentFilter = filter
}

func (b *documentTaskBuilder) build() (documentTask, error) {
	documentTask := documentTask{}

	documentTask.DocumentFilter = b.documentFilter

	return documentTask, b.err
}
