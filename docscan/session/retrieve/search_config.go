package retrieve

// SearchConfig is the base type which search configs must satisfy
type SearchConfig interface {
	Type() string
	Categories() []string
}
