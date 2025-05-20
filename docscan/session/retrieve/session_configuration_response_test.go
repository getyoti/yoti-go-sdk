package retrieve

import (
	"encoding/json"
	"testing"

	"gotest.tools/v3/assert"
)

func TestNewSessionConfigurationResponse_Success(t *testing.T) {
	payload := SessionConfigurationResponse{
		ClientSessionTokenTtl: 3600,
		SessionId:             "abc123",
		RequestedChecks:       []string{"ID_DOCUMENT"},
		Capture:               &CaptureResponse{}, // assuming zero value is acceptable
	}

	data, err := json.Marshal(payload)
	assert.NilError(t, err)

	result, err := NewSessionConfigurationResponse(data)
	assert.NilError(t, err)
	assert.Equal(t, result.ClientSessionTokenTtl, 3600)
	assert.Equal(t, result.SessionId, "abc123")
	assert.DeepEqual(t, result.RequestedChecks, []string{"ID_DOCUMENT"})
	assert.Assert(t, result.Capture != nil)
}

func TestNewSessionConfigurationResponse_MissingTTL(t *testing.T) {
	jsonData := `{
		"session_id": "abc123",
		"requested_checks": ["ID_DOCUMENT"],
		"capture": {}
	}`

	_, err := NewSessionConfigurationResponse([]byte(jsonData))
	assert.ErrorContains(t, err, "client_session_token_ttl must be a positive integer")
}

func TestNewSessionConfigurationResponse_MissingSessionID(t *testing.T) {
	jsonData := `{
		"client_session_token_ttl": 3600,
		"requested_checks": ["ID_DOCUMENT"],
		"capture": {}
	}`

	_, err := NewSessionConfigurationResponse([]byte(jsonData))
	assert.ErrorContains(t, err, "session_id must be a non-empty string")
}

func TestNewSessionConfigurationResponse_MissingRequestedChecks(t *testing.T) {
	jsonData := `{
		"client_session_token_ttl": 3600,
		"session_id": "abc123",
		"capture": {}
	}`

	_, err := NewSessionConfigurationResponse([]byte(jsonData))
	assert.ErrorContains(t, err, "requested_checks must be a non-empty array")
}

func TestSessionConfigurationResponse_Getters(t *testing.T) {
	resp := &SessionConfigurationResponse{
		ClientSessionTokenTtl: 900,
		SessionId:             "test-session",
		RequestedChecks:       []string{"FACE_CAPTURE"},
		Capture:               &CaptureResponse{},
	}

	assert.Equal(t, resp.GetClientSessionTokenTtl(), 900)
	assert.Equal(t, resp.GetSessionId(), "test-session")
	assert.DeepEqual(t, resp.GetRequestedChecks(), []string{"FACE_CAPTURE"})
	assert.Assert(t, resp.GetCapture() != nil)
}
