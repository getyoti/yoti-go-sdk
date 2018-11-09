package attribute

import (
	"encoding/json"
	"fmt"

	"github.com/getyoti/yoti-go-sdk/anchor"
	"github.com/getyoti/yoti-go-sdk/yotiprotoattr_v3"
)

//JSONAttribute is a Yoti attribute which returns an interface as its value
type JSONAttribute struct {
	*yotiprotoattr_v3.Attribute // Value returns the value of a JSON attribute in the form of an interface
	Value                       interface{}
	Anchors                     []*anchor.Anchor
	Sources                     []*anchor.Anchor
	Verifiers                   []*anchor.Anchor
}

//NewJSON creates a new JSON attribute
func NewJSON(a *yotiprotoattr_v3.Attribute) (*JSONAttribute, error) {
	interfaceValue, err := UnmarshallJSON(a.Value)
	if err != nil {
		err = fmt.Errorf("Unable to parse JSON value: %q. Error: %q", a.Value, err)
		return nil, err
	}

	parsedAnchors := anchor.ParseAnchors(a.Anchors)

	return &JSONAttribute{
		Attribute: &yotiprotoattr_v3.Attribute{
			Name:        a.Name,
			ContentType: a.ContentType,
		},
		Value:     interfaceValue,
		Anchors:   parsedAnchors,
		Sources:   anchor.GetSources(parsedAnchors),
		Verifiers: anchor.GetVerifiers(parsedAnchors),
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
