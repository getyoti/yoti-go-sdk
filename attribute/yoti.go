package attribute

import (
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

var mimeTypeMap = map[AttrType]string{
	AttrTypeTime:   "text/plain; charset=UTF-8",
	AttrTypeString: "text/plain; charset=UTF-8",
	AttrTypeJPEG:   "image/jpeg",
	AttrTypePNG:    "image/png",
	AttrTypeJSON:   "application/json; charset=UTF-8",
}

// GetMIMEType returns the MIME type of this piece of Yoti user information. For more information see:
// https://en.wikipedia.org/wiki/Media_type
func GetMIMEType(attributeType AttrType) (result string) {
	if val, ok := mimeTypeMap[attributeType]; ok {
		return val
	}

	log.Printf("Unable to find a matching MIME type for value type %q", attributeType.String())
	return
}
