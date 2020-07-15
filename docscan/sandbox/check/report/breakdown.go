package report

import (
	"github.com/getyoti/yoti-go-sdk/v3/validate"
	"github.com/getyoti/yoti-go-sdk/v3/yotierror"
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
	err := validate.NotEmpty(subCheck, "sub check cannot be empty")
	if err != nil {
		b.err = yotierror.MultiError{This: err, Next: b.err}
	}
	b.subCheck = subCheck
	return b
}

// WithResult sets the result of a Breakdown
func (b *breakdownBuilder) WithResult(result string) *breakdownBuilder {
	err := validate.NotEmpty(result, "result cannot be empty")
	if err != nil {
		b.err = yotierror.MultiError{This: err, Next: b.err}
	}
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
	return Breakdown{
		SubCheck: b.subCheck,
		Result:   b.result,
		Details:  b.details,
	}, b.err
}
