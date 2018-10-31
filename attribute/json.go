package attribute

import (
	"encoding/json"
	"fmt"
)

//JSON is a Yoti attribute which returns an interface as its value
type JSON struct {
	Attribute
	// Value returns the value of a JSON attribute in the form of an interface
	Value interface{}
}

//NewJSON creates a new JSON attribute
func NewJSON(a *Attribute) (*JSON, error) {
	interfaceValue, err := UnmarshallJSON(a.Value)
	if err != nil {
		err = fmt.Errorf("Unable to parse JSON value: %q. Error: %q", a.Value, err)
		return nil, err
	}

	return &JSON{
		Attribute: Attribute{
			Name:    a.Name,
			Type:    AttrTypeJSON,
			Anchors: a.Anchors,
		},
		Value: interfaceValue,
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
