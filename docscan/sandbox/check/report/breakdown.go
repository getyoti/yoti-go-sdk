package report

import (
	"github.com/getyoti/yoti-go-sdk/v3/validate"
)

// Breakdown describes a breakdown on check
type Breakdown struct {
	SubCheck string   `json:"sub_check"`
	Result   string   `json:"result"`
	Details  []detail `json:"details"`
}

type breakdownBuilder struct {
	subCheck string
	result   string
	details  []detail
	err      error
}

// Detail is an individual breakdown detail
type detail struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func NewBreakdownBuilder() *breakdownBuilder {
	return &breakdownBuilder{
		details: []detail{},
	}
}

// WithSubCheck sets the sub check of a Breakdown
func (b *breakdownBuilder) WithSubCheck(subCheck string) *breakdownBuilder {
	b.subCheck = subCheck
	return b
}

// WithResult sets the result of a Breakdown
func (b *breakdownBuilder) WithResult(result string) *breakdownBuilder {
	b.result = result
	return b
}

// WithDetail sets the Detail of a Breakdown
func (b *breakdownBuilder) WithDetail(name string, value string) *breakdownBuilder {
	b.details = append(b.details, detail{
		Name:  name,
		Value: value,
	})
	return b
}

func (b *breakdownBuilder) Build() (Breakdown, error) {
	breakdown := Breakdown{
		SubCheck: b.subCheck,
		Result:   b.result,
		Details:  b.details,
	}

	subCheckErr := validate.NotEmpty(breakdown.SubCheck, "Sub Check cannot be empty")
	if subCheckErr != nil {
		return breakdown, subCheckErr
	}

	resultErr := validate.NotEmpty(breakdown.Result, "Result cannot be empty")
	if resultErr != nil {
		return breakdown, resultErr
	}

	return breakdown, b.err
}
