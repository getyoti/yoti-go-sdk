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
	value media.Media
}

// NewImage creates a new Image attribute
func NewImage(a *yotiprotoattr.Attribute) (*ImageAttribute, error) {
	imageValue, err := parseImageValue(a.ContentType, a.Value)
	if err != nil {
		return nil, err
	}

	parsedAnchors := anchor.ParseAnchors(a.Anchors)

	return &ImageAttribute{
		attributeDetails: attributeDetails{
			name:        a.Name,
			contentType: a.ContentType.String(),
			anchors:     parsedAnchors,
			id:          &a.EphemeralId,
		},
		value: imageValue,
	}, nil
}

// Value returns the value of the ImageAttribute as media.Media
func (a *ImageAttribute) Value() media.Media {
	return a.value
}

func parseImageValue(contentType yotiprotoattr.ContentType, byteValue []byte) (media.Media, error) {
	switch contentType {
	case yotiprotoattr.ContentType_JPEG:
		return media.JPEGImage(byteValue), nil

	case yotiprotoattr.ContentType_PNG:
		return media.PNGImage(byteValue), nil

	default:
		return nil, errors.New("cannot create Image with unsupported type")
	}
}
