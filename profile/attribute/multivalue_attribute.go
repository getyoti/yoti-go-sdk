package attribute

import (
	"fmt"

	"github.com/getyoti/yoti-go-sdk/v3/profile/attribute/anchor"
	"github.com/getyoti/yoti-go-sdk/v3/yotiprotoattr"
	"github.com/golang/protobuf/proto"
)

// MultiValueAttribute is a Yoti attribute which returns a multi-valued attribute
type MultiValueAttribute struct {
	attributeDetails
	items []*Item
}

// NewMultiValue creates a new MultiValue attribute
func NewMultiValue(a *yotiprotoattr.Attribute) (*MultiValueAttribute, error) {
	attributeItems, err := parseMultiValue(a.Value)

	if err != nil {
		return nil, err
	}

	return &MultiValueAttribute{
		attributeDetails: attributeDetails{
			name:        a.Name,
			contentType: a.ContentType.String(),
			anchors:     anchor.ParseAnchors(a.Anchors),
		},
		items: attributeItems,
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
