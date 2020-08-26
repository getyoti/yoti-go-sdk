package extension

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/getyoti/yoti-go-sdk/v3/profile/attribute"
	"gotest.tools/v3/assert"
)

func createDefinitionByName(name string) attribute.Definition {
	return attribute.NewAttributeDefinition(name)
}

func ExampleThirdPartyAttributeExtension() {
	attributeDefinition := attribute.NewAttributeDefinition("some_value")

	datetime, err := time.Parse("2006-01-02T15:04:05.000Z", "2019-10-30T12:10:09.458Z")
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	extension, err := (&ThirdPartyAttributeExtensionBuilder{}).
		WithExpiryDate(&datetime).
		WithDefinition(attributeDefinition).
		Build()

	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, _ := extension.MarshalJSON()
	fmt.Println(string(data))
	// Output: {"type":"THIRD_PARTY_ATTRIBUTE","content":{"expiry_date":"2019-10-30T12:10:09.458Z","definitions":[{"name":"some_value"}]}}
}

func TestWithDefinitionShouldAddToList(t *testing.T) {
	datetime, err := time.Parse("2006-01-02T15:04:05.000Z", "2019-10-30T12:10:09.458Z")
	assert.NilError(t, err)

	definitionList := []attribute.Definition{
		createDefinitionByName("some_attribute"),
		createDefinitionByName("some_other_attribute"),
	}

	someOtherDefinition := createDefinitionByName("wanted_definition")

	extension, err := (&ThirdPartyAttributeExtensionBuilder{}).
		WithExpiryDate(&datetime).
		WithDefinitions(definitionList).
		WithDefinition(someOtherDefinition).
		Build()

	assert.NilError(t, err)
	assert.Equal(t, len(extension.definitions), 3)
	assert.Equal(t, extension.definitions[0].Name(), "some_attribute")
	assert.Equal(t, extension.definitions[1].Name(), "some_other_attribute")
	assert.Equal(t, extension.definitions[2].Name(), "wanted_definition")
}

func TestWithDefinitionsShouldOverwriteList(t *testing.T) {
	datetime, err := time.Parse("2006-01-02T15:04:05.000Z", "2019-10-30T12:10:09.458Z")
	assert.NilError(t, err)

	definitionList := []attribute.Definition{
		createDefinitionByName("some_attribute"),
		createDefinitionByName("some_other_attribute"),
	}

	someOtherDefinition := createDefinitionByName("wanted_definition")

	extension, err := (&ThirdPartyAttributeExtensionBuilder{}).
		WithExpiryDate(&datetime).
		WithDefinition(someOtherDefinition).
		WithDefinitions(definitionList).
		Build()

	assert.NilError(t, err)
	assert.Equal(t, len(extension.definitions), 2)
	assert.Equal(t, extension.definitions[0].Name(), "some_attribute")
	assert.Equal(t, extension.definitions[1].Name(), "some_other_attribute")
}

var expiryDates = []struct {
	in       time.Time
	expected string
}{
	{time.Date(2051, 01, 13, 19, 50, 53, 1, time.UTC), "2051-01-13T19:50:53.000Z"},
	{time.Date(2026, 02, 02, 22, 04, 05, 123, time.UTC), "2026-02-02T22:04:05.000Z"},
	{time.Date(2051, 03, 13, 19, 50, 53, 9999, time.UTC), "2051-03-13T19:50:53.000Z"},
	{time.Date(2051, 04, 13, 19, 50, 53, 999999, time.UTC), "2051-04-13T19:50:53.000Z"},
	{time.Date(2026, 01, 31, 22, 04, 05, 1232567, time.UTC), "2026-01-31T22:04:05.001Z"},
	{time.Date(2026, 01, 31, 22, 04, 05, 17777777, time.UTC), "2026-01-31T22:04:05.017Z"},
	{time.Date(2026, 07, 31, 22, 04, 05, 000777777, time.UTC), "2026-07-31T22:04:05.000Z"},
	{time.Date(2026, 01, 02, 22, 04, 05, 123456789, time.UTC), "2026-01-02T22:04:05.123Z"},
	{time.Date(2028, 10, 02, 10, 00, 00, 0, time.FixedZone("UTC-5", -5*60*60)), "2028-10-02T15:00:00.000Z"},
	{time.Date(2028, 10, 02, 10, 00, 00, 0, time.FixedZone("UTC+11", 11*60*60)), "2028-10-01T23:00:00.000Z"},
	{time.Unix(1734567899, 0), "2024-12-19T00:24:59.000Z"},
	{time.Unix(2234567891, 0), "2040-10-23T01:18:11.000Z"},
}

func TestExpiryDatesAreFormattedCorrectly(t *testing.T) {
	attributeDefinition := attribute.NewAttributeDefinition("some_value")

	for _, date := range expiryDates {
		extension, err := (&ThirdPartyAttributeExtensionBuilder{}).
			WithExpiryDate(&date.in).
			WithDefinition(attributeDefinition).
			Build()

		assert.NilError(t, err)

		marshalledJson, _ := extension.MarshalJSON()

		attributeIssuanceDetailsJson := unmarshalJSONIntoMap(t, marshalledJson)

		content := attributeIssuanceDetailsJson["content"].(map[string]interface{})
		result := content["expiry_date"]

		if result != date.expected {
			t.Errorf("got %q, want %q", result, date.expected)
		}
	}
}

func unmarshalJSONIntoMap(t *testing.T, byteValue []byte) (result map[string]interface{}) {
	var unmarshalled interface{}
	err := json.Unmarshal(byteValue, &unmarshalled)
	assert.NilError(t, err)

	return unmarshalled.(map[string]interface{})
}
