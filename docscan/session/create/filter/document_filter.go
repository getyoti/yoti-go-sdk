package filter

// RequestedDocumentFilter filters for a required document, allowing specification of restrictive parameters
type RequestedDocumentFilter interface {
	Type() string
}
