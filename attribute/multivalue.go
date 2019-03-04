package attribute

import (
	"log"

	"github.com/getyoti/yoti-go-sdk/yotiprotoattr"
	"github.com/golang/protobuf/proto"
)

func parseMultiValue(bytes []byte) (multiValueStruct *yotiprotoattr.MultiValue) {
	if err := proto.Unmarshal(bytes, multiValueStruct); err != nil {
		log.Printf("Unable to parse MULTI_VALUE value: %q. Error: %q", string(bytes), err)
		return nil
	}

	return multiValueStruct
}
