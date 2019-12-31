package attribute

import (
	"testing"

	"gotest.tools/assert"
)

func TestStringAttribute_NewThirdPartyAttribute(t *testing.T) {
	protoAttribute := createAttributeFromTestFile(t, "../test/fixtures/test_attribute_third_party.txt")

	stringAttribute := NewString(protoAttribute)

	assert.Equal(t, stringAttribute.Value(), "test-third-party-attribute-0")
	assert.Equal(t, stringAttribute.Name(), "com.thirdparty.id")

	assert.Equal(t, stringAttribute.Sources()[0].Value()[0], "THIRD_PARTY")
	assert.Equal(t, stringAttribute.Sources()[0].SubType(), "orgName")

	assert.Equal(t, stringAttribute.Verifiers()[0].Value()[0], "THIRD_PARTY")
	assert.Equal(t, stringAttribute.Verifiers()[0].SubType(), "orgName")
}
