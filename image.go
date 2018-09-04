package yoti

import "encoding/base64"

//Deprecated: Will be removed in 3.0.0. GetContentType returns the MIME type of this piece of Yoti user information. For more information see:
// https://en.wikipedia.org/wiki/Media_type
func (image *Image) GetContentType() string {
	switch image.Type {
	case AttrTypeJPEG:
		return "image/jpeg"

	case AttrTypePNG:
		return "image/png"

	default:
		return ""
	}
}

//Deprecated: Will be removed in 3.0.0. URL Image encoded in a base64 URL
func (image *Image) URL() string {
	base64EncodedImage := base64.StdEncoding.EncodeToString(image.Data)
	return "data:" + image.GetContentType() + ";base64;," + base64EncodedImage
}
