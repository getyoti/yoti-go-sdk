package attribute

import (
	"github.com/getyoti/yoti-go-sdk/anchor"
)

//String is a Yoti attribute which returns a string as its value
type String struct {
	Name    string
	Value   string
	Type    AttrType
	Anchors []*anchor.Anchor
}

//NewString creates a new String attribute
func NewString(a *Attribute) *String {
	return &String{
		Name:    a.Name,
		Value:   string(a.Value),
		Type:    AttrTypeString,
		Anchors: a.Anchors,
	}
}
