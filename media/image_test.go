package media

import (
	"encoding/base64"
	"testing"

	"gotest.tools/v3/assert"
)

func TestImage_Base64Selfie_Png(t *testing.T) {
	imageBase64Value, media := createImage(ImageTypePng)
	expectedDataUrl := "data:image/png;base64," + imageBase64Value

	assert.Equal(t, expectedDataUrl, media.Base64URL())
}

func TestImage_Base64Selfie_Jpeg(t *testing.T) {
	imageBase64Value, media := createImage(ImageTypeJpeg)
	expectedDataUrl := "data:image/jpeg;base64," + imageBase64Value

	assert.Equal(t, expectedDataUrl, media.Base64URL())
}

func createImage(MIMEType string) (string, *Image) {
	imageBytes := []byte("value")
	imageBase64Value := base64.StdEncoding.EncodeToString(imageBytes)
	media := NewImage(MIMEType, imageBytes)

	return imageBase64Value, media
}
