package filter

import (
	"encoding/json"
	"fmt"
	"testing"

	"gotest.tools/v3/assert"
)

func TestDocumentFilterBuilder_WithCountryCode(t *testing.T) {
	filter, err := NewDocumentFilterBuilder().
		WithDocumentType("some_type").
		Build()

	assert.NilError(t, err)
	assert.Equal(t, filter.DocumentTypes[0], "some_type")
}

func TestDocumentFilterBuilder_WithDocumentType(t *testing.T) {
	filter, err := NewDocumentFilterBuilder().
		WithCountryCode("some_country").
		Build()

	assert.NilError(t, err)
	assert.Equal(t, filter.CountryCodes[0], "some_country")
}

func ExampleDocumentFilterBuilder() {
	filter, err := NewDocumentFilterBuilder().
		WithCountryCode("some_country").
		WithDocumentType("some_type").
		Build()

	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, err := json.Marshal(filter)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	fmt.Println(string(data))
	// Output: {"document_types":["some_type"],"country_codes":["some_country"]}
}

func ExampleDocumentFilterBuilder_multipleCountriesAndDocumentTypes() {
	filter, err := NewDocumentFilterBuilder().
		WithCountryCode("some_country").
		WithCountryCode("some_other_country").
		WithDocumentType("some_type").
		WithDocumentType("some_other_type").
		Build()

	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, err := json.Marshal(filter)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	fmt.Println(string(data))
	// Output: {"document_types":["some_type","some_other_type"],"country_codes":["some_country","some_other_country"]}
}

func ExampleDocumentFilterBuilder_countriesAndDocumentTypes() {
	filter, err := NewDocumentFilterBuilder().
		WithCountryCodes([]string{"some_country", "some_other_country"}).
		WithDocumentTypes([]string{"some_type", "some_other_type"}).
		Build()

	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, err := json.Marshal(filter)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	fmt.Println(string(data))
	// Output: {"document_types":["some_type","some_other_type"],"country_codes":["some_country","some_other_country"]}
}
