package media

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestGeneric_Base64URL_create(t *testing.T) {
	databytes := []byte("value")
	mime := "foo/bar"

	result := NewGeneric(mime, databytes)

	expectedDataUrl := "data:" + mime + ";base64," + imageBase64Value

	assert.Equal(t, expectedDataUrl, result.Base64URL())
}

func TestNewMedia(t *testing.T) {
	databytes := []byte("value")

	v := NewMedia(ImageTypeJPEG, databytes)
	_, ok := v.(JPEGImage)
	assert.Assert(t, ok)

	v = NewMedia(ImageTypePNG, databytes)
	_, ok = v.(PNGImage)
	assert.Assert(t, ok)

	v = NewMedia("foo/bar", databytes)
	_, ok = v.(Generic)
	assert.Assert(t, ok)
}
