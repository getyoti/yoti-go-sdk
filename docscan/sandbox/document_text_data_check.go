package sandbox

type documentTextDataCheck struct {
	Result documentTextDataCheckResult `json:"result"`
	documentCheck
}

type documentTextDataCheckBuilder struct {
	documentCheckBuilder
	documentFields map[string]string
	err            error
}

type documentTextDataCheckResult struct {
	checkResult
	DocumentFields map[string]string `json:"document_fields"`
}

func NewDocumentTextDataCheckBuilder() *documentTextDataCheckBuilder {
	return &documentTextDataCheckBuilder{}
}

func (b *documentTextDataCheckBuilder) WithRecommendation(recommendation recommendation) *documentTextDataCheckBuilder {
	b.documentCheckBuilder.withRecommendation(recommendation)
	return b
}

func (b *documentTextDataCheckBuilder) WithBreakdown(breakdown breakdown) *documentTextDataCheckBuilder {
	b.documentCheckBuilder.withBreakdown(breakdown)
	return b
}

func (b *documentTextDataCheckBuilder) WithDocumentFilter(filter documentFilter) *documentTextDataCheckBuilder {
	b.documentCheckBuilder.withDocumentFilter(filter)
	return b
}

func (b *documentTextDataCheckBuilder) WithDocumentField(key string, value string) *documentTextDataCheckBuilder {
	if b.documentFields == nil {
		b.documentFields = make(map[string]string)
	}
	b.documentFields[key] = value
	return b
}

func (b *documentTextDataCheckBuilder) Build() (documentTextDataCheck, error) {
	documentTextDataCheck := documentTextDataCheck{}

	documentCheck, err := b.documentCheckBuilder.build()
	if err != nil {
		return documentTextDataCheck, err
	}

	documentTextDataCheck.documentCheck = documentCheck
	documentTextDataCheck.Result = documentTextDataCheckResult{
		checkResult:    documentCheck.Result,
		DocumentFields: b.documentFields,
	}

	return documentTextDataCheck, b.err
}
