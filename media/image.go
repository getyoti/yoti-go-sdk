package media

import (
	"errors"

	"github.com/getyoti/yoti-go-sdk/v3/yotiprotoattr"
)

const (
	// ImageTypeJpeg JPEG format
	ImageTypeJpeg string = "jpeg"
	// ImageTypePng PNG format
	ImageTypePng string = "png"
)

// Image format of the image and the image data
type Image struct {
	Type string
	Data []byte
}

// Base64URL is the Image encoded as a base64 URL
func (image *Image) Base64URL() string {
	mediaValue := Value{
		Type:    "image",
		SubType: image.Type,
		Data:    image.Data,
	}

	return mediaValue.Base64URL()
}

// ParseImageValue wraps image data into an image struct
func ParseImageValue(contentType yotiprotoattr.ContentType, byteValue []byte) (*Image, error) {
	var imageType string

	switch contentType {
	case yotiprotoattr.ContentType_JPEG:
		imageType = ImageTypeJpeg

	case yotiprotoattr.ContentType_PNG:
		imageType = ImageTypePng

	default:
		return nil, errors.New("cannot create Image with unsupported type")
	}

	return &Image{
		Type: imageType,
		Data: byteValue,
	}, nil
}
