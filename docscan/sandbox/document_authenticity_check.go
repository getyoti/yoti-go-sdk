package sandbox

type documentAuthenticityCheck struct {
	check
	documentCheck
}

type documentAuthenticityCheckBuilder struct {
	checkBuilder
	documentCheckBuilder
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

func (b *documentAuthenticityCheckBuilder) Build() documentAuthenticityCheck {
	report := checkReport{
		Recommendation: b.recommendation,
		Breakdown:      b.breakdowns,
	}
	result := checkResult{
		Report: report,
	}
	return documentAuthenticityCheck{
		check{
			Result: result,
		},
		documentCheck{
			DocumentFilter: b.documentFilter,
		},
	}
}
