package check

import (
	"encoding/json"
	"fmt"
	"os"
)

func ExampleRequestedFaceMatchCheckBuilder() {
	check, err := NewRequestedFaceMatchCheckBuilder().
		WithManualCheckNever().
		Build()
	if err != nil {
		fmt.Fprintf(os.Stdout, "error: %s", err.Error())
		return
	}

	data, _ := json.Marshal(check)
	fmt.Println(string(data))
	// Output: {"type":"ID_DOCUMENT_FACE_MATCH","config":{"manual_check":"NEVER"}}
}
