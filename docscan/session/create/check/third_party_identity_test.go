package check

import (
	"encoding/json"
	"fmt"
)

func ExampleRequestedThirdPartyIdentityCheck() {
	thirdPartyCheck, err := NewRequestedThirdPartyIdentityCheckBuilder().Build()

	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	data, err := json.Marshal(thirdPartyCheck)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	fmt.Println(string(data))
	// Output: {"type":"THIRD_PARTY_IDENTITY","config":{}}
}
