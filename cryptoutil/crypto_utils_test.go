package cryptoutil

import (
	"crypto/aes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"os"
	"testing"

	"github.com/getyoti/yoti-go-sdk/v3/file"
	"github.com/getyoti/yoti-go-sdk/v3/util"
	"github.com/getyoti/yoti-go-sdk/v3/yotiprotocom"
	"google.golang.org/protobuf/proto"
	"gotest.tools/v3/assert"
)

const (
	wrappedKey       = "kyHPjq2+Y48cx+9yS/XzmW09jVUylSdhbP+3Q9Tc9p6bCEnyfa8vj38AIu744RzzE+Dc4qkSF21VfzQKtJVILfOXu5xRc7MYa5k3zWhjiesg/gsrv7J4wDyyBpHIJB8TWXnubYMbSYQJjlsfwyxE9kGe0YI08pRo2Tiht0bfR5Z/YrhAk4UBvjp84D+oyug/1mtGhKphA4vgPhQ9/y2wcInYxju7Q6yzOsXGaRUXR38Tn2YmY9OBgjxiTnhoYJFP1X9YJkHeWMW0vxF1RHxgIVrpf7oRzdY1nq28qzRg5+wC7cjRpS2i/CKUAo0oVG4pbpXsaFhaTewStVC7UFtA77JHb3EnF4HcSWMnK5FM7GGkL9MMXQenh11NZHKPWXpux0nLZ6/vwffXZfsiyTIcFL/NajGN8C/hnNBljoQ+B3fzWbjcq5ueUOPwARZ1y38W83UwMynzkud/iEdHLaZIu4qUCRkfSxJg7Dc+O9/BdiffkOn2GyFmNjVeq754DCUypxzMkjYxokedN84nK13OU4afVyC7t5DDxAK/MqAc69NCBRLqMi5f8BMeOZfMcSWPGC9a2Qu8VgG125TuZT4+wIykUhGyj3Bb2/fdPsxwuKFR+E0uqs0ZKvcv1tkNRRtKYBqTacgGK9Yoehg12cyLrITLdjU1fmIDn4/vrhztN5w="
	b64EncryptedData = "ChCZAib1TBm9Q5GYfFrS1ep9EnAwQB5shpAPWLBgZgFgt6bCG3S5qmZHhrqUbQr3yL6yeLIDwbM7x4nuT/MYp+LDXgmFTLQNYbDTzrEzqNuO2ZPn9Kpg+xpbm9XtP7ZLw3Ep2BCmSqtnll/OdxAqLb4DTN4/wWdrjnFC+L/oQEECu646"
	encryptedToken   = "b6H19bUCJhwh6WqQX_sEHWX9RP-A_ANr1fkApwA4Dp2nJQFAjrF9e6YCXhNBpAIhfHnN0iXubyXxXZMNwNMSQ5VOxkqiytrvPykfKQWHC6ypSbfy0ex8ihndaAXG5FUF-qcU8QaFPMy6iF3x0cxnY0Ij0kZj0Ng2t6oiNafb7AhT-VGXxbFbtZu1QF744PpWMuH0LVyBsAa5N5GJw2AyBrnOh67fWMFDKTJRziP5qCW2k4h5vJfiYr_EOiWKCB1d_zINmUm94ZffGXxcDAkq-KxhN1ZuNhGlJ2fKcFh7KxV0BqlUWPsIEiwS0r9CJ2o1VLbEs2U_hCEXaqseEV7L29EnNIinEPVbL4WR7vkF6zQCbK_cehlk2Qwda-VIATqupRO5grKZN78R9lBitvgilDaoE7JB_VFcPoljGQ48kX0wje1mviX4oJHhuO8GdFITS5LTbojGVQWT7LUNgAUe0W0j-FLHYYck3v84OhWTqads5_jmnnLkp9bdJSRuJF0e8pNdePnn2lgF-GIcyW_0kyGVqeXZrIoxnObLpF-YeUteRBKTkSGFcy7a_V_DLiJMPmH8UXDLOyv8TVt3ppzqpyUrLN2JVMbL5wZ4oriL2INEQKvw_boDJjZDGeRlu5m1y7vGDNBRDo64-uQM9fRUULPw-YkABNwC0DeShswzT00="
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

	assert.Error(t, err, "invalid key: not RSA private key")
}

func TestCryptoutil_ParseRSAKey_InvalidKeyShouldFail(t *testing.T) {
	testPEM := []byte(`-----BEGIN RSA PRIVATE KEY-----
VGVzdCBTdHJpbmc=
-----END RSA PRIVATE KEY-----`)

	_, err := ParseRSAKey(testPEM)

	assert.Error(t, err, "invalid key: bad RSA private key")
}

func TestCryptoutil_ParseRSAKey_InvalidShouldFail(t *testing.T) {
	var testPEM []byte

	_, err := ParseRSAKey(testPEM)

	assert.Error(t, err, "invalid key: not PEM-encoded")
}

func TestCryptoutil_DecipherAes(t *testing.T) {
	unwrappedKey, err := UnwrapKey(wrappedKey, getKey())
	if err != nil {
		return
	}

	encryptedBytes, err := util.Base64ToBytes(b64EncryptedData)
	if err != nil || len(encryptedBytes) == 0 {
		assert.NilError(t, err)
	}
	encryptedData := &yotiprotocom.EncryptedData{}
	err = proto.Unmarshal(encryptedBytes, encryptedData)
	assert.NilError(t, err)

	bytes, err := DecipherAes(unwrappedKey, encryptedData.Iv, encryptedData.CipherText)
	assert.NilError(t, err)
	assert.Check(t, bytes != nil)
}

func TestCryptoutil_DecipherAes_EmptyCiphertextShouldError(t *testing.T) {
	_, err := DecipherAes([]byte{}, []byte{}, []byte{})
	assert.Check(t, err != nil)
}

func TestCryptoutil_DecipherAes_CiphertextNotMatchingBlocksizeShouldError(t *testing.T) {
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

func TestCryptoutil_DecryptToken(t *testing.T) {
	_, err := DecryptToken(encryptedToken, getKey())
	assert.NilError(t, err)
}

func TestCryptoutil_DecryptToken_InvalidToken(t *testing.T) {
	_, err := DecryptToken("c29tZS10b2tlbg==", getKey())
	assert.Error(t, err, "crypto/rsa: decryption error")
}

func TestCryptoutil_DecryptToken_InvalidBase64ShouldError(t *testing.T) {
	_, err := DecryptToken("Not a Token", nil)
	assert.Check(t, err != nil)
}

func TestCryptoutil_UnwrapKey(t *testing.T) {
	unwrappedKey, err := UnwrapKey(wrappedKey, getKey())

	assert.NilError(t, err)
	assert.Check(t, unwrappedKey != nil)
}

func TestCryptoutil_UnwrapKey_InvalidBase64ShouldError(t *testing.T) {
	_, err := UnwrapKey("Not b64 encoded", nil)
	assert.Check(t, err != nil)
}

func getKey() (key *rsa.PrivateKey) {
	keyBytes, err := ioutil.ReadFile("../test/test-key.pem")
	if err != nil {
		panic("Error reading the test key: " + err.Error())
	}

	key, err = ParseRSAKey(keyBytes)
	if err != nil {
		panic("Error parsing the test key: " + err.Error())
	}

	return key
}
