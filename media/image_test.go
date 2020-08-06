package media

import (
	"testing"

	"gotest.tools/v3/assert"
)

const (
	imageBase64Value = "dmFsdWU="
)

func TestImage_Base64Selfie_CreateImage(t *testing.T) {
	imageBytes := []byte("value")

	result := NewImage(ImageTypePng, imageBytes)
	expectedDataUrl := "data:image/png;base64," + imageBase64Value

	assert.Equal(t, expectedDataUrl, result.Base64URL())
}

func TestImage_Base64Selfie_CreateJpegImage(t *testing.T) {
	imageBytes := []byte("value")

	result := NewImage(ImageTypeJpeg, imageBytes)
	expectedDataUrl := "data:image/jpeg;base64," + imageBase64Value

	assert.Equal(t, expectedDataUrl, result.Base64URL())
}

func TestImage_Base64Selfie_CreatePngImage(t *testing.T) {
	imageBytes := []byte("value")

	result := NewImage(ImageTypeJpeg, imageBytes)
	expectedDataUrl := "data:image/jpeg;base64," + imageBase64Value

	assert.Equal(t, expectedDataUrl, result.Base64URL())
}
