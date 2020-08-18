package retrieve

// ReportResponse represents a report for a given check
type ReportResponse struct {
	Recommendation *RecommendationResponse `json:"recommendation"`
	Breakdown      []*BreakdownResponse    `json:"breakdown"`
}
