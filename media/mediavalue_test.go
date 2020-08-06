package media

import (
	"encoding/base64"
	"testing"

	"gotest.tools/v3/assert"
)

func TestMedia_Base64URL_PNG(t *testing.T) {
	mediaBase64Value, media := createMedia(ImageTypePNG)
	expectedDataUrl := "data:image/png;base64," + mediaBase64Value

	assert.Equal(t, expectedDataUrl, media.Base64URL())
}

func TestMedia_Base64URL_JPEG(t *testing.T) {
	mediaBase64Value, media := createMedia(ImageTypeJPEG)
	expectedDataUrl := "data:image/jpeg;base64," + mediaBase64Value

	assert.Equal(t, expectedDataUrl, media.Base64URL())
}

func TestMedia_Base64URL_PDF(t *testing.T) {
	mediaBase64Value, media := createMedia("application/pdf")
	expectedDataUrl := "data:application/pdf;base64," + mediaBase64Value

	assert.Equal(t, expectedDataUrl, media.Base64URL())
}

func createMedia(contentType string) (string, *Value) {
	imageBytes := []byte("value")
	mediaBase64Value := base64.StdEncoding.EncodeToString(imageBytes)

	media := &Value{
		MIMEType: contentType,
		Data:     imageBytes,
	}
	return mediaBase64Value, media
}
