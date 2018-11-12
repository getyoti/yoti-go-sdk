package attribute

import (
	"errors"

	"github.com/getyoti/yoti-go-sdk/anchor"
	"github.com/getyoti/yoti-go-sdk/yotiprotoattr_v3"
)

//ImageAttribute is a Yoti attribute which returns an image as its value
type ImageAttribute struct {
	*yotiprotoattr_v3.Attribute
	Value     *Image
	anchors   []*anchor.Anchor
	sources   []*anchor.Anchor
	verifiers []*anchor.Anchor
}

//NewImage creates a new Image attribute
func NewImage(a *yotiprotoattr_v3.Attribute) (*ImageAttribute, error) {
	var imageType string

	switch a.ContentType {
	case yotiprotoattr_v3.ContentType_JPEG:
		imageType = ImageTypeJpeg

	case yotiprotoattr_v3.ContentType_PNG:
		imageType = ImageTypePng

	default:
		return nil, errors.New("Cannot create ImageAttribute with unsupported type")
	}

	parsedAnchors := anchor.ParseAnchors(a.Anchors)

	return &ImageAttribute{
		Attribute: &yotiprotoattr_v3.Attribute{
			Name:        a.Name,
			ContentType: a.ContentType,
		},
		Value: &Image{
			Data: a.Value,
			Type: imageType,
		},
		anchors:   parsedAnchors,
		sources:   anchor.GetSources(parsedAnchors),
		verifiers: anchor.GetVerifiers(parsedAnchors),
	}, nil
}

// Anchors are the metadata associated with an attribute. They describe
// how an attribute has been provided to Yoti (SOURCE Anchor) and how
// it has been verified (VERIFIER Anchor).
func (a *ImageAttribute) Anchors() []*anchor.Anchor {
	return a.anchors
}

// Sources returns the anchors which identify how and when an attribute value was acquired.
func (a *ImageAttribute) Sources() []*anchor.Anchor {
	return a.sources
}

// Verifiers returns the anchors which identify how and when an attribute value was verified by another provider.
func (a *ImageAttribute) Verifiers() []*anchor.Anchor {
	return a.verifiers
}
