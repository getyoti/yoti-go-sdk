package sandbox

import "strconv"

// Attribute describes an attribute on a sandbox profile
type Attribute struct {
	Name       string   `json:"name"`
	Value      string   `json:"value"`
	Derivation string   `json:"derivation"`
	Optional   string   `json:"optional"`
	Anchors    []Anchor `json:"anchors"`
}

// Derivation is a builder for derivation strings
type Derivation struct {
	value string
}

// WithName sets the value of a Sandbox Attribute
func (attr Attribute) WithName(name string) Attribute {
	attr.Name = name
	return attr
}

// WithAnchor sets the value of a Sandbox Attribute
func (attr Attribute) WithValue(value string) Attribute {
	attr.Value = value
	return attr
}

// WithAnchor sets the Anchor of a Sandbox Attribute
func (attr Attribute) WithAnchor(anchor Anchor) Attribute {
	attr.Anchors = append(attr.Anchors, anchor)
	return attr
}

// ToString returns the string representation for a derivation
func (derivation Derivation) ToString() string {
	return derivation.value
}

// AgeOver builds an age over age derivation
func (derivation Derivation) AgeOver(age int) Derivation {
	derivation.value = "age_over:" + strconv.Itoa(age)
	return derivation
}

// AgeUnder builds an age under age derivation
func (derivation Derivation) AgeUnder(age int) Derivation {
	derivation.value = "age_under:" + strconv.Itoa(age)
	return derivation
}
