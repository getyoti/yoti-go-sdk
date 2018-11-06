package attribute

//ImageType Image format
type ImageType int

const (
	//ImageTypeJpeg JPEG format
	ImageTypeJpeg ImageType = 1 + iota
	//ImageTypePng PNG format
	ImageTypePng
	//ImageTypeOther Other image formats
	ImageTypeOther
)

//Image format of the image and the image data
type Image struct {
	Type ImageType
	Data []byte
}
