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

func (b *ThirdPartyIdentityCheckBuilder) WithBreakdown(breakdown *report.Breakdown) *ThirdPartyIdentityCheckBuilder {
	b.checkBuilder.withBreakdown(breakdown)
	return b
}

func (b *ThirdPartyIdentityCheckBuilder) WithRecommendation(recommendation *report.Recommendation) *ThirdPartyIdentityCheckBuilder {
	b.checkBuilder.withRecommendation(recommendation)
	return b
}
