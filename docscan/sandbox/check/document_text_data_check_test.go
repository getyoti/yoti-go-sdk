package check

import (
	"encoding/json"
	"fmt"

	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox/check/report"
	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox/filter"
)

func Example_documentTextDataCheckBuilder() {
	breakdown, err := report.NewBreakdownBuilder().
		WithResult("some_result").
		WithSubCheck("some_check").
		Build()
	if err != nil {
		return
	}

	recommendation, err := report.NewRecommendationBuilder().
		WithValue("some_value").
		Build()
	if err != nil {
		return
	}

	filter, err := filter.NewDocumentFilterBuilder().Build()
	if err != nil {
		return
	}

	check, err := NewDocumentTextDataCheckBuilder().
		WithBreakdown(breakdown).
		WithRecommendation(recommendation).
		WithDocumentFilter(filter).
		WithDocumentField("some", "field").
		Build()
	if err != nil {
		return
	}

	data, _ := json.Marshal(check)
	fmt.Println(string(data))
	// Output: {"result":{"report":{"recommendation":{"value":"some_value"},"breakdown":[{"sub_check":"some_check","result":"some_result","details":[]}]},"document_fields":{"some":"field"}},"document_filter":{"document_types":[],"country_codes":[]}}
}
