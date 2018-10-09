package attribute

import (
	"fmt"
	"log"

	"github.com/getyoti/yoti-go-sdk/anchor"
)

// Attribute represents a small piece of information about a Yoti user such as a photo of the user or the
// user's date of birth.
type Attribute struct {
	Name    string
	Value   []byte
	Type    AttrType
	Anchors []*anchor.Anchor
}

//AttrType format of the attribute
type AttrType uint

const (
	//AttrTypeTime date format
	AttrTypeTime AttrType = iota + 1
	//AttrTypeString text format
	AttrTypeString
	//AttrTypeJPEG JPEG format
	AttrTypeJPEG
	//AttrTypePNG PNG fornmat
	AttrTypePNG
	//AttrTypeJSON JSON fornmat
	AttrTypeJSON
	//AttrTypeBool Boolean format
	AttrTypeBool
	//AttrTypeInterface generic interface format
	AttrTypeInterface
)

// GetMIMEType returns the MIME type of this piece of Yoti user information. For more information see:
// https://en.wikipedia.org/wiki/Media_type
func GetMIMEType(attributeType fmt.Stringer) (result string) {
	switch attributeType {
	case AttrTypeTime:
		result = "text/plain; charset=UTF-8"

	case AttrTypeString:
		result = "text/plain; charset=UTF-8"

	case AttrTypeJPEG:
		result = "image/jpeg"

	case AttrTypePNG:
		result = "image/png"

	case AttrTypeJSON:
		result = "application/json; charset=UTF-8"

	default:
		log.Printf("Unable to find a matching MIME type for value type %q", attributeType.String())
	}
	return
}
