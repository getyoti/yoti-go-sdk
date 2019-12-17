package yoti

import (
	"fmt"
)

func ExampleWantedAnchorBuilder() {
	aadhaarAnchor, err := (&WantedAnchorBuilder{}).New().
		WithValue("NATIONAL_ID").
		WithSubType("AADHAAR").
		Build()
	if err != nil {
		return
	}

	aadhaarJSON, _ := aadhaarAnchor.MarshalJSON()

	fmt.Println("Aadhaar:", string(aadhaarJSON))
	// Output: Aadhaar: {"name":"NATIONAL_ID","sub_type":"AADHAAR"}
}
