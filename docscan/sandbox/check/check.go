package check

import (
	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox/check/report"
)

type check struct {
	Result checkResult `json:"result"`
}

type checkBuilder struct {
	recommendation report.Recommendation
	breakdowns     []report.Breakdown
	err            error
}

type checkResult struct {
	Report checkReport `json:"report"`
}

type checkReport struct {
	Recommendation report.Recommendation `json:"recommendation,omitempty"`
	Breakdown      []report.Breakdown    `json:"breakdown,omitempty"`
}

func (b *checkBuilder) withRecommendation(recommendation report.Recommendation) {
	b.recommendation = recommendation
}

func (b *checkBuilder) withBreakdown(breakdown report.Breakdown) {
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
