package attribute

import (
	"testing"

	"github.com/getyoti/yoti-go-sdk/v3/test"
	"github.com/getyoti/yoti-go-sdk/v3/yotiprotoattr"
	"github.com/golang/protobuf/proto"
	"gotest.tools/v3/assert"
)

func createAttributeFromTestFile(t *testing.T, filename string) *yotiprotoattr.Attribute {
	attributeBytes := test.DecodeTestFile(t, filename)

	attributeStruct := &yotiprotoattr.Attribute{}

	err2 := proto.Unmarshal(attributeBytes, attributeStruct)
	assert.NilError(t, err2)

	return attributeStruct
}
