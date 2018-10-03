package attribute

import (
	"encoding/base64"

	"github.com/getyoti/yoti-go-sdk/anchor"
)

//Image is a Yoti attribute which returns an image as its value
type Image struct {
	Name    string
	Value   []byte
	Type    AttrType
	Anchors []*anchor.Anchor
}

//NewImage creates a new Image attribute
func NewImage(a *Attribute) *Image {
	return &Image{
		Name:    a.Name,
		Value:   a.Value,
		Type:    a.Type,
		Anchors: a.Anchors,
	}
}

// Base64URL is the Image encoded as a base64 URL
func (imageAttribute *Image) Base64URL() (string, error) {
	base64EncodedImage := base64.StdEncoding.EncodeToString(imageAttribute.Value)
	contentType := GetMIMEType(imageAttribute.Type)

	return "data:" + contentType + ";base64;," + base64EncodedImage, nil
}
