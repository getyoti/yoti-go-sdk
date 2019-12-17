package yoti

import (
	"fmt"
)

func ExampleWantedAttributeBuilder_WithName() {
	builder := (&WantedAttributeBuilder{}).New().WithName("TEST NAME")
	attribute, err := builder.Build()
	if err != nil {
		return
	}
	fmt.Println(attribute.name)
	// Output: TEST NAME
}

func ExampleWantedAttributeBuilder_WithDerivation() {
	attribute, err := (&WantedAttributeBuilder{}).New().
		WithDerivation("TEST DERIVATION").
		WithName("TEST NAME").
		Build()
	if err != nil {
		return
	}
	fmt.Println(attribute.derivation)
	// Output: TEST DERIVATION
}

func ExampleWantedAttributeBuilder_WithConstraint() {
	constraint, err := (&SourceConstraintBuilder{}).New().
		Build()
	if err != nil {
		return
	}
	attribute, err := (&WantedAttributeBuilder{}).New().
		WithName("TEST NAME").
		WithConstraint(&constraint).
		Build()
	if err != nil {
		return
	}

	json, _ := attribute.MarshalJSON()
	fmt.Println(string(json))
	// Output: {"name":"TEST NAME","constraints":[{"type":"SOURCE","preferred_sources":{"anchors":[],"soft_preference":false}}]}
}

func ExampleWantedAttributeBuilder_WithAcceptSelfAsserted() {
	attribute, err := (&WantedAttributeBuilder{}).New().
		WithName("TEST NAME").
		WithAcceptSelfAsserted(true).
		Build()
	if err != nil {
		return
	}

	json, _ := attribute.MarshalJSON()
	fmt.Println(string(json))
	// Output: {"name":"TEST NAME","accept_self_asserted":true}
}
