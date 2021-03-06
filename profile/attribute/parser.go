package attribute

import (
	"fmt"
	"strconv"
	"time"

	"github.com/getyoti/yoti-go-sdk/v3/yotiprotoattr"
)

func parseValue(contentType yotiprotoattr.ContentType, byteValue []byte) (interface{}, error) {
	switch contentType {
	case yotiprotoattr.ContentType_DATE:
		parsedTime, err := time.Parse("2006-01-02", string(byteValue))

		if err == nil {
			return &parsedTime, nil
		}

		return nil, fmt.Errorf("unable to parse date value: %q. Error: %q", string(byteValue), err)

	case yotiprotoattr.ContentType_JSON:
		unmarshalledJSON, err := unmarshallJSON(byteValue)

		if err == nil {
			return unmarshalledJSON, nil
		}

		return nil, fmt.Errorf("unable to parse JSON value: %q. Error: %q", string(byteValue), err)

	case yotiprotoattr.ContentType_STRING:
		return string(byteValue), nil

	case yotiprotoattr.ContentType_MULTI_VALUE:
		return parseMultiValue(byteValue)

	case yotiprotoattr.ContentType_INT:
		var stringValue = string(byteValue)
		intValue, err := strconv.Atoi(stringValue)
		if err == nil {
			return intValue, nil
		}

		return nil, fmt.Errorf("unable to parse INT value: %q. Error: %q", string(byteValue), err)

	case yotiprotoattr.ContentType_JPEG,
		yotiprotoattr.ContentType_PNG:
		return parseImageValue(contentType, byteValue)

	case yotiprotoattr.ContentType_UNDEFINED:
		return string(byteValue), nil

	default:
		return string(byteValue), nil
	}
}
