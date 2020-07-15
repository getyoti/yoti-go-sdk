package check

import (
	"encoding/json"
	"fmt"

	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox/check/report"
)

func Example_zoomLivenessCheckBuilder() {
	breakdown, err := report.NewBreakdownBuilder().Build()
	if err != nil {
		return
	}

	recommendation, err := report.NewRecommendationBuilder().Build()
	if err != nil {
		return
	}

	check, err := NewZoomLivenessCheckBuilder().
		WithBreakdown(breakdown).
		WithRecommendation(recommendation).
		Build()
	if err != nil {
		return
	}

	data, _ := json.Marshal(check)
	fmt.Println(string(data))
	// Output: {"result":{"report":{"recommendation":{"value":"","reason":"","recovery_suggestion":""},"breakdown":[{"sub_check":"","result":"","details":null}]}},"liveness_type":"ZOOM"}
}
