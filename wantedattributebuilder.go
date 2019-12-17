package yoti

import (
	"encoding/json"
)

type constraintInterface interface {
	MarshalJSON() ([]byte, error)
	isConstraint() bool // This function is not used but makes inheritance explicit
}

// WantedAttributeBuilder generates the payload for specifying a single wanted
// attribute as part of a dynamic scenario
type WantedAttributeBuilder struct {
	attr WantedAttribute
}

// WantedAttribute represents a wanted attribute in a dynamic sharing policy
type WantedAttribute struct {
	name               string
	derivation         string
	constraints        []constraintInterface
	acceptSelfAsserted bool
}

// New initialises the internal state of a WantedAttributeBuilder so that
// it can be used
func (builder *WantedAttributeBuilder) New() *WantedAttributeBuilder {
	builder.attr.name = ""
	builder.attr.derivation = ""
	builder.attr.constraints = make([]constraintInterface, 0)
	builder.attr.acceptSelfAsserted = false
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

// WithConstraint adds a constraint to a wanted attribute
func (builder *WantedAttributeBuilder) WithConstraint(constraint constraintInterface) *WantedAttributeBuilder {
	builder.attr.constraints = append(builder.attr.constraints, constraint)
	return builder
}

// WithAcceptSelfAsserted enables self-asserted user details such as from Aadhar
func (builder *WantedAttributeBuilder) WithAcceptSelfAsserted(accept bool) *WantedAttributeBuilder {
	builder.attr.acceptSelfAsserted = accept
	return builder
}

// Build generates the wanted attribute's specification
func (builder *WantedAttributeBuilder) Build() (WantedAttribute, error) {
	err := notEmpty(builder.attr.name, "Wanted attribute names must not be empty")
	return builder.attr, err
}

// MarshalJSON returns the JSON encoding
func (attr *WantedAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Name               string                `json:"name,omitempty"`
		Derivation         string                `json:"derivation,omitempty"`
		Constraints        []constraintInterface `json:"constraints,omitempty"`
		AcceptSelfAsserted bool                  `json:"accept_self_asserted,omitempty"`
	}{
		Name:               attr.name,
		Derivation:         attr.derivation,
		Constraints:        attr.constraints,
		AcceptSelfAsserted: attr.acceptSelfAsserted,
	})
}
