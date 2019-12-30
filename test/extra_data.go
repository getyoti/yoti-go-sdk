package test

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"testing"

	"github.com/getyoti/yoti-go-sdk/v2/yotiprotocom"
	"github.com/getyoti/yoti-go-sdk/v2/yotiprotoshare"
	"github.com/golang/protobuf/proto"
	"gotest.tools/assert"
)

// CreateExtraDataContent creates an encrypted base64 encoded ExtraData
// attribute for use as a test fixture
func CreateExtraDataContent(t *testing.T, pemBytes []byte, protoExtraData *yotiprotoshare.ExtraData, wrappedReceiptKey string) string {
	outBytes, err := proto.Marshal(protoExtraData)
	assert.NilError(t, err)

	keyBytes, _ := pem.Decode(pemBytes)
	key, err := x509.ParsePKCS1PrivateKey(keyBytes.Bytes)
	assert.NilError(t, err)
	cipherBytes, err := base64.StdEncoding.DecodeString(wrappedReceiptKey)
	assert.NilError(t, err)
	unwrappedKey, err := rsa.DecryptPKCS1v15(rand.Reader, key, cipherBytes)
	assert.NilError(t, err)
	cipherBlock, err := aes.NewCipher(unwrappedKey)
	assert.NilError(t, err)

	padLength := cipherBlock.BlockSize() - len(outBytes)%cipherBlock.BlockSize()
	outBytes = append(outBytes, bytes.Repeat([]byte{byte(padLength)}, padLength)...)

	iv := make([]byte, cipherBlock.BlockSize())
	encrypter := cipher.NewCBCEncrypter(cipherBlock, iv)
	encrypter.CryptBlocks(outBytes, outBytes)

	outProto := &yotiprotocom.EncryptedData{
		CipherText: outBytes,
		Iv:         iv,
	}
	outBytes, err = proto.Marshal(outProto)
	assert.NilError(t, err)

	return base64.StdEncoding.EncodeToString(outBytes)
}
