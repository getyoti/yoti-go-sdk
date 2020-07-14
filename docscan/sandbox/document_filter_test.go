package sandbox

import (
	"encoding/json"
	"fmt"
	"gotest.tools/v3/assert"
	"testing"
)

func Test_documentFilterBuilder_WithCountryCode(t *testing.T) {
	filter := NewDocumentFilterBuilder().
		WithDocumentType("some_type").
		Build()

	assert.Equal(t, filter.DocumentTypes[0], "some_type")
}

func Test_documentFilterBuilder_WithDocumentType(t *testing.T) {
	filter := NewDocumentFilterBuilder().
		WithCountryCode("some_country").
		Build()

	assert.Equal(t, filter.CountryCodes[0], "some_country")
}

func Example_documentFilterBuilder_WithDocumentType_WithCountryCode() {
	filter := NewDocumentFilterBuilder().
		WithCountryCode("some_country").
		WithDocumentType("some_type").
		Build()

	data, _ := json.Marshal(filter)
	fmt.Println(string(data))
	// Output: {"document_types":["some_type"],"country_codes":["some_country"]}
}

func Example_documentFilterBuilder_WithDocumentType_WithCountryCode_Multiple() {
	filter := NewDocumentFilterBuilder().
		WithCountryCode("some_country").
		WithCountryCode("some_other_country").
		WithDocumentType("some_type").
		WithDocumentType("some_other_type").
		Build()

	data, _ := json.Marshal(filter)
	fmt.Println(string(data))
	// Output: {"document_types":["some_type","some_other_type"],"country_codes":["some_country","some_other_country"]}
}
