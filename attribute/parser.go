package attribute

import (
	"fmt"
	"time"

	"github.com/getyoti/yoti-go-sdk/v2/yotiprotoattr"
)

func parseValue(contentType yotiprotoattr.ContentType, byteValue []byte) (interface{}, error) {
	switch contentType {
	case yotiprotoattr.ContentType_DATE:
		parsedTime, err := time.Parse("2006-01-02", string(byteValue))

		if err == nil {
		} else {
			return nil, fmt.Errorf("Unable to parse date value: %q. Error: %q", string(byteValue), err)
		}
		return &parsedTime, nil

	case yotiprotoattr.ContentType_JSON:
		unmarshalledJSON, err := UnmarshallJSON(byteValue)

		if err == nil {
			return unmarshalledJSON, nil
		} else {
			return nil, fmt.Errorf("Unable to parse JSON value: %q. Error: %q", string(byteValue), err)
		}

	case yotiprotoattr.ContentType_STRING:
		return string(byteValue), nil

	case yotiprotoattr.ContentType_MULTI_VALUE:
		return parseMultiValue(byteValue)

	case yotiprotoattr.ContentType_JPEG,
		yotiprotoattr.ContentType_PNG,
		yotiprotoattr.ContentType_UNDEFINED:
		return byteValue, nil

	default:
		return byteValue, nil
	}
}
