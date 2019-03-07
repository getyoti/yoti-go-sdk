package attribute

import (
	"log"
	"time"

	"github.com/getyoti/yoti-go-sdk/yotiprotoattr"
)

func parseAttribute(contentType yotiprotoattr.ContentType, byteValue []byte) (result interface{}) {
	switch contentType {
	case yotiprotoattr.ContentType_DATE:
		parsedTime, err := time.Parse("2006-01-02", string(byteValue))
		if err == nil {
			result = &parsedTime
		} else {
			log.Printf("Unable to parse date value: %q. Error: %q", string(byteValue), err)
		}

	case yotiprotoattr.ContentType_JSON:
		unmarshalledJSON, err := UnmarshallJSON(byteValue)

		if err == nil {
			result = unmarshalledJSON
		} else {
			log.Printf("Unable to parse JSON value: %q. Error: %q", string(byteValue), err)
		}

	case yotiprotoattr.ContentType_STRING:
		result = string(byteValue)

	case yotiprotoattr.ContentType_MULTI_VALUE:
		result = parseMultiValue(byteValue)

	case yotiprotoattr.ContentType_JPEG,
		yotiprotoattr.ContentType_PNG,
		yotiprotoattr.ContentType_UNDEFINED:
		result = byteValue

	default:
		result = byteValue
	}

	return result
}
