package sandbox

type documentAuthenticityCheck struct {
	check
	documentCheck
}

type documentAuthenticityCheckBuilder struct {
	checkBuilder
	documentCheckBuilder
	err error
}

func NewDocumentAuthenticityCheckBuilder() *documentAuthenticityCheckBuilder {
	return &documentAuthenticityCheckBuilder{}
}

func (b *documentAuthenticityCheckBuilder) WithRecommendation(recommendation recommendation) *documentAuthenticityCheckBuilder {
	b.checkBuilder.withRecommendation(recommendation)
	return b
}

func (b *documentAuthenticityCheckBuilder) WithBreakdown(breakdown breakdown) *documentAuthenticityCheckBuilder {
	b.checkBuilder.withBreakdown(breakdown)
	return b
}

func (b *documentAuthenticityCheckBuilder) WithDocumentFilter(filter documentFilter) *documentAuthenticityCheckBuilder {
	b.documentCheckBuilder.withDocumentFilter(filter)
	return b
}

func (b *documentAuthenticityCheckBuilder) Build() (documentAuthenticityCheck, error) {
	documentAuthenticityCheck := documentAuthenticityCheck{}

	check, err := b.checkBuilder.build()
	if err != nil {
		return documentAuthenticityCheck, err
	}

	documentCheck, err := b.documentCheckBuilder.build()
	if err != nil {
		return documentAuthenticityCheck, err
	}

	documentAuthenticityCheck.check = check
	documentAuthenticityCheck.documentCheck = documentCheck

	return documentAuthenticityCheck, b.err
}
