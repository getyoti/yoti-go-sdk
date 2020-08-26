package dynamic

import (
	"fmt"
)

func ExampleWantedAttributeBuilder_WithName() {
	builder := (&WantedAttributeBuilder{}).WithName("TEST NAME")
	attribute, err := builder.Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	fmt.Println(attribute.name)
	// Output: TEST NAME
}

func ExampleWantedAttributeBuilder_WithDerivation() {
	attribute, err := (&WantedAttributeBuilder{}).
		WithDerivation("TEST DERIVATION").
		WithName("TEST NAME").
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	fmt.Println(attribute.derivation)
	// Output: TEST DERIVATION
}

func ExampleWantedAttributeBuilder_WithConstraint() {
	constraint, err := (&SourceConstraintBuilder{}).
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	attribute, err := (&WantedAttributeBuilder{}).
		WithName("TEST NAME").
		WithConstraint(&constraint).
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	marshalledJSON, _ := attribute.MarshalJSON()
	fmt.Println(string(marshalledJSON))
	// Output: {"name":"TEST NAME","constraints":[{"type":"SOURCE","preferred_sources":{"anchors":[],"soft_preference":false}}],"accept_self_asserted":false}
}

func ExampleWantedAttributeBuilder_WithAcceptSelfAsserted() {
	attribute, err := (&WantedAttributeBuilder{}).
		WithName("TEST NAME").
		WithAcceptSelfAsserted(true).
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	marshalledJSON, _ := attribute.MarshalJSON()
	fmt.Println(string(marshalledJSON))
	// Output: {"name":"TEST NAME","accept_self_asserted":true}
}

func ExampleWantedAttributeBuilder_WithAcceptSelfAsserted_false() {
	attribute, err := (&WantedAttributeBuilder{}).
		WithName("TEST NAME").
		WithAcceptSelfAsserted(false).
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	marshalledJSON, _ := attribute.MarshalJSON()
	fmt.Println(string(marshalledJSON))
	// Output: {"name":"TEST NAME","accept_self_asserted":false}
}
