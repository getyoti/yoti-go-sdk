package attribute

import (
	"log"
	"time"

	"github.com/getyoti/yoti-go-sdk/anchor"
	"github.com/getyoti/yoti-go-sdk/yotiprotoattr_v3"
)

//GenericAttribute is a Yoti attribute which returns a generic value
type GenericAttribute struct {
	*yotiprotoattr_v3.Attribute
	Value     interface{}
	anchors   []*anchor.Anchor
	sources   []*anchor.Anchor
	verifiers []*anchor.Anchor
}

//NewGeneric creates a new generic attribute
func NewGeneric(a *yotiprotoattr_v3.Attribute) *GenericAttribute {
	var value interface{}

	switch a.ContentType {
	case yotiprotoattr_v3.ContentType_DATE:
		parsedTime, err := time.Parse("2006-01-02", string(a.Value))
		if err == nil {
			value = &parsedTime
		} else {
			log.Printf("Unable to parse date value: %q. Error: %q", string(a.Value), err)
		}

	case yotiprotoattr_v3.ContentType_JSON:
		unmarshalledJSON, err := UnmarshallJSON(a.Value)

		if err == nil {
			value = unmarshalledJSON
		} else {
			log.Printf("Unable to parse JSON value: %q. Error: %q", string(a.Value), err)
		}

	case yotiprotoattr_v3.ContentType_STRING:
		value = string(a.Value)

	case yotiprotoattr_v3.ContentType_JPEG,
		yotiprotoattr_v3.ContentType_PNG,
		yotiprotoattr_v3.ContentType_UNDEFINED:
		value = a.Value

	default:
		value = a.Value
	}

	parsedAnchors := anchor.ParseAnchors(a.Anchors)

	return &GenericAttribute{
		Attribute: &yotiprotoattr_v3.Attribute{
			Name:        a.Name,
			ContentType: a.ContentType,
		},
		Value:     value,
		anchors:   parsedAnchors,
		sources:   anchor.GetSources(parsedAnchors),
		verifiers: anchor.GetVerifiers(parsedAnchors),
	}
}

// Anchors are the metadata associated with an attribute. They describe
// how an attribute has been provided to Yoti (SOURCE Anchor) and how
// it has been verified (VERIFIER Anchor).
func (a *GenericAttribute) Anchors() []*anchor.Anchor {
	return a.anchors
}

// Sources returns the anchors which identify how and when an attribute value was acquired.
func (a *GenericAttribute) Sources() []*anchor.Anchor {
	return a.sources
}

// Verifiers returns the anchors which identify how and when an attribute value was verified by another provider.
func (a *GenericAttribute) Verifiers() []*anchor.Anchor {
	return a.verifiers
}
