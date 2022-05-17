package check

// RequestedCAMatchingStrategy is the base type which other CA matching strategies must satisfy
type RequestedCAMatchingStrategy interface {
	Type() string
}
