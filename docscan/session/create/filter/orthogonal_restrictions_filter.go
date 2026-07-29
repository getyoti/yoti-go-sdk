package filter

import "encoding/json"

// RequestedOrthogonalRestrictionsFilter filters for a required document, allowing specification of restrictive parameters
type RequestedOrthogonalRestrictionsFilter struct {
	countryRestriction     *CountryRestriction
	typeRestriction        *TypeRestriction
	allowExpiredDocuments  *bool
	allowNonLatinDocuments *bool
	allowDigitalIDs        *bool
	allowedProviders       []*DigitalIDProvider
}

// Type returns the type of the RequestedOrthogonalRestrictionsFilter
func (r RequestedOrthogonalRestrictionsFilter) Type() string {
	return orthogonalRestriction
}

// MarshalJSON returns the JSON encoding
func (r RequestedOrthogonalRestrictionsFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type                   string               `json:"type"`
		CountryRestriction     *CountryRestriction  `json:"country_restriction,omitempty"`
		TypeRestriction        *TypeRestriction     `json:"type_restriction,omitempty"`
		AllowExpiredDocuments  *bool                `json:"allow_expired_documents,omitempty"`
		AllowNonLatinDocuments *bool                `json:"allow_non_latin_documents,omitempty"`
		AllowDigitalIDs        *bool                `json:"allow_digital_ids,omitempty"`
		AllowedProviders       []*DigitalIDProvider `json:"allowed_providers,omitempty"`
	}{
		CountryRestriction:     r.countryRestriction,
		TypeRestriction:        r.typeRestriction,
		Type:                   r.Type(),
		AllowExpiredDocuments:  r.allowExpiredDocuments,
		AllowNonLatinDocuments: r.allowNonLatinDocuments,
		AllowDigitalIDs:        r.allowDigitalIDs,
		AllowedProviders:       r.allowedProviders,
	})
}

// RequestedOrthogonalRestrictionsFilterBuilder builds a RequestedOrthogonalRestrictionsFilter
type RequestedOrthogonalRestrictionsFilterBuilder struct {
	countryRestriction     *CountryRestriction
	typeRestriction        *TypeRestriction
	allowExpiredDocuments  *bool
	allowNonLatinDocuments *bool
	allowDigitalIDs        *bool
	allowedProviders       []*DigitalIDProvider
}

// NewRequestedOrthogonalRestrictionsFilterBuilder creates a new RequestedOrthogonalRestrictionsFilterBuilder
func NewRequestedOrthogonalRestrictionsFilterBuilder() *RequestedOrthogonalRestrictionsFilterBuilder {
	return &RequestedOrthogonalRestrictionsFilterBuilder{
		countryRestriction:     nil,
		typeRestriction:        nil,
		allowExpiredDocuments:  nil,
		allowNonLatinDocuments: nil,
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

// WithNonLatinDocuments sets a bool value to allowNonLatinDocuments on filter
func (b *RequestedOrthogonalRestrictionsFilterBuilder) WithNonLatinDocuments(allowNonLatinDocuments bool) *RequestedOrthogonalRestrictionsFilterBuilder {
	b.allowNonLatinDocuments = &allowNonLatinDocuments
	return b
}

// WithExpiredDocuments sets a bool value to allowExpiredDocuments on filter
func (b *RequestedOrthogonalRestrictionsFilterBuilder) WithExpiredDocuments(allowExpiredDocuments bool) *RequestedOrthogonalRestrictionsFilterBuilder {
	b.allowExpiredDocuments = &allowExpiredDocuments
	return b
}

// WithAllowDigitalIDs sets a bool value to allowDigitalIDs on the filter
func (b *RequestedOrthogonalRestrictionsFilterBuilder) WithAllowDigitalIDs(allowDigitalIDs bool) *RequestedOrthogonalRestrictionsFilterBuilder {
	b.allowDigitalIDs = &allowDigitalIDs
	return b
}

// WithAllowedProviders sets the list of allowed digital ID providers on the filter, replacing any previously set value
func (b *RequestedOrthogonalRestrictionsFilterBuilder) WithAllowedProviders(allowedProviders []*DigitalIDProvider) *RequestedOrthogonalRestrictionsFilterBuilder {
	b.allowedProviders = allowedProviders
	return b
}

// WithAllowedProvider adds a single allowed digital ID provider to the filter
func (b *RequestedOrthogonalRestrictionsFilterBuilder) WithAllowedProvider(allowedProvider *DigitalIDProvider) *RequestedOrthogonalRestrictionsFilterBuilder {
	b.allowedProviders = append(b.allowedProviders, allowedProvider)
	return b
}

// Build creates a new RequestedOrthogonalRestrictionsFilter
func (b *RequestedOrthogonalRestrictionsFilterBuilder) Build() (*RequestedOrthogonalRestrictionsFilter, error) {
	return &RequestedOrthogonalRestrictionsFilter{
		countryRestriction:     b.countryRestriction,
		typeRestriction:        b.typeRestriction,
		allowExpiredDocuments:  b.allowExpiredDocuments,
		allowNonLatinDocuments: b.allowNonLatinDocuments,
		allowDigitalIDs:        b.allowDigitalIDs,
		allowedProviders:       b.allowedProviders,
	}, nil
}
