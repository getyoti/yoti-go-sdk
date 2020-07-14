package sandbox

type check struct {
	Result checkResult `json:"result"`
}

type checkBuilder struct {
	recommendation recommendation
	breakdowns     []breakdown
	err            error
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

func (b *checkBuilder) build() (check, error) {
	report := checkReport{
		Recommendation: b.recommendation,
		Breakdown:      b.breakdowns,
	}
	result := checkResult{
		Report: report,
	}
	return check{
		Result: result,
	}, b.err
}
