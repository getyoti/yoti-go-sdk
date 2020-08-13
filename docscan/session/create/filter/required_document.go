package filter

// RequiredDocument is a document to be required for the session
type RequiredDocument interface {
	Type() string
	MarshalJSON() ([]byte, error)
}
