package check

import (
	"encoding/json"
	"fmt"

	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox/request/check/report"
	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox/request/filter"
)

func ExampleSupplementaryDocumentTextDataCheckBuilder() {
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

	check, err := NewSupplementaryDocumentTextDataCheckBuilder().
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

func ExampleSupplementaryDocumentTextDataCheckBuilder_WithDocumentFields() {
	check, err := NewSupplementaryDocumentTextDataCheckBuilder().
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

func ExampleSupplementaryDocumentTextDataCheckBuilder_minimal() {
	check, err := NewSupplementaryDocumentTextDataCheckBuilder().Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, _ := json.Marshal(check)
	fmt.Println(string(data))
	// Output: {"result":{"report":{}}}
}
