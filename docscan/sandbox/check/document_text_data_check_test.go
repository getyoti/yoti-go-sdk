package check

import (
	"encoding/json"
	"fmt"

	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox"
	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox/check/report"
)

func Example_documentTextDataCheckBuilder() {
	breakdown, _ := report.NewBreakdownBuilder().Build()
	recommendation, _ := report.NewRecommendationBuilder().Build()
	filter, _ := sandbox.NewDocumentFilterBuilder().Build()

	check, _ := NewDocumentTextDataCheckBuilder().
		WithBreakdown(breakdown).
		WithRecommendation(recommendation).
		WithDocumentFilter(filter).
		WithDocumentField("some", "field").
		Build()

	data, _ := json.Marshal(check)
	fmt.Println(string(data))
	// Output: {"result":{"report":{"recommendation":{"value":"","reason":"","recovery_suggestion":""},"breakdown":[{"sub_check":"","result":"","details":null}]},"document_fields":{"some":"field"}},"document_filter":{"document_types":null,"country_codes":null}}
}
