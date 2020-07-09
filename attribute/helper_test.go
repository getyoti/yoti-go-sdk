package attribute

import (
	"testing"

	"github.com/getyoti/yoti-go-sdk/v2/test"
	"github.com/getyoti/yoti-go-sdk/v2/yotiprotoattr"
	"github.com/golang/protobuf/proto"
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func createAttributeFromTestFile(t *testing.T, filename string) *yotiprotoattr.Attribute {
	attributeBytes := test.DecodeTestFile(t, filename)

	attributeStruct := &yotiprotoattr.Attribute{}

	err2 := proto.Unmarshal(attributeBytes, attributeStruct)

	assert.Assert(t, is.Nil(err2))

	return attributeStruct
}
