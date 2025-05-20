package retrieve

import (
	"encoding/json"
	"testing"

	"gotest.tools/v3/assert"
)

func TestFaceCaptureResourceResponse_Getters(t *testing.T) {
	resp := &FaceCaptureResourceResponse{
		ID:     "face-resource-id",
		Frames: 3,
	}

	assert.Equal(t, resp.GetID(), "face-resource-id")
	assert.Equal(t, resp.GetFrames(), 3)
}

func TestFaceCaptureResourceResponse_UnmarshalJSON(t *testing.T) {
	jsonData := `{
		"id": "resource-123",
		"frames": 5
	}`

	var resp FaceCaptureResourceResponse
	err := json.Unmarshal([]byte(jsonData), &resp)
	assert.NilError(t, err)

	assert.Equal(t, resp.ID, "resource-123")
	assert.Equal(t, resp.Frames, 5)
}
