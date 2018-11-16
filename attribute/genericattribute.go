package attribute

import (
	"log"
	"time"

	"github.com/getyoti/yoti-go-sdk/anchor"
	"github.com/getyoti/yoti-go-sdk/yotiprotoattr"
)

//GenericAttribute is a Yoti attribute which returns a generic value
type GenericAttribute struct {
	*yotiprotoattr.Attribute
	value     interface{}
	anchors   []*anchor.Anchor
	sources   []*anchor.Anchor
	verifiers []*anchor.Anchor
}

//NewGeneric creates a new generic attribute
func NewGeneric(a *yotiprotoattr.Attribute) *GenericAttribute {
	var value interface{}

	switch a.ContentType {
	case yotiprotoattr.ContentType_DATE:
		parsedTime, err := time.Parse("2006-01-02", string(a.Value))
		if err == nil {
			value = &parsedTime
		} else {
			log.Printf("Unable to parse date value: %q. Error: %q", string(a.Value), err)
		}

	case yotiprotoattr.ContentType_JSON:
		unmarshalledJSON, err := UnmarshallJSON(a.Value)

		if err == nil {
			value = unmarshalledJSON
		} else {
			log.Printf("Unable to parse JSON value: %q. Error: %q", string(a.Value), err)
		}

	case yotiprotoattr.ContentType_STRING:
		value = string(a.Value)

	case yotiprotoattr.ContentType_JPEG,
		yotiprotoattr.ContentType_PNG,
		yotiprotoattr.ContentType_UNDEFINED:
		value = a.Value

	default:
		value = a.Value
	}

	parsedAnchors := anchor.ParseAnchors(a.Anchors)

	return &GenericAttribute{
		Attribute: &yotiprotoattr.Attribute{
			Name:        a.Name,
			ContentType: a.ContentType,
		},
		value:     value,
		anchors:   parsedAnchors,
		sources:   anchor.GetSources(parsedAnchors),
		verifiers: anchor.GetVerifiers(parsedAnchors),
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
	return a.sources
}

// Verifiers returns the anchors which identify how and when an attribute value was verified by another provider.
func (a *GenericAttribute) Verifiers() []*anchor.Anchor {
	return a.verifiers
}
