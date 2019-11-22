package yoti

import (
	"fmt"
	"gotest.tools/assert"
	"testing"
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
	attribute := (&WantedAttributeBuilder{}).New().WithName("attr").WithConstraint(&constraint).Build()

	json, _ := attribute.MarshalJSON()
	fmt.Println(string(json))
	// Output: {"name":"attr","constraints":[{"type":"SOURCE","preferred_sources":{"anchors":[],"soft_preference":false}}]}
}

func ExampleWantedAttributeBuilder_WithAcceptSelfAsserted() {
	attribute := (&WantedAttributeBuilder{}).New().WithName("attr").WithAcceptSelfAsserted(true).Build()

	json, _ := attribute.MarshalJSON()
	fmt.Println(string(json))
	// Output: {"name":"attr","accept_self_asserted":true}
}

func TestWantedAttributeBuilderShouldRejectEmptyName(t *testing.T) {
	attribute := (&WantedAttributeBuilder{}).New().Build()
	json, err := attribute.MarshalJSON()
	assert.Check(t, err != nil)
	assert.Equal(t, err.Error(), "Wanted attribute names must not be empty")
	assert.Equal(t, len(json), 0)
}
