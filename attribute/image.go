package attribute

import (
	"encoding/base64"
)

//Image is a Yoti attribute which returns an image as its value
type Image struct {
	Attribute
	Value []byte
}

//NewImage creates a new Image attribute
func NewImage(a *Attribute) *Image {
	return &Image{
		Attribute: Attribute{
			Name:    a.Name,
			Type:    a.Type,
			Anchors: a.Anchors,
		},
		Value: a.Value, //TODO: why is this bytes?
	}
}

// Base64URL is the Image encoded as a base64 URL
func (imageAttribute *Image) Base64URL() (string, error) {
	base64EncodedImage := base64.StdEncoding.EncodeToString(imageAttribute.Value)
	contentType := GetMIMEType(imageAttribute.Type)

	return "data:" + contentType + ";base64;," + base64EncodedImage, nil
}
