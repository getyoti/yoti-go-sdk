package attribute

import (
	"github.com/getyoti/yoti-go-sdk/anchor"
)

//ImageAttribute is a Yoti attribute which returns an image as its value
type ImageAttribute struct {
	Name    string
	Value   *Image
	Type    AttrType
	Anchors []*anchor.Anchor
}

//NewImageAttribute creates a new Image attribute
func NewImageAttribute(a *Attribute) *ImageAttribute {
	return &ImageAttribute{
		Name: a.Name,
		Value: &Image{
			Data: a.Value,
		},
		Type:    a.Type,
		Anchors: a.Anchors,
	}
}
