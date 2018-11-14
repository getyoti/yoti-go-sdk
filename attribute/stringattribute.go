package attribute

import (
	"github.com/getyoti/yoti-go-sdk/anchor"
	"github.com/getyoti/yoti-go-sdk/yotiprotoattr_v3"
)

//StringAttribute is a Yoti attribute which returns a string as its value
type StringAttribute struct {
	*yotiprotoattr_v3.Attribute
	value     string
	anchors   []*anchor.Anchor
	sources   []*anchor.Anchor
	verifiers []*anchor.Anchor
}

//NewString creates a new String attribute
func NewString(a *yotiprotoattr_v3.Attribute) *StringAttribute {
	parsedAnchors := anchor.ParseAnchors(a.Anchors)

	return &StringAttribute{
		Attribute: &yotiprotoattr_v3.Attribute{
			Name:        a.Name,
			ContentType: a.ContentType,
		},
		value:     string(a.Value),
		anchors:   parsedAnchors,
		sources:   anchor.GetSources(parsedAnchors),
		verifiers: anchor.GetVerifiers(parsedAnchors),
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
	return a.sources
}

// Verifiers returns the anchors which identify how and when an attribute value was verified by another provider.
func (a *StringAttribute) Verifiers() []*anchor.Anchor {
	return a.verifiers
}
