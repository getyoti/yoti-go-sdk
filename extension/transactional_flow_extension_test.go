package extension

import (
	"encoding/json"
	"fmt"
)

func ExampleTransactionalFlowExtension() {
	content := "SOME CONTENT"

	extension, err := (&TransactionalFlowExtensionBuilder{}).
		WithContent(content).
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, err := json.Marshal(extension)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	fmt.Println(string(data))
	// Output: {"type":"TRANSACTIONAL_FLOW","content":"SOME CONTENT"}
}
