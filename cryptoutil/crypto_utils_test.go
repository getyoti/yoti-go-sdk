package cryptoutil

import (
	"crypto/aes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
	"testing"

	"github.com/getyoti/yoti-go-sdk/v3/file"
	"gotest.tools/v3/assert"
)

func TestCryptoUtil_ParseRSAKey(t *testing.T) {
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	assert.NilError(t, err)
	block := pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}
	keyFileName := "tmpKey.pem"
	keyFile, err := os.Create(keyFileName)
	assert.NilError(t, err)

	err = pem.Encode(keyFile, &block)
	assert.NilError(t, err)

	err = keyFile.Close()
	assert.NilError(t, err)

	var keyBytes []byte
	keyBytes, err = file.ReadFile(keyFileName)
	assert.NilError(t, err)

	key, err = ParseRSAKey(keyBytes)
	assert.NilError(t, err)
	assert.Check(t, key != nil)
}

func TestCryptoutil_ParseRSAKey_PublicKeyShouldFail(t *testing.T) {
	testPEM := []byte(`-----BEGIN RSA PUBLIC KEY-----
VGVzdCBTdHJpbmc=
-----END RSA PUBLIC KEY-----`)

	_, err := ParseRSAKey(testPEM)

	assert.Error(t, err, "invalid Key: not RSA private key")
}

func TestCryptoutil_ParseRSAKey_InvalidKeyShouldFail(t *testing.T) {
	testPEM := []byte(`-----BEGIN RSA PRIVATE KEY-----
VGVzdCBTdHJpbmc=
-----END RSA PRIVATE KEY-----`)

	_, err := ParseRSAKey(testPEM)

	assert.Error(t, err, "invalid Key: bad RSA private key")
}

func TestCryptoutil_decipherAes_EmptyCiphertextShouldError(t *testing.T) {
	_, err := DecipherAes([]byte{}, []byte{}, []byte{})
	assert.Check(t, err != nil)
}

func TestCryptoutil_decipherAes_CiphertextNotMatchingBlocksizeShouldError(t *testing.T) {
	_, err := DecipherAes(
		make([]byte, 16),
		[]byte{},
		make([]byte, aes.BlockSize+1),
	)
	assert.Check(t, err != nil)
}

func TestCryptoutil_pkcs7Unpad_InvalidBlocksizeShouldError(t *testing.T) {
	_, err := pkcs7Unpad([]byte{}, -1)
	assert.Error(t, err, "blocksize -1 is not valid for padding removal")
}

func TestCryptoutil_pkcs7Unpad_EmptyByteArrayShouldError(t *testing.T) {
	_, err := pkcs7Unpad([]byte{}, 1)
	assert.Error(t, err, "cannot remove padding on empty byte array")
}

func TestCryptoutil_pkcs7Unpad_CiphertextNotMultipleOfBlocksizeShouldError(t *testing.T) {
	_, err := pkcs7Unpad([]byte{0, 0, 0}, 2)
	assert.Error(t, err, "ciphertext is not a multiple of the block size")
}

func TestCryptoutil_pkcs7Unpad_CiphertextNotPaddedShouldError(t *testing.T) {
	_, err := pkcs7Unpad([]byte{0xF, 0x0, 0x0}, 3)
	assert.Error(t, err, "ciphertext is not padded with PKCS#7 padding")
}

func TestCryptoutil_pkcs7Unpad_CiphertextPaddingIncorrect(t *testing.T) {
	_, err := pkcs7Unpad([]byte{0x1, 0x1, 0x1, 0xF, 0x0, 0x0, 0xB, 0xA, 0x7}, 3)
	assert.Error(t, err, "ciphertext is not padded with PKCS#7 padding")
}

func TestCryptoutil_decryptToken_InvalidBase64ShouldError(t *testing.T) {
	_, err := DecryptToken("Not a Token", nil)
	assert.Check(t, err != nil)
}

func TestCryptoutil_unwrapKey_InvalidBase64ShouldError(t *testing.T) {
	_, err := UnwrapKey("Not b64 encoded", nil)
	assert.Check(t, err != nil)
}
