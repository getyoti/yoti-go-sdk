package yoti

import (
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
)

func loadRsaKey(keyBytes []byte) (*rsa.PrivateKey, error) {
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

func decryptRsa(cipherBytes []byte, key *rsa.PrivateKey) ([]byte, error) {
	return rsa.DecryptPKCS1v15(rand.Reader, key, cipherBytes)
}

func decipherAes(key, iv, cipherBytes []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte{}, err
	}

	// CBC mode always works in whole blocks.
	if (len(cipherBytes) % aes.BlockSize) != 0 {
		return []byte{}, errors.New("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	decipheredBytes := make([]byte, len(cipherBytes))

	mode.CryptBlocks(decipheredBytes, cipherBytes)

	return pkcs7Unpad(decipheredBytes, aes.BlockSize)
}

func pkcs7Unpad(ciphertext []byte, blocksize int) (result []byte, err error) {
	if blocksize <= 0 {
		err = fmt.Errorf("blocksize %d is not valid for padding removal", blocksize)
		return
	}
	if ciphertext == nil || len(ciphertext) == 0 {
		err = errors.New("Cannot remove padding on empty byte array")
		return
	}
	if len(ciphertext)%blocksize != 0 {
		err = errors.New("ciphertext is not a multiple of the block size")
		return
	}

	c := ciphertext[len(ciphertext)-1]
	n := int(c)
	if n == 0 || n > len(ciphertext) {
		err = errors.New("ciphertext is not padded with PKCS#7 padding")
		return
	}

	// verify all padding bytes are correct
	for i := 0; i < n; i++ {
		if ciphertext[len(ciphertext)-n+i] != c {
			err = errors.New("ciphertext is not padded with PKCS#7 padding")
			return
		}
	}
	return ciphertext[:len(ciphertext)-n], nil
}

func signDigest(digest []byte, key *rsa.PrivateKey) ([]byte, error) {
	hashed := sha256.Sum256(digest)

	signedDigest, err := rsa.SignPKCS1v15(rand.Reader, key, crypto.SHA256, hashed[:])
	if err != nil {
		return []byte{}, err
	}

	return signedDigest, nil
}

func getDerEncodedPublicKey(key *rsa.PrivateKey) (result string, err error) {
	var derEncodedBytes []byte
	if derEncodedBytes, err = x509.MarshalPKIXPublicKey(key.Public()); err != nil {
		return
	}

	result = bytesToBase64(derEncodedBytes)
	return
}

func generateNonce() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	uuid := fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

	return uuid, nil
}
