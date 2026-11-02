package create_test

import (
	"encoding/json"
	"testing"

	"github.com/getyoti/yoti-go-sdk/v3/docscan/session/create"
	"gotest.tools/v3/assert"
)

func TestSessionOverviewPdfPayload_AllOptionsSet(t *testing.T) {
	payload := create.NewSessionOverviewPdfPayload().
		WithIncludeImages(true).
		WithIncludeExtractedText(false).
		WithIncludeDeviceTracking(true)

	bytes, err := json.Marshal(payload)
	assert.NilError(t, err)
	assert.Equal(t, `{"include_images":true,"include_extracted_text":false,"include_device_tracking":true}`, string(bytes))
}

func TestSessionOverviewPdfPayload_PartialOptionsSet(t *testing.T) {
	payload := create.NewSessionOverviewPdfPayload().WithIncludeImages(true)

	bytes, err := json.Marshal(payload)
	assert.NilError(t, err)
	assert.Equal(t, `{"include_images":true}`, string(bytes))
}

func TestSessionOverviewPdfPayload_NoOptionsSet(t *testing.T) {
	payload := create.NewSessionOverviewPdfPayload()

	bytes, err := json.Marshal(payload)
	assert.NilError(t, err)
	assert.Equal(t, `{}`, string(bytes))
}
