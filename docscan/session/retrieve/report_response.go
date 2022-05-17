package retrieve

// ReportResponse represents a report for a given check
type ReportResponse struct {
	Recommendation RecommendationResponse `json:"recommendation"`
	Breakdown      []BreakdownResponse    `json:"breakdown"`
}

// ReportResponseWithSummary represents a report for a given check
type ReportResponseWithSummary struct {
	WatchlistSummary WatchlistSummary `json:"watchlist_summary"`
}
