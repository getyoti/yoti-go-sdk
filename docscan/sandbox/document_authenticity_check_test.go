package sandbox

import (
	"encoding/json"
	"fmt"
)

func ExampleNewDocumentAuthenticityCheckBuilder() {
	breakdown := NewBreakdownBuilder().Build()
	recommendation := NewRecommendationBuilder().Build()

	check := NewDocumentAuthenticityCheckBuilder().
		WithBreakdown(breakdown).
		WithRecommendation(recommendation).
		Build()

	data, _ := json.Marshal(check)
	fmt.Println(string(data))
	// Output: {"result":{"report":{"recommendation":{"value":"","reason":"","recovery_suggestion":""},"breakdown":[{"sub_check":"","result":"","details":null}]}}}
}
