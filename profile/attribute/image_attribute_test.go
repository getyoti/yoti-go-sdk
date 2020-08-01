package attribute

import (
	"encoding/base64"
	"testing"

	"github.com/getyoti/yoti-go-sdk/v3/consts"
	"github.com/getyoti/yoti-go-sdk/v3/yotiprotoattr"
	"gotest.tools/v3/assert"
)

func TestImageAttribute_Image_Png(t *testing.T) {
	attributeName := consts.AttrSelfie
	byteValue := []byte("value")

	var attributeImage = &yotiprotoattr.Attribute{
		Name:        attributeName,
		Value:       byteValue,
		ContentType: yotiprotoattr.ContentType_PNG,
		Anchors:     []*yotiprotoattr.Anchor{},
	}

	selfie, err := NewImage(attributeImage)
	assert.NilError(t, err)

	assert.DeepEqual(t, selfie.Value().Data, byteValue)
}

func TestImageAttribute_Image_Jpeg(t *testing.T) {
	attributeName := consts.AttrSelfie
	byteValue := []byte("value")

	var attributeImage = &yotiprotoattr.Attribute{
		Name:        attributeName,
		Value:       byteValue,
		ContentType: yotiprotoattr.ContentType_JPEG,
		Anchors:     []*yotiprotoattr.Anchor{},
	}

	selfie, err := NewImage(attributeImage)
	assert.NilError(t, err)

	assert.DeepEqual(t, selfie.Value().Data, byteValue)
}

func TestImageAttribute_Image_Default(t *testing.T) {
	attributeName := consts.AttrSelfie
	byteValue := []byte("value")

	var attributeImage = &yotiprotoattr.Attribute{
		Name:        attributeName,
		Value:       byteValue,
		ContentType: yotiprotoattr.ContentType_PNG,
		Anchors:     []*yotiprotoattr.Anchor{},
	}
	selfie, err := NewImage(attributeImage)
	assert.NilError(t, err)

	assert.DeepEqual(t, selfie.Value().Data, byteValue)
}

func TestImageAttribute_Base64Selfie_Png(t *testing.T) {
	attributeName := consts.AttrSelfie
	imageBytes := []byte("value")

	var attributeImage = &yotiprotoattr.Attribute{
		Name:        attributeName,
		Value:       imageBytes,
		ContentType: yotiprotoattr.ContentType_PNG,
		Anchors:     []*yotiprotoattr.Anchor{},
	}

	selfie, err := NewImage(attributeImage)
	assert.NilError(t, err)

	base64ImageExpectedValue := base64.StdEncoding.EncodeToString(imageBytes)

	expectedBase64Selfie := "data:image/png;base64," + base64ImageExpectedValue

	base64Selfie := selfie.Value().Base64URL()

	assert.Equal(t, base64Selfie, expectedBase64Selfie)
}

func TestImageAttribute_Base64URL_Jpeg(t *testing.T) {
	attributeName := consts.AttrSelfie
	imageBytes := []byte("value")

	var attributeImage = &yotiprotoattr.Attribute{
		Name:        attributeName,
		Value:       imageBytes,
		ContentType: yotiprotoattr.ContentType_JPEG,
		Anchors:     []*yotiprotoattr.Anchor{},
	}

	selfie, err := NewImage(attributeImage)
	assert.NilError(t, err)

	base64ImageExpectedValue := base64.StdEncoding.EncodeToString(imageBytes)

	expectedBase64Selfie := "data:image/jpeg;base64," + base64ImageExpectedValue

	base64Selfie := selfie.Value().Base64URL()

	assert.Equal(t, base64Selfie, expectedBase64Selfie)
}
