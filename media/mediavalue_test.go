package media

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestGeneric_Base64URL_create(t *testing.T) {
	dataBytes := []byte("value")
	mime := "foo/bar"

	result := NewGeneric(mime, dataBytes)

	expectedDataURL := "data:" + mime + ";base64," + imageBase64Value

	assert.Equal(t, expectedDataURL, result.Base64URL())
	assert.DeepEqual(t, dataBytes, result.Data())
}

func TestNewMedia(t *testing.T) {
	dataBytes := []byte("value")

	v := NewMedia(ImageTypeJPEG, dataBytes)
	_, ok := v.(JPEGImage)
	assert.Assert(t, ok)

	v = NewMedia(ImageTypePNG, dataBytes)
	_, ok = v.(PNGImage)
	assert.Assert(t, ok)

	v = NewMedia("foo/bar", dataBytes)
	_, ok = v.(Generic)
	assert.Assert(t, ok)
}
