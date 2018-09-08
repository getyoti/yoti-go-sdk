package yoti

import (
	"encoding/base64"
	"log"
	"strconv"

	"github.com/getyoti/yoti-go-sdk/anchor"
	"github.com/getyoti/yoti-go-sdk/attribute"
	"github.com/getyoti/yoti-go-sdk/yotiprotoattr_v3"
)

func bytesToUtf8(bytes []byte) string {
	return string(bytes)
}

func bytesToBase64(bytes []byte) string {
	return base64.StdEncoding.EncodeToString(bytes)
}

/* UrlSafe Base64 uses '-' and '_' instead of '+' and '/' respectively so it can be passed
 * as a url parameter without extra encoding.
 */
func bytesToUrlsafeBase64(bytes []byte) string {
	return base64.URLEncoding.EncodeToString(bytes)
}

func utfToBytes(utf8 string) []byte {
	return []byte(utf8)
}

func base64ToBytes(base64Str string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(base64Str)
}

/* UrlSafe Base64 uses '-' and '_' instead of '+' and '/' respectively so it can be passed
 * as a url parameter without extra encoding.
 */
func urlSafeBase64ToBytes(urlSafeBase64 string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(urlSafeBase64)
}

func bytesToBool(bytes []byte) (result bool, err error) {
	result, err = strconv.ParseBool(string(bytes))
	if err != nil {
		log.Printf(
			"Unable to parse bytes to bool. Error: %s", err)
		return false, err
	}

	return result, nil
}

func convertAttribute(protoAttribute *yotiprotoattr_v3.Attribute) *attribute.Attribute {
	processedAnchors := anchor.ParseAnchors(protoAttribute.Anchors)

	var attrType attribute.AttrType

	switch protoAttribute.ContentType {
	case yotiprotoattr_v3.ContentType_STRING:
		attrType = attribute.AttrTypeString

	case yotiprotoattr_v3.ContentType_DATE:
		attrType = attribute.AttrTypeTime

	case yotiprotoattr_v3.ContentType_JPEG:
		attrType = attribute.AttrTypeJPEG

	case yotiprotoattr_v3.ContentType_PNG:
		attrType = attribute.AttrTypePNG

	case yotiprotoattr_v3.ContentType_JSON:
		attrType = attribute.AttrTypeJSON

	case yotiprotoattr_v3.ContentType_UNDEFINED:
	default:
		attrType = attribute.AttrTypeInterface
	}

	return &attribute.Attribute{
		Name:    protoAttribute.Name,
		Value:   protoAttribute.Value,
		Type:    attrType,
		Anchors: processedAnchors,
	}
}
