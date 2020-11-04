package retrieve

import (
	"encoding/json"
	"testing"

	"github.com/getyoti/yoti-go-sdk/v3/docscan/constants"
	"gotest.tools/v3/assert"
)

func TestTaskResponse_UnmarshalJSON(t *testing.T) {
	checks := []*GeneratedCheckResponse{
		{
			Type: constants.IDDocumentTextDataCheck,
			ID:   "some-id",
		},
		{
			Type: constants.SupplementaryDocumentTextDataCheck,
			ID:   "supplementary-id",
		},
		{
			Type: "OTHER_TYPE",
			ID:   "other-id",
		},
	}

	taskResponse := TaskResponse{
		GeneratedChecks: checks,
	}
	marshalled, err := json.Marshal(&taskResponse)
	assert.NilError(t, err)

	var result TaskResponse
	err = json.Unmarshal(marshalled, &result)
	assert.NilError(t, err)

	assert.Equal(t, 1, len(result.GeneratedTextDataChecks()))
	assert.Equal(t, "some-id", result.GeneratedTextDataChecks()[0].ID)

	assert.Equal(t, 1, len(result.generatedTextDataChecks))
	assert.Equal(t, "some-id", result.generatedTextDataChecks[0].ID)

	assert.Equal(t, 1, len(result.generatedSupplementaryTextDataChecks))
	assert.Equal(t, "supplementary-id", result.generatedSupplementaryTextDataChecks[0].ID)
}

func TestTaskResponse_UnmarshalJSON_Invalid(t *testing.T) {
	var result TaskResponse
	err := result.UnmarshalJSON([]byte("some-invalid-json"))
	assert.ErrorContains(t, err, "invalid character")
}
