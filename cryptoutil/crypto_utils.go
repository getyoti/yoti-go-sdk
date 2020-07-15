package cryptoutil

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"

	"github.com/getyoti/yoti-go-sdk/v3/util"
)

// ParseRSAKey parses a PKCS1 private key from bytes
func ParseRSAKey(keyBytes []byte) (*rsa.PrivateKey, error) {
	// Extract the PEM-encoded data
	block, _ := pem.Decode(keyBytes)

	if block == nil {
		return nil, errors.New("Invalid Key: not PEM-encoded")
	}

	if block.Type != "RSA PRIVATE KEY" {
		return nil, errors.New("Invalid Key: not RSA private key")
	}

	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, errors.New("Invalid Key: bad RSA private key")
	}

	return key, nil
}

func decryptRsa(cipherBytes []byte, key *rsa.PrivateKey) ([]byte, error) {
	return rsa.DecryptPKCS1v15(rand.Reader, key, cipherBytes)
}

// DecipherAes deciphers AES-encrypted bytes
func DecipherAes(key, iv, cipherBytes []byte) ([]byte, error) {
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
	if len(ciphertext) == 0 {
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

// DecryptToken decrypts an RSA-encrypted token, using the specified RSA private key
func DecryptToken(encryptedConnectToken string, key *rsa.PrivateKey) (result string, err error) {
	// token was encoded as a urlsafe base64 so it can be transferred in a url
	var cipherBytes []byte
	if cipherBytes, err = util.UrlSafeBase64ToBytes(encryptedConnectToken); err != nil {
		return "", err
	}

	var decipheredBytes []byte
	if decipheredBytes, err = decryptRsa(cipherBytes, key); err != nil {
		return "", err
	}

	return util.BytesToUtf8(decipheredBytes), nil
}

// UnwrapKey unwraps an RSA private key
func UnwrapKey(wrappedKey string, key *rsa.PrivateKey) (result []byte, err error) {
	var cipherBytes []byte
	if cipherBytes, err = util.Base64ToBytes(wrappedKey); err != nil {
		return nil, err
	}
	return decryptRsa(cipherBytes, key)
}
