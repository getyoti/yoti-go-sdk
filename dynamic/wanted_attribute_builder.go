package dynamic

import (
	"errors"
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
	Name               string                `json:"name"`
	Derivation         string                `json:"derivation,omitempty"`
	Constraints        []constraintInterface `json:"constraints,omitempty"`
	AcceptSelfAsserted bool                  `json:"accept_self_asserted"`
}

// WithName sets the name of the wanted attribute
func (builder *WantedAttributeBuilder) WithName(name string) *WantedAttributeBuilder {
	builder.attr.Name = name
	return builder
}

// WithDerivation sets the derivation
func (builder *WantedAttributeBuilder) WithDerivation(derivation string) *WantedAttributeBuilder {
	builder.attr.Derivation = derivation
	return builder
}

// WithConstraint adds a constraint to a wanted attribute
func (builder *WantedAttributeBuilder) WithConstraint(constraint constraintInterface) *WantedAttributeBuilder {
	builder.attr.Constraints = append(builder.attr.Constraints, constraint)
	return builder
}

// WithAcceptSelfAsserted allows self-asserted user details, such as those from Aadhar
func (builder *WantedAttributeBuilder) WithAcceptSelfAsserted(accept bool) *WantedAttributeBuilder {
	builder.attr.AcceptSelfAsserted = accept
	return builder
}

// Build generates the wanted attribute's specification
func (builder *WantedAttributeBuilder) Build() (WantedAttribute, error) {
	if builder.attr.Constraints == nil {
		builder.attr.Constraints = make([]constraintInterface, 0)
	}

	var err error
	if len(builder.attr.Name) == 0 {
		err = errors.New("Wanted attribute names must not be empty")
	}

	return builder.attr, err
}
