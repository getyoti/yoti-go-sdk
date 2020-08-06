package media

import (
	"encoding/base64"
	"fmt"
)

type Media interface {
	Base64URL() string
	MIMEType() MIMEType
	Data() []byte
}

func base64URL(mimeType MIMEType, data []byte) string {
	base64EncodedImage := base64.StdEncoding.EncodeToString(data)
	return fmt.Sprintf("data:%s;base64,%s", mimeType, base64EncodedImage)
}
