package report

import (
	"encoding/json"
	"fmt"
	"testing"

	"gotest.tools/v3/assert"
)

func TestRecommendationBuilder(t *testing.T) {
	recommendation, err := NewRecommendationBuilder().
		WithValue("some_value").
		WithReason("some_reason").
		WithRecoverySuggestion("some_suggestion").
		Build()

	assert.NilError(t, err)
	assert.Equal(t, recommendation.Reason, "some_reason")
	assert.Equal(t, recommendation.Value, "some_value")
	assert.Equal(t, recommendation.RecoverySuggestion, "some_suggestion")
}

func TestRecommendationBuilder_ShouldRequireValue(t *testing.T) {
	_, err := NewRecommendationBuilder().Build()

	assert.Error(t, err, "Value cannot be empty")
}

func ExampleNewRecommendationBuilder() {
	recommendation, err := NewRecommendationBuilder().
		WithReason("some_reason").
		WithValue("some_value").
		WithRecoverySuggestion("some_suggestion").
		Build()
	if err != nil {
		return
	}

	data, _ := json.Marshal(recommendation)
	fmt.Println(string(data))
	// Output: {"value":"some_value","reason":"some_reason","recovery_suggestion":"some_suggestion"}
}

func ExampleNewRecommendationBuilder_minimal() {
	recommendation, err := NewRecommendationBuilder().
		WithValue("some_value").
		Build()
	if err != nil {
		return
	}

	data, _ := json.Marshal(recommendation)
	fmt.Println(string(data))
	// Output: {"value":"some_value"}
}
