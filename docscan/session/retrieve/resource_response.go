package retrieve

// ResourceResponse represents a resource, with associated tasks
type ResourceResponse struct {
	ID    string          `json:"id"`
	Tasks []*TaskResponse `json:"tasks"`
}
