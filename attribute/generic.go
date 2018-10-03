package attribute

import (
	"github.com/getyoti/yoti-go-sdk/anchor"
)

//Generic is a Yoti attribute which returns a generic value
type Generic struct {
	Name    string
	Value   string
	Data    []byte
	Type    AttrType
	Anchors []*anchor.Anchor
}

//NewGeneric creates a new generic attribute
func NewGeneric(a *Attribute) *Generic {
	return &Generic{
		Name:    a.Name,
		Value:   string(a.Value),
		Data:    a.Value,
		Type:    AttrTypeInterface,
		Anchors: a.Anchors,
	}
}
