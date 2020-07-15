package report

import (
	"encoding/json"
	"fmt"
	"gotest.tools/v3/assert"
	"testing"
)

func Test_recommendationBuilder(t *testing.T) {
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

func Test_recommendationBuilder_ShouldRequireValue(t *testing.T) {
	_, err := NewRecommendationBuilder().Build()

	assert.Error(t, err, "Value cannot be empty")
}

func Example_recommendationBuilder() {
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

func Example_recommendationBuilder_Minimal() {
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
