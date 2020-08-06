package media

import (
	"encoding/base64"
)

// Value information about the media item, it's MIME type and binary data
type Value struct {
	MIMEType string
	Data     []byte
}

// Base64URL is the base64 data URL for the media
func (v *Value) Base64URL() string {
	return base64URL(v.MIMEType, v.Data)
}

func base64URL(MIMEType string, data []byte) string {
	base64EncodedImage := base64.StdEncoding.EncodeToString(data)
	return "data:" + MIMEType + ";base64," + base64EncodedImage
}
