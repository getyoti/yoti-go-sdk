package cryptoutil

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// ParseRSAKey parses a PKCS1 private key from bytes
func ParseRSAKey(keyBytes []byte) (*rsa.PrivateKey, error) {
	// Extract the PEM-encoded data
	block, _ := pem.Decode(keyBytes)

	if block == nil {
		return nil, errors.New("not PEM-encoded")
	}

	if block.Type != "RSA PRIVATE KEY" {
		return nil, errors.New("not RSA private key")
	}

	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, errors.New("bad RSA private key")
	}

	return key, nil
}
