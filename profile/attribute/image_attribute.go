package attribute

import (
	"errors"

	"github.com/getyoti/yoti-go-sdk/v3/media"
	"github.com/getyoti/yoti-go-sdk/v3/profile/attribute/anchor"
	"github.com/getyoti/yoti-go-sdk/v3/yotiprotoattr"
)

// ImageAttribute is a Yoti attribute which returns an image as its value
type ImageAttribute struct {
	attributeDetails
	value *media.Image
}

// NewImage creates a new Image attribute
func NewImage(a *yotiprotoattr.Attribute) (*ImageAttribute, error) {
	imageValue, err := parseImageValue(a.ContentType, a.Value)
	parsedAnchors := anchor.ParseAnchors(a.Anchors)

	if err != nil {
		return nil, err
	}

	return &ImageAttribute{
		attributeDetails: attributeDetails{
			name:        a.Name,
			contentType: a.ContentType.String(),
			anchors:     parsedAnchors,
		},
		value: imageValue,
	}, nil
}

// Value returns the value of the ImageAttribute as *Image
func (a *ImageAttribute) Value() *media.Image {
	return a.value
}

func parseImageValue(contentType yotiprotoattr.ContentType, byteValue []byte) (*media.Image, error) {
	var imageType string

	switch contentType {
	case yotiprotoattr.ContentType_JPEG:
		imageType = media.ImageTypeJpeg

	case yotiprotoattr.ContentType_PNG:
		imageType = media.ImageTypePng

	default:
		return nil, errors.New("cannot create Image with unsupported type")
	}

	imageValue := media.NewImage(imageType, byteValue)

	return &imageValue, nil
}
