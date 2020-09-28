package filter

import (
	"encoding/json"
)

// RequiredIDDocument details a required identity document
type RequiredIDDocument struct {
	Filter RequestedDocumentFilter
}

// Type returns the type of the identity document
func (i *RequiredIDDocument) Type() string {
	return identityDocument
}

// MarshalJSON returns the JSON encoding
func (i *RequiredIDDocument) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type   string                  `json:"type"`
		Filter RequestedDocumentFilter `json:"filter,omitempty"`
	}{
		Type:   i.Type(),
		Filter: i.Filter,
	})
}

// NewRequiredIDDocumentBuilder creates a new RequiredIDDocumentBuilder
func NewRequiredIDDocumentBuilder() *RequiredIDDocumentBuilder {
	return &RequiredIDDocumentBuilder{}
}

// RequiredIDDocumentBuilder builds a RequiredIDDocument
type RequiredIDDocumentBuilder struct {
	filter RequestedDocumentFilter
}

// WithFilter sets the filter on the required ID document
func (r RequiredIDDocumentBuilder) WithFilter(filter RequestedDocumentFilter) RequiredIDDocumentBuilder {
	r.filter = filter
	return r
}

// Build builds the RequiredIDDocument struct
func (r RequiredIDDocumentBuilder) Build() (*RequiredIDDocument, error) {
	return &RequiredIDDocument{
		Filter: r.filter,
	}, nil
}
