package media

const (
	// ImageTypeJpeg JPEG format
	ImageTypeJpeg string = "image/jpeg"

	// ImageTypePng PNG format
	ImageTypePng string = "image/png"
)

// Image format of the image and the image data
type Image struct {
	Value
}

// NewImage creates a new image with the specified MIME type and binary data
func NewImage(MIMEType string, data []byte) *Image {
	return &Image{
		Value: Value{
			MIMEType: MIMEType,
			Data:     data,
		}}
}

// NewJpegImage creates a new JPEG image with the specified binary data
func NewJpegImage(data []byte) *Image {
	return NewImage(ImageTypeJpeg, data)
}

// NewPngImage creates a new PNG image with the specified binary data
func NewPngImage(data []byte) *Image {
	return NewImage(ImageTypePng, data)
}

// Base64URL is the Image encoded as a base64 URL
func (i *Image) Base64URL() string {
	return i.Value.Base64URL()
}
