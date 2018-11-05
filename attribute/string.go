package attribute

import (
	"github.com/getyoti/yoti-go-sdk/anchor"
	"github.com/getyoti/yoti-go-sdk/yotiprotoattr_v3"
)

//String is a Yoti attribute which returns a string as its value
type String struct {
	*yotiprotoattr_v3.Attribute
	Value   string
	Anchors []*anchor.Anchor
}

//NewString creates a new String attribute
func NewString(a *yotiprotoattr_v3.Attribute) *String {
	return &String{
		Attribute: &yotiprotoattr_v3.Attribute{
			Name:        a.Name,
			ContentType: a.ContentType,
		},
		Value:   string(a.Value),
		Anchors: anchor.ParseAnchors(a.Anchors),
	}
}
