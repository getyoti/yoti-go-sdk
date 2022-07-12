package retrieve

type SearchConfig struct {
	Type             string                     `json:"type"`
	Categories       []string                   `json:"categories"`
	RemoveDeceased   bool                       `json:"remove_deceased"`
	ShareURL         bool                       `json:"share_url"`
	Sources          CASourcesResponse          `json:"sources"`
	MatchingStrategy CAMatchingStrategyResponse `json:"matching_strategy"`
	APIKey           string                     `json:"api_key"`
	Monitoring       bool                       `json:"monitoring"`
	Tags             map[string]string          `json:"tags"`
	ClientRef        string                     `json:"client_ref"`
}

type CAMatchingStrategyResponse struct {
	Type       string  `json:"type"`
	ExactMatch string  `json:"exact_match"`
	Fuzziness  float64 `json:"fuzziness"`
}
