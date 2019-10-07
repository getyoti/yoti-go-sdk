package sandbox

import (
	"fmt"
	"time"
)

func ExampleAttribute_WithAnchor() {
	attribute := Attribute{
		Name:  "AttributeName",
		Value: "Value",
	}.WithAnchor(SourceAnchor("", time.Unix(1234567890, 0), ""))
	fmt.Println(attribute)
	// Output: {AttributeName Value   [{SOURCE   2009-02-13 23:31:30 +0000 GMT}]}
}

func ExampleDerivation_AgeOver() {
	attribute := Attribute{
		Name:       "date_of_birth",
		Value:      "Value",
		Derivation: Derivation{}.AgeOver(18).ToString(),
	}
	fmt.Println(attribute)
	// Output: {date_of_birth Value age_over:18  []}
}

func ExampleDerivation_AgeUnder() {
	attribute := Attribute{
		Name:       "date_of_birth",
		Value:      "Value",
		Derivation: Derivation{}.AgeUnder(14).ToString(),
	}
	fmt.Println(attribute)
	// Output: {date_of_birth Value age_under:14  []}
}
