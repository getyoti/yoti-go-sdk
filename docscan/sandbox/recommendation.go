package sandbox

// Recommendation describes a recommendation on check
type recommendation struct {
	Value              string `json:"value"`
	Reason             string `json:"reason"`
	RecoverySuggestion string `json:"recovery_suggestion"`
}

type recommendationBuilder struct {
	value              string
	reason             string
	recoverySuggestion string
}

func NewRecommendationBuilder() *recommendationBuilder {
	return &recommendationBuilder{}
}

// WithReason sets the reason of a Recommendation
func (b *recommendationBuilder) WithReason(reason string) *recommendationBuilder {
	b.reason = reason
	return b
}

// WithValue sets the value of a Recommendation
func (b *recommendationBuilder) WithValue(value string) *recommendationBuilder {
	b.value = value
	return b
}

// WithRecoverySuggestion sets the recovery suggestion of a Recommendation
func (b *recommendationBuilder) WithRecoverySuggestion(recoverySuggestion string) *recommendationBuilder {
	b.recoverySuggestion = recoverySuggestion
	return b
}

func (b *recommendationBuilder) Build() (recommendation, error) {
	return recommendation{
		Value:              b.value,
		Reason:             b.reason,
		RecoverySuggestion: b.recoverySuggestion,
	}, nil
}
