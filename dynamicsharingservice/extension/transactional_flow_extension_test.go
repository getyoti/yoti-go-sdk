package extension

import (
	"fmt"
)

func ExampleTransactionalFlowExtension() {
	content := "SOME CONTENT"

	extension := (&TransactionalFlowExtensionBuilder{}).New().
		WithContent(content).
		Build()

	data, _ := extension.MarshalJSON()
	fmt.Println(string(data))
	// Output: {"type":"TRANSACTIONAL_FLOW","content":"SOME CONTENT"}
}
