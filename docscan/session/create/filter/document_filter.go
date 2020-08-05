package filter

// RequestedDocumentFilter filters for a required document, allowing specification of restrictive parameters
type RequestedDocumentFilter struct {
	DocumentTypes []string `json:"document_types"`
	CountryCodes  []string `json:"country_codes"`
}

// RequestedDocumentFilterBuilder builds a DocumentFilter
type RequestedDocumentFilterBuilder struct {
	documentTypes []string
	countryCodes  []string
}

// NewRequestedDocumentFilterBuilder creates a new RequestedDocumentFilterBuilder
func NewRequestedDocumentFilterBuilder() *RequestedDocumentFilterBuilder {
	return &RequestedDocumentFilterBuilder{
		documentTypes: []string{},
		countryCodes:  []string{},
	}
}

// WithCountryCode adds a country code to the filter
func (b *RequestedDocumentFilterBuilder) WithCountryCode(countryCode string) *RequestedDocumentFilterBuilder {
	b.countryCodes = append(b.countryCodes, countryCode)
	return b
}

// WithCountryCode sets the country codes on the filter
func (b *RequestedDocumentFilterBuilder) WithCountryCodes(countryCodes []string) *RequestedDocumentFilterBuilder {
	b.countryCodes = countryCodes
	return b
}

// WithDocumentType adds a document type to the filter
func (b *RequestedDocumentFilterBuilder) WithDocumentType(documentType string) *RequestedDocumentFilterBuilder {
	b.documentTypes = append(b.documentTypes, documentType)
	return b
}

// Build creates a new RequestedDocumentFilter
func (b *RequestedDocumentFilterBuilder) Build() (*RequestedDocumentFilter, error) {
	return &RequestedDocumentFilter{
		b.documentTypes,
		b.countryCodes,
	}, nil
}
