package sandbox

import (
	"strings"

	"github.com/getyoti/yoti-go-sdk/v3/media"
)

// DocumentImages describes a Document Images attribute on a sandbox profile
type DocumentImages struct {
	Images []*media.Image
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
	pngImage := media.NewPNGImage(imageContent)

	d.Images = append(d.Images, pngImage)

	return d
}

// WithJpegImage adds a JPEG image to the slice of document images
func (d DocumentImages) WithJpegImage(imageContent []byte) DocumentImages {
	jpegImage := media.NewJPEGImage(imageContent)

	d.Images = append(d.Images, jpegImage)

	return d
}
