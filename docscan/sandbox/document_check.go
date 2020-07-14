package sandbox

type documentCheck struct {
	DocumentFilter documentFilter `json:"document_filter"`
}

type documentCheckBuilder struct {
	documentFilter documentFilter
}

func (b *documentCheckBuilder) withDocumentFilter(filter documentFilter) {
	b.documentFilter = filter
}
