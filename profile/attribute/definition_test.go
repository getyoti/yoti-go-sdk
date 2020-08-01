package attribute

import (
	"encoding/json"
	"fmt"
)

func ExampleDefinition_MarshalJSON() {
	exampleDefinition := NewAttributeDefinition("exampleDefinition")
	json, _ := json.Marshal(exampleDefinition)
	fmt.Println(string(json))
	// Output: {"name":"exampleDefinition"}
}
