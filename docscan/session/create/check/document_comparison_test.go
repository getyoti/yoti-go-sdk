package check

import (
	"encoding/json"
	"fmt"
)

func ExampleNewRequestedIDDocumentComparisonCheckBuilder() {
	check, err := NewRequestedIDDocumentComparisonCheckBuilder().Build()

	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, err := json.Marshal(check)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	fmt.Println(string(data))
	// Output: {"type":"ID_DOCUMENT_COMPARISON","config":{}}
}
