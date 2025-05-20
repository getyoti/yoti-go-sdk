package retrieve

import (
	"encoding/json"
	"errors"
)

// CaptureResponse should be defined elsewhere
// type CaptureResponse struct { ... }

type SessionConfigurationResponse struct {
	ClientSessionTokenTtl int              `json:"client_session_token_ttl"`
	SessionId             string           `json:"session_id"`
	RequestedChecks       []string         `json:"requested_checks"`
	Capture               *CaptureResponse `json:"capture"`
}

// NewSessionConfigurationResponse creates a new SessionConfigurationResponse from JSON payload bytes
func NewSessionConfigurationResponse(payload []byte) (*SessionConfigurationResponse, error) {
	var resp SessionConfigurationResponse
	if err := json.Unmarshal(payload, &resp); err != nil {
		return nil, err
	}

	// Validate required fields similar to JS Validation.isX
	if resp.ClientSessionTokenTtl == 0 {
		return nil, errors.New("client_session_token_ttl must be a non-zero number")
	}
	if resp.SessionId == "" {
		return nil, errors.New("session_id must be a non-empty string")
	}
	if resp.RequestedChecks == nil {
		return nil, errors.New("requested_checks must be an array")
	}
	// Assuming CaptureResponse struct has its own validation if needed

	return &resp, nil
}

// Getters, matching JS class methods (optional in Go, as fields are accessible)
// but provided here for API similarity

func (r *SessionConfigurationResponse) GetClientSessionTokenTtl() int {
	return r.ClientSessionTokenTtl
}

func (r *SessionConfigurationResponse) GetSessionId() string {
	return r.SessionId
}

func (r *SessionConfigurationResponse) GetRequestedChecks() []string {
	return r.RequestedChecks
}

func (r *SessionConfigurationResponse) GetCapture() *CaptureResponse {
	return r.Capture
}
