package attribute

import (
	"encoding/base64"
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

// GetMIMEType returns the MIME type of this piece of Yoti user information. For more information see:
// https://en.wikipedia.org/wiki/Media_type
func GetMIMEType(imageType string) string {
	return fmt.Sprintf("image/%v", imageType)
}

// Base64URL is the Image encoded as a base64 URL
func (image *Image) Base64URL() string {
	base64EncodedImage := base64.StdEncoding.EncodeToString(image.Data)
	contentType := GetMIMEType(image.Type)

	return "data:" + contentType + ";base64;," + base64EncodedImage
}
