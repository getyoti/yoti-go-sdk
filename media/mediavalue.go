package media

import (
	"encoding/base64"
)

// Image format of the image and the image data
type Value struct {
	MimeType string
	Data     []byte
}

// Base64URL is the base64 data URL for the media
func (v *Value) Base64URL() string {
	return base64URL(v.MimeType, v.Data)
}

func base64URL(mimeType string, data []byte) string {
	base64EncodedImage := base64.StdEncoding.EncodeToString(data)
	return "data:" + mimeType + ";base64," + base64EncodedImage
}
