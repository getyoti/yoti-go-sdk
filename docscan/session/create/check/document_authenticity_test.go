package check

import (
	"encoding/json"
	"fmt"
)

func ExampleRequestedDocumentAuthenticityCheckBuilder() {
	check, err := NewRequestedDocumentAuthenticityCheckBuilder().Build()

	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, _ := json.Marshal(check)
	fmt.Println(string(data))
	// Output: {"type":"ID_DOCUMENT_AUTHENTICITY","config":{}}
}
