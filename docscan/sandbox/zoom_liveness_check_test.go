package sandbox

import (
	"encoding/json"
	"fmt"
)

func Example_zoomLivenessCheckBuilder() {
	breakdown, _ := NewBreakdownBuilder().Build()
	recommendation, _ := NewRecommendationBuilder().Build()

	check, _ := NewZoomLivenessCheckBuilder().
		WithBreakdown(breakdown).
		WithRecommendation(recommendation).
		Build()

	data, _ := json.Marshal(check)
	fmt.Println(string(data))
	// Output: {"result":{"report":{"recommendation":{"value":"","reason":"","recovery_suggestion":""},"breakdown":[{"sub_check":"","result":"","details":null}]}},"liveness_type":"ZOOM"}
}
