package sandbox

import (
	"encoding/json"
	"fmt"
	"gotest.tools/v3/assert"
	"testing"
)

func TestRecommendationBuilder_WithReason(t *testing.T) {
	recommendation := NewRecommendationBuilder().
		WithReason("some_reason").
		Build()

	assert.Equal(t, recommendation.Reason, "some_reason")
}

func TestRecommendationBuilder_WithValue(t *testing.T) {
	recommendation := NewRecommendationBuilder().
		WithValue("some_value").
		Build()

	assert.Equal(t, recommendation.Value, "some_value")
}

func TestRecommendationBuilder_WithRecoverySuggestion(t *testing.T) {
	recommendation := NewRecommendationBuilder().
		WithRecoverySuggestion("some_suggestion").
		Build()

	assert.Equal(t, recommendation.RecoverySuggestion, "some_suggestion")
}

func ExampleNewRecommendationBuilder() {
	recommendation := NewRecommendationBuilder().
		WithReason("some_reason").
		WithValue("some_value").
		WithRecoverySuggestion("some_suggestion").
		Build()

	data, _ := json.Marshal(recommendation)
	fmt.Println(string(data))
	// Output: {"value":"some_value","reason":"some_reason","recovery_suggestion":"some_suggestion"}
}
