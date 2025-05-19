package check

// RequestedFaceComparisonConfig is the configuration for a face comparison check.
type RequestedFaceComparisonConfig struct {
	ManualCheck string `json:"manual_check,omitempty"`
}

// Config interface implementation marker
func (RequestedFaceComparisonConfig) isRequestedCheckConfig() {}
