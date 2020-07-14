package sandbox

import (
	"encoding/json"
	"fmt"
)

func Example_documentAuthenticityCheckBuilder() {
	breakdown := NewBreakdownBuilder().Build()
	recommendation := NewRecommendationBuilder().Build()
	filter := NewDocumentFilterBuilder().Build()

	check := NewDocumentAuthenticityCheckBuilder().
		WithBreakdown(breakdown).
		WithRecommendation(recommendation).
		WithDocumentFilter(filter).
		Build()

	data, _ := json.Marshal(check)
	fmt.Println(string(data))
	// Output: {"result":{"report":{"recommendation":{"value":"","reason":"","recovery_suggestion":""},"breakdown":[{"sub_check":"","result":"","details":null}]}},"document_filter":{"document_types":null,"country_codes":null}}
}
