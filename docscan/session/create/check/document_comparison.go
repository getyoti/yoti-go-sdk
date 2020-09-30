package check

import (
	"encoding/json"

	"github.com/getyoti/yoti-go-sdk/v3/docscan/constants"
)

// RequestedIDDocumentComparisonCheck requests creation of a Document Comparison Check
type RequestedIDDocumentComparisonCheck struct {
	config RequestedIDDocumentComparisonConfig
}

// Type is the type of the Requested Check
func (c *RequestedIDDocumentComparisonCheck) Type() string {
	return constants.IDDocumentComparison
}

// Config is the configuration of the Requested Check
func (c *RequestedIDDocumentComparisonCheck) Config() RequestedCheckConfig {
	return RequestedCheckConfig(
		c.config,
	)
}

// MarshalJSON returns the JSON encoding
func (c *RequestedIDDocumentComparisonCheck) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type   string               `json:"type"`
		Config RequestedCheckConfig `json:"config,omitempty"`
	}{
		Type:   c.Type(),
		Config: c.Config(),
	})
}

// RequestedIDDocumentComparisonConfig is the configuration applied when creating a Document Comparison Check
type RequestedIDDocumentComparisonConfig struct {
}

// RequestedIDDocumentComparisonCheckBuilder builds a RequestedIDDocumentComparisonCheck
type RequestedIDDocumentComparisonCheckBuilder struct {
	config RequestedIDDocumentComparisonConfig
}

// NewRequestedIDDocumentComparisonCheckBuilder creates a new DocumentComparisonCheckBuilder
func NewRequestedIDDocumentComparisonCheckBuilder() *RequestedIDDocumentComparisonCheckBuilder {
	return &RequestedIDDocumentComparisonCheckBuilder{}
}

// Build builds the RequestedIDDocumentComparisonCheck
func (b *RequestedIDDocumentComparisonCheckBuilder) Build() (*RequestedIDDocumentComparisonCheck, error) {
	return &RequestedIDDocumentComparisonCheck{
		config: b.config,
	}, nil
}
