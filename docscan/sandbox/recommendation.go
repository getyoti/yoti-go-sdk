package sandbox

// Recommendation describes a recommendation on check
type Recommendation struct {
	Value              string `json:"value"`
	Reason             string `json:"reason"`
	RecoverySuggestion string `json:"recovery_suggestion"`
}

// WithReason sets the reason of a Recommendation
func (recommendation Recommendation) WithReason(reason string) Recommendation {
	recommendation.Reason = reason
	return recommendation
}

// WithValue sets the value of a Recommendation
func (recommendation Recommendation) WithValue(value string) Recommendation {
	recommendation.Value = value
	return recommendation
}

// WithRecoverySuggestion sets the recovery suggestion of a Recommendation
func (recommendation Recommendation) WithRecoverySuggestion(recoverySuggestion string) Recommendation {
	recommendation.RecoverySuggestion = recoverySuggestion
	return recommendation
}
