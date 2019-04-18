package attribute

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/getyoti/yoti-go-sdk/v2/yotiprotoattr"
)

func parseValue(contentType yotiprotoattr.ContentType, byteValue []byte) (interface{}, error) {
	switch contentType {
	case yotiprotoattr.ContentType_DATE:
		parsedTime, err := time.Parse("2006-01-02", string(byteValue))

		if err == nil {
			return &parsedTime, nil
		}

		return nil, fmt.Errorf("Unable to parse date value: %q. Error: %q", string(byteValue), err)

	case yotiprotoattr.ContentType_JSON:
		unmarshalledJSON, err := UnmarshallJSON(byteValue)

		if err == nil {
			return unmarshalledJSON, nil
		}

		return nil, fmt.Errorf("Unable to parse JSON value: %q. Error: %q", string(byteValue), err)

	case yotiprotoattr.ContentType_STRING:
		return string(byteValue), nil

	case yotiprotoattr.ContentType_MULTI_VALUE:
		return parseMultiValue(byteValue)

	case yotiprotoattr.ContentType_INT:
		var stringValue = string(byteValue)
		int, err := strconv.Atoi(stringValue)
		if err == nil {
			return int, nil
		}

		return nil, fmt.Errorf("Unable to parse INT value: %q. Error: %q", string(byteValue), err)

	case yotiprotoattr.ContentType_JPEG,
		yotiprotoattr.ContentType_PNG:
		return byteValue, nil

	default:
		log.Printf("Unknown type '%s', attempting to parse it as a String", contentType)
		return string(byteValue), nil
	}
}
