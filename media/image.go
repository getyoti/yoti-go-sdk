package media

import "fmt"

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
func (i *Image) Base64URL() string {
	mimeType := fmt.Sprintf("image/%s", i.Type)
	return base64URL(mimeType, i.Data)
}
