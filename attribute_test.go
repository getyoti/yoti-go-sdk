package yoti

import (
	"encoding/base64"
	"testing"
	"time"

	"github.com/getyoti/yoti-go-sdk/v2/anchor"
	"github.com/getyoti/yoti-go-sdk/v2/attribute"
	"github.com/getyoti/yoti-go-sdk/v2/yotiprotoattr"
	"github.com/golang/protobuf/proto"
	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

func assertIsExpectedDocumentImagesAttribute(t *testing.T, actualDocumentImages []*attribute.Image, anchor *anchor.Anchor) {

	assert.Equal(t, len(actualDocumentImages), 2, "This Document Images attribute should have two images")

	assertIsExpectedImage(t, actualDocumentImages[0], "jpeg", "vWgD//2Q==")
	assertIsExpectedImage(t, actualDocumentImages[1], "jpeg", "38TVEH/9k=")

	expectedValue := "NATIONAL_ID"
	assert.Equal(t, anchor.Value()[0], expectedValue)

	expectedSubType := "STATE_ID"
	assert.Equal(t, anchor.SubType(), expectedSubType)
}

func assertIsExpectedImage(t *testing.T, image *attribute.Image, imageType string, expectedBase64URLLast10 string) {
	assert.Equal(t, image.Type, imageType)

	actualBase64URL := image.Base64URL()

	ActualBase64URLLast10Chars := actualBase64URL[len(actualBase64URL)-10:]

	assert.Equal(t, ActualBase64URLLast10Chars, expectedBase64URLLast10)
}

func marshallMultiValue(t *testing.T, multiValue *yotiprotoattr.MultiValue) []byte {
	marshalled, err := proto.Marshal(multiValue)

	assert.Assert(t, is.Nil(err))

	return marshalled
}

func createMultiValueAttribute(t *testing.T, multiValueItemSlice []*yotiprotoattr.MultiValue_Value) (*attribute.MultiValueAttribute, error) {
	var multiValueStruct = &yotiprotoattr.MultiValue{
		Values: multiValueItemSlice,
	}

	var marshalledMultiValueData = marshallMultiValue(t, multiValueStruct)
	attributeName := "nestedMultiValue"

	var protoAttribute = &yotiprotoattr.Attribute{
		Name:        attributeName,
		Value:       marshalledMultiValueData,
		ContentType: yotiprotoattr.ContentType_MULTI_VALUE,
		Anchors:     []*yotiprotoattr.Anchor{},
	}

	return attribute.NewMultiValue(protoAttribute)
}

func TestAttributeImage_Image_Png(t *testing.T) {
	attributeName := AttrConstSelfie
	byteValue := []byte("value")

	var attributeImage = &yotiprotoattr.Attribute{
		Name:        attributeName,
		Value:       byteValue,
		ContentType: yotiprotoattr.ContentType_PNG,
		Anchors:     []*yotiprotoattr.Anchor{},
	}

	result := createProfileWithSingleAttribute(attributeImage)
	selfie := result.Selfie()

	assert.DeepEqual(t, selfie.Value().Data, byteValue)
}

func TestAttributeImage_Image_Jpeg(t *testing.T) {
	attributeName := AttrConstSelfie
	byteValue := []byte("value")

	var attributeImage = &yotiprotoattr.Attribute{
		Name:        attributeName,
		Value:       byteValue,
		ContentType: yotiprotoattr.ContentType_JPEG,
		Anchors:     []*yotiprotoattr.Anchor{},
	}

	result := createProfileWithSingleAttribute(attributeImage)
	selfie := result.Selfie()

	assert.DeepEqual(t, selfie.Value().Data, byteValue)
}

func TestAttributeImage_Image_Default(t *testing.T) {
	attributeName := AttrConstSelfie
	byteValue := []byte("value")

	var attributeImage = &yotiprotoattr.Attribute{
		Name:        attributeName,
		Value:       byteValue,
		ContentType: yotiprotoattr.ContentType_PNG,
		Anchors:     []*yotiprotoattr.Anchor{},
	}
	result := createProfileWithSingleAttribute(attributeImage)
	selfie := result.Selfie()

	assert.DeepEqual(t, selfie.Value().Data, byteValue)
}
func TestAttributeImage_Base64Selfie_Png(t *testing.T) {
	attributeName := AttrConstSelfie
	imageBytes := []byte("value")

	var attributeImage = &yotiprotoattr.Attribute{
		Name:        attributeName,
		Value:       imageBytes,
		ContentType: yotiprotoattr.ContentType_PNG,
		Anchors:     []*yotiprotoattr.Anchor{},
	}

	result := createProfileWithSingleAttribute(attributeImage)

	base64ImageExpectedValue := base64.StdEncoding.EncodeToString(imageBytes)

	expectedBase64Selfie := "data:image/png;base64;," + base64ImageExpectedValue

	base64Selfie := result.Selfie().Value().Base64URL()

	assert.Equal(t, base64Selfie, expectedBase64Selfie)
}

func TestAttributeImage_Base64URL_Jpeg(t *testing.T) {
	attributeName := AttrConstSelfie
	imageBytes := []byte("value")

	var attributeImage = &yotiprotoattr.Attribute{
		Name:        attributeName,
		Value:       imageBytes,
		ContentType: yotiprotoattr.ContentType_JPEG,
		Anchors:     []*yotiprotoattr.Anchor{},
	}

	result := createProfileWithSingleAttribute(attributeImage)

	base64ImageExpectedValue := base64.StdEncoding.EncodeToString(imageBytes)

	expectedBase64Selfie := "data:image/jpeg;base64;," + base64ImageExpectedValue

	base64Selfie := result.Selfie().Value().Base64URL()

	assert.Equal(t, base64Selfie, expectedBase64Selfie)
}
func TestAttribute_DateOeafBirth(t *testing.T) {
	protoAttribute := createAttributeFromTestFile(t, "fixtures/test_attribute_date_of_birth.txt")

	dateOfBirthAttribute, err := attribute.NewTime(protoAttribute)

	assert.Assert(t, is.Nil(err))

	expectedDateOfBirth := time.Date(1970, time.December, 01, 0, 0, 0, 0, time.UTC)
	actualDateOfBirth := dateOfBirthAttribute.Value()

	assert.Assert(t, actualDateOfBirth.Equal(expectedDateOfBirth))
}

func TestAttribute_NewImageSlice(t *testing.T) {
	protoAttribute := createAttributeFromTestFile(t, "fixtures/test_attribute_multivalue.txt")

	documentImagesAttribute, err := attribute.NewImageSlice(protoAttribute)

	assert.Assert(t, is.Nil(err))

	assertIsExpectedDocumentImagesAttribute(t, documentImagesAttribute.Value(), documentImagesAttribute.Anchors()[0])
}

func TestAttribute_ImageSliceNotCreatedWithNonMultiValueType(t *testing.T) {
	attributeValueString := "value"
	attributeValue := []byte(attributeValueString)

	var attr = &yotiprotoattr.Attribute{
		Name:        attributeName,
		Value:       attributeValue,
		ContentType: yotiprotoattr.ContentType_STRING,
		Anchors:     []*yotiprotoattr.Anchor{},
	}

	_, err := attribute.NewImageSlice(attr)

	assert.Assert(t, err != nil, "Expected error when creating image slice from attribute which isn't of multi-value type")
}

func TestAttribute_MultiValueNotCreatedWithNonMultiValueType(t *testing.T) {
	attributeValueString := "value"
	attributeValue := []byte(attributeValueString)

	var attr = &yotiprotoattr.Attribute{
		Name:        attributeName,
		Value:       attributeValue,
		ContentType: yotiprotoattr.ContentType_STRING,
		Anchors:     []*yotiprotoattr.Anchor{},
	}

	_, err := attribute.NewMultiValue(attr)

	assert.Assert(t, err != nil, "Expected error when creating multi value from attribute which isn't of multi-value type")
}

func TestAttribute_NewMultiValue(t *testing.T) {
	protoAttribute := createAttributeFromTestFile(t, "fixtures/test_attribute_multivalue.txt")

	multiValueAttribute, err := attribute.NewMultiValue(protoAttribute)

	assert.Assert(t, is.Nil(err))

	var documentImagesAttributeItems []*attribute.Image = attribute.CreateImageSlice(multiValueAttribute.Value())

	assertIsExpectedDocumentImagesAttribute(t, documentImagesAttributeItems, multiValueAttribute.Anchors()[0])
}

func TestAttribute_InvalidMultiValueNotReturned(t *testing.T) {
	var invalidMultiValueItem = &yotiprotoattr.MultiValue_Value{
		ContentType: yotiprotoattr.ContentType_DATE,
		Data:        []byte("invalid"),
	}

	var stringMultiValueItem = &yotiprotoattr.MultiValue_Value{
		ContentType: yotiprotoattr.ContentType_STRING,
		Data:        []byte("string"),
	}

	var multiValueItemSlice = []*yotiprotoattr.MultiValue_Value{invalidMultiValueItem, stringMultiValueItem}

	var multiValueStruct = &yotiprotoattr.MultiValue{
		Values: multiValueItemSlice,
	}

	var marshalledMultiValueData = marshallMultiValue(t, multiValueStruct)
	attributeName := "nestedMultiValue"

	var protoAttribute = &yotiprotoattr.Attribute{
		Name:        attributeName,
		Value:       marshalledMultiValueData,
		ContentType: yotiprotoattr.ContentType_MULTI_VALUE,
		Anchors:     []*yotiprotoattr.Anchor{},
	}

	profile := createProfileWithSingleAttribute(protoAttribute)

	assert.Assert(t, is.Nil(profile.GetAttribute(attributeName)))
}

func TestNewGeneric_ShouldParseUnknownTypeAsString(t *testing.T) {
	value := []byte("value")
	parsed := attribute.NewGeneric(&yotiprotoattr.Attribute{
		ContentType: yotiprotoattr.ContentType_UNDEFINED,
		Value:       value,
	})

	stringValue, ok := parsed.Value().(string)
	assert.Check(t, ok)

	assert.Equal(t, stringValue, string(value))
}

func TestAttribute_NestedMultiValue(t *testing.T) {
	var innerMultiValueProtoValue []byte = createAttributeFromTestFile(t, "fixtures/test_attribute_multivalue.txt").Value

	var stringMultiValueItem = &yotiprotoattr.MultiValue_Value{
		ContentType: yotiprotoattr.ContentType_STRING,
		Data:        []byte("string"),
	}

	var multiValueItem = &yotiprotoattr.MultiValue_Value{
		ContentType: yotiprotoattr.ContentType_MULTI_VALUE,
		Data:        innerMultiValueProtoValue,
	}

	var multiValueItemSlice = []*yotiprotoattr.MultiValue_Value{stringMultiValueItem, multiValueItem}

	multiValueAttribute, err := createMultiValueAttribute(t, multiValueItemSlice)

	assert.Assert(t, is.Nil(err))

	for key, value := range multiValueAttribute.Value() {
		switch key {
		case 0:
			value0 := value.GetValue()

			assert.Equal(t, value0.(string), "string")
		case 1:
			value1 := value.GetValue()

			innerItems, ok := value1.([]*attribute.Item)
			assert.Assert(t, ok)

			for innerKey, item := range innerItems {
				switch innerKey {
				case 0:
					assertIsExpectedImage(t, item.GetValue().(*attribute.Image), "jpeg", "vWgD//2Q==")

				case 1:
					assertIsExpectedImage(t, item.GetValue().(*attribute.Image), "jpeg", "38TVEH/9k=")
				}
			}
		}
	}
}

func TestAttribute_MultiValueGenericGetter(t *testing.T) {
	protoAttribute := createAttributeFromTestFile(t, "fixtures/test_attribute_multivalue.txt")
	profile := createProfileWithSingleAttribute(protoAttribute)

	multiValueAttribute := profile.GetAttribute(AttrConstDocumentImages)

	// We need to cast, since GetAttribute always returns generic attributes
	multiValueAttributeValue := multiValueAttribute.Value().([]*attribute.Item)
	imageSlice := attribute.CreateImageSlice(multiValueAttributeValue)

	assertIsExpectedDocumentImagesAttribute(t, imageSlice, multiValueAttribute.Anchors()[0])
}
func TestNewThirdPartyAttribute(t *testing.T) {
	protoAttribute := createAttributeFromTestFile(t, "fixtures/test_attribute_third_party.txt")

	stringAttribute := attribute.NewString(protoAttribute)

	assert.Equal(t, stringAttribute.Value(), "test-third-party-attribute-0")
	assert.Equal(t, stringAttribute.Name(), "com.thirdparty.id")

	assert.Equal(t, stringAttribute.Sources()[0].Value()[0], "THIRD_PARTY")
	assert.Equal(t, stringAttribute.Sources()[0].SubType(), "orgName")

	assert.Equal(t, stringAttribute.Verifiers()[0].Value()[0], "THIRD_PARTY")
	assert.Equal(t, stringAttribute.Verifiers()[0].SubType(), "orgName")
}
