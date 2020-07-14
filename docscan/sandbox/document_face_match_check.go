package sandbox

type documentFaceMatchCheck struct {
	check
	documentCheck
}

type documentFaceMatchCheckBuilder struct {
	checkBuilder
	documentCheckBuilder
	err error
}

func NewDocumentFaceMatchCheckBuilder() *documentFaceMatchCheckBuilder {
	return &documentFaceMatchCheckBuilder{}
}

func (b *documentFaceMatchCheckBuilder) WithRecommendation(recommendation recommendation) *documentFaceMatchCheckBuilder {
	b.checkBuilder.withRecommendation(recommendation)
	return b
}

func (b *documentFaceMatchCheckBuilder) WithBreakdown(breakdown breakdown) *documentFaceMatchCheckBuilder {
	b.checkBuilder.withBreakdown(breakdown)
	return b
}

func (b *documentFaceMatchCheckBuilder) WithDocumentFilter(filter documentFilter) *documentFaceMatchCheckBuilder {
	b.documentCheckBuilder.withDocumentFilter(filter)
	return b
}

func (b *documentFaceMatchCheckBuilder) Build() (documentFaceMatchCheck, error) {
	documentAuthenticityCheck := documentFaceMatchCheck{}

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
