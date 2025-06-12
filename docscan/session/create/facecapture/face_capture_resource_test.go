package facecapture

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestNewCreateFaceCaptureResourcePayload(t *testing.T) {
	requirementID := "test-requirement-id"
	payload := NewCreateFaceCaptureResourcePayload(requirementID)

	assert.Assert(t, payload != nil)
	assert.Equal(t, payload.RequirementID, requirementID)
}
