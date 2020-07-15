package check

import (
	"encoding/json"
	"fmt"

	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox/check/report"
)

func Example_zoomLivenessCheckBuilder() {
	breakdown, _ := report.NewBreakdownBuilder().Build()
	recommendation, _ := report.NewRecommendationBuilder().Build()

	check, _ := NewZoomLivenessCheckBuilder().
		WithBreakdown(breakdown).
		WithRecommendation(recommendation).
		Build()

	data, _ := json.Marshal(check)
	fmt.Println(string(data))
	// Output: {"result":{"report":{"recommendation":{"value":"","reason":"","recovery_suggestion":""},"breakdown":[{"sub_check":"","result":"","details":null}]}},"liveness_type":"ZOOM"}
}
