package attribute

import (
	"log"

	"github.com/getyoti/yoti-go-sdk/yotiprotoattr_v3"
)

var mimeTypeMap = map[yotiprotoattr_v3.ContentType]string{
	yotiprotoattr_v3.ContentType_DATE:   "text/plain; charset=UTF-8",
	yotiprotoattr_v3.ContentType_STRING: "text/plain; charset=UTF-8",
	yotiprotoattr_v3.ContentType_JPEG:   "image/jpeg",
	yotiprotoattr_v3.ContentType_PNG:    "image/png",
	yotiprotoattr_v3.ContentType_JSON:   "application/json; charset=UTF-8",
}

// GetMIMEType returns the MIME type of this piece of Yoti user information. For more information see:
// https://en.wikipedia.org/wiki/Media_type
func GetMIMEType(attributeType yotiprotoattr_v3.ContentType) (result string) {
	if val, ok := mimeTypeMap[attributeType]; ok {
		return val
	}

	log.Printf("Unable to find a matching MIME type for value type %q", attributeType.String())
	return
}
