package attribute

import (
	"log"

	"github.com/getyoti/yoti-go-sdk/v2/anchor"
	"github.com/getyoti/yoti-go-sdk/v2/yotiprotoattr"
)

// GenericAttribute is a Yoti attribute which returns a generic value
type GenericAttribute struct {
	attribute *yotiprotoattr.Attribute
	value     interface{}
	anchors   []*anchor.Anchor
}

// NewGeneric creates a new generic attribute
func NewGeneric(a *yotiprotoattr.Attribute) *GenericAttribute {
	value, err := parseValue(a.ContentType, a.Value)

	if err != nil {
		log.Printf("Error creating new generic attribute: `%s`", err)
		return nil
	}

	var parsedAnchors []*anchor.Anchor = anchor.ParseAnchors(a.Anchors)

	return &GenericAttribute{
		attribute: &yotiprotoattr.Attribute{
			Name:        a.Name,
			ContentType: a.ContentType,
		},
		value:   value,
		anchors: parsedAnchors,
	}
}

// Value returns the value of the GenericAttribute as an interface
func (a *GenericAttribute) Value() interface{} {
	return a.value
}

// Name returns the name as a string
func (a *GenericAttribute) Name() string {
	return a.attribute.Name
}

// ContentType returns the content type as a string
func (a *GenericAttribute) ContentType() string {
	return a.attribute.ContentType.String()
}

// Anchors are the metadata associated with an attribute. They describe
// how an attribute has been provided to Yoti (SOURCE Anchor) and how
// it has been verified (VERIFIER Anchor).
func (a *GenericAttribute) Anchors() []*anchor.Anchor {
	return a.anchors
}

// Sources returns the anchors which identify how and when an attribute value was acquired.
func (a *GenericAttribute) Sources() []*anchor.Anchor {
	return anchor.GetSources(a.anchors)
}

// Verifiers returns the anchors which identify how and when an attribute value was verified by another provider.
func (a *GenericAttribute) Verifiers() []*anchor.Anchor {
	return anchor.GetVerifiers(a.anchors)
}
