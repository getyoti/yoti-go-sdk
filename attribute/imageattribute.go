package attribute

import (
	"github.com/getyoti/yoti-go-sdk/v2/anchor"
	"github.com/getyoti/yoti-go-sdk/v2/yotiprotoattr"
)

// ImageAttribute is a Yoti attribute which returns an image as its value
type ImageAttribute struct {
	Details
	value *Image
}

// NewImage creates a new Image attribute
func NewImage(a *yotiprotoattr.Attribute) (*ImageAttribute, error) {
	imageValue, err := ParseImageValue(a.ContentType, a.Value)
	parsedAnchors := anchor.ParseAnchors(a.Anchors)

	if err != nil {
		return nil, err
	}

	return &ImageAttribute{
		Details: Details{
			name:        a.Name,
			contentType: a.ContentType.String(),
			anchors:     parsedAnchors,
		},
		value: imageValue,
	}, nil
}

// Value returns the value of the ImageAttribute as *Image
func (a *ImageAttribute) Value() *Image {
	return a.value
}
