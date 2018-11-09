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
	Anchors   []*anchor.Anchor
	Sources   []*anchor.Anchor
	Verifiers []*anchor.Anchor
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
		Anchors:   parsedAnchors,
		Sources:   anchor.GetSources(parsedAnchors),
		Verifiers: anchor.GetVerifiers(parsedAnchors),
	}
}
