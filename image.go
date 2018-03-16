package yoti

import "encoding/base64"

type imageType int

const (
	//ImageTypeJpeg JPEG format
	ImageTypeJpeg imageType = 1 + iota
	//ImageTypePng PNG format
	ImageTypePng
)

//Image format of the image and the image data
type Image struct {
	Type imageType
	Data []byte
}

// GetContentType returns the MIME type of this piece of Yoti user information. For more information see:
// https://en.wikipedia.org/wiki/Media_type
func (image *Image) GetContentType() string {
	switch image.Type {
	case ImageTypeJpeg:
		return "image/jpeg"

	case ImageTypePng:
		return "image/png"

	default:
		return ""
	}
}

//URL Image encoded in a base64 URL
func (image *Image) URL() string {
	base64EncodedImage := base64.StdEncoding.EncodeToString(image.Data)
	return "data:" + image.GetContentType() + ";base64;," + base64EncodedImage
}
