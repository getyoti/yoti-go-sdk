package check

import (
	"encoding/json"
	"fmt"

	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox/request/check/report"
	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox/request/filter"
)

func ExampleDocumentTextDataCheckBuilder() {
	breakdown, err := report.NewBreakdownBuilder().
		WithResult("some_result").
		WithSubCheck("some_check").
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	recommendation, err := report.NewRecommendationBuilder().
		WithValue("some_value").
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	docFilter, err := filter.NewDocumentFilterBuilder().Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	check, err := NewDocumentTextDataCheckBuilder().
		WithBreakdown(breakdown).
		WithRecommendation(recommendation).
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
	// Output: {"result":{"report":{"recommendation":{"value":"some_value"},"breakdown":[{"sub_check":"some_check","result":"some_result","details":[]}]},"document_fields":{"some-key":"some-value","some-other-key":{"some-nested-key":"some-nested-value"}}},"document_filter":{"document_types":[],"country_codes":[]}}
}

func ExampleDocumentTextDataCheckBuilder_WithDocumentFields() {
	check, err := NewDocumentTextDataCheckBuilder().
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
	// Output: {"result":{"report":{},"document_fields":{"some-key":"some-value","some-other-key":{"some-nested-key":"some-nested-value"}}}}
}

func ExampleDocumentTextDataCheckBuilder_minimal() {
	check, err := NewDocumentTextDataCheckBuilder().Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, _ := json.Marshal(check)
	fmt.Println(string(data))
	// Output: {"result":{"report":{}}}
}
