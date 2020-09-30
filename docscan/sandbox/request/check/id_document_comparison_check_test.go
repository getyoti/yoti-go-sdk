package check

import (
	"encoding/json"
	"fmt"

	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox/request/check/report"
	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox/request/filter"
)

func ExampleIDDocumentComparisonCheckBuilder() {
	breakdown, err := report.NewBreakdownBuilder().
		WithResult("some_result").
		WithSubCheck("some_check").
		WithDetail("some_name", "some_value").
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

	docFilter, err := filter.NewDocumentFilterBuilder().
		WithCountryCode("AUS").
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	secondaryDocFilter, err := filter.NewDocumentFilterBuilder().
		WithCountryCode("FJI").
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	idDocumentComparisonCheck, err := NewIDDocumentComparisonCheckBuilder().
		WithBreakdown(breakdown).
		WithRecommendation(recommendation).
		WithDocumentFilter(docFilter).
		WithSecondaryDocumentFilter(secondaryDocFilter).
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, _ := json.Marshal(idDocumentComparisonCheck)
	fmt.Println(string(data))
	// Output: {"result":{"report":{"recommendation":{"value":"some_value"},"breakdown":[{"sub_check":"some_check","result":"some_result","details":[{"name":"some_name","value":"some_value"}]}]}},"document_filter":{"document_types":[],"country_codes":["AUS"]},"secondary_document_filter":{"document_types":[],"country_codes":["FJI"]}}
}

func ExampleIDDocumentComparisonCheckBuilder_minimal() {
	idDocumentComparisonCheck, err := NewIDDocumentComparisonCheckBuilder().Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, _ := json.Marshal(idDocumentComparisonCheck)
	fmt.Println(string(data))
	// Output: {"result":{"report":{}}}
}
