package retrieve

// ResourceResponse represents a resource, with associated tasks
type ResourceResponse struct {
	ID    string          `json:"id"`
	Tasks []*TaskResponse `json:"tasks"`
}

// GetID returns the ID of the resource.
func (r *ResourceResponse) GetID() string {
	if r == nil {
		return ""
	}
	return r.ID
}
