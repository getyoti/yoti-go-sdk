package task

import (
	"encoding/json"
	"fmt"
)

func ExampleTextDataExtractionReasonBuilder() {
	reason, err := NewTextDataExtractionReasonBuilder().
		ForQuality().
		WithDetail("some-detail").
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, err := json.Marshal(reason)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	fmt.Println(string(data))
	// Output: {"value":"QUALITY","detail":"some-detail"}
}

func ExampleTextDataExtractionReasonBuilder_ForQuality() {
	reason, err := NewTextDataExtractionReasonBuilder().
		ForQuality().
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, err := json.Marshal(reason)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	fmt.Println(string(data))
	// Output: {"value":"QUALITY"}
}
func ExampleTextDataExtractionReasonBuilder_ForUserError() {
	reason, err := NewTextDataExtractionReasonBuilder().
		ForUserError().
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, err := json.Marshal(reason)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	fmt.Println(string(data))
	// Output: {"value":"USER_ERROR"}
}
