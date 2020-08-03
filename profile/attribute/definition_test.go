package attribute

import (
	"encoding/json"
	"fmt"
)

func ExampleDefinition_MarshalJSON() {
	exampleDefinition := NewAttributeDefinition("exampleDefinition")
	marshalledJSON, _ := json.Marshal(exampleDefinition)
	fmt.Println(string(marshalledJSON))
	// Output: {"name":"exampleDefinition"}
}
