package media

const (
	// ImageTypeJPEG JPEG format
	ImageTypeJPEG string = "image/jpeg"

	// ImageTypePNG PNG format
	ImageTypePNG string = "image/png"
)

// PNGImage holds the binary data of a PNG image.
type PNGImage []byte

// Base64URL is PNG image encoded as a base64 URL.
func (i PNGImage) Base64URL() string {
	return base64URL(i.MIME(), i)
}

// MIME returns the MIME type for PNG images.
func (PNGImage) MIME() string {
	return ImageTypePNG
}

// Data returns the PNG image as raw bytes.
func (i PNGImage) Data() []byte {
	return []byte(i)
}

// JPEGImage holds the binary data of a JPEG image.
type JPEGImage []byte

// Base64URL is JPEG image encoded as a base64 URL.
func (i JPEGImage) Base64URL() string {
	return base64URL(i.MIME(), i)
}

// MIME returns the MIME type for JPEG images.
func (JPEGImage) MIME() string {
	return ImageTypeJPEG
}

// Data returns the JPEG image as raw bytes.
func (i JPEGImage) Data() []byte {
	return []byte(i)
}
