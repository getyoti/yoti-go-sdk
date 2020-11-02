package attribute

import (
	"encoding/json"
	"fmt"
)

func ExampleDefinition_MarshalJSON() {
	exampleDefinition := NewAttributeDefinition("exampleDefinition")
	marshalledJSON, err := json.Marshal(exampleDefinition)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	fmt.Println(string(marshalledJSON))
	// Output: {"name":"exampleDefinition"}
}
