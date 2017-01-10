package yoti

import (
	"encoding/base64"
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
