package attribute

import (
	"strconv"
	"strings"

	"github.com/getyoti/yoti-go-sdk/v3/yotiprotoattr"
)

// AgeVerification encapsulates the result of a single age verification
// as part of a share
type AgeVerification struct {
	Age       int
	CheckType string
	Result    bool
	Attribute *yotiprotoattr.Attribute
}

// New constructs an AgeVerification from a protobuffer
func (AgeVerification) New(attr *yotiprotoattr.Attribute) (value AgeVerification, err error) {
	split := strings.Split(attr.Name, ":")
	value.Age, err = strconv.Atoi(split[1])
	value.CheckType = split[0]

	if string(attr.Value) == "true" {
		value.Result = true
	} else {
		value.Result = false
	}

	value.Attribute = attr

	return
}
