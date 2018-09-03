package yoti

import (
	"log"
)

//AttributeJSON is a Yoti attribute which returns a string as its value
type AttributeJSON struct {
	attribute
}

func newAttributeJSON(byteValue []byte, anchors []Anchor, name string, attrType AttrType) (result AttributeJSON) {
	if attrType != AttrTypeJSON {
		log.Printf("Cannot create JSON attribute with non-JSON type: %q", attrType.String())
		return
	}

	aj := attribute{
		anchors: anchors,
		name:    name,
	}

	aj.Type = attrType
	aj.Value = byteValue

	return AttributeJSON{
		attribute: aj,
	}
}

// Anchors are the metadata associated with an attribute. They describe how an attribute has been provided
// to Yoti (SOURCE Anchor) and how it has been verified (VERIFIER Anchor)
func (aj AttributeJSON) Anchors() []Anchor {
	return aj.anchors
}

// AttrValue represents the value associated with a Yoti Attribute: the attribute type and the byte value
func (aj AttributeJSON) AttrValue() AttrValue {
	return aj.attribute.AttrValue
}

// Name is the name of the attribute
func (aj AttributeJSON) Name() string {
	return aj.attribute.name
}

// Interface returns the value of an attribute in the form of an interface
func (aj AttributeJSON) Interface() interface{} {
	json, err := unmarshallJSON(aj.Value)

	if err != nil {
		log.Printf("Unable to parse JSON value: %q. Error: %q", aj.Value, err)
	}

	return json
}
