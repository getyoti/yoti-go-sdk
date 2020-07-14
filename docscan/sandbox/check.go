package sandbox

type check struct {
	Result checkResult `json:"result"`
}

type checkBuilder struct {
	recommendation recommendation
	breakdowns     []breakdown
}

type checkResult struct {
	Report checkReport `json:"report"`
}

type checkReport struct {
	Recommendation recommendation `json:"recommendation"`
	Breakdown      []breakdown    `json:"breakdown"`
}

func (b *checkBuilder) withRecommendation(recommendation recommendation) {
	b.recommendation = recommendation
}

func (b *checkBuilder) withBreakdown(breakdown breakdown) {
	b.breakdowns = append(b.breakdowns, breakdown)
}
