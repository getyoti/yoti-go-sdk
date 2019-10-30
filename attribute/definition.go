package attribute

import (
	"encoding/json"
)

// AttributeDefinition contains information about the attribute(s) issued by a third party.
type AttributeDefinition struct {
	name string
}

// Name of the attribute to be issued.
func (a AttributeDefinition) Name() string {
	return a.name
}

// MarshalJSON returns encoded json
func (a AttributeDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Name string `json:"name"`
	}{
		Name: a.name,
	})
}

// NewAttributeDefinition returns a new AttributeDefinition
func NewAttributeDefinition(s string) AttributeDefinition {
	return AttributeDefinition{
		name: s,
	}
}
