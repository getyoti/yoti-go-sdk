package filter

import (
	"encoding/json"
	"fmt"
	"testing"

	"gotest.tools/v3/assert"
)

func TestDocumentFilterBuilder_WithCountryCode(t *testing.T) {
	filter, _ := NewDocumentFilterBuilder().
		WithDocumentType("some_type").
		Build()

	assert.Equal(t, filter.DocumentTypes[0], "some_type")
}

func TestDocumentFilterBuilder_WithDocumentType(t *testing.T) {
	filter, _ := NewDocumentFilterBuilder().
		WithCountryCode("some_country").
		Build()

	assert.Equal(t, filter.CountryCodes[0], "some_country")
}

func ExampleDocumentFilterBuilder() {
	filter, _ := NewDocumentFilterBuilder().
		WithCountryCode("some_country").
		WithDocumentType("some_type").
		Build()

	data, _ := json.Marshal(filter)
	fmt.Println(string(data))
	// Output: {"document_types":["some_type"],"country_codes":["some_country"]}
}

func ExampleDocumentFilterBuilder_multipleCountriesAndDocumentTypes() {
	filter, _ := NewDocumentFilterBuilder().
		WithCountryCode("some_country").
		WithCountryCode("some_other_country").
		WithDocumentType("some_type").
		WithDocumentType("some_other_type").
		Build()

	data, _ := json.Marshal(filter)
	fmt.Println(string(data))
	// Output: {"document_types":["some_type","some_other_type"],"country_codes":["some_country","some_other_country"]}
}
