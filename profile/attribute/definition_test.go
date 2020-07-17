package attribute

import (
	"encoding/json"
	"fmt"
)

func ExampleAttributeDefinition_MarshalJSON() {
	exampleDefinition := NewAttributeDefinition("exampleDefinition")
	json, _ := json.Marshal(exampleDefinition)
	fmt.Println(string(json))
	// Output: {"name":"exampleDefinition"}
}
