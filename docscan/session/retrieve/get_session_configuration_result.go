package retrieve

import (
	"encoding/json"
	"fmt"
)

// GetSessionConfigurationResult models the response for /sessions/{sessionId}/configuration
// Mirrors SessionConfigurationResponse from Node.js

type GetSessionConfigurationResult struct {
	ClientSessionTokenTTL int             `json:"client_session_token_ttl"`
	SessionID             string          `json:"session_id"`
	RequestedChecks       []string        `json:"requested_checks"`
	Capture               CaptureResponse `json:"capture"`
}

// CaptureResponse contains details about capture resource requirements
// e.g. face capture

type CaptureResponse struct {
	FaceCaptureResourceRequirements []FaceCaptureResourceRequirement `json:"face_capture_resource_requirements"`
}

type FaceCaptureResourceRequirement struct {
	ID string `json:"id"`
}

// Utility Getters (optional)

func (r *GetSessionConfigurationResult) GetClientSessionTokenTTL() int {
	return r.ClientSessionTokenTTL
}

func (r *GetSessionConfigurationResult) GetSessionID() string {
	return r.SessionID
}

func (r *GetSessionConfigurationResult) GetRequestedChecks() []string {
	return r.RequestedChecks
}

func (r *GetSessionConfigurationResult) GetCapture() CaptureResponse {
	return r.Capture
}

func (r *GetSessionConfigurationResult) String() string {
	b, _ := json.MarshalIndent(r, "", "  ")
	return fmt.Sprintf("SessionConfigurationResult: %s", string(b))
}
