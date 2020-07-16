package profile

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestYotiClient_ParseIsAgeVerifiedValue_True(t *testing.T) {
	trueValue := []byte("true")

	isAgeVerified, err := parseIsAgeVerifiedValue(trueValue)

	assert.Assert(t, is.Nil(err), "Failed to parse IsAgeVerified value")
	assert.Check(t, *isAgeVerified)
}

func TestYotiClient_ParseIsAgeVerifiedValue_False(t *testing.T) {
	falseValue := []byte("false")

	isAgeVerified, err := parseIsAgeVerifiedValue(falseValue)

	assert.Assert(t, is.Nil(err), "Failed to parse IsAgeVerified value")
	assert.Check(t, !*isAgeVerified)

}
func TestYotiClient_ParseIsAgeVerifiedValue_InvalidValueThrowsError(t *testing.T) {
	invalidValue := []byte("invalidBool")

	_, err := parseIsAgeVerifiedValue(invalidValue)

	assert.Assert(t, err != nil)
}
