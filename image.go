package yoti

type ImageType int

const (
	ImageType_Jpeg ImageType = 1 + iota
	ImageType_Png
)

type Image struct {
	Type ImageType
	Data []byte
	URL  string
}

func (image Image) GetContentType() string {
	switch image.Type {
	case ImageType_Jpeg:
		return "image/jpeg"

	case ImageType_Png:
		return "image/png"

	default:
		return ""
	}
}
