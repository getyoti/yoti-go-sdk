package retrieve

import (
	"encoding/json"
	"errors"
)

type SessionConfigurationResponse struct {
	ClientSessionTokenTtl int              `json:"client_session_token_ttl"`
	SessionId             string           `json:"session_id"`
	RequestedChecks       []string         `json:"requested_checks"`
	Capture               *CaptureResponse `json:"capture"`
}

// NewSessionConfigurationResponse creates a new SessionConfigurationResponse from JSON payload bytes,
// validating mandatory fields.
func NewSessionConfigurationResponse(payload []byte) (*SessionConfigurationResponse, error) {
	var resp SessionConfigurationResponse
	if err := json.Unmarshal(payload, &resp); err != nil {
		return nil, err
	}

	if resp.ClientSessionTokenTtl <= 0 {
		return nil, errors.New("client_session_token_ttl must be a positive integer")
	}
	if resp.SessionId == "" {
		return nil, errors.New("session_id must be a non-empty string")
	}
	if resp.RequestedChecks == nil || len(resp.RequestedChecks) == 0 {
		return nil, errors.New("requested_checks must be a non-empty array")
	}

	return &resp, nil
}

// Getter methods for each field

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
