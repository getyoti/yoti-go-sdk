package yoti

import (
	"fmt"
)

//JSONAttribute is a Yoti attribute which returns an interface as its value
type JSONAttribute struct {
	Name string
	// Value returns the value of a JSON attribute in the form of an interface
	Value   interface{}
	Type    AttrType
	Anchors []*Anchor
	Err     error
}

func newJSONAttribute(a *Attribute) *JSONAttribute {
	interfaceValue, err := unmarshallJSON(a.Value)
	if err != nil {
		err = fmt.Errorf("Unable to parse JSON value: %q. Error: %q", a.Value, err)
	}

	return &JSONAttribute{
		Name:    a.Name,
		Value:   interfaceValue,
		Type:    AttrTypeJSON,
		Anchors: a.Anchors,
		Err:     err,
	}
}
