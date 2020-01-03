package attribute

import (
	"testing"

	"github.com/getyoti/yoti-go-sdk/v2/anchor"
	"github.com/getyoti/yoti-go-sdk/v2/yotiprotoattr"
	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

func assertIsExpectedImage(t *testing.T, image *Image, imageType string, expectedBase64URLLast10 string) {
	assert.Equal(t, image.Type, imageType)

	actualBase64URL := image.Base64URL()

	ActualBase64URLLast10Chars := actualBase64URL[len(actualBase64URL)-10:]

	assert.Equal(t, ActualBase64URLLast10Chars, expectedBase64URLLast10)
}

func assertIsExpectedDocumentImagesAttribute(t *testing.T, actualDocumentImages []*Image, anchor *anchor.Anchor) {

	assert.Equal(t, len(actualDocumentImages), 2, "This Document Images attribute should have two images")

	assertIsExpectedImage(t, actualDocumentImages[0], "jpeg", "vWgD//2Q==")
	assertIsExpectedImage(t, actualDocumentImages[1], "jpeg", "38TVEH/9k=")

	expectedValue := "NATIONAL_ID"
	assert.Equal(t, anchor.Value()[0], expectedValue)

	expectedSubType := "STATE_ID"
	assert.Equal(t, anchor.SubType(), expectedSubType)
}

func TestAttribute_NewImageSlice(t *testing.T) {
	protoAttribute := createAttributeFromTestFile(t, "../test/fixtures/test_attribute_multivalue.txt")

	documentImagesAttribute, err := NewImageSlice(protoAttribute)

	assert.Assert(t, is.Nil(err))

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
