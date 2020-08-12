package dynamic

import (
	"fmt"
)

func ExampleWantedAnchorBuilder() {
	aadhaarAnchor, err := (&WantedAnchorBuilder{}).
		WithValue("NATIONAL_ID").
		WithSubType("AADHAAR").
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	aadhaarJSON, _ := aadhaarAnchor.MarshalJSON()

	fmt.Println("Aadhaar:", string(aadhaarJSON))
	// Output: Aadhaar: {"name":"NATIONAL_ID","sub_type":"AADHAAR"}
}
