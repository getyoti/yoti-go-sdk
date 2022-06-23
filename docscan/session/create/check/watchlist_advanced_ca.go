package check

type RequestedWatchlistAdvancedCAConfig struct {
	Type             string                      `json:"type,omitempty"`
	RemoveDeceased   bool                        `json:"remove_deceased,omitempty"`
	ShareUrl         bool                        `json:"share_url,omitempty"`
	Sources          RequestedCASources          `json:"sources,omitempty"`
	MatchingStrategy RequestedCAMatchingStrategy `json:"matching_strategy,omitempty"`
}
