package attribute

import (
	"testing"

	"github.com/getyoti/yoti-go-sdk/v3/yotiprotoattr"
	"github.com/golang/protobuf/proto"
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func marshallMultiValue(t *testing.T, multiValue *yotiprotoattr.MultiValue) []byte {
	marshalled, err := proto.Marshal(multiValue)

	assert.NilError(t, err)

	return marshalled
}

func createMultiValueAttribute(t *testing.T, multiValueItemSlice []*yotiprotoattr.MultiValue_Value) (*MultiValueAttribute, error) {
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

	return NewMultiValue(protoAttribute)
}

func TestAttribute_MultiValueNotCreatedWithNonMultiValueType(t *testing.T) {
	attributeName := "attributeName"
	attributeValueString := "value"
	attributeValue := []byte(attributeValueString)

	var attr = &yotiprotoattr.Attribute{
		Name:        attributeName,
		Value:       attributeValue,
		ContentType: yotiprotoattr.ContentType_STRING,
		Anchors:     []*yotiprotoattr.Anchor{},
	}

	_, err := NewMultiValue(attr)

	assert.Assert(t, err != nil, "Expected error when creating multi value from attribute which isn't of multi-value type")
}

func TestAttribute_NewMultiValue(t *testing.T) {
	protoAttribute := createAttributeFromTestFile(t, "../../test/fixtures/test_attribute_multivalue.txt")

	multiValueAttribute, err := NewMultiValue(protoAttribute)

	assert.NilError(t, err)

	var documentImagesAttributeItems = CreateImageSlice(multiValueAttribute.Value())

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

	multiValueAttr, err := NewMultiValue(protoAttribute)
	assert.Check(t, err != nil)

	assert.Assert(t, is.Nil(multiValueAttr))
}

func TestAttribute_NestedMultiValue(t *testing.T) {
	var innerMultiValueProtoValue = createAttributeFromTestFile(t, "../../test/fixtures/test_attribute_multivalue.txt").Value

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

	assert.NilError(t, err)

	for key, value := range multiValueAttribute.Value() {
		switch key {
		case 0:
			value0 := value.GetValue()

			assert.Equal(t, value0.(string), "string")
		case 1:
			value1 := value.GetValue()

			innerItems, ok := value1.([]*Item)
			assert.Assert(t, ok)

			for innerKey, item := range innerItems {
				switch innerKey {
				case 0:
					assertIsExpectedImage(t, item.GetValue().(*Image), "jpeg", "vWgD//2Q==")

				case 1:
					assertIsExpectedImage(t, item.GetValue().(*Image), "jpeg", "38TVEH/9k=")
				}
			}
		}
	}
}

func TestAttribute_MultiValueGenericGetter(t *testing.T) {
	protoAttribute := createAttributeFromTestFile(t, "../../test/fixtures/test_attribute_multivalue.txt")
	multiValueAttribute, err := NewMultiValue(protoAttribute)
	assert.NilError(t, err)

	// We need to cast, since GetAttribute always returns generic attributes
	multiValueAttributeValue := multiValueAttribute.Value()
	imageSlice := CreateImageSlice(multiValueAttributeValue)

	assertIsExpectedDocumentImagesAttribute(t, imageSlice, multiValueAttribute.Anchors()[0])
}
