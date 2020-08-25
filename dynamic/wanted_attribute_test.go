package dynamic

import (
	"encoding/json"
	"fmt"
)

func ExampleWantedAttributeBuilder_WithName() {
	builder := (&WantedAttributeBuilder{}).WithName("TEST NAME")
	attribute, err := builder.Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	fmt.Println(attribute.Name)
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

	fmt.Println(attribute.Derivation)
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

	json, _ := json.Marshal(attribute)
	fmt.Println(string(json))
	// Output: {"name":"TEST NAME","constraints":[{"type":"SOURCE","preferred_sources":{"anchors":[],"soft_preference":false}}]}
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

	json, _ := json.Marshal(attribute)
	fmt.Println(string(json))
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

	json, _ := json.Marshal(attribute)
	fmt.Println(string(json))
	// Output: {"name":"TEST NAME","accept_self_asserted":false}
}
