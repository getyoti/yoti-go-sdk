package filter

// DocumentFilter represents a document filter for checks and tasks
type DocumentFilter struct {
	DocumentTypes []string `json:"document_types"`
	CountryCodes  []string `json:"country_codes"`
}

// DocumentFilterBuilder builds a DocumentFilter
type DocumentFilterBuilder struct {
	documentTypes []string
	countryCodes  []string
}

// NewDocumentFilterBuilder creates a new DocumentFilterBuilder
func NewDocumentFilterBuilder() *DocumentFilterBuilder {
	return &DocumentFilterBuilder{
		documentTypes: []string{},
		countryCodes:  []string{},
	}
}

// WithCountryCode adds a country code to the filter
func (b *DocumentFilterBuilder) WithCountryCode(countryCode string) *DocumentFilterBuilder {
	b.countryCodes = append(b.countryCodes, countryCode)
	return b
}

// WithDocumentType adds a document type to the filter
func (b *DocumentFilterBuilder) WithDocumentType(documentType string) *DocumentFilterBuilder {
	b.documentTypes = append(b.documentTypes, documentType)
	return b
}

// Build creates a new DocumentFilter
func (b *DocumentFilterBuilder) Build() (DocumentFilter, error) {
	return DocumentFilter{
		DocumentTypes: b.documentTypes,
		CountryCodes:  b.countryCodes,
	}, nil
}
