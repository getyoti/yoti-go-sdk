package filter

// RequestedDocumentRestriction represents a document Restriction for checks and tasks
type RequestedDocumentRestriction struct {
	DocumentTypes []string `json:"document_types"`
	CountryCodes  []string `json:"country_codes"`
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

// WithCountryCodes sets the country codes of the Restriction
func (b *RequestedDocumentRestrictionBuilder) WithCountryCodes(countryCodes []string) *RequestedDocumentRestrictionBuilder {
	b.countryCodes = countryCodes
	return b
}

// WithDocumentType adds a document type to the Restriction
func (b *RequestedDocumentRestrictionBuilder) WithDocumentType(documentType string) *RequestedDocumentRestrictionBuilder {
	b.documentTypes = append(b.documentTypes, documentType)
	return b
}

// Build creates a new RequestedDocumentRestriction
func (b *RequestedDocumentRestrictionBuilder) Build() (*RequestedDocumentRestriction, error) {
	return &RequestedDocumentRestriction{
		DocumentTypes: b.documentTypes,
		CountryCodes:  b.countryCodes,
	}, nil
}
