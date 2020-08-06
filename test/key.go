package test

import (
	"crypto/rsa"
	"io/ioutil"

	"github.com/getyoti/yoti-go-sdk/v3/cryptoutil"
)

// GetValidKey returns a parsed RSA Private Key from a test key
func GetValidKey(filepath string) (key *rsa.PrivateKey) {
	keyBytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic("Error reading the test key: " + err.Error())
	}

	key, err = cryptoutil.ParseRSAKey(keyBytes)
	if err != nil {
		panic("Error parsing the test key: " + err.Error())
	}

	return key
}
