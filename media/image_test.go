package media

import (
	"testing"

	"gotest.tools/v3/assert"
)

const (
	imageBase64Value = "dmFsdWU="
)

func TestImage_Base64URL_CreateJpegImage(t *testing.T) {
	imageBytes := []byte("value")

	result := JPEGImage(imageBytes)
	expectedDataURL := "data:image/jpeg;base64," + imageBase64Value

	assert.Equal(t, expectedDataURL, result.Base64URL())
}

func TestImage_Base64URL_CreatePngImage(t *testing.T) {
	imageBytes := []byte("value")

	result := PNGImage(imageBytes)
	expectedDataURL := "data:image/png;base64," + imageBase64Value

	assert.Equal(t, expectedDataURL, result.Base64URL())
}
