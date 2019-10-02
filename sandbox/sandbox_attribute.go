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

type Derivation struct {
	value string
}

// AddAnchor adds an Anchor to a Sandbox Attribute
func (attr *Attribute) AddAnchor(anchor Anchor) *Attribute {
	attr.Anchors = append(attr.Anchors, anchor)
	return attr
}

func (derivation Derivation) ToString() string {
	return derivation.value
}

func (derivation Derivation) AgeOver(age int) Derivation {
	derivation.value = "age_over:" + strconv.Itoa(age)
	return derivation
}

func (derivation Derivation) AgeUnder(age int) Derivation {
	derivation.value = "age_under:" + strconv.Itoa(age)
	return derivation
}
