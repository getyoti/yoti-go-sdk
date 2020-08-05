package retrieve

// BreakdownResponse represents one breakdown item for a given check
type BreakdownResponse struct {
	SubCheck string             `json:"sub_check"`
	Result   string             `json:"result"`
	Details  []*DetailsResponse `json:"details"`
}
