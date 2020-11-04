package objective

import (
	"encoding/json"
	"fmt"
)

func ExampleProofOfAddressObjectiveBuilder() {
	objective, err := NewProofOfAddressObjectiveBuilder().
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, err := json.Marshal(objective)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	fmt.Println(string(data))
	// Output: {"type":"PROOF_OF_ADDRESS"}
}
