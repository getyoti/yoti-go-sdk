package requests

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strconv"
	"time"
)

// SignedRequestStrategy implements AuthStrategy using Yoti's signed-request
// authentication. It generates a digest signature using the RSA private key
// and includes nonce, timestamp, and sdkID as query parameters.
type SignedRequestStrategy struct {
	key         *rsa.PrivateKey
	clientSdkID string
}

// NewSignedRequestStrategy creates a new SignedRequestStrategy with the given
// RSA private key and SDK ID.
func NewSignedRequestStrategy(key *rsa.PrivateKey, clientSdkID string) (*SignedRequestStrategy, error) {
	if key == nil {
		return nil, fmt.Errorf("private key must not be nil")
	}
	if clientSdkID == "" {
		return nil, fmt.Errorf("client SDK ID must not be empty")
	}
	return &SignedRequestStrategy{key: key, clientSdkID: clientSdkID}, nil
}

// CreateAuthHeaders generates the X-Yoti-Auth-Digest and X-Yoti-Auth-Id headers
// by signing the request digest with the RSA private key.
func (s *SignedRequestStrategy) CreateAuthHeaders(httpMethod, endpoint string, body []byte) (map[string][]string, error) {
	digest := generateDigest(httpMethod, endpoint, body)

	signedDigest, err := signDigest(s.key, []byte(digest))
	if err != nil {
		return nil, fmt.Errorf("failed to sign digest: %v", err)
	}

	return map[string][]string{
		"X-Yoti-Auth-Digest": {signedDigest},
		"X-Yoti-Auth-Id":     {s.clientSdkID},
	}, nil
}

// CreateQueryParams returns the nonce, timestamp, and sdkID query parameters
// required for signed-request authentication.
func (s *SignedRequestStrategy) CreateQueryParams() (map[string]string, error) {
	nonce, err := generateNonce()
	if err != nil {
		return nil, fmt.Errorf("failed to generate nonce: %v", err)
	}

	return map[string]string{
		"nonce":     nonce,
		"timestamp": generateTimestamp(),
		"sdkID":     s.clientSdkID,
	}, nil
}

func generateDigest(httpMethod, endpoint string, body []byte) string {
	if body != nil {
		return fmt.Sprintf(
			"%s&%s&%s",
			httpMethod,
			endpoint,
			base64.StdEncoding.EncodeToString(body),
		)
	}
	return fmt.Sprintf("%s&%s", httpMethod, endpoint)
}

func signDigest(key *rsa.PrivateKey, digest []byte) (string, error) {
	hash := sha256.Sum256(digest)
	signed, err := rsa.SignPKCS1v15(rand.Reader, key, crypto.SHA256, hash[:])
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signed), nil
}

func generateNonce() (string, error) {
	nonce := make([]byte, 16)
	_, err := rand.Read(nonce)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%X-%X-%X-%X-%X", nonce[0:4], nonce[4:6], nonce[6:8], nonce[8:10], nonce[10:]), nil
}

func generateTimestamp() string {
	return strconv.FormatInt(time.Now().Unix()*1000, 10)
}
