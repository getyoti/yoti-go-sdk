package requests

import (
	"crypto/rsa"
	"fmt"
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
// by signing the request digest with the RSA private key. The digest
// generation and signing are delegated to SignedRequest, which already
// implements the digest/signature scheme required by the Yoti API.
func (s *SignedRequestStrategy) CreateAuthHeaders(httpMethod, endpoint string, body []byte) (map[string][]string, error) {
	msg := SignedRequest{Key: s.key, HTTPMethod: httpMethod, Body: body}
	digest := msg.generateDigest(endpoint)

	signedDigest, err := msg.signDigest([]byte(digest))
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
	nonce, err := getNonce()
	if err != nil {
		return nil, fmt.Errorf("failed to generate nonce: %v", err)
	}

	return map[string]string{
		"nonce":     nonce,
		"timestamp": getTimestamp(),
		"sdkID":     s.clientSdkID,
	}, nil
}
