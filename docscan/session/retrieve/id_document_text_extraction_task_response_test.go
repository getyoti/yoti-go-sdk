package retrieve

import (
	"testing"

	"github.com/getyoti/yoti-go-sdk/v3/docscan/constants"
	"gotest.tools/v3/assert"
)

func TestTextExtractionTaskResponse_GeneratedTextDataChecks(t *testing.T) {
	var checks []*GeneratedTextDataCheckResponse
	checks = append(
		checks,
		&GeneratedTextDataCheckResponse{
			&GeneratedCheckResponse{
				Type: constants.IDDocumentTextDataCheck,
				ID:   "some-id",
			},
		},
	)

	taskResponse := &TextExtractionTaskResponse{
		TaskResponse: &TaskResponse{
			generatedTextDataChecks: checks,
		},
	}

	assert.Equal(t, 1, len(taskResponse.GeneratedTextDataChecks()))
	assert.Equal(t, "some-id", taskResponse.GeneratedTextDataChecks()[0].ID)
}
