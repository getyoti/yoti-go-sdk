package sandbox

type documentAuthenticityCheck struct {
	check
}

type documentAuthenticityCheckBuilder struct {
	checkBuilder
}

func NewDocumentAuthenticityCheckBuilder() *documentAuthenticityCheckBuilder {
	return &documentAuthenticityCheckBuilder{}
}

func (b *documentAuthenticityCheckBuilder) WithRecommendation(recommendation recommendation) *documentAuthenticityCheckBuilder {
	b.checkBuilder.WithRecommendation(recommendation)
	return b
}

func (b *documentAuthenticityCheckBuilder) WithBreakdown(breakdown breakdown) *documentAuthenticityCheckBuilder {
	b.checkBuilder.WithBreakdown(breakdown)
	return b
}

func (b *documentAuthenticityCheckBuilder) Build() documentAuthenticityCheck {
	report := checkReport{
		Recommendation: b.Recommendation,
		Breakdown:      b.Breakdowns,
	}
	result := checkResult{
		Report: report,
	}
	return documentAuthenticityCheck{check{
		Result: result,
	}}
}
