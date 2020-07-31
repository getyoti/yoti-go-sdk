package media

import (
	"encoding/base64"
	"fmt"
)

// Image format of the image and the image data
type Value struct {
	Type    string
	SubType string
	Data    []byte
}

// getMIMEType returns the MIME type of this piece of Yoti media
func getMIMEType(mediaType, subType string) string {
	return fmt.Sprintf("%s/%s", mediaType, subType)
}

// Base64URL is the Image encoded as a base64 URL
func (m *Value) Base64URL() string {
	base64EncodedImage := base64.StdEncoding.EncodeToString(m.Data)
	contentType := getMIMEType(m.Type, m.SubType)

	return "data:" + contentType + ";base64," + base64EncodedImage
}
