package attribute

import (
	"errors"
	"fmt"

	"github.com/getyoti/yoti-go-sdk/v3/media"
	"github.com/getyoti/yoti-go-sdk/v3/profile/attribute/anchor"
	"github.com/getyoti/yoti-go-sdk/v3/yotiprotoattr"
)

// ImageSliceAttribute is a Yoti attribute which returns a slice of images as its value
type ImageSliceAttribute struct {
	attributeDetails
	value []media.Media
}

// NewImageSlice creates a new ImageSlice attribute
func NewImageSlice(a *yotiprotoattr.Attribute) (*ImageSliceAttribute, error) {
	if a.ContentType != yotiprotoattr.ContentType_MULTI_VALUE {
		return nil, errors.New("creating an Image Slice attribute with content types other than MULTI_VALUE is not supported")
	}

	parsedMultiValue, err := parseMultiValue(a.Value)

	if err != nil {
		return nil, err
	}

	var imageSliceValue []media.Media
	if parsedMultiValue != nil {
		imageSliceValue, err = CreateImageSlice(parsedMultiValue)
		if err != nil {
			return nil, err
		}
	}

	return &ImageSliceAttribute{
		attributeDetails: attributeDetails{
			name:        a.Name,
			contentType: a.ContentType.String(),
			anchors:     anchor.ParseAnchors(a.Anchors),
		},
		value: imageSliceValue,
	}, nil
}

// CreateImageSlice takes a slice of Items, and converts them into a slice of images
func CreateImageSlice(items []*Item) (result []media.Media, err error) {
	for _, item := range items {

		switch i := item.Value.(type) {
		case media.PNGImage:
			result = append(result, i)
		case media.JPEGImage:
			result = append(result, i)
		default:
			return nil, fmt.Errorf("unexpected item type %T", i)
		}
	}

	return result, nil
}

// Value returns the value of the ImageSliceAttribute
func (a *ImageSliceAttribute) Value() []media.Media {
	return a.value
}
