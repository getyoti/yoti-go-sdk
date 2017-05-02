package yoti

import "encoding/base64"

type ImageType int

const (
	ImageType_Jpeg ImageType = 1 + iota
	ImageType_Png
)

type Image struct {
	Type ImageType
	Data []byte
}

func (image *Image) GetContentType() string {
	switch image.Type {
	case ImageType_Jpeg:
		return "image/jpeg"

	case ImageType_Png:
		return "image/png"

	default:
		return ""
	}
}

func (image *Image) URL() string {
	base64EncodedImage := base64.StdEncoding.EncodeToString(image.Data)
	return "data:" + image.GetContentType() + ";base64;," + base64EncodedImage
}
