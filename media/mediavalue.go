package media

import (
	"encoding/base64"
	"fmt"
)

// Media holds a piece of binary data.
type Media interface {
	// Base64URL is the media encoded as a base64 URL.
	Base64URL() string

	// MIMETYPE returns the media's MIME type.
	MIMEType() MIMEType

	// Data returns the media's raw data.
	Data() []byte
}

func base64URL(mimeType MIMEType, data []byte) string {
	base64EncodedImage := base64.StdEncoding.EncodeToString(data)
	return fmt.Sprintf("data:%s;base64,%s", mimeType, base64EncodedImage)
}
