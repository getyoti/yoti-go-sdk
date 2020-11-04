package task

import (
	"encoding/json"
	"fmt"

	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox/request/filter"
)

func ExampleSupplementaryDocumentTextDataExtractionTaskBuilder() {
	docFilter, err := filter.NewDocumentFilterBuilder().Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	task, err := NewSupplementaryDocumentTextDataExtractionTaskBuilder().
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

	data, err := json.Marshal(task)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	fmt.Println(string(data))
	// Output: {"document_filter":{"document_types":[],"country_codes":[]},"result":{"document_fields":{"some-key":"some-value","some-other-key":{"some-nested-key":"some-nested-value"}}}}
}

func ExampleSupplementaryDocumentTextDataExtractionTaskBuilder_WithDocumentFields() {
	docFilter, err := filter.NewDocumentFilterBuilder().Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	task, err := NewSupplementaryDocumentTextDataExtractionTaskBuilder().
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

	data, err := json.Marshal(task)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	fmt.Println(string(data))
	// Output: {"document_filter":{"document_types":[],"country_codes":[]},"result":{"document_fields":{"some-key":"some-value","some-other-key":{"some-nested-key":"some-nested-value"}}}}
}

func ExampleSupplementaryDocumentTextDataExtractionTaskBuilder_WithDetectedCountry() {
	task, err := NewSupplementaryDocumentTextDataExtractionTaskBuilder().
		WithDetectedCountry("some-country").
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, err := json.Marshal(task)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	fmt.Println(string(data))
	// Output: {"result":{"detected_country":"some-country"}}
}

func ExampleSupplementaryDocumentTextDataExtractionTaskBuilder_WithRecommendation() {
	recommendation, err := NewTextDataExtractionRecommendationBuilder().
		ForProgress().
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	task, err := NewSupplementaryDocumentTextDataExtractionTaskBuilder().
		WithRecommendation(recommendation).
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, err := json.Marshal(task)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	fmt.Println(string(data))
	// Output: {"result":{"recommendation":{"value":"PROGRESS"}}}
}

func ExampleSupplementaryDocumentTextDataExtractionTaskBuilder_minimal() {
	task, err := NewSupplementaryDocumentTextDataExtractionTaskBuilder().Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, err := json.Marshal(task)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	fmt.Println(string(data))
	// Output: {"result":{}}
}
