package check

// RequestedCASources is the base type which other CA sources must satisfy
type RequestedCASources interface {
	Type() string
}
