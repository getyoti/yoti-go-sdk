package sandbox

type check struct {
	Result checkResult `json:"result"`
}

type checkBuilder struct {
	Recommendation recommendation
	Breakdowns     []breakdown
}

type checkResult struct {
	Report checkReport `json:"report"`
}

type checkReport struct {
	Recommendation recommendation `json:"recommendation"`
	Breakdown      []breakdown    `json:"breakdown"`
}

func (b *checkBuilder) WithRecommendation(recommendation recommendation) *checkBuilder {
	b.Recommendation = recommendation
	return b
}

func (b *checkBuilder) WithBreakdown(breakdown breakdown) *checkBuilder {
	b.Breakdowns = append(b.Breakdowns, breakdown)
	return b
}
