package yoti

import (
	"fmt"
)

func ExampleWantedAttributeBuilder_WithName() {
	builder := (&WantedAttributeBuilder{}).New().WithName("TEST NAME")
	attribute := builder.Build()
	fmt.Println(attribute.name)
	// Output: TEST NAME
}

func ExampleWantedAttributeBuilder_WithDerivation() {
	attribute := (&WantedAttributeBuilder{}).New().WithDerivation("TEST DERIVATION").Build()
	fmt.Println(attribute.derivation)
	// Output: TEST DERIVATION
}
