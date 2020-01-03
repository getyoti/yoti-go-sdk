package yoti

import (
	"crypto/aes"
	"testing"

	"gotest.tools/assert"
)

func TestCrypto_loadRsaKey_PublicKeyShouldFail(t *testing.T) {
	testPEM := []byte(`-----BEGIN RSA PUBLIC KEY-----
VGVzdCBTdHJpbmc=
-----END RSA PUBLIC KEY-----`)

	_, err := loadRsaKey(testPEM)

	assert.Error(t, err, "Invalid Key: not RSA private key")
}

func TestCrypto_loadRsaKey_InvalidKeyShouldFail(t *testing.T) {
	testPEM := []byte(`-----BEGIN RSA PRIVATE KEY-----
VGVzdCBTdHJpbmc=
-----END RSA PRIVATE KEY-----`)

	_, err := loadRsaKey(testPEM)

	assert.Error(t, err, "Invalid Key: bad RSA private key")
}

func TestCrypto_decipherAes_EmptyCiphertextShouldError(t *testing.T) {
	_, err := decipherAes([]byte{}, []byte{}, []byte{})
	assert.Check(t, err != nil)
}

func TestCrypto_decipherAes_CiphertextNotMatchingBlocksizeShouldError(t *testing.T) {
	_, err := decipherAes(
		make([]byte, 16),
		[]byte{},
		make([]byte, aes.BlockSize+1),
	)
	assert.Check(t, err != nil)
}

func TestCrypto_pkcs7Unpad_InvalidBlocksizeShouldError(t *testing.T) {
	_, err := pkcs7Unpad([]byte{}, -1)
	assert.Error(t, err, "blocksize -1 is not valid for padding removal")
}

func TestCrypto_pkcs7Unpad_EmptyByteArrayShouldError(t *testing.T) {
	_, err := pkcs7Unpad([]byte{}, 1)
	assert.Error(t, err, "Cannot remove padding on empty byte array")
}

func TestCrypto_pkcs7Unpad_CiphertextNotMultipleOfBlocksizeShouldError(t *testing.T) {
	_, err := pkcs7Unpad([]byte{0, 0, 0}, 2)
	assert.Error(t, err, "ciphertext is not a multiple of the block size")
}

func TestCrypto_pkcs7Unpad_CiphertextNotPaddedShouldError(t *testing.T) {
	_, err := pkcs7Unpad([]byte{0xF, 0x0, 0x0}, 3)
	assert.Error(t, err, "ciphertext is not padded with PKCS#7 padding")
}

func TestCrypto_pkcs7Unpad_CiphertextPaddingIncorrect(t *testing.T) {
	_, err := pkcs7Unpad([]byte{0x1, 0x1, 0x1, 0xF, 0x0, 0x0, 0xB, 0xA, 0x7}, 3)
	assert.Error(t, err, "ciphertext is not padded with PKCS#7 padding")
}

func TestCrypto_decryptToken_InvalidBase64ShouldError(t *testing.T) {
	_, err := decryptToken("Not a Token", nil)
	assert.Check(t, err != nil)
}

func TestCrypto_unwrapKey_InvalidBase64ShouldError(t *testing.T) {
	_, err := unwrapKey("Not b64 encoded", nil)
	assert.Check(t, err != nil)
}
