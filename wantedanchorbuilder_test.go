package yoti

import (
	"fmt"
)

func ExampleWantedAnchorBuilder() {
	aadhaarAnchor := (&WantedAnchorBuilder{}).New().
		WithValue("NATIONAL_ID").
		WithSubType("AADHAAR").
		Build()

	aadhaarJSON, _ := aadhaarAnchor.MarshalJSON()

	fmt.Println("Aadhaar:", string(aadhaarJSON))
	// Output: Aadhaar: {"name":"NATIONAL_ID","sub_type":"AADHAAR"}
}
