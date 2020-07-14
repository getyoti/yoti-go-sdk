package cryptoutil

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
	"testing"

	"github.com/getyoti/yoti-go-sdk/v3/file"
	"gotest.tools/v3/assert"
)

func TestCryptoUtils_ParseRSAKey(t *testing.T) {
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
