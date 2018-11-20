package attribute

import (
	"github.com/getyoti/yoti-go-sdk/v2/anchor"
	"github.com/getyoti/yoti-go-sdk/v2/yotiprotoattr"
)

// StringAttribute is a Yoti attribute which returns a string as its value
type StringAttribute struct {
	*yotiprotoattr.Attribute
	value   string
	anchors []*anchor.Anchor
}

// NewString creates a new String attribute
func NewString(a *yotiprotoattr.Attribute) *StringAttribute {
	parsedAnchors := anchor.ParseAnchors(a.Anchors)

	return &StringAttribute{
		Attribute: &yotiprotoattr.Attribute{
			Name:        a.Name,
			ContentType: a.ContentType,
		},
		value:   string(a.Value),
		anchors: parsedAnchors,
	}
}

// Value returns the value of the StringAttribute as a string
func (a *StringAttribute) Value() string {
	return a.value
}

// Anchors are the metadata associated with an attribute. They describe
// how an attribute has been provided to Yoti (SOURCE Anchor) and how
// it has been verified (VERIFIER Anchor).
func (a *StringAttribute) Anchors() []*anchor.Anchor {
	return a.anchors
}

// Sources returns the anchors which identify how and when an attribute value was acquired.
func (a *StringAttribute) Sources() []*anchor.Anchor {
	return anchor.GetSources(a.anchors)
}

// Verifiers returns the anchors which identify how and when an attribute value was verified by another provider.
func (a *StringAttribute) Verifiers() []*anchor.Anchor {
	return anchor.GetVerifiers(a.anchors)
}
