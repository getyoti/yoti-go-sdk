package extension

import (
	"encoding/json"
	"fmt"
	"os"
)

func ExampleTransactionalFlowExtension() {
	content := "SOME CONTENT"

	extension, err := (&TransactionalFlowExtensionBuilder{}).
		WithContent(content).
		Build()
	if err != nil {
		fmt.Fprintf(os.Stdout, "error: %s", err.Error())
		return
	}

	data, _ := json.Marshal(extension)
	fmt.Println(string(data))
	// Output: {"type":"TRANSACTIONAL_FLOW","content":"SOME CONTENT"}
}
