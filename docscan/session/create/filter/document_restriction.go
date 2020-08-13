package filter

// RequestedDocumentRestriction represents a document filter for checks and tasks
type RequestedDocumentRestriction struct {
	DocumentTypes []string `json:"document_types,omitempty"`
	CountryCodes  []string `json:"country_codes,omitempty"`
}

// RequestedDocumentRestrictionBuilder builds a RequestedDocumentRestriction
type RequestedDocumentRestrictionBuilder struct {
	documentTypes []string
	countryCodes  []string
}

// NewRequestedDocumentRestrictionBuilder creates a new RequestedDocumentRestrictionBuilder
func NewRequestedDocumentRestrictionBuilder() *RequestedDocumentRestrictionBuilder {
	return &RequestedDocumentRestrictionBuilder{
		documentTypes: []string{},
		countryCodes:  []string{},
	}
}

// WithCountryCodes sets the country codes of the filter
func (b *RequestedDocumentRestrictionBuilder) WithCountryCodes(countryCodes []string) *RequestedDocumentRestrictionBuilder {
	b.countryCodes = countryCodes
	return b
}

// WithDocumentTypes sets the document types of the filter
func (b *RequestedDocumentRestrictionBuilder) WithDocumentTypes(documentTypes []string) *RequestedDocumentRestrictionBuilder {
	b.documentTypes = documentTypes
	return b
}

// Build creates a new RequestedDocumentRestriction
func (b *RequestedDocumentRestrictionBuilder) Build() (*RequestedDocumentRestriction, error) {
	return &RequestedDocumentRestriction{
		DocumentTypes: b.documentTypes,
		CountryCodes:  b.countryCodes,
	}, nil
}
