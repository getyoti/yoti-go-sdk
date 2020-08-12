package task

import (
	"encoding/json"
	"fmt"

	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox/request/filter"
)

func ExampleDocumentTextDataExtractionTaskBuilder() {
	docFilter, err := filter.NewDocumentFilterBuilder().Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	check, err := NewDocumentTextDataExtractionTaskBuilder().
		WithDocumentFilter(docFilter).
		WithDocumentField("some-key", "some-value").
		WithDocumentField("some-other-key", map[string]string{
			"some-nested-key": "some-nested-value",
		}).
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, _ := json.Marshal(check)
	fmt.Println(string(data))
	// Output: {"document_filter":{"document_types":[],"country_codes":[]},"result":{"document_fields":{"some-key":"some-value","some-other-key":{"some-nested-key":"some-nested-value"}}}}
}

func ExampleDocumentTextDataExtractionTaskBuilder_WithDocumentFields() {
	docFilter, err := filter.NewDocumentFilterBuilder().Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	check, err := NewDocumentTextDataExtractionTaskBuilder().
		WithDocumentFilter(docFilter).
		WithDocumentFields(map[string]interface{}{
			"some-key": "some-value",
			"some-other-key": map[string]string{
				"some-nested-key": "some-nested-value",
			},
		}).
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, _ := json.Marshal(check)
	fmt.Println(string(data))
	// Output: {"document_filter":{"document_types":[],"country_codes":[]},"result":{"document_fields":{"some-key":"some-value","some-other-key":{"some-nested-key":"some-nested-value"}}}}
}

func ExampleDocumentTextDataExtractionTaskBuilder_minimal() {
	check, err := NewDocumentTextDataExtractionTaskBuilder().Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, _ := json.Marshal(check)
	fmt.Println(string(data))
	// Output: {"result":{}}
}
