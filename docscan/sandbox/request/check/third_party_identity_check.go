package check

import "github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox/request/check/report"

// ThirdPartyIdentityCheck defines a sandbox check with a third party credit reporting agency
type ThirdPartyIdentityCheck struct {
	*check
}

// ThirdPartyIdentityCheckBuilder builds a ThirdPartyIdentityCheck
type ThirdPartyIdentityCheckBuilder struct {
	checkBuilder
}

func (c *ThirdPartyIdentityCheck) Report() checkReport {
	return c.check.Result.Report
}

// NewThirdPartyIdentityCheckBuilder creates a new ThirdPartyIdentityCheckBuilder
func NewThirdPartyIdentityCheckBuilder() *ThirdPartyIdentityCheckBuilder {
	return &ThirdPartyIdentityCheckBuilder{}
}

// Build creates a new ThirdPartyIdentityCheck
func (b *ThirdPartyIdentityCheckBuilder) Build() (*ThirdPartyIdentityCheck, error) {
	tpiCheck := ThirdPartyIdentityCheck{
		check: b.checkBuilder.build(),
	}

	return &tpiCheck, nil
}

func (b *ThirdPartyIdentityCheckBuilder) WithBreakdown(breakdown *report.Breakdown) {
	b.checkBuilder.withBreakdown(breakdown)
}

func (b *ThirdPartyIdentityCheckBuilder) WithRecommendation(recommendation *report.Recommendation) {
	b.checkBuilder.withRecommendation(recommendation)
}
