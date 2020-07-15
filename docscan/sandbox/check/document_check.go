package check

import (
	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox/filter"
)

type documentCheck struct {
	check
	DocumentFilter filter.DocumentFilter `json:"document_filter"`
}

type documentCheckBuilder struct {
	checkBuilder
	documentFilter filter.DocumentFilter
	err            error
}

func (b *documentCheckBuilder) withDocumentFilter(filter filter.DocumentFilter) {
	b.documentFilter = filter
}

func (b *documentCheckBuilder) build() (documentCheck, error) {
	documentCheck := documentCheck{}

	check, err := b.checkBuilder.build()
	if err != nil {
		return documentCheck, err
	}

	documentCheck.check = check
	documentCheck.DocumentFilter = b.documentFilter

	return documentCheck, b.err
}
