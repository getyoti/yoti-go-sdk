package filter

import "encoding/json"

// RequestedDocumentRestrictionsFilter filters for a required document, allowing specification of restrictive parameters
type RequestedDocumentRestrictionsFilter struct {
	inclusion              string
	documents              []*RequestedDocumentRestriction
	allowExpiredDocuments  *bool
	allowNonLatinDocuments *bool
	allowDigitalIDs        *bool
	allowedProviders       []*DigitalIDProvider
}

// Type is the type of the document restriction filter
func (r RequestedDocumentRestrictionsFilter) Type() string {
	return documentRestriction
}

// MarshalJSON returns the JSON encoding
func (r *RequestedDocumentRestrictionsFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type                   string                          `json:"type"`
		Inclusion              string                          `json:"inclusion"`
		Documents              []*RequestedDocumentRestriction `json:"documents"`
		AllowExpiredDocuments  *bool                           `json:"allow_expired_documents,omitempty"`
		AllowNonLatinDocuments *bool                           `json:"allow_non_latin_documents,omitempty"`
		AllowDigitalIDs        *bool                           `json:"allow_digital_ids,omitempty"`
		AllowedProviders       []*DigitalIDProvider            `json:"allowed_providers,omitempty"`
	}{
		Type:                   r.Type(),
		Inclusion:              r.inclusion,
		Documents:              r.documents,
		AllowExpiredDocuments:  r.allowExpiredDocuments,
		AllowNonLatinDocuments: r.allowNonLatinDocuments,
		AllowDigitalIDs:        r.allowDigitalIDs,
		AllowedProviders:       r.allowedProviders,
	})
}

// RequestedDocumentRestrictionsFilterBuilder builds a RequestedDocumentRestrictionsFilter
type RequestedDocumentRestrictionsFilterBuilder struct {
	inclusion              string
	documents              []*RequestedDocumentRestriction
	allowExpiredDocuments  *bool
	allowNonLatinDocuments *bool
	allowDigitalIDs        *bool
	allowedProviders       []*DigitalIDProvider
}

// NewRequestedDocumentRestrictionsFilterBuilder creates a new RequestedDocumentRestrictionsFilterBuilder
func NewRequestedDocumentRestrictionsFilterBuilder() *RequestedDocumentRestrictionsFilterBuilder {
	return &RequestedDocumentRestrictionsFilterBuilder{
		documents: []*RequestedDocumentRestriction{},
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
func (b *RequestedDocumentRestrictionsFilterBuilder) WithDocumentRestriction(docRestriction *RequestedDocumentRestriction) *RequestedDocumentRestrictionsFilterBuilder {
	b.documents = append(b.documents, docRestriction)
	return b
}

// WithExpiredDocuments sets a bool value to allowExpiredDocuments on filter
func (b *RequestedDocumentRestrictionsFilterBuilder) WithExpiredDocuments(allowExpiredDocuments bool) *RequestedDocumentRestrictionsFilterBuilder {
	b.allowExpiredDocuments = &allowExpiredDocuments
	return b
}

// WithExpiredDocuments sets a bool value to allowExpiredDocuments on filter
func (b *RequestedDocumentRestrictionsFilterBuilder) WithAllowNonLatinDocuments(allowNonLatinDocuments bool) *RequestedDocumentRestrictionsFilterBuilder {
	b.allowNonLatinDocuments = &allowNonLatinDocuments
	return b
}

// WithAllowDigitalIDs sets a bool value to allowDigitalIDs on the filter
func (b *RequestedDocumentRestrictionsFilterBuilder) WithAllowDigitalIDs(allowDigitalIDs bool) *RequestedDocumentRestrictionsFilterBuilder {
	b.allowDigitalIDs = &allowDigitalIDs
	return b
}

// WithAllowedProviders sets the list of allowed digital ID providers on the filter, replacing any previously set value
func (b *RequestedDocumentRestrictionsFilterBuilder) WithAllowedProviders(allowedProviders []*DigitalIDProvider) *RequestedDocumentRestrictionsFilterBuilder {
	b.allowedProviders = allowedProviders
	return b
}

// WithAllowedProvider adds a single allowed digital ID provider to the filter
func (b *RequestedDocumentRestrictionsFilterBuilder) WithAllowedProvider(allowedProvider *DigitalIDProvider) *RequestedDocumentRestrictionsFilterBuilder {
	b.allowedProviders = append(b.allowedProviders, allowedProvider)
	return b
}

// Build creates a new RequestedDocumentRestrictionsFilter
func (b *RequestedDocumentRestrictionsFilterBuilder) Build() (*RequestedDocumentRestrictionsFilter, error) {
	return &RequestedDocumentRestrictionsFilter{
		inclusion:              b.inclusion,
		documents:              b.documents,
		allowExpiredDocuments:  b.allowExpiredDocuments,
		allowNonLatinDocuments: b.allowNonLatinDocuments,
		allowDigitalIDs:        b.allowDigitalIDs,
		allowedProviders:       b.allowedProviders,
	}, nil
}
