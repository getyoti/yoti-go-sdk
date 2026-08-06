package filter

import "encoding/json"

// RequiredShareCode details a required share code resource
type RequiredShareCode struct {
	Issuer string
	Scheme string
}

// Type returns the type identifier
func (r *RequiredShareCode) Type() string {
	return "SHARE_CODE"
}

// MarshalJSON returns the JSON encoding
func (r *RequiredShareCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Issuer string `json:"issuer"`
		Scheme string `json:"scheme"`
	}{
		Issuer: r.Issuer,
		Scheme: r.Scheme,
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

// WithIssuer sets the issuer
func (r *RequiredShareCodeBuilder) WithIssuer(issuer string) *RequiredShareCodeBuilder {
	r.issuer = issuer
	return r
}

// WithScheme sets the scheme
func (r *RequiredShareCodeBuilder) WithScheme(scheme string) *RequiredShareCodeBuilder {
	r.scheme = scheme
	return r
}

// Build builds the RequiredShareCode struct
func (r *RequiredShareCodeBuilder) Build() (*RequiredShareCode, error) {
	return &RequiredShareCode{
		Issuer: r.issuer,
		Scheme: r.scheme,
	}, nil
}
