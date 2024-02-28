package attribute

import (
	"testing"

	"github.com/getyoti/yoti-go-sdk/v3/yotiprotoattr"
	"gotest.tools/v3/assert"
)

func TestNewGeneric_ShouldParseUnknownTypeAsString(t *testing.T) {
	value := []byte("value")
	protoAttr := yotiprotoattr.Attribute{
		ContentType: yotiprotoattr.ContentType_UNDEFINED,
		Value:       value,
	}
	parsed := NewGeneric(&protoAttr)

	stringValue, ok := parsed.Value().(string)
	assert.Check(t, ok)

	assert.Equal(t, stringValue, string(value))
}

func TestGeneric_ContentType(t *testing.T) {
	attribute := GenericAttribute{
		attributeDetails: attributeDetails{
			contentType: "contentType",
		},
	}

	assert.Equal(t, attribute.ContentType(), "contentType")
}

func TestNewGeneric_ShouldReturnNilForInvalidProtobuf(t *testing.T) {
	invalid := NewGeneric(&yotiprotoattr.Attribute{
		ContentType: yotiprotoattr.ContentType_JSON,
	})
	assert.Check(t, invalid == nil)
}
