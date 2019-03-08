package attribute

import (
	"github.com/getyoti/yoti-go-sdk/anchor"
	"github.com/getyoti/yoti-go-sdk/yotiprotoattr"
)

// GenericAttribute is a Yoti attribute which returns a generic value
type GenericAttribute struct {
	*yotiprotoattr.Attribute
	value   interface{}
	anchors []*anchor.Anchor
}

// NewGeneric creates a new generic attribute
func NewGeneric(a *yotiprotoattr.Attribute) *GenericAttribute {
	var value interface{} = parseValue(a.ContentType, a.Value)
	var parsedAnchors []*anchor.Anchor = anchor.ParseAnchors(a.Anchors)

	return &GenericAttribute{
		Attribute: &yotiprotoattr.Attribute{
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
