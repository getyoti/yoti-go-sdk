package sandbox

type task struct {
	Result taskResult `json:"result"`
}

type taskBuilder struct {
	recommendation recommendation
	breakdowns     []breakdown
	err            error
}

type taskResult struct {
	Report taskReport `json:"report"`
}

type taskReport struct {
	Recommendation recommendation `json:"recommendation"`
	Breakdown      []breakdown    `json:"breakdown"`
}

func (b *taskBuilder) withRecommendation(recommendation recommendation) {
	b.recommendation = recommendation
}

func (b *taskBuilder) withBreakdown(breakdown breakdown) {
	b.breakdowns = append(b.breakdowns, breakdown)
}

func (b *taskBuilder) build() (task, error) {
	report := taskReport{
		Recommendation: b.recommendation,
		Breakdown:      b.breakdowns,
	}
	result := taskResult{
		Report: report,
	}
	return task{
		Result: result,
	}, b.err
}
