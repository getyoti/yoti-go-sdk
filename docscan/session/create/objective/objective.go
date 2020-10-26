package objective

type Objective interface {
	Type() string
	MarshalJSON() ([]byte, error)
}
