package sandbox

import (
	"strings"

	"github.com/getyoti/yoti-go-sdk/v3/profile/attribute"
)

// DocumentImages describes a Document Images attribute on a sandbox profile
type DocumentImages struct {
	Images []attribute.Image
}

func (d DocumentImages) getValue() string {
	var imageUrls []string

	for _, i := range d.Images {
		imageUrls = append(imageUrls, i.Base64URL())
	}

	return strings.Join(imageUrls, "&")
}

// WithPngImage adds a PNG image to the slice of document images
func (d DocumentImages) WithPngImage(imageContent []byte) DocumentImages {
	pngImage := attribute.Image{
		Type: attribute.ImageTypePng,
		Data: imageContent,
	}

	d.Images = append(d.Images, pngImage)

	return d
}

// WithJpegImage adds a JPEG image to the slice of document images
func (d DocumentImages) WithJpegImage(imageContent []byte) DocumentImages {
	jpegImage := attribute.Image{
		Type: attribute.ImageTypeJpeg,
		Data: imageContent,
	}

	d.Images = append(d.Images, jpegImage)

	return d
}
