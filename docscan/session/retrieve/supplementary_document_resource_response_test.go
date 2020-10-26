package retrieve

import (
	"encoding/json"
	"testing"

	"gotest.tools/v3/assert"
)

func TestSupplementaryDocumentResourceResponse_UnmarshalJSON(t *testing.T) {
	idDocumentResource := &SupplementaryDocumentResourceResponse{
		ResourceResponse: &ResourceResponse{
			ID: "",
			Tasks: []*TaskResponse{
				{
					ID:   "some-id",
					Type: "SUPPLEMENTARY_DOCUMENT_TEXT_DATA_EXTRACTION",
				},
				{
					ID:   "other-id",
					Type: "OTHER_TASK_TYPE",
				},
			},
		},
	}

	marshalledResource, err := json.Marshal(idDocumentResource)
	assert.NilError(t, err)

	var result SupplementaryDocumentResourceResponse
	err = json.Unmarshal(marshalledResource, &result)
	assert.NilError(t, err)

	assert.Equal(t, 2, len(result.ResourceResponse.Tasks))

	assert.Equal(t, 1, len(result.TextExtractionTasks()))
	assert.Equal(t, "some-id", result.TextExtractionTasks()[0].ID)
}

func TestSupplementaryDocumentResourceResponse_UnmarshalJSON_Invalid(t *testing.T) {
	var result SupplementaryDocumentResourceResponse
	err := result.UnmarshalJSON([]byte("some-invalid-json"))
	assert.ErrorContains(t, err, "invalid character")
}
