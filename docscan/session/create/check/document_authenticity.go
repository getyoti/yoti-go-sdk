package check

import (
	"encoding/json"

	"github.com/getyoti/yoti-go-sdk/v3/docscan/constants"
)

// RequestedDocumentAuthenticityCheck requests creation of a Document Authenticity Check
type RequestedDocumentAuthenticityCheck struct {
	config RequestedDocumentAuthenticityConfig
}

// Type is the type of the Requested Check
func (c RequestedDocumentAuthenticityCheck) Type() string {
	return string(constants.Authenticity)
}

// Config is the configuration of the Requested Check
func (c RequestedDocumentAuthenticityCheck) Config() RequestedCheckConfig {
	return RequestedCheckConfig(
		c.config,
	)
}

// MarshalJSON returns the JSON encoding
func (c RequestedDocumentAuthenticityCheck) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type   string               `json:"type"`
		Config RequestedCheckConfig `json:"config,omitempty"`
	}{
		Type:   c.Type(),
		Config: c.Config(),
	})
}

// RequestedDocumentAuthenticityConfig is the configuration applied when creating a Document Authenticity Check
type RequestedDocumentAuthenticityConfig struct {
}

// RequestedDocumentAuthenticityCheckBuilder builds a RequestedDocumentAuthenticityCheck
type RequestedDocumentAuthenticityCheckBuilder struct {
	config RequestedDocumentAuthenticityConfig
}

// NewRequestedDocumentAuthenticityCheckBuilder creates a new DocumentAuthenticityCheckBuilder
func NewRequestedDocumentAuthenticityCheckBuilder() *RequestedDocumentAuthenticityCheckBuilder {
	return &RequestedDocumentAuthenticityCheckBuilder{}
}

// Build builds the RequestedDocumentAuthenticityCheck
func (b *RequestedDocumentAuthenticityCheckBuilder) Build() (*RequestedDocumentAuthenticityCheck, error) {
	return &RequestedDocumentAuthenticityCheck{
		config: b.config,
	}, nil
}
