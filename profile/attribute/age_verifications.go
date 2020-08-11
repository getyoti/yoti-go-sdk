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

// NewAgeVerification constructs an AgeVerification from a protobuffer
func NewAgeVerification(attr *yotiprotoattr.Attribute) (verification AgeVerification, err error) {
	split := strings.Split(attr.Name, ":")
	verification.Age, err = strconv.Atoi(split[1])
	verification.CheckType = split[0]

	if string(attr.Value) == "true" {
		verification.Result = true
	} else {
		verification.Result = false
	}

	verification.Attribute = attr

	return
}
