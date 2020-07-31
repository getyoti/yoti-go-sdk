package check

import (
	"encoding/json"
	"fmt"
	"os"
)

func ExampleRequestedDocumentAuthenticityCheckBuilder() {
	check, err := NewRequestedDocumentAuthenticityCheckBuilder().Build()

	if err != nil {
		fmt.Fprintf(os.Stdout, "error: %s", err.Error())
		return
	}

	data, _ := json.Marshal(check)
	fmt.Println(string(data))
	// Output: {"type":"ID_DOCUMENT_AUTHENTICITY","config":{}}

	// TODO: check we want config={}?
}
