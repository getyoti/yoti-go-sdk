package filter

import "encoding/json"

// RequestedOrthogonalRestrictionsFilter filters for a required document, allowing specification of restrictive parameters
type RequestedOrthogonalRestrictionsFilter struct {
	countryRestriction    *CountryRestriction
	typeRestriction       *TypeRestriction
	allowExpiredDocuments *bool
}

// Type returns the type of the RequestedOrthogonalRestrictionsFilter
func (r RequestedOrthogonalRestrictionsFilter) Type() string {
	return orthogonalRestriction
}

// MarshalJSON returns the JSON encoding
func (r RequestedOrthogonalRestrictionsFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type                  string              `json:"type"`
		CountryRestriction    *CountryRestriction `json:"country_restriction,omitempty"`
		TypeRestriction       *TypeRestriction    `json:"type_restriction,omitempty"`
		AllowExpiredDocuments *bool               `json:"allow_expired_documents,omitempty"`
	}{
		CountryRestriction:    r.countryRestriction,
		TypeRestriction:       r.typeRestriction,
		Type:                  r.Type(),
		AllowExpiredDocuments: r.allowExpiredDocuments,
	})
}

// RequestedOrthogonalRestrictionsFilterBuilder builds a RequestedOrthogonalRestrictionsFilter
type RequestedOrthogonalRestrictionsFilterBuilder struct {
	countryRestriction    *CountryRestriction
	typeRestriction       *TypeRestriction
	allowExpiredDocuments *bool
}

// NewRequestedOrthogonalRestrictionsFilterBuilder creates a new RequestedOrthogonalRestrictionsFilterBuilder
func NewRequestedOrthogonalRestrictionsFilterBuilder() *RequestedOrthogonalRestrictionsFilterBuilder {
	return &RequestedOrthogonalRestrictionsFilterBuilder{
		countryRestriction:    nil,
		typeRestriction:       nil,
		allowExpiredDocuments: nil,
	}
}

// WithIncludedCountries sets an "INCLUDE" slice of country codes on the filter
func (b *RequestedOrthogonalRestrictionsFilterBuilder) WithIncludedCountries(countryCodes []string) *RequestedOrthogonalRestrictionsFilterBuilder {
	b.countryRestriction = &CountryRestriction{
		includeList,
		countryCodes,
	}
	return b
}

// WithExcludedCountries sets an "EXCLUDE" slice of country codes on the filter
func (b *RequestedOrthogonalRestrictionsFilterBuilder) WithExcludedCountries(countryCodes []string) *RequestedOrthogonalRestrictionsFilterBuilder {
	b.countryRestriction = &CountryRestriction{
		excludeList,
		countryCodes,
	}
	return b
}

// WithIncludedDocumentTypes sets an "INCLUDE" slice of document types on the filter
func (b *RequestedOrthogonalRestrictionsFilterBuilder) WithIncludedDocumentTypes(documentTypes []string) *RequestedOrthogonalRestrictionsFilterBuilder {
	b.typeRestriction = &TypeRestriction{
		includeList,
		documentTypes,
	}
	return b
}

// WithExcludedDocumentTypes sets an "EXCLUDE" slice of document types on the filter
func (b *RequestedOrthogonalRestrictionsFilterBuilder) WithExcludedDocumentTypes(documentTypes []string) *RequestedOrthogonalRestrictionsFilterBuilder {
	b.typeRestriction = &TypeRestriction{
		excludeList,
		documentTypes,
	}
	return b
}

func (b *RequestedOrthogonalRestrictionsFilterBuilder) WithExpiredDocuments(allowExpiredDocuments bool) *RequestedOrthogonalRestrictionsFilterBuilder {
	b.allowExpiredDocuments = &allowExpiredDocuments
	return b
}

// Build creates a new RequestedOrthogonalRestrictionsFilter
func (b *RequestedOrthogonalRestrictionsFilterBuilder) Build() (*RequestedOrthogonalRestrictionsFilter, error) {
	return &RequestedOrthogonalRestrictionsFilter{
		b.countryRestriction,
		b.typeRestriction,
		b.allowExpiredDocuments,
	}, nil
}
