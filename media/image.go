package media

const (
	// ImageTypeJPEG JPEG format
	ImageTypeJPEG string = "image/jpeg"

	// ImageTypePNG PNG format
	ImageTypePNG string = "image/png"
)

// PNGImage holds the binary data of a PNG image.
type PNGImage []byte

func (i PNGImage) Base64URL() string {
	return base64URL(i.MIME(), i)
}

func (PNGImage) MIME() string {
	return ImageTypePNG
}

func (i PNGImage) Data() []byte {
	return []byte(i)
}

// JPEGImage holds the binary data of a JPEG image.
type JPEGImage []byte

func (i JPEGImage) Base64URL() string {
	return base64URL(i.MIME(), i)
}

func (JPEGImage) MIME() string {
	return ImageTypeJPEG
}

func (i JPEGImage) Data() []byte {
	return []byte(i)
}
