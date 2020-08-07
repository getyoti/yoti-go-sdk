package extension

import (
	"fmt"
	"os"
)

func ExampleExtension() {
	content := "SOME CONTENT"
	extType := "SOME_TYPE"
	extension, err := (&Builder{}).WithContent(content).WithType(extType).Build()
	if err != nil {
		fmt.Fprintf(os.Stdout, "error: %s", err.Error())
		return
	}

	data, _ := extension.MarshalJSON()
	fmt.Println(string(data))
	// Output: {"type":"SOME_TYPE","content":"SOME CONTENT"}
}
