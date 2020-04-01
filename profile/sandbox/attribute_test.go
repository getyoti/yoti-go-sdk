package sandbox

import (
	"fmt"
	"testing"
	"time"

	"gotest.tools/v3/assert"
)

func ExampleAttribute_WithAnchor() {
	time.Local = time.UTC
	attribute := Attribute{}.WithAnchor(SourceAnchor("", time.Unix(1234567890, 0), ""))
	fmt.Print(attribute)
	// Output: {    [{SOURCE   2009-02-13 23:31:30 +0000 UTC}]}
}

func TestAttribute_WithName(t *testing.T) {
	attribute := Attribute{}.WithName("attribute_name")

	assert.Equal(t, attribute.Name, "attribute_name")
}

func TestAttribute_WithValue(t *testing.T) {
	attribute := Attribute{}.WithValue("Value")

	assert.Equal(t, attribute.Value, "Value")
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
