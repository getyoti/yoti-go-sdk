package attribute

import (
	"github.com/getyoti/yoti-go-sdk/anchor"
	"github.com/getyoti/yoti-go-sdk/yotiprotoattr_v3"
)

//ImageAttribute is a Yoti attribute which returns an image as its value
type ImageAttribute struct {
	*yotiprotoattr_v3.Attribute
	Value   *Image
	Anchors []*anchor.Anchor
}

//NewImage creates a new Image attribute
func NewImage(a *yotiprotoattr_v3.Attribute) *ImageAttribute {
	var imageType ImageType

	switch a.ContentType {
	case yotiprotoattr_v3.ContentType_JPEG:
		imageType = ImageTypeJpeg

	case yotiprotoattr_v3.ContentType_PNG:
		imageType = ImageTypePng

	default:
		imageType = ImageTypeOther
	}

	return &ImageAttribute{
		Attribute: &yotiprotoattr_v3.Attribute{
			Name:        a.Name,
			ContentType: a.ContentType,
		},
		Value: &Image{
			Data: a.Value,
			Type: imageType,
		},
		Anchors: anchor.ParseAnchors(a.Anchors),
	}
}
