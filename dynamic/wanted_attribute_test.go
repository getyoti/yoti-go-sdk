package dynamic

import (
	"encoding/json"
	"fmt"
	"testing"

	"gotest.tools/v3/assert"
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

	marshalledJSON, err := attribute.MarshalJSON()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

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

	marshalledJSON, err := attribute.MarshalJSON()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

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

	marshalledJSON, err := attribute.MarshalJSON()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	fmt.Println(string(marshalledJSON))
	// Output: {"name":"TEST NAME","accept_self_asserted":false}
}

func ExampleWantedAttributeBuilder_optional_true() {
	attribute, err := (&WantedAttributeBuilder{}).
		WithName("TEST NAME").
		WithAcceptSelfAsserted(false).
		Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	attribute.Optional = true

	marshalledJSON, err := attribute.MarshalJSON()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	fmt.Println(string(marshalledJSON))
	// Output: {"name":"TEST NAME","accept_self_asserted":false,"optional":true}
}

func TestWantedAttributeBuilder_Optional_IsOmittedByDefault(t *testing.T) {
	attribute, err := (&WantedAttributeBuilder{}).
		WithName("TEST NAME").
		Build()
	if err != nil {
		t.Errorf("error: %s", err.Error())
	}

	marshalledJSON, err := attribute.MarshalJSON()
	if err != nil {
		t.Errorf("error: %s", err.Error())
	}

	attributeMap := unmarshalJSONIntoMap(t, marshalledJSON)

	optional := attributeMap["optional"]

	if optional != nil {
		t.Errorf("expected `optional` to be nil, but was: '%v'", optional)
	}
}

func unmarshalJSONIntoMap(t *testing.T, byteValue []byte) (result map[string]interface{}) {
	var unmarshalled interface{}
	err := json.Unmarshal(byteValue, &unmarshalled)
	assert.NilError(t, err)

	return unmarshalled.(map[string]interface{})
}
