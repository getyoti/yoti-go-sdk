package config

// SessionConfigurationResponse represents the configuration of a Doc Scan session.
type SessionConfigurationResponse struct {
	ClientSessionTokenTTL int              `json:"client_session_token_ttl"`
	SessionID             string           `json:"session_id"`
	RequestedChecks       []string         `json:"requested_checks"`
	Capture               *CaptureResponse `json:"capture"`
}

// GetClientSessionTokenTTL returns the amount of time remaining in seconds until the session expires.
func (s *SessionConfigurationResponse) GetClientSessionTokenTTL() int {
	return s.ClientSessionTokenTTL
}

// GetSessionID returns the session ID that the configuration belongs to.
func (s *SessionConfigurationResponse) GetSessionID() string {
	return s.SessionID
}

// GetRequestedChecks returns a list of strings, signifying the checks that have been requested in the session.
func (s *SessionConfigurationResponse) GetRequestedChecks() []string {
	return s.RequestedChecks
}

// GetCapture returns information about what needs to be captured to fulfil the session's requirements.
func (s *SessionConfigurationResponse) GetCapture() *CaptureResponse {
	return s.Capture
}

// CaptureResponse represents the capture requirements for the session.
type CaptureResponse struct {
	BiometricConsent  string              `json:"biometric_consent"`
	RequiredResources []*RequiredResource `json:"required_resources"`
}

// RequiredResource represents a resource that needs to be captured.
type RequiredResource struct {
	Type                  string              `json:"type"`
	ID                    string              `json:"id,omitempty"`
	State                 string              `json:"state,omitempty"`
	AllowedSources        []*AllowedSource    `json:"allowed_sources,omitempty"`
	SupportedCountries    []*SupportedCountry `json:"supported_countries,omitempty"`
	AllowedCaptureMethods string              `json:"allowed_capture_methods,omitempty"`
	RequestedTasks        []*RequestedTask    `json:"requested_tasks,omitempty"`
	AttemptsRemaining     map[string]int      `json:"attempts_remaining,omitempty"`
	DocumentTypes         []string            `json:"document_types,omitempty"`
	CountryCodes          []string            `json:"country_codes,omitempty"`
	Objective             *Objective          `json:"objective,omitempty"`
	LivenessType          string              `json:"liveness_type,omitempty"`
}

// AllowedSource represents a source that is allowed for a resource.
type AllowedSource struct {
	Type string `json:"type"`
}

// SupportedCountry represents a country and its supported documents.
type SupportedCountry struct {
	Code               string               `json:"code"`
	SupportedDocuments []*SupportedDocument `json:"supported_documents"`
}

// SupportedDocument represents a document type supported for a country.
type SupportedDocument struct {
	Type string `json:"type"`
}

// RequestedTask represents a task that has been requested for a resource.
type RequestedTask struct {
	Type  string `json:"type"`
	State string `json:"state,omitempty"`
}

// Objective represents the objective for a supplementary document.
type Objective struct {
	Type string `json:"type"`
}
