package retrieve

type CASourcesResponse struct {
	Type          string   `json:"type"`
	SearchProfile string   `json:"search_profile"`
	Types         []string `json:"types"`
}
