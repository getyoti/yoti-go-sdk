package sandbox

type documentFaceMatchCheck struct {
	check
	documentCheck
}

type documentFaceMatchCheckBuilder struct {
	checkBuilder
	documentCheckBuilder
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
	report := checkReport{
		Recommendation: b.recommendation,
		Breakdown:      b.breakdowns,
	}
	result := checkResult{
		Report: report,
	}
	return documentFaceMatchCheck{
		check{
			Result: result,
		},
		documentCheck{
			DocumentFilter: b.documentFilter,
		},
	}, nil
}
