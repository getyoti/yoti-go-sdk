package attribute

import (
	"encoding/json"
)

// Definition contains information about the attribute(s) issued by a third party.
type Definition struct {
	name string
}

// Name of the attribute to be issued.
func (a Definition) Name() string {
	return a.name
}

// MarshalJSON returns encoded json
func (a Definition) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Name string `json:"name"`
	}{
		Name: a.name,
	})
}

// NewAttributeDefinition returns a new AttributeDefinition
func NewAttributeDefinition(s string) Definition {
	return Definition{
		name: s,
	}
}
