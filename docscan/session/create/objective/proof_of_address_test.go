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

	data, _ := json.Marshal(objective)
	fmt.Println(string(data))
	// Output: {"type":"PROOF_OF_ADDRESS"}
}
