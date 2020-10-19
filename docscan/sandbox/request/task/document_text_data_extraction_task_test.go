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

	task, err := NewDocumentTextDataExtractionTaskBuilder().
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

	data, _ := json.Marshal(task)
	fmt.Println(string(data))
	// Output: {"document_filter":{"document_types":[],"country_codes":[]},"result":{"document_fields":{"some-key":"some-value","some-other-key":{"some-nested-key":"some-nested-value"}}}}
}

func ExampleDocumentTextDataExtractionTaskBuilder_WithDocumentFields() {
	docFilter, err := filter.NewDocumentFilterBuilder().Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	task, err := NewDocumentTextDataExtractionTaskBuilder().
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

	data, _ := json.Marshal(task)
	fmt.Println(string(data))
	// Output: {"document_filter":{"document_types":[],"country_codes":[]},"result":{"document_fields":{"some-key":"some-value","some-other-key":{"some-nested-key":"some-nested-value"}}}}
}

func ExampleDocumentTextDataExtractionTaskBuilder_WithDocumentIDPhoto() {
	task, err := NewDocumentTextDataExtractionTaskBuilder().
		WithDocumentIDPhoto("some-content-type", []byte{0xDE, 0xAD, 0xBE, 0xEF}).
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, _ := json.Marshal(task)
	fmt.Println(string(data))
	// Output: {"result":{"document_id_photo":{"content_type":"some-content-type","data":"3q2+7w=="}}}
}

func ExampleDocumentTextDataExtractionTaskBuilder_WithDetectedCountry() {
	task, err := NewDocumentTextDataExtractionTaskBuilder().
		WithDetectedCountry("some-country").
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, _ := json.Marshal(task)
	fmt.Println(string(data))
	// Output: {"result":{"detected_country":"some-country"}}
}

func ExampleDocumentTextDataExtractionTaskBuilder_WithRecommendation() {
	recommendation, err := NewTextDataExtractionRecommendationBuilder().
		ForProgress().
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	task, err := NewDocumentTextDataExtractionTaskBuilder().
		WithRecommendation(recommendation).
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, _ := json.Marshal(task)
	fmt.Println(string(data))
	// Output: {"result":{"recommendation":{"value":"PROGRESS"}}}
}

func ExampleDocumentTextDataExtractionTaskBuilder_minimal() {
	task, err := NewDocumentTextDataExtractionTaskBuilder().Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, _ := json.Marshal(task)
	fmt.Println(string(data))
	// Output: {"result":{}}
}
