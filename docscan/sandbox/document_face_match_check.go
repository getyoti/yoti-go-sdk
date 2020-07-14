package sandbox

type documentFaceMatchCheck struct {
	documentCheck
}

type documentFaceMatchCheckBuilder struct {
	documentCheckBuilder
	err error
}

func NewDocumentFaceMatchCheckBuilder() *documentFaceMatchCheckBuilder {
	return &documentFaceMatchCheckBuilder{}
}

func (b *documentFaceMatchCheckBuilder) WithRecommendation(recommendation recommendation) *documentFaceMatchCheckBuilder {
	b.documentCheckBuilder.withRecommendation(recommendation)
	return b
}

func (b *documentFaceMatchCheckBuilder) WithBreakdown(breakdown breakdown) *documentFaceMatchCheckBuilder {
	b.documentCheckBuilder.withBreakdown(breakdown)
	return b
}

func (b *documentFaceMatchCheckBuilder) WithDocumentFilter(filter documentFilter) *documentFaceMatchCheckBuilder {
	b.documentCheckBuilder.withDocumentFilter(filter)
	return b
}

func (b *documentFaceMatchCheckBuilder) Build() (documentFaceMatchCheck, error) {
	documentFaceMatchCheck := documentFaceMatchCheck{}

	documentCheck, err := b.documentCheckBuilder.build()
	if err != nil {
		return documentFaceMatchCheck, err
	}

	documentFaceMatchCheck.documentCheck = documentCheck

	return documentFaceMatchCheck, b.err
}
