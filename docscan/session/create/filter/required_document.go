package filter

import "github.com/getyoti/yoti-go-sdk/v3/docscan/constants"

// RequiredDocument defines a document to be provided by the user
type RequiredDocument struct {
	Type string `json:"type"`
}

// RequiredIDDocument defines an identity document to be provided by the user
type RequiredIDDocument struct {
	*RequiredDocument
	Filter *RequestedDocumentFilter `json:"filter"`
}

// NewRequiredIDDocument creates a new RequiredIDDocument struct
func NewRequiredIDDocument(filter *RequestedDocumentFilter) *RequiredIDDocument {
	return &RequiredIDDocument{
		RequiredDocument: &RequiredDocument{Type: constants.IDDocument},
		Filter:           filter,
	}
}

// RequiredIDDocumentBuilder builds a RequiredIDDocument
type RequiredIDDocumentBuilder struct {
	Filter *RequestedDocumentFilter
}

// WithFilter applies a document filter to an ID Document
func (b *RequiredIDDocumentBuilder) WithFilter(filter *RequestedDocumentFilter) *RequiredIDDocumentBuilder {
	b.Filter = filter
	return b
}

// Build builds the RequiredIDDocument
func (b *RequiredIDDocumentBuilder) Build() (*RequiredIDDocument, error) {
	return NewRequiredIDDocument(b.Filter), nil
}
