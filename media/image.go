package media

type MIMEType string

const (
	// ImageTypeJPEG JPEG format
	ImageTypeJPEG MIMEType = "image/jpeg"

	// ImageTypePNG PNG format
	ImageTypePNG MIMEType = "image/png"
)

// PNGImage holds the binary data of a PNG image.
type PNGImage []byte

func (i PNGImage) Base64URL() string {
	return base64URL(i.MIMEType(), i)
}

func (PNGImage) MIMEType() MIMEType {
	return ImageTypePNG
}

// JPEGImage holds the binary data of a JPEG image.
type JPEGImage []byte

func (i JPEGImage) Base64URL() string {
	return base64URL(i.MIMEType(), i)
}

func (JPEGImage) MIMEType() MIMEType {
	return ImageTypeJPEG
}
