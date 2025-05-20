package facecapture

import (
	"strings"
	"testing"

	"gotest.tools/v3/assert"
)

func TestNewUploadFaceCaptureImagePayload(t *testing.T) {
	contentType := "image/png"
	content := []byte{1, 2, 3}

	payload := NewUploadFaceCaptureImagePayload(contentType, content)

	assert.Equal(t, payload.ImageContentType, contentType)
	assert.DeepEqual(t, payload.ImageContents, content)
}

func TestPrepare_ValidPayload(t *testing.T) {
	contentType := "image/jpeg"
	content := []byte{0x01, 0x02, 0x03}

	payload := NewUploadFaceCaptureImagePayload(contentType, content)
	err := payload.Prepare()

	assert.NilError(t, err)
	assert.Assert(t, payload.body != nil)
	assert.Assert(t, payload.writer != nil)

	// Multipart body should contain the image bytes
	bodyStr := payload.body.String()
	assert.Assert(t, strings.Contains(bodyStr, "Content-Disposition"))
	assert.Assert(t, strings.Contains(bodyStr, contentType))
	assert.Assert(t, strings.Contains(bodyStr, string(content)))
}

func TestPrepare_EmptyContentType(t *testing.T) {
	payload := NewUploadFaceCaptureImagePayload("", []byte{1, 2, 3})
	err := payload.Prepare()

	assert.ErrorContains(t, err, "ImageContentType must not be empty")
}

func TestPrepare_EmptyContents(t *testing.T) {
	payload := NewUploadFaceCaptureImagePayload("image/png", []byte{})
	err := payload.Prepare()

	assert.ErrorContains(t, err, "ImageContents must not be empty")
}

func TestHeaders_ReturnsContentType(t *testing.T) {
	payload := NewUploadFaceCaptureImagePayload("image/png", []byte{1, 2, 3})
	err := payload.Prepare()
	assert.NilError(t, err)

	headers := payload.Headers()
	contentTypes, ok := headers["Content-Type"]
	assert.Assert(t, ok)
	assert.Assert(t, len(contentTypes) > 0)
	assert.Assert(t, strings.HasPrefix(contentTypes[0], "multipart/form-data; boundary="))
}

func TestMultipartFormBody_ReturnsBody(t *testing.T) {
	payload := NewUploadFaceCaptureImagePayload("image/png", []byte{1, 2, 3})
	err := payload.Prepare()
	assert.NilError(t, err)

	body := payload.MultipartFormBody()
	assert.Assert(t, body != nil)
	assert.Assert(t, body.Len() > 0)
}
