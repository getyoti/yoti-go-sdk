package attribute

import (
	"github.com/getyoti/yoti-go-sdk/v3/profile/attribute/anchor"
)

// attributeDetails is embedded in each attribute for fields common to all
// attributes
type attributeDetails struct {
	name        string
	contentType string
	anchors     []*anchor.Anchor
}

// Name gets the attribute name
func (a attributeDetails) Name() string {
	return a.name
}

// ContentType gets the attribute's content type description
func (a attributeDetails) ContentType() string {
	return a.contentType
}

// Anchors are the metadata associated with an attribute. They describe
// how an attribute has been provided to Yoti (SOURCE Anchor) and how
// it has been verified (VERIFIER Anchor).
func (a attributeDetails) Anchors() []*anchor.Anchor {
	return a.anchors
}

// Sources returns the anchors which identify how and when an attribute value
// was acquired.
func (a attributeDetails) Sources() []*anchor.Anchor {
	return anchor.GetSources(a.anchors)
}

// Verifiers returns the anchors which identify how and when an attribute value
// was verified by another provider.
func (a attributeDetails) Verifiers() []*anchor.Anchor {
	return anchor.GetVerifiers(a.anchors)
}
