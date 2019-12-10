package attribute

import (
	"errors"
	"log"

	"github.com/getyoti/yoti-go-sdk/v2/anchor"
	"github.com/getyoti/yoti-go-sdk/v2/yotiprotoattr"
)

// ImageSliceAttribute is a Yoti attribute which returns a slice of images as its value
type ImageSliceAttribute struct {
	Details
	value []*Image
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
		Details: Details{
			name:        a.Name,
			contentType: a.ContentType.String(),
			anchors:     anchor.ParseAnchors(a.Anchors),
		},
		value: imageSliceValue,
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
