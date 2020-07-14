package sandbox

import (
	"encoding/json"
	"fmt"
)

func Example_documentTextDataExtractionTaskBuilder() {
	breakdown, _ := NewBreakdownBuilder().Build()
	recommendation, _ := NewRecommendationBuilder().Build()
	filter, _ := NewDocumentFilterBuilder().Build()

	check, _ := NewDocumentTextDataExtractionTaskBuilder().
		WithBreakdown(breakdown).
		WithRecommendation(recommendation).
		WithDocumentFilter(filter).
		WithDocumentField("some", "field").
		Build()

	data, _ := json.Marshal(check)
	fmt.Println(string(data))
	// Output: {"result":{"report":{"recommendation":{"value":"","reason":"","recovery_suggestion":""},"breakdown":[{"sub_check":"","result":"","details":null}]},"document_fields":{"some":"field"}},"document_filter":{"document_types":null,"country_codes":null}}
}
