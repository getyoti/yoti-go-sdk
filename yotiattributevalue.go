package yoti

import "log"

//AttributeType format of the attribute
type AttributeType int

const (
	//AttributeTypeDate date format
	AttributeTypeDate AttributeType = 1 + iota
	//AttributeTypeText text format
	AttributeTypeText
	//AttributeTypeJPEG JPEG format
	AttributeTypeJPEG
	//AttributeTypePNG PNG fornmat
	AttributeTypePNG
	//AttributeTypeJSON JSON fornmat
	AttributeTypeJSON
)

// AttributeValue represents a small piece of information about a Yoti user such as a photo of the user or the
// user's date of birth.
type AttributeValue struct {
	// Type represents the format of the piece of user data, whether it is a date, a piece of text or a picture
	//
	// Note the potential values for this variable are stored in constants with names beginning with
	// 'AttributeType'. These include:
	// 	yoti.AttributeTypeDate
	// 	yoti.AttributeTypeText
	// 	yoti.AttributeTypeJPEG
	// 	yoti.AttributeTypePNG
	// 	yoti.AttributeTypeJSON
	Type  AttributeType
	Value []byte
}

// GetContentType returns the MIME type of this piece of Yoti user information. For more information see:
// https://en.wikipedia.org/wiki/Media_type
func (val AttributeValue) GetContentType() (result string) {
	switch val.Type {
	case AttributeTypeDate:
		result = "text/plain; charset=UTF-8"

	case AttributeTypeText:
		result = "text/plain; charset=UTF-8"

	case AttributeTypeJPEG:
		result = "image/jpeg"

	case AttributeTypePNG:
		result = "image/png"

	case AttributeTypeJSON:
		result = "application/json; charset=UTF-8"

	default:
		log.Printf("Unable to find a matching MIME type for value type %q", val.Type)
	}
	return
}
