package extension

import (
	"fmt"
)

func ExampleExtension() {
	content := "SOME CONTENT"
	extType := "SOME_TYPE"
	extension := (&ExtensionBuilder{}).New().WithContent(content).WithType(extType).Build()

	data, _ := extension.MarshalJSON()
	fmt.Println(string(data))
}
