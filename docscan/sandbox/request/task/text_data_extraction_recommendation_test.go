package task

import (
	"encoding/json"
	"fmt"
)

func ExampleTextDataExtractionRecommendationBuilder() {
	reason, err := NewTextDataExtractionReasonBuilder().
		ForQuality().
		Build()

	recommendation, err := NewTextDataExtractionRecommendationBuilder().
		ForShouldTryAgain().
		WithReason(reason).
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, _ := json.Marshal(recommendation)
	fmt.Println(string(data))
	// Output: {"value":"SHOULD_TRY_AGAIN","reason":{"value":"QUALITY"}}
}

func ExampleTextDataExtractionRecommendationBuilder_ForProgress() {
	recommendation, err := NewTextDataExtractionRecommendationBuilder().
		ForProgress().
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, _ := json.Marshal(recommendation)
	fmt.Println(string(data))
	// Output: {"value":"PROGRESS"}
}

func ExampleTextDataExtractionRecommendationBuilder_ForShouldTryAgain() {
	recommendation, err := NewTextDataExtractionRecommendationBuilder().
		ForShouldTryAgain().
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, _ := json.Marshal(recommendation)
	fmt.Println(string(data))
	// Output: {"value":"SHOULD_TRY_AGAIN"}
}

func ExampleTextDataExtractionRecommendationBuilder_ForMustTryAgain() {
	recommendation, err := NewTextDataExtractionRecommendationBuilder().
		ForMustTryAgain().
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, _ := json.Marshal(recommendation)
	fmt.Println(string(data))
	// Output: {"value":"MUST_TRY_AGAIN"}
}
