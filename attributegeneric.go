package yoti

//AttributeGeneric is a Yoti attribute which returns a generic value, which must be explicitly cast
type AttributeGeneric struct {
	attribute
}

func newAttributeGeneric(byteValue []byte, anchors []Anchor, name string, attrType AttrType) (result AttributeGeneric) {
	ag := attribute{
		anchors: anchors,
		name:    name,
	}

	ag.Type = attrType
	ag.Value = byteValue

	return AttributeGeneric{
		attribute: ag,
	}
}

// Anchors are the metadata associated with an attribute. They describe how an attribute has been provided
// to Yoti (SOURCE Anchor) and how it has been verified (VERIFIER Anchor)
func (ab AttributeGeneric) Anchors() []Anchor {
	return ab.anchors
}

// AttrValue represents the value associated with a Yoti Attribute: the attribute type and the byte value
func (ab AttributeGeneric) AttrValue() AttrValue {
	return ab.attribute.AttrValue
}

// Name is the name of the attribute
func (ab AttributeGeneric) Name() string {
	return ab.attribute.name
}

// Value is the string value of the attribute
func (ab AttributeGeneric) Value() string {
	return string(ab.attribute.Value)
}
