package yoti

type attributeType int

const (
	//AttributeTypeDate date format
	AttributeTypeDate attributeType = 1 + iota
	//AttributeTypeText text format
	AttributeTypeText
	//AttributeTypeJpeg JPEG format
	AttributeTypeJpeg
	//AttributeTypePng PNG fornmat
	AttributeTypePng
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
	// 	yoti.AttributeTypeJpeg
	// 	yoti.AttributeTypePng
	Type  attributeType
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

	case AttributeTypeJpeg:
		result = "image/jpeg"

	case AttributeTypePng:
		result = "image/png"
	}
	return
}
