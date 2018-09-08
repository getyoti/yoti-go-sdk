package yoti

import (
	"encoding/base64"
	"log"
	"strconv"

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

func convertAttribute(protoAttribute *yotiprotoattr_v3.Attribute) *Attribute {
	processedAnchors := parseAnchors(protoAttribute.Anchors)

	var attrType AttrType

	switch protoAttribute.ContentType {
	case yotiprotoattr_v3.ContentType_STRING:
		attrType = AttrTypeString

	case yotiprotoattr_v3.ContentType_DATE:
		attrType = AttrTypeTime

	case yotiprotoattr_v3.ContentType_JPEG:
		attrType = AttrTypeJPEG

	case yotiprotoattr_v3.ContentType_PNG:
		attrType = AttrTypePNG

	case yotiprotoattr_v3.ContentType_JSON:
		attrType = AttrTypeJSON

	case yotiprotoattr_v3.ContentType_UNDEFINED:
	default:
		attrType = AttrTypeInterface
	}

	return &Attribute{
		Name:    protoAttribute.Name,
		Value:   protoAttribute.Value,
		Type:    attrType,
		Anchors: processedAnchors,
	}
}
