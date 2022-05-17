package retrieve

type RawResults struct {
	Media MediaResponse `json:"media"`
}

type WatchlistSummaryResponse struct {
	TotalHits              int          `json:"total_hits"`
	RawResults             RawResults   `json:"raw_results"`
	AssociatedCountryCodes []string     `json:"associated_country_codes"`
	SearchConfig           SearchConfig `json:"search_config"`
}
