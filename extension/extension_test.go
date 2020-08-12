package extension

import (
	"fmt"
)

func ExampleExtension() {
	content := "SOME CONTENT"
	extType := "SOME_TYPE"
	extension, err := (&Builder{}).WithContent(content).WithType(extType).Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, _ := extension.MarshalJSON()
	fmt.Println(string(data))
	// Output: {"type":"SOME_TYPE","content":"SOME CONTENT"}
}
