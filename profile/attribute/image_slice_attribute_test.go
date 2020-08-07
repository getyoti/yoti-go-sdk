package attribute

import (
	"testing"

	"github.com/getyoti/yoti-go-sdk/v3/media"
	"github.com/getyoti/yoti-go-sdk/v3/profile/attribute/anchor"
	"github.com/getyoti/yoti-go-sdk/v3/yotiprotoattr"
	"gotest.tools/v3/assert"
)

func assertIsExpectedImage(t *testing.T, image media.Media, imageMIMEType string, expectedBase64URLLast10 string) {
	assert.Equal(t, image.MIME(), imageMIMEType)

	actualBase64URL := image.Base64URL()

	ActualBase64URLLast10Chars := actualBase64URL[len(actualBase64URL)-10:]

	assert.Equal(t, ActualBase64URLLast10Chars, expectedBase64URLLast10)
}

func assertIsExpectedDocumentImagesAttribute(t *testing.T, actualDocumentImages []media.Media, anchor *anchor.Anchor) {

	assert.Equal(t, len(actualDocumentImages), 2, "This Document Images attribute should have two images")

	assertIsExpectedImage(t, actualDocumentImages[0], media.ImageTypeJPEG, "vWgD//2Q==")
	assertIsExpectedImage(t, actualDocumentImages[1], media.ImageTypeJPEG, "38TVEH/9k=")

	expectedValue := "NATIONAL_ID"
	assert.Equal(t, anchor.Value(), expectedValue)

	expectedSubType := "STATE_ID"
	assert.Equal(t, anchor.SubType(), expectedSubType)
}

func TestAttribute_NewImageSlice(t *testing.T) {
	protoAttribute := createAttributeFromTestFile(t, "../../test/fixtures/test_attribute_multivalue.txt")

	documentImagesAttribute, err := NewImageSlice(protoAttribute)

	assert.NilError(t, err)

	assertIsExpectedDocumentImagesAttribute(t, documentImagesAttribute.Value(), documentImagesAttribute.Anchors()[0])
}

func TestAttribute_ImageSliceNotCreatedWithNonMultiValueType(t *testing.T) {
	attributeName := "attributeName"
	attributeValueString := "value"
	attributeValue := []byte(attributeValueString)

	var attr = &yotiprotoattr.Attribute{
		Name:        attributeName,
		Value:       attributeValue,
		ContentType: yotiprotoattr.ContentType_STRING,
		Anchors:     []*yotiprotoattr.Anchor{},
	}

	_, err := NewImageSlice(attr)

	assert.Assert(t, err != nil, "Expected error when creating image slice from attribute which isn't of multi-value type")
}
