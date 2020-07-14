package sandbox

type documentCheck struct {
	check
	DocumentFilter documentFilter `json:"document_filter"`
}

type documentCheckBuilder struct {
	checkBuilder
	documentFilter documentFilter
	err            error
}

func (b *documentCheckBuilder) withDocumentFilter(filter documentFilter) {
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
