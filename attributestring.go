package yoti

import "log"

//AttributeString is a Yoti attribute which returns a string as its value
type AttributeString struct {
	attribute
}

func newAttributeString(byteValue []byte, anchors []Anchor, name string, attrType AttrType) (result AttributeString) {
	if attrType != AttrTypeString {
		log.Printf("Cannot create string attribute with non-string type: %q", attrType.String())
		return
	}

	as := attribute{
		anchors: anchors,
		name:    name,
	}

	as.Type = attrType
	as.Value = byteValue

	return AttributeString{
		attribute: as,
	}
}

// Anchors are the metadata associated with an attribute. They describe how an attribute has been provided
// to Yoti (SOURCE Anchor) and how it has been verified (VERIFIER Anchor)
func (as AttributeString) Anchors() []Anchor {
	return as.anchors
}

// AttrValue represents the value associated with a Yoti Attribute: the attribute type and the byte value
func (as AttributeString) AttrValue() AttrValue {
	return as.attribute.AttrValue
}

// Name is the name of the attribute
func (as AttributeString) Name() string {
	return as.attribute.name
}

// Value returns the value of an attribute in the form of byte slice
func (as AttributeString) Value() []byte {
	return as.attribute.Value
}

func (as AttributeString) String() string {
	return string(as.Value())
}
