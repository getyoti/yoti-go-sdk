package sandbox

// Breakdown describes a breakdown on check
type Breakdown struct {
	SubCheck string   `json:"sub_check"`
	Result   string   `json:"result"`
	Details  []Detail `json:"details"`
}

// Detail is an individual breakdown detail
type Detail struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// WithSubCheck sets the sub check of a Breakdown
func (breakdown Breakdown) WithSubCheck(subCheck string) Breakdown {
	breakdown.SubCheck = subCheck
	return breakdown
}

// WithResult sets the result of a Breakdown
func (breakdown Breakdown) WithResult(result string) Breakdown {
	breakdown.Result = result
	return breakdown
}

// WithDetail sets the Detail of a Breakdown
func (breakdown Breakdown) WithDetail(name string, value string) Breakdown {
	breakdown.Details = append(breakdown.Details, Detail{
		Name:  name,
		Value: value,
	})
	return breakdown
}
