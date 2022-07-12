package check

import (
	"encoding/json"

	"github.com/getyoti/yoti-go-sdk/v3/docscan/constants"
)

// RequestedThirdPartyIdentityCheck requests creation of a third party CRA check
type RequestedThirdPartyIdentityCheck struct {
	config RequestedThirdPartyIdentityCheckConfig
}

// Type is the type of the requested check
func (c *RequestedThirdPartyIdentityCheck) Type() string {
	return constants.ThirdPartyIdentityCheck
}

// Config is the configuration of the requested check
func (c *RequestedThirdPartyIdentityCheck) Config() RequestedCheckConfig {
	return RequestedCheckConfig(c.config)
}

// MarshalJSON returns the JSON encoding
func (c *RequestedThirdPartyIdentityCheck) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type   string               `json:"type"`
		Config RequestedCheckConfig `json:"config,omitempty"`
	}{
		Type:   c.Type(),
		Config: c.Config(),
	})
}

// RequestedThirdPartyIdentityCheckConfig is the configuration applied when creating
// a third party identity check
type RequestedThirdPartyIdentityCheckConfig struct {
}

// RequestedThirdPartyIdentityCheckBuilder builds a RequestedThirdPartyIdentityCheck
type RequestedThirdPartyIdentityCheckBuilder struct {
	config RequestedThirdPartyIdentityCheckConfig
}

// NewRequestedThirdPartyIdentityCheckBuilder creates a new builder for RequestedThirdPartyIdentityCheck
func NewRequestedThirdPartyIdentityCheckBuilder() *RequestedThirdPartyIdentityCheckBuilder {
	return &RequestedThirdPartyIdentityCheckBuilder{}
}

// Build builds the RequestedThirdPartyIdentityCheck
func (b *RequestedThirdPartyIdentityCheckBuilder) Build() (*RequestedThirdPartyIdentityCheck, error) {
	return &RequestedThirdPartyIdentityCheck{
		config: b.config,
	}, nil
}
