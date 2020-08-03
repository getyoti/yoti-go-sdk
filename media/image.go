package media

// ImageType is the MIME type of the image
type ImageType string

const (
	// ImageTypeJpeg JPEG format
	ImageTypeJpeg ImageType = "image/jpeg"

	// ImageTypePng PNG format
	ImageTypePng ImageType = "image/png"
)

// Image format of the image and the image data
type Image struct {
	Type ImageType
	Data []byte
}

// Base64URL is the Image encoded as a base64 URL
func (i *Image) Base64URL() string {
	return base64URL(string(i.Type), i.Data)
}
