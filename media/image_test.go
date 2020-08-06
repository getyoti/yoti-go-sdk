package media

import (
	"testing"

	"gotest.tools/v3/assert"
)

const (
	imageBase64Value = "dmFsdWU="
)

func TestImage_Base64URL_CreateImage(t *testing.T) {
	imageBytes := []byte("value")

	result := NewImage(ImageTypePNG, imageBytes)
	expectedDataUrl := "data:image/png;base64," + imageBase64Value

	assert.Equal(t, expectedDataUrl, result.Base64URL())
}

func TestImage_Base64URL_CreateJpegImage(t *testing.T) {
	imageBytes := []byte("value")

	result := NewJPEGImage(imageBytes)
	expectedDataUrl := "data:image/jpeg;base64," + imageBase64Value

	assert.Equal(t, expectedDataUrl, result.Base64URL())
}

func TestImage_Base64URL_CreatePngImage(t *testing.T) {
	imageBytes := []byte("value")

	result := NewPNGImage(imageBytes)
	expectedDataUrl := "data:image/png;base64," + imageBase64Value

	assert.Equal(t, expectedDataUrl, result.Base64URL())
}