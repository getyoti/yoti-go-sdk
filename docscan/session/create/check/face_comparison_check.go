package check

import (
	"encoding/json"

	"github.com/getyoti/yoti-go-sdk/v3/docscan/constants"
)

// RequestedFaceComparisonCheck represents a face comparison check request.
type RequestedFaceComparisonCheck struct {
	config RequestedFaceComparisonConfig
}

// Type returns the type of the check.
func (c *RequestedFaceComparisonCheck) Type() string {
	return constants.FaceComparison
}

// Config returns the configuration of the check.
func (c *RequestedFaceComparisonCheck) Config() RequestedCheckConfig {
	return c.config
}

// MarshalJSON encodes the struct into JSON.
func (c *RequestedFaceComparisonCheck) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type   string               `json:"type"`
		Config RequestedCheckConfig `json:"config,omitempty"`
	}{
		Type:   c.Type(),
		Config: c.Config(),
	})
}
