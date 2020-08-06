package attribute

import (
	"fmt"
	"testing"

	"github.com/getyoti/yoti-go-sdk/v3/yotiprotoattr"
	"gotest.tools/v3/assert"
)

func ExampleNewJSON() {
	proto := yotiprotoattr.Attribute{
		Name:        "exampleJSON",
		Value:       []byte(`{"foo":"bar"}`),
		ContentType: yotiprotoattr.ContentType_JSON,
	}
	attribute, err := NewJSON(&proto)
	if err != nil {
		return
	}
	fmt.Println(attribute.Value())
	// Output: map[foo:bar]
}

func TestNewJSON_ShouldReturnNilForInvalidJSON(t *testing.T) {
	proto := yotiprotoattr.Attribute{
		Name:        "exampleJSON",
		Value:       []byte("Not a json document"),
		ContentType: yotiprotoattr.ContentType_JSON,
	}
	attribute, err := NewJSON(&proto)
	assert.Check(t, attribute == nil)
	assert.ErrorContains(t, err, "unable to parse JSON value")
}

func TestYotiClient_UnmarshallJSONValue_InvalidValueThrowsError(t *testing.T) {
	invalidStructuredAddress := []byte("invalidBool")

	_, err := unmarshallJSON(invalidStructuredAddress)

	assert.Assert(t, err != nil)
}

func TestYotiClient_UnmarshallJSONValue_ValidValue(t *testing.T) {
	const (
		countryIso  = "IND"
		nestedValue = "NestedValue"
	)

	var structuredAddress = []byte(`
	{
		"address_format": 2,
		"building": "House No.86-A",
		"state": "Punjab",
		"postal_code": "141012",
		"country_iso": "` + countryIso + `",
		"country": "India",
		"formatted_address": "House No.86-A\nRajgura Nagar\nLudhina\nPunjab\n141012\nIndia",
		"1":
		{
			"1-1":
			{
			  "1-1-1": "` + nestedValue + `"
			}
		}
	}
	`)

	parsedStructuredAddress, err := unmarshallJSON(structuredAddress)
	assert.NilError(t, err, "Failed to parse structured address")

	actualCountryIso := parsedStructuredAddress["country_iso"]

	assert.Equal(t, countryIso, actualCountryIso)
}
