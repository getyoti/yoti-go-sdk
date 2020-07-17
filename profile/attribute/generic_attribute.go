package attribute

import (
	"log"

	"github.com/getyoti/yoti-go-sdk/v3/anchor"
	"github.com/getyoti/yoti-go-sdk/v3/yotiprotoattr"
)

// GenericAttribute is a Yoti attribute which returns a generic value
type GenericAttribute struct {
	attributeDetails
	value interface{}
}

// NewGeneric creates a new generic attribute
func NewGeneric(a *yotiprotoattr.Attribute) *GenericAttribute {
	value, err := parseValue(a.ContentType, a.Value)

	if err != nil {
		log.Printf("Error creating new generic attribute: `%s`", err)
		return nil
	}

	var parsedAnchors = anchor.ParseAnchors(a.Anchors)

	return &GenericAttribute{
		attributeDetails: attributeDetails{
			name:        a.Name,
			contentType: a.ContentType.String(),
			anchors:     parsedAnchors,
		},
		value: value,
	}
}

// Value returns the value of the GenericAttribute as an interface
func (a *GenericAttribute) Value() interface{} {
	return a.value
}
