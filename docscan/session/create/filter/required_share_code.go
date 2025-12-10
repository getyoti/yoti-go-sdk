package filter

import (
	"encoding/json"
)

// RequiredShareCode details a required share code
type RequiredShareCode struct {
	Issuer string
	Scheme string
}

// Type returns the type of the share code
func (s *RequiredShareCode) Type() string {
	return shareCode
}

// MarshalJSON returns the JSON encoding
func (s *RequiredShareCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type   string `json:"type"`
		Issuer string `json:"issuer,omitempty"`
		Scheme string `json:"scheme,omitempty"`
	}{
		Type:   s.Type(),
		Issuer: s.Issuer,
		Scheme: s.Scheme,
	})
}

// NewRequiredShareCodeBuilder creates a new RequiredShareCodeBuilder
func NewRequiredShareCodeBuilder() *RequiredShareCodeBuilder {
	return &RequiredShareCodeBuilder{}
}

// RequiredShareCodeBuilder builds a RequiredShareCode
type RequiredShareCodeBuilder struct {
	issuer string
	scheme string
}

// WithIssuer sets the issuer on the required share code
func (r RequiredShareCodeBuilder) WithIssuer(issuer string) RequiredShareCodeBuilder {
	r.issuer = issuer
	return r
}

// WithScheme sets the scheme on the required share code
func (r RequiredShareCodeBuilder) WithScheme(scheme string) RequiredShareCodeBuilder {
	r.scheme = scheme
	return r
}

// Build builds the RequiredShareCode struct
func (r RequiredShareCodeBuilder) Build() (*RequiredShareCode, error) {
	return &RequiredShareCode{
		Issuer: r.issuer,
		Scheme: r.scheme,
	}, nil
}
