package media

const (
	// ImageTypeJPEG JPEG format
	ImageTypeJPEG string = "image/jpeg"

	// ImageTypePNG PNG format
	ImageTypePNG string = "image/png"
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

// NewJPEGImage creates a new JPEG image with the specified binary data
func NewJPEGImage(data []byte) *Image {
	return NewImage(ImageTypeJPEG, data)
}

// NewPNGImage creates a new PNG image with the specified binary data
func NewPNGImage(data []byte) *Image {
	return NewImage(ImageTypePNG, data)
}

// Base64URL is the Image encoded as a base64 URL
func (i *Image) Base64URL() string {
	return i.Value.Base64URL()
}
