package retrieve

// RecommendationResponse represents the recommendation given for a check
type RecommendationResponse struct {
	Value              string `json:"value"`
	Reason             string `json:"reason"`
	RecoverySuggestion string `json:"recovery_suggestion"`
}
