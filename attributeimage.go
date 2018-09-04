package yoti

import (
	"encoding/base64"
	"log"
)

//AttributeImage is a Yoti attribute which returns an image as its value
type AttributeImage struct {
	attribute
}

//Image format of the image and the image data
type Image struct {
	Type AttrType
	Data []byte
}

func newAttributeImage(byteValue []byte, anchors []Anchor, name string, attrType AttrType) (result AttributeImage) {
	if attrType != AttrTypeJPEG && attrType != AttrTypePNG {
		log.Printf("Cannot create image attribute with non-image type: %q", attrType.String())
		return
	}

	return AttributeImage{
		attribute: attribute{
			anchors: anchors,
			name:    name,
			AttrValue: AttrValue{
				Type:  attrType,
				Value: byteValue,
			},
		},
	}
}

// Anchors are the metadata associated with an attribute. They describe how an attribute has been provided
// to Yoti (SOURCE Anchor) and how it has been verified (VERIFIER Anchor)
func (ai AttributeImage) Anchors() []Anchor {
	return ai.anchors
}

// AttrValue represents the value associated with a Yoti Attribute: the attribute type and the byte value
func (ai AttributeImage) AttrValue() AttrValue {
	return ai.attribute.AttrValue
}

// Name is the name of the attribute
func (ai AttributeImage) Name() string {
	return ai.attribute.name
}

// Base64Selfie is the Image encoded as a base64 URL
func (ai AttributeImage) Base64Selfie() (result string) {
	base64EncodedImage := base64.StdEncoding.EncodeToString(ai.Value)

	var imageType string
	if ai.Type == AttrTypeJPEG {
		imageType = "jpeg"
	} else if ai.Type == AttrTypePNG {
		imageType = "png"
	} else {
		log.Printf("Unable to create base64 URL for type: %q", ai.Type.String())
		return
	}

	return "data:image/" + imageType + ";base64;," + base64EncodedImage
}

// Image returns the value of an attribute in the form of a Yoti Image object
func (ai AttributeImage) Image() *Image {
	return &Image{
		Type: ai.Type,
		Data: ai.Value,
	}
}
