package check

import (
	"encoding/json"
	"fmt"
)

func ExampleRequestedFaceMatchCheckBuilder() {
	check, err := NewRequestedFaceMatchCheckBuilder().
		WithManualCheckNever().
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, _ := json.Marshal(check)
	fmt.Println(string(data))
	// Output: {"type":"ID_DOCUMENT_FACE_MATCH","config":{"manual_check":"NEVER"}}
}
