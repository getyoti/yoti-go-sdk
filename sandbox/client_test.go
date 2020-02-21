package sandbox

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"os"
	"testing"

	"gotest.tools/assert"
)

func TestClient_LoadPEMFile(t *testing.T) {
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

	client := &Client{}
	err = client.LoadPEMFile(keyFileName)
	assert.NilError(t, err)
	assert.Check(t, client.Key != nil)
}

func TestClient_LoadPEMFile_ShouldFailForFileNotFound(t *testing.T) {
	MissingFileName := "/tmp/file_not_found"
	client := &Client{}
	err := client.LoadPEMFile(MissingFileName)
	assert.Check(t, err != nil)
}

func TestClient_LoadPEMFile_ShouldFailForInvalidFile(t *testing.T) {
	InvalidFileName := "/tmp/invalid_file"
	err := ioutil.WriteFile(InvalidFileName, []byte("Not a PEM"), 0644)
	assert.NilError(t, err)

	client := &Client{}
	err = client.LoadPEMFile(InvalidFileName)
	assert.Check(t, err != nil)
}
