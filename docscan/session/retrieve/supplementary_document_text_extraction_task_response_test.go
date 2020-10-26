package retrieve

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestSupplementaryDocumentTextExtractionTaskResponse_GeneratedTextDataChecks(t *testing.T) {
	var checks []*GeneratedSupplementaryTextDataCheckResponse
	checks = append(
		checks,
		&GeneratedSupplementaryTextDataCheckResponse{
			&GeneratedCheckResponse{
				ID: "some-id",
			},
		},
	)

	taskResponse := &SupplementaryDocumentTextExtractionTaskResponse{
		TaskResponse: &TaskResponse{
			generatedSupplementaryTextDataChecks: checks,
		},
	}

	assert.Equal(t, 1, len(taskResponse.GeneratedTextDataChecks()))
	assert.Equal(t, "some-id", taskResponse.GeneratedTextDataChecks()[0].ID)
}
