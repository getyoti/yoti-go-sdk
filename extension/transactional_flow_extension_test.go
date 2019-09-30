package extension

import (
	"encoding/json"
	"fmt"
)

func ExampleTransactionalFlowExtension() {
	content := "SOME CONTENT"

	extension := (&TransactionalFlowExtensionBuilder{}).New().
		WithContent(content).
		Build()

	data, _ := json.Marshal(extension)
	fmt.Println(string(data))
	// Output: {"type":"TRANSACTIONAL_FLOW","content":"SOME CONTENT"}
}
