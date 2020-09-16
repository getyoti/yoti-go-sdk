package retrieve

import (
	"encoding/json"
	"testing"

	"github.com/getyoti/yoti-go-sdk/v3/docscan/constants"
	"gotest.tools/v3/assert"
)

func TestTaskResponse_UnmarshalJSON(t *testing.T) {
	generatedTextDataCheck := GeneratedCheckResponse{
		Type: constants.IDDocumentTextDataCheck,
		ID:   "some-id",
	}

	var checks []*GeneratedCheckResponse
	checks = append(checks, &GeneratedCheckResponse{Type: "OTHER_TYPE", ID: "other-id"})
	checks = append(checks, &generatedTextDataCheck)

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
}

func TestTaskResponse_UnmarshalJSON_Invalid(t *testing.T) {
	var result TaskResponse
	err := result.UnmarshalJSON([]byte("some-invalid-json"))
	assert.ErrorContains(t, err, "invalid character")
}
