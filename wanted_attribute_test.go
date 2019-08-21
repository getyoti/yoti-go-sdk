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

func ExampleWantedAttributeBuilder_WithConstraint() {
	constraint := (&SourceConstraintBuilder{}).New().Build()
	attribute := (&WantedAttributeBuilder{}).New().WithConstraint(&constraint).Build()

	json, _ := attribute.MarshalJSON()
	fmt.Println(string(json))
	// Output: {"constraints":[{"type":"SOURCE","preferred_sources":{"anchors":[],"soft_preference":false}}]}
}

func ExampleWantedAttributeBuilder_WithOptional() {
	attribute := (&WantedAttributeBuilder{}).New().
		WithName(AttrConstGender).
		WithOptional(true).
		Build()

	json, _ := attribute.MarshalJSON()
	fmt.Println(string(json))
	// Output: {"name":"gender","optional":true}
}

func ExampleWantedAttributeBuilder_WithAcceptSelfAsserted() {
	attribute := (&WantedAttributeBuilder{}).New().WithAcceptSelfAsserted(true).Build()

	json, _ := attribute.MarshalJSON()
	fmt.Println(string(json))
	// Output: {"accept_self_asserted":true}
}
