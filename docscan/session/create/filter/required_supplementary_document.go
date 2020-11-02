package filter

import (
	"encoding/json"

	"github.com/getyoti/yoti-go-sdk/v3/docscan/session/create/objective"
)

// RequiredSupplementaryDocument details a required supplementary document
type RequiredSupplementaryDocument struct {
	Filter        RequestedDocumentFilter
	DocumentTypes []string
	CountryCodes  []string
	Objective     objective.Objective
}

// Type returns the type of the supplementary document
func (r *RequiredSupplementaryDocument) Type() string {
	return supplementaryDocument
}

// MarshalJSON returns the JSON encoding
func (r *RequiredSupplementaryDocument) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type          string                  `json:"type"`
		Filter        RequestedDocumentFilter `json:"filter,omitempty"`
		CountryCodes  []string                `json:"country_codes,omitempty"`
		DocumentTypes []string                `json:"document_types,omitempty"`
		Objective     objective.Objective     `json:"objective,omitempty"`
	}{
		Type:          r.Type(),
		Filter:        r.Filter,
		CountryCodes:  r.CountryCodes,
		DocumentTypes: r.DocumentTypes,
		Objective:     r.Objective,
	})
}

// NewRequiredSupplementaryDocumentBuilder creates a new RequiredSupplementaryDocumentBuilder
func NewRequiredSupplementaryDocumentBuilder() *RequiredSupplementaryDocumentBuilder {
	return &RequiredSupplementaryDocumentBuilder{}
}

// RequiredSupplementaryDocumentBuilder builds a RequiredSupplementaryDocument
type RequiredSupplementaryDocumentBuilder struct {
	filter        RequestedDocumentFilter
	documentTypes []string
	countryCodes  []string
	objective     objective.Objective
}

// WithFilter sets the filter on the required supplementary document
func (r *RequiredSupplementaryDocumentBuilder) WithFilter(filter RequestedDocumentFilter) *RequiredSupplementaryDocumentBuilder {
	r.filter = filter
	return r
}

// WithCountryCodes sets the country codes on the required supplementary document
func (r *RequiredSupplementaryDocumentBuilder) WithCountryCodes(countryCodes []string) *RequiredSupplementaryDocumentBuilder {
	r.countryCodes = countryCodes
	return r
}

// WithDocumentTypes sets the document types on the required supplementary document
func (r *RequiredSupplementaryDocumentBuilder) WithDocumentTypes(documentTypes []string) *RequiredSupplementaryDocumentBuilder {
	r.documentTypes = documentTypes
	return r
}

// WithObjective sets the objective for the required supplementary document
func (r *RequiredSupplementaryDocumentBuilder) WithObjective(objective objective.Objective) *RequiredSupplementaryDocumentBuilder {
	r.objective = objective
	return r
}

// Build builds the RequiredSupplementaryDocument struct
func (r *RequiredSupplementaryDocumentBuilder) Build() (*RequiredSupplementaryDocument, error) {
	return &RequiredSupplementaryDocument{
		Filter:        r.filter,
		DocumentTypes: r.documentTypes,
		CountryCodes:  r.countryCodes,
		Objective:     r.objective,
	}, nil
}
