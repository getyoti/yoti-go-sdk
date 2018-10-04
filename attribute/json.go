package attribute

import (
	"encoding/json"
	"fmt"

	"github.com/getyoti/yoti-go-sdk/anchor"
)

//JSON is a Yoti attribute which returns an interface as its value
type JSON struct {
	Name string
	// Value returns the value of a JSON attribute in the form of an interface
	Value   interface{}
	Type    AttrType
	Anchors []*anchor.Anchor
}

//NewJSON creates a new JSON attribute
func NewJSON(a *Attribute) (*JSON, error) {
	interfaceValue, err := UnmarshallJSON(a.Value)
	if err != nil {
		err = fmt.Errorf("Unable to parse JSON value: %q. Error: %q", a.Value, err)
		return nil, err
	}

	return &JSON{
		Name:    a.Name,
		Value:   interfaceValue,
		Type:    AttrTypeJSON,
		Anchors: a.Anchors,
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
