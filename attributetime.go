package yoti

import (
	"log"
	"time"
)

//AttributeTime is a Yoti attribute which returns a time as its value
type AttributeTime struct {
	attribute
}

func newAttributeTime(byteValue []byte, anchors []Anchor, name string, attrType AttrType) (result AttributeTime) {
	if attrType != AttrTypeTime {
		log.Printf("Cannot create time attribute with non-time type: %q", attrType.String())
		return
	}

	return AttributeTime{
		attribute: attribute{
			anchors: anchors,
			name:    name,
			AttrValue: AttrValue{
				Type:  attrType,
				Value: byteValue,
			},
		},
	}
}

// Anchors are the metadata associated with an attribute. They describe how an attribute has been provided
// to Yoti (SOURCE Anchor) and how it has been verified (VERIFIER Anchor)
func (at AttributeTime) Anchors() []Anchor {
	return at.anchors
}

// AttrValue represents the value associated with a Yoti Attribute: the attribute type and the byte value
func (at AttributeTime) AttrValue() AttrValue {
	return at.attribute.AttrValue
}

// Name is the name of the attribute
func (at AttributeTime) Name() string {
	return at.attribute.name
}

// Time returns the value of an attribute in time format
func (at AttributeTime) Time() *time.Time {
	parsedTime, err := time.Parse("2006-01-02", string(at.Value))
	if err != nil {
		log.Printf("Unable to parse time value of: %q. Error: %q", at.Value, err)
	}

	return &parsedTime
}
