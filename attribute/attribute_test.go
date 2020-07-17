package attribute

import (
	"testing"
	"time"

	"gotest.tools/v3/assert"
)

func TestNewThirdPartyAttribute(t *testing.T) {
	protoAttribute := createAttributeFromTestFile(t, "../test/fixtures/test_attribute_third_party.txt")

	stringAttribute := NewString(protoAttribute)

	assert.Equal(t, stringAttribute.Value(), "test-third-party-attribute-0")
	assert.Equal(t, stringAttribute.Name(), "com.thirdparty.id")

	assert.Equal(t, stringAttribute.Sources()[0].Value(), "THIRD_PARTY")
	assert.Equal(t, stringAttribute.Sources()[0].SubType(), "orgName")

	assert.Equal(t, stringAttribute.Verifiers()[0].Value(), "THIRD_PARTY")
	assert.Equal(t, stringAttribute.Verifiers()[0].SubType(), "orgName")
}

func TestAttribute_DateOfBirth(t *testing.T) {
	protoAttribute := createAttributeFromTestFile(t, "../test/fixtures/test_attribute_date_of_birth.txt")

	dateOfBirthAttribute, err := NewDate(protoAttribute)

	assert.NilError(t, err)

	expectedDateOfBirth := time.Date(1970, time.December, 01, 0, 0, 0, 0, time.UTC)
	actualDateOfBirth := dateOfBirthAttribute.Value()

	assert.Assert(t, actualDateOfBirth.Equal(expectedDateOfBirth))
}
