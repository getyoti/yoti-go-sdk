package filter

// RequestedOrthogonalRestrictionsFilter filters for a required document, allowing specification of restrictive parameters
type RequestedOrthogonalRestrictionsFilter struct {
	CountryRestriction *CountryRestriction `json:"country_restriction"`
	TypeRestriction    *TypeRestriction    `json:"type_restriction"`
}

// RequestedOrthogonalRestrictionsFilterBuilder builds a RequestedOrthogonalRestrictionsFilter
type RequestedOrthogonalRestrictionsFilterBuilder struct {
	countryRestriction *CountryRestriction
	typeRestriction    *TypeRestriction
}

// NewRequestedOrthogonalRestrictionsFilterBuilder creates a new RequestedOrthogonalRestrictionsFilterBuilder
func NewRequestedOrthogonalRestrictionsFilterBuilder() *RequestedOrthogonalRestrictionsFilterBuilder {
	return &RequestedOrthogonalRestrictionsFilterBuilder{
		countryRestriction: &CountryRestriction{},
		typeRestriction:    &TypeRestriction{},
	}
}

// WithIncludeCountryList sets an "INCLUDE" list of country codes on the filter
func (b *RequestedOrthogonalRestrictionsFilterBuilder) WithIncludeCountryList(countryCodes []string) *RequestedOrthogonalRestrictionsFilterBuilder {
	b.countryRestriction = &CountryRestriction{
		includeList,
		countryCodes,
	}
	return b
}

// WithExcludeCountryList sets an "EXCLUDE" list of country codes on the filter
func (b *RequestedOrthogonalRestrictionsFilterBuilder) WithExcludeCountryList(countryCodes []string) *RequestedOrthogonalRestrictionsFilterBuilder {
	b.countryRestriction = &CountryRestriction{
		excludeList,
		countryCodes,
	}
	return b
}

// WithIncludeDocumentTypeList sets an "INCLUDE" list of document types on the filter
func (b *RequestedOrthogonalRestrictionsFilterBuilder) WithIncludeDocumentTypeList(documentTypes []string) *RequestedOrthogonalRestrictionsFilterBuilder {
	b.typeRestriction = &TypeRestriction{
		includeList,
		documentTypes,
	}
	return b
}

// WithExcludeDocumentTypeList sets an "EXCLUDE" list of document types on the filter
func (b *RequestedOrthogonalRestrictionsFilterBuilder) WithExcludeDocumentTypeList(documentTypes []string) *RequestedOrthogonalRestrictionsFilterBuilder {
	b.typeRestriction = &TypeRestriction{
		excludeList,
		documentTypes,
	}
	return b
}

// Build creates a new RequestedOrthogonalRestrictionsFilter
func (b *RequestedOrthogonalRestrictionsFilterBuilder) Build() (*RequestedOrthogonalRestrictionsFilter, error) {
	return &RequestedOrthogonalRestrictionsFilter{
		b.countryRestriction,
		b.typeRestriction,
	}, nil
}
