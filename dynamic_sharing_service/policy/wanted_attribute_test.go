package policy

import (
	"fmt"
)

func ExampleSetName() {
	builder := (&WantedAttributeBuilder{}).New().WithName("TEST NAME")
	attribute := builder.Build()
	fmt.Println(attribute.name)
	// Output: TEST NAME
}

func ExampleSetDerivation() {
	attribute := (&WantedAttributeBuilder{}).New().WithDerivation("TEST DERIVATION").Build()
	fmt.Println(attribute.derivation)
	// Output: TEST DERIVATION
}
