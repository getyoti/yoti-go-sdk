package sandbox

type documentFilter struct {
	DocumentTypes []string `json:"document_types"`
	CountryCodes  []string `json:"country_codes"`
}

type documentFilterBuilder struct {
	documentTypes []string
	countryCodes  []string
}

func NewDocumentFilterBuilder() *documentFilterBuilder {
	return &documentFilterBuilder{}
}

func (b *documentFilterBuilder) WithCountryCode(countryCode string) *documentFilterBuilder {
	b.countryCodes = append(b.countryCodes, countryCode)
	return b
}

func (b *documentFilterBuilder) WithDocumentType(documentType string) *documentFilterBuilder {
	b.documentTypes = append(b.documentTypes, documentType)
	return b
}

func (b *documentFilterBuilder) Build() documentFilter {
	return documentFilter{
		DocumentTypes: b.documentTypes,
		CountryCodes:  b.countryCodes,
	}
}
