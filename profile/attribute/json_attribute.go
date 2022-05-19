package attribute

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/getyoti/yoti-go-sdk/v3/profile/attribute/anchor"
	"github.com/getyoti/yoti-go-sdk/v3/yotiprotoattr"
)

// JSONAttribute is a Yoti attribute which returns an interface as its value
type JSONAttribute struct {
	attributeDetails
	// value returns the value of a JSON attribute in the form of an interface
	value map[string]interface{}
}

// NewJSON creates a new JSON attribute
func NewJSON(a *yotiprotoattr.Attribute) (*JSONAttribute, error) {
	var interfaceValue map[string]interface{}
	decoder := json.NewDecoder(bytes.NewReader(a.Value))
	decoder.UseNumber()
	err := decoder.Decode(&interfaceValue)
	if err != nil {
		err = fmt.Errorf("unable to parse JSON value: %q. Error: %q", a.Value, err)
		return nil, err
	}

	parsedAnchors := anchor.ParseAnchors(a.Anchors)

	return &JSONAttribute{
		attributeDetails: attributeDetails{
			name:        a.Name,
			contentType: a.ContentType.String(),
			anchors:     parsedAnchors,
			id:          &a.EphemeralId,
		},
		value: interfaceValue,
	}, nil
}

// unmarshallJSON unmarshalls JSON into an interface
func unmarshallJSON(byteValue []byte) (result map[string]interface{}, err error) {
	var unmarshalledJSON map[string]interface{}
	err = json.Unmarshal(byteValue, &unmarshalledJSON)

	if err != nil {
		return nil, err
	}

	return unmarshalledJSON, err
}

// Value returns the value of the JSONAttribute as an interface.
func (a *JSONAttribute) Value() map[string]interface{} {
	return a.value
}
