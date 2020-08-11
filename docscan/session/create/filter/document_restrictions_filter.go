package filter

// RequestedDocumentRestrictionsFilter filters for a required document, allowing specification of restrictive parameters
type RequestedDocumentRestrictionsFilter struct {
	Inclusion string                         `json:"inclusion"`
	Documents []RequestedDocumentRestriction `json:"documents"`
}

// RequestedDocumentRestrictionsFilterBuilder builds a RequestedDocumentRestrictionsFilter
type RequestedDocumentRestrictionsFilterBuilder struct {
	inclusion string
	documents []RequestedDocumentRestriction
}

// NewRequestedDocumentRestrictionsFilterBuilder creates a new RequestedDocumentRestrictionsFilterBuilder
func NewRequestedDocumentRestrictionsFilterBuilder() *RequestedDocumentRestrictionsFilterBuilder {
	return &RequestedDocumentRestrictionsFilterBuilder{
		documents: []RequestedDocumentRestriction{},
	}
}

// ForIncludeList sets the type restriction to INCLUDE the document restrictions
func (b *RequestedDocumentRestrictionsFilterBuilder) ForIncludeList() *RequestedDocumentRestrictionsFilterBuilder {
	b.inclusion = includeList
	return b
}

// ForExcludeList sets the type restriction to EXCLUDE the document restrictions
func (b *RequestedDocumentRestrictionsFilterBuilder) ForExcludeList() *RequestedDocumentRestrictionsFilterBuilder {
	b.inclusion = excludeList
	return b
}

// WithDocumentRestriction adds a document restriction to the filter
func (b *RequestedDocumentRestrictionsFilterBuilder) WithDocumentRestriction(docRestriction RequestedDocumentRestriction) *RequestedDocumentRestrictionsFilterBuilder {
	b.documents = append(b.documents, docRestriction)
	return b
}

// Build creates a new RequestedDocumentRestrictionsFilter
func (b *RequestedDocumentRestrictionsFilterBuilder) Build() (*RequestedDocumentRestrictionsFilter, error) {
	return &RequestedDocumentRestrictionsFilter{
		b.inclusion,
		b.documents,
	}, nil
}
