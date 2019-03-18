package attribute

import (
	"errors"
	"log"

	"github.com/getyoti/yoti-go-sdk/v2/anchor"
	"github.com/getyoti/yoti-go-sdk/v2/yotiprotoattr"
)

// ImageSliceAttribute is a Yoti attribute which returns a slice of images as its value
type ImageSliceAttribute struct {
	*yotiprotoattr.Attribute
	value   []*Image
	anchors []*anchor.Anchor
}

// NewImageSlice creates a new ImageSlice attribute
func NewImageSlice(a *yotiprotoattr.Attribute) (*ImageSliceAttribute, error) {
	if a.ContentType != yotiprotoattr.ContentType_MULTI_VALUE {
		return nil, errors.New("Creating an Image Slice attribute with content types other than MULTI_VALUE is not supported")
	}

	parsedMultiValue, err := parseMultiValue(a.Value)

	if err != nil {
		return nil, err
	}

	var imageSliceValue []*Image
	if parsedMultiValue != nil {
		imageSliceValue = CreateImageSlice(parsedMultiValue)
	}

	return &ImageSliceAttribute{
		Attribute: &yotiprotoattr.Attribute{
			Name:        a.Name,
			ContentType: a.ContentType,
		},
		value:   imageSliceValue,
		anchors: anchor.ParseAnchors(a.Anchors),
	}, nil
}

// CreateImageSlice takes a slice of Items, and converts them into a slice of images
func CreateImageSlice(items []*Item) (result []*Image) {
	for _, item := range items {

		imageValue, err := ParseImageValue(item.GetContentType(), item.GetValue().([]byte))

		if err != nil {
			log.Printf("error parsing image value. ContentType: %s, data: %v, error: %v", item.contentType, item.value.([]byte), err)
			continue
		}

		result = append(result, imageValue)
	}

	return result
}

// Value returns the value of the ImageSliceAttribute as a string
func (a *ImageSliceAttribute) Value() []*Image {
	return a.value
}

// Anchors are the metadata associated with an attribute. They describe
// how an attribute has been provided to Yoti (SOURCE Anchor) and how
// it has been verified (VERIFIER Anchor).
func (a *ImageSliceAttribute) Anchors() []*anchor.Anchor {
	return a.anchors
}

// Sources returns the anchors which identify how and when an attribute value was acquired.
func (a *ImageSliceAttribute) Sources() []*anchor.Anchor {
	return anchor.GetSources(a.anchors)
}

// Verifiers returns the anchors which identify how and when an attribute value was verified by another provider.
func (a *ImageSliceAttribute) Verifiers() []*anchor.Anchor {
	return anchor.GetVerifiers(a.anchors)
}
