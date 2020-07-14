package sandbox

type documentAuthenticityCheck struct {
	documentCheck
}

type documentAuthenticityCheckBuilder struct {
	documentCheckBuilder
	err error
}

func NewDocumentAuthenticityCheckBuilder() *documentAuthenticityCheckBuilder {
	return &documentAuthenticityCheckBuilder{}
}

func (b *documentAuthenticityCheckBuilder) WithRecommendation(recommendation recommendation) *documentAuthenticityCheckBuilder {
	b.documentCheckBuilder.withRecommendation(recommendation)
	return b
}

func (b *documentAuthenticityCheckBuilder) WithBreakdown(breakdown breakdown) *documentAuthenticityCheckBuilder {
	b.documentCheckBuilder.withBreakdown(breakdown)
	return b
}

func (b *documentAuthenticityCheckBuilder) WithDocumentFilter(filter documentFilter) *documentAuthenticityCheckBuilder {
	b.documentCheckBuilder.withDocumentFilter(filter)
	return b
}

func (b *documentAuthenticityCheckBuilder) Build() (documentAuthenticityCheck, error) {
	documentAuthenticityCheck := documentAuthenticityCheck{}

	documentCheck, err := b.documentCheckBuilder.build()
	if err != nil {
		return documentAuthenticityCheck, err
	}

	documentAuthenticityCheck.documentCheck = documentCheck

	return documentAuthenticityCheck, b.err
}
