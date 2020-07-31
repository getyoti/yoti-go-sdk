package check

import (
	"encoding/json"
	"fmt"
	"os"
)

func ExampleRequestedLivenessCheckBuilder() {
	check, err := NewRequestedLivenessCheckBuilder().
		ForZoomLiveness().
		WithMaxRetries(9).
		Build()
	if err != nil {
		fmt.Fprintf(os.Stdout, "error: %s", err.Error())
		return
	}

	data, _ := json.Marshal(check)
	fmt.Println(string(data))
	// Output: {"type":"LIVENESS","config":{"max_retries":9,"liveness_type":"ZOOM"}}
}
