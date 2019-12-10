package attribute

import (
	"github.com/getyoti/yoti-go-sdk/v2/anchor"
)

// Details is embedded in each attribute for fields common to all
// attributes
type Details struct {
	name        string
	contentType string
	anchors     []*anchor.Anchor
}

// Name gets the attribute name
func (a Details) Name() string {
	return a.name
}

// ContentType gets the attribute's content type description
func (a Details) ContentType() string {
	return a.contentType
}

// Anchors are the metadata associated with an attribute. They describe
// how an attribute has been provided to Yoti (SOURCE Anchor) and how
// it has been verified (VERIFIER Anchor).
func (a Details) Anchors() []*anchor.Anchor {
	return a.anchors
}

// Sources returns the anchors which identify how and when an attribute value
// was acquired.
func (a Details) Sources() []*anchor.Anchor {
	return anchor.GetSources(a.anchors)
}

// Verifiers returns the anchors which identify how and when an attribute value
// was verified by another provider.
func (a Details) Verifiers() []*anchor.Anchor {
	return anchor.GetVerifiers(a.anchors)
}
