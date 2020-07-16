package check

import (
	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox/request/filter"
)

type documentCheck struct {
	check
	DocumentFilter *filter.DocumentFilter `json:"document_filter,omitempty"`
}

type documentCheckBuilder struct {
	checkBuilder
	documentFilter *filter.DocumentFilter
}

func (b *documentCheckBuilder) withDocumentFilter(filter filter.DocumentFilter) {
	b.documentFilter = &filter
}

func (b *documentCheckBuilder) build() documentCheck {
	return documentCheck{
		check:          b.checkBuilder.build(),
		DocumentFilter: b.documentFilter,
	}
}
