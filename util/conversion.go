package util

import (
	"encoding/base64"
)

// Base64ToBytes converts a base64 string to bytes
func Base64ToBytes(base64Str string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(base64Str)
}

// UrlSafeBase64ToBytes UrlSafe Base64 uses '-' and '_', instead of '+' and '/' respectively, so it can be passed
// as a url parameter without extra encoding.
func UrlSafeBase64ToBytes(urlSafeBase64 string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(urlSafeBase64)
}
