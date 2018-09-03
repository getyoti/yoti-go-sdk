package yoti

// Attribute represents a small piece of information about a Yoti user such as a photo of the user or the
// user's date of birth.
type Attribute interface {
	AttrValue() AttrValue
	// Name is the name of the Yoti attribute
	Name() string
	Anchors() []Anchor
}

type attribute struct {
	AttrValue
	name    string
	anchors []Anchor
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

// AttrValue represents the value associated with a Yoti Attribute.
type AttrValue struct {
	// Type represents the format of the piece of user data, whether it is a date, a piece of text or a picture
	//
	// Note the potential values for this variable are stored in constants with names beginning with
	// 'AttrType'. These include:
	//  yoti.AttrTypeTime
	//  yoti.AttrTypeString
	//  yoti.AttrTypeJPEG
	//  yoti.AttrTypePNG
	//  yoti.AttrTypeJSON
	//  yoti.AttrTypeBool
	//  yoti.AttrTypeInterface
	Type  AttrType
	Value []byte
}
