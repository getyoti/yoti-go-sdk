package attribute

import (
	"encoding/base64"

	"github.com/getyoti/yoti-go-sdk/anchor"
	"github.com/getyoti/yoti-go-sdk/yotiprotoattr_v3"
)

//ImageAttribute is a Yoti attribute which returns an image as its value
type ImageAttribute struct {
	*yotiprotoattr_v3.Attribute
	Value   []byte
	Anchors []*anchor.Anchor
}

//NewImage creates a new Image attribute
func NewImage(a *yotiprotoattr_v3.Attribute) *ImageAttribute {
	return &ImageAttribute{
		Attribute: &yotiprotoattr_v3.Attribute{
			Name:        a.Name,
			ContentType: a.ContentType,
		},
		Value:   a.Value,
		Anchors: anchor.ParseAnchors(a.Anchors),
	}
}

// Base64URL is the Image encoded as a base64 URL
func (imageAttribute *ImageAttribute) Base64URL() (string, error) {
	base64EncodedImage := base64.StdEncoding.EncodeToString(imageAttribute.Value)
	contentType := GetMIMEType(imageAttribute.ContentType)

	return "data:" + contentType + ";base64;," + base64EncodedImage, nil
}
