package sandbox

// Breakdown describes a breakdown on check
type breakdown struct {
	SubCheck string   `json:"sub_check"`
	Result   string   `json:"result"`
	Details  []detail `json:"details"`
}

type breakdownBuilder struct {
	subCheck string
	result   string
	details  []detail
}

// Detail is an individual breakdown detail
type detail struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func NewBreakdownBuilder() *breakdownBuilder {
	return &breakdownBuilder{}
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

func (b *breakdownBuilder) Build() (breakdown, error) {
	return breakdown{
		SubCheck: b.subCheck,
		Result:   b.result,
		Details:  b.details,
	}, nil
}
