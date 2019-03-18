package attribute

import (
	"fmt"

	"github.com/getyoti/yoti-go-sdk/v2/anchor"
	"github.com/getyoti/yoti-go-sdk/v2/yotiprotoattr"
	"github.com/golang/protobuf/proto"
)

// MultiValueAttribute is a Yoti attribute which returns a multi-valued attribute
type MultiValueAttribute struct {
	*yotiprotoattr.Attribute
	items   []*Item
	anchors []*anchor.Anchor
}

// NewMultiValue creates a new MultiValue attribute
func NewMultiValue(a *yotiprotoattr.Attribute) (*MultiValueAttribute, error) {
	attributeItems, err := parseMultiValue(a.Value)

	if err != nil {
		return nil, err
	}

	return &MultiValueAttribute{
		Attribute: &yotiprotoattr.Attribute{
			Name:        a.Name,
			ContentType: a.ContentType,
		},
		items:   attributeItems,
		anchors: anchor.ParseAnchors(a.Anchors),
	}, nil
}

// parseMultiValue recursively unmarshals and converts Multi Value bytes into a slice of Items
func parseMultiValue(data []byte) ([]*Item, error) {
	var attributeItems []*Item
	protoMultiValueStruct, err := unmarshallMultiValue(data)

	if err != nil {
		return nil, err
	}

	for _, multiValueItem := range protoMultiValueStruct.Values {
		var value *Item
		if multiValueItem.ContentType == yotiprotoattr.ContentType_MULTI_VALUE {
			parsedInnerMultiValueItems, err := parseMultiValue(multiValueItem.Data)

			if err != nil {
				return nil, fmt.Errorf("Unable to parse multi-value data: %v", err)
			}

			value = &Item{
				contentType: yotiprotoattr.ContentType_MULTI_VALUE,
				value:       parsedInnerMultiValueItems,
			}
		} else {
			itemValue, err := parseValue(multiValueItem.ContentType, multiValueItem.Data)

			if err != nil {
				return nil, fmt.Errorf("Unable to parse data within a multi-value attribute. Content type: %q, data: %q, error: %v",
					multiValueItem.ContentType, multiValueItem.Data, err)
			}

			value = &Item{
				contentType: multiValueItem.ContentType,
				value:       itemValue,
			}
		}
		attributeItems = append(attributeItems, value)
	}

	return attributeItems, nil
}

func unmarshallMultiValue(bytes []byte) (*yotiprotoattr.MultiValue, error) {
	multiValueStruct := &yotiprotoattr.MultiValue{}

	if err := proto.Unmarshal(bytes, multiValueStruct); err != nil {
		return nil, fmt.Errorf("Unable to parse MULTI_VALUE value: %q. Error: %q", string(bytes), err)
	}

	return multiValueStruct, nil
}

// Value returns the value of the MultiValueAttribute as a string
func (a *MultiValueAttribute) Value() []*Item {
	return a.items
}

// Anchors are the metadata associated with an attribute. They describe
// how an attribute has been provided to Yoti (SOURCE Anchor) and how
// it has been verified (VERIFIER Anchor).
func (a *MultiValueAttribute) Anchors() []*anchor.Anchor {
	return a.anchors
}

// Sources returns the anchors which identify how and when an attribute value was acquired.
func (a *MultiValueAttribute) Sources() []*anchor.Anchor {
	return anchor.GetSources(a.anchors)
}

// Verifiers returns the anchors which identify how and when an attribute value was verified by another provider.
func (a *MultiValueAttribute) Verifiers() []*anchor.Anchor {
	return anchor.GetVerifiers(a.anchors)
}
