package extension

import (
	"fmt"
)

func ExampleExtension() {
	content := "SOME CONTENT"
	extType := "SOME_TYPE"
	extension, err := (&Builder{}).WithContent(content).WithType(extType).Build()
	if err != nil {
		return
	}

	data, _ := extension.MarshalJSON()
	fmt.Println(string(data))
	// Output: {"type":"SOME_TYPE","content":"SOME CONTENT"}
}
