package sandbox

import (
	"encoding/json"
	"fmt"
	"gotest.tools/v3/assert"
	"testing"
)

func TestRecommendation_WithReason(t *testing.T) {
	recommendation := Recommendation{}.WithReason("some_reason")

	assert.Equal(t, recommendation.Reason, "some_reason")
}

func TestRecommendation_WithValue(t *testing.T) {
	recommendation := Recommendation{}.WithValue("some_value")

	assert.Equal(t, recommendation.Value, "some_value")
}

func TestRecommendation_WithRecoverySuggestion(t *testing.T) {
	recommendation := Recommendation{}.WithRecoverySuggestion("some_suggestion")

	assert.Equal(t, recommendation.RecoverySuggestion, "some_suggestion")
}

func ExampleRecommendation() {
	recommendation := Recommendation{}.
		WithReason("some_reason").
		WithValue("some_value").
		WithRecoverySuggestion("some_suggestion")

	data, _ := json.Marshal(recommendation)
	fmt.Println(string(data))
	// Output: {"value":"some_value","reason":"some_reason","recovery_suggestion":"some_suggestion"}
}
