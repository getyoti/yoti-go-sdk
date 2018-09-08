package attribute

import (
	"github.com/getyoti/yoti-go-sdk/anchor"
)

//GenericAttribute is a Yoti attribute which returns a generic value
type GenericAttribute struct {
	Name    string
	Value   string
	Data    []byte
	Type    AttrType
	Anchors []*anchor.Anchor
}

//NewGenericAttribute creates a new generic attribute
func NewGenericAttribute(a *Attribute) *GenericAttribute {
	return &GenericAttribute{
		Name:    a.Name,
		Value:   string(a.Value),
		Data:    a.Value,
		Type:    AttrTypeInterface,
		Anchors: a.Anchors,
	}
}
