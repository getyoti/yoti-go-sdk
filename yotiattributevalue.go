package yoti

type AttributeType int

const (
	AttributeType_Date AttributeType = 1 + iota
	AttributeType_Text
	AttributeType_Jpeg
	AttributeType_Png
)

// YotiAttributeValue represents a small piece of information about a Yoti user such as a photo of the user or the
// user's date of birth.
type YotiAttributeValue struct {
	// Type represents the format of the piece of user data, whether it is a date, a piece of text or a picture
	//
	// Note the potential values for this variable are stored in constants with names beginning with
	// 'AttributeType'. These include:
	// 	yoti.AttributeType_Date
	// 	yoti.AttributeType_Text
	// 	yoti.AttributeType_Jpeg
	// 	yoti.AttributeType_Png
	Type  AttributeType
	Value []byte
}

// GetContentType returns the MIME type of this piece of Yoti user information. For more information see:
// https://en.wikipedia.org/wiki/Media_type
func (val YotiAttributeValue) GetContentType() (result string) {
	switch val.Type {
	case AttributeType_Date:
		result = "text/plain; charset=UTF-8"

	case AttributeType_Text:
		result = "text/plain; charset=UTF-8"

	case AttributeType_Jpeg:
		result = "image/jpeg"

	case AttributeType_Png:
		result = "image/png"
	}
	return
}
