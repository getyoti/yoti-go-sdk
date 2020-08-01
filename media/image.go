package media

import (
	"fmt"
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
		MimeType: fmt.Sprintf("image/%s", image.Type),
		Data:     image.Data,
	}

	return mediaValue.Base64URL()
}
