package task

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox/request/filter"
)

func ExampleDocumentTextDataExtractionTaskBuilder() {
	docFilter, err := filter.NewDocumentFilterBuilder().Build()
	if err != nil {
		fmt.Fprintf(os.Stdout, "error: %s", err.Error())
		return
	}

	check, err := NewDocumentTextDataExtractionTaskBuilder().
		WithDocumentFilter(docFilter).
		WithDocumentField("some", "field").
		WithDocumentField("other", "field").
		Build()
	if err != nil {
		fmt.Fprintf(os.Stdout, "error: %s", err.Error())
		return
	}

	data, _ := json.Marshal(check)
	fmt.Println(string(data))
	// Output: {"document_filter":{"document_types":[],"country_codes":[]},"result":{"document_fields":{"other":"field","some":"field"}}}
}

func ExampleDocumentTextDataExtractionTaskBuilder_WithDocumentFields() {
	docFilter, err := filter.NewDocumentFilterBuilder().Build()
	if err != nil {
		fmt.Fprintf(os.Stdout, "error: %s", err.Error())
		return
	}

	check, err := NewDocumentTextDataExtractionTaskBuilder().
		WithDocumentFilter(docFilter).
		WithDocumentFields(map[string]string{
			"some": "field",
		}).
		Build()
	if err != nil {
		fmt.Fprintf(os.Stdout, "error: %s", err.Error())
		return
	}

	data, _ := json.Marshal(check)
	fmt.Println(string(data))
	// Output: {"document_filter":{"document_types":[],"country_codes":[]},"result":{"document_fields":{"some":"field"}}}
}

func ExampleDocumentTextDataExtractionTaskBuilder_minimal() {
	check, err := NewDocumentTextDataExtractionTaskBuilder().Build()
	if err != nil {
		fmt.Fprintf(os.Stdout, "error: %s", err.Error())
		return
	}

	data, _ := json.Marshal(check)
	fmt.Println(string(data))
	// Output: {"result":{}}
}
