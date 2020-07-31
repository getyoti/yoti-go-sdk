package report


import (
	"errors"
)

// Breakdown describes a breakdown on check
type Breakdown struct {
	SubCheck string    `json:"sub_check"`
	Result   string    `json:"result"`
	Details  []*detail `json:"details"`
}

// BreakdownBuilder builds a Breakdown
type BreakdownBuilder struct {
	subCheck string
	result   string
	details  []*detail
}

// Detail is an individual breakdown detail
type detail struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// NewBreakdownBuilder creates a new BreakdownBuilder
func NewBreakdownBuilder() *BreakdownBuilder {
	return &BreakdownBuilder{
		details: []*detail{},
	}
}

// WithSubCheck sets the sub check of a Breakdown
func (b *BreakdownBuilder) WithSubCheck(subCheck string) *BreakdownBuilder {
	b.subCheck = subCheck
	return b
}

// WithResult sets the result of a Breakdown
func (b *BreakdownBuilder) WithResult(result string) *BreakdownBuilder {
	b.result = result
	return b
}

// WithDetail sets the Detail of a Breakdown
func (b *BreakdownBuilder) WithDetail(name string, value string) *BreakdownBuilder {
	b.details = append(b.details, &detail{
		Name:  name,
		Value: value,
	})
	return b
}

// Build creates a new Breakdown
func (b *BreakdownBuilder) Build() (*Breakdown, error) {
	breakdown := &Breakdown{
		SubCheck: b.subCheck,
		Result:   b.result,
		Details:  b.details,
	}

	if len(breakdown.SubCheck) == 0 {
		return nil, errors.New("Sub Check cannot be empty")
	}

	if len(breakdown.Result) == 0 {
		return nil, errors.New("Result cannot be empty")
	}

	return breakdown, nil
}
