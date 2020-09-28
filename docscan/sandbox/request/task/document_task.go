package task

import (
	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox/request/filter"
)

type documentTask struct {
	DocumentFilter *filter.DocumentFilter `json:"document_filter,omitempty"`
}

type documentTaskBuilder struct {
	documentFilter *filter.DocumentFilter
}

func (b *documentTaskBuilder) withDocumentFilter(filter *filter.DocumentFilter) {
	b.documentFilter = filter
}

func (b *documentTaskBuilder) build() *documentTask {
	return &documentTask{
		DocumentFilter: b.documentFilter,
	}
}
