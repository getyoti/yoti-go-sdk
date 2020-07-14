package sandbox

type documentCheck struct {
	DocumentFilter documentFilter `json:"document_filter"`
}

type documentCheckBuilder struct {
	documentFilter documentFilter
	err            error
}

func (b *documentCheckBuilder) withDocumentFilter(filter documentFilter) {
	b.documentFilter = filter
}

func (b *documentCheckBuilder) build() (documentCheck, error) {
	return documentCheck{
		DocumentFilter: b.documentFilter,
	}, b.err
}
