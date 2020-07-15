package report

import (
	"github.com/getyoti/yoti-go-sdk/v3/validate"
)

// Recommendation describes a recommendation on check
type Recommendation struct {
	Value              string `json:"value"`
	Reason             string `json:"reason,omitempty"`
	RecoverySuggestion string `json:"recovery_suggestion,omitempty"`
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

func (b *recommendationBuilder) Build() (Recommendation, error) {
	recommendation := Recommendation{
		Value:              b.value,
		Reason:             b.reason,
		RecoverySuggestion: b.recoverySuggestion,
	}

	valueErr := validate.NotEmpty(recommendation.Value, "Value cannot be empty")
	if valueErr != nil {
		return recommendation, valueErr
	}

	return recommendation, nil
}
