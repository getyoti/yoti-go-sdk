package attribute

import (
	"encoding/base64"
	"log"
)

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

var mimeTypeMap = map[ImageType]string{
	ImageTypeJpeg: "image/jpeg",
	ImageTypePng:  "image/png",
}

// GetMIMEType returns the MIME type of this piece of Yoti user information. For more information see:
// https://en.wikipedia.org/wiki/Media_type
func GetMIMEType(imageType ImageType) (result string) {
	if val, ok := mimeTypeMap[imageType]; ok {
		return val
	}

	log.Printf("Unable to find a matching MIME type for value type %q", imageType)
	return
}

// Base64URL is the Image encoded as a base64 URL
func (image *Image) Base64URL() (string, error) {
	base64EncodedImage := base64.StdEncoding.EncodeToString(image.Data)
	contentType := GetMIMEType(image.Type)

	return "data:" + contentType + ";base64;," + base64EncodedImage, nil
}
