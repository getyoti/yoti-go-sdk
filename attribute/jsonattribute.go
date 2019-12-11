package attribute

import (
	"encoding/json"
	"fmt"

	"github.com/getyoti/yoti-go-sdk/v2/anchor"
	"github.com/getyoti/yoti-go-sdk/v2/yotiprotoattr"
)

// JSONAttribute is a Yoti attribute which returns an interface as its value
type JSONAttribute struct {
	attributeDetails
	// Value returns the value of a JSON attribute in the form of an interface
	value interface{}
}

// NewJSON creates a new JSON attribute
func NewJSON(a *yotiprotoattr.Attribute) (*JSONAttribute, error) {
	interfaceValue, err := UnmarshallJSON(a.Value)
	if err != nil {
		err = fmt.Errorf("Unable to parse JSON value: %q. Error: %q", a.Value, err)
		return nil, err
	}

	parsedAnchors := anchor.ParseAnchors(a.Anchors)

	return &JSONAttribute{
		attributeDetails: attributeDetails{
			name:        a.Name,
			contentType: a.ContentType.String(),
			anchors:     parsedAnchors,
		},
		value: interfaceValue,
	}, nil
}

// UnmarshallJSON unmarshalls JSON into an interface
func UnmarshallJSON(byteValue []byte) (result interface{}, err error) {
	var unmarshalledJSON interface{}
	err = json.Unmarshal(byteValue, &unmarshalledJSON)

	if err != nil {
		return nil, err
	}

	return unmarshalledJSON, err
}

// Value returns the value of the JSONAttribute as an interface.
func (a *JSONAttribute) Value() interface{} {
	return a.value
}
