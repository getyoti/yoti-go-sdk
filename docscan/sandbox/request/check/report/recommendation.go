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

// RecommendationBuilder builds a Recommendation
type RecommendationBuilder struct {
	value              string
	reason             string
	recoverySuggestion string
}

// NewRecommendationBuilder creates a new RecommendationBuilder
func NewRecommendationBuilder() *RecommendationBuilder {
	return &RecommendationBuilder{}
}

// WithReason sets the reason of a Recommendation
func (b *RecommendationBuilder) WithReason(reason string) *RecommendationBuilder {
	b.reason = reason
	return b
}

// WithValue sets the value of a Recommendation
func (b *RecommendationBuilder) WithValue(value string) *RecommendationBuilder {
	b.value = value
	return b
}

// WithRecoverySuggestion sets the recovery suggestion of a Recommendation
func (b *RecommendationBuilder) WithRecoverySuggestion(recoverySuggestion string) *RecommendationBuilder {
	b.recoverySuggestion = recoverySuggestion
	return b
}

// Build creates a new Recommendation
func (b *RecommendationBuilder) Build() (Recommendation, error) {
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
