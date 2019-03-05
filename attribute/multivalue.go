package attribute

import (
	"log"

	"github.com/getyoti/yoti-go-sdk/yotiprotoattr"
	"github.com/golang/protobuf/proto"
)

// ParseMultiValue unmarshals bytes into a MultiValue struct
func ParseMultiValue(bytes []byte) *yotiprotoattr.MultiValue {
	multiValueStruct := &yotiprotoattr.MultiValue{}

	if err := proto.Unmarshal(bytes, multiValueStruct); err != nil {
		log.Printf("Unable to parse MULTI_VALUE value: %q. Error: %q", string(bytes), err)
		return nil
	}

	return multiValueStruct
}
