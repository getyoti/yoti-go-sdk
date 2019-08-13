package yoti

import (
	"encoding/json"
)

const ()

// WantedAttributeBuilder generates the payload for specifying a single wanted
// attribute as part of a dynamic scenario
type WantedAttributeBuilder struct {
	attr WantedAttribute
}

// WantedAttribute represents a wanted attribute in a dynamic sharing policy
type WantedAttribute struct {
	name       string
	derivation string
}

// New initialises the internal state of a WantedAttributeBuilder so that
// it can be used
func (builder *WantedAttributeBuilder) New() *WantedAttributeBuilder {
	builder.attr.name = ""
	builder.attr.derivation = ""
	return builder
}

// WithName sets the name of the wanted attribute
func (builder *WantedAttributeBuilder) WithName(name string) *WantedAttributeBuilder {
	builder.attr.name = name
	return builder
}

// WithDerivation sets the derivation
func (builder *WantedAttributeBuilder) WithDerivation(derivation string) *WantedAttributeBuilder {
	builder.attr.derivation = derivation
	return builder
}

// Build generates the wanted attribute's specification
func (builder *WantedAttributeBuilder) Build() WantedAttribute {
	attr := builder.attr
	return attr
}

// MarshalJSON ...
func (attr *WantedAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Name       string `json:"name,omitempty"`
		Derivation string `json:"derivation,omitempty"`
	}{
		Name:       attr.name,
		Derivation: attr.derivation,
	})
}
