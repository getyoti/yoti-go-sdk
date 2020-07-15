package report

import (
	"encoding/json"
	"fmt"
	"gotest.tools/v3/assert"
	"testing"
)

func Test_recommendationBuilder_WithReason(t *testing.T) {
	recommendation, _ := NewRecommendationBuilder().
		WithReason("some_reason").
		Build()

	assert.Equal(t, recommendation.Reason, "some_reason")
}

func Test_recommendationBuilder_WithValue(t *testing.T) {
	recommendation, _ := NewRecommendationBuilder().
		WithValue("some_value").
		Build()

	assert.Equal(t, recommendation.Value, "some_value")
}

func Test_recommendationBuilder_WithRecoverySuggestion(t *testing.T) {
	recommendation, _ := NewRecommendationBuilder().
		WithRecoverySuggestion("some_suggestion").
		Build()

	assert.Equal(t, recommendation.RecoverySuggestion, "some_suggestion")
}

func Example_recommendationBuilder() {
	recommendation, _ := NewRecommendationBuilder().
		WithReason("some_reason").
		WithValue("some_value").
		WithRecoverySuggestion("some_suggestion").
		Build()

	data, _ := json.Marshal(recommendation)
	fmt.Println(string(data))
	// Output: {"value":"some_value","reason":"some_reason","recovery_suggestion":"some_suggestion"}
}
