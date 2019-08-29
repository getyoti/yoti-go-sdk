package requests

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// SignedMessage is a builder for constructing a http.Request with Yoti signing
type SignedMessage struct {
	Key        *rsa.PrivateKey
	HTTPMethod string
	BaseURL    string
	Port       int
	Endpoint   string
	Headers    map[string][]string
	Body       []byte
}

func (msg *SignedMessage) signDigest(digest []byte) (string, error) {
	hash := sha256.Sum256(digest)
	signed, err := rsa.SignPKCS1v15(rand.Reader, msg.Key, crypto.SHA256, hash[:])
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signed), nil
}

func getTimestamp() string {
	return strconv.FormatInt(time.Now().Unix()*1000, 10)
}

func getNonce() (string, error) {
	nonce := make([]byte, 16)
	_, err := rand.Read(nonce)
	return fmt.Sprintf("%X-%X-%X-%X-%X", nonce[0:4], nonce[4:6], nonce[6:8], nonce[8:10], nonce[10:]), err
}

// Request builds a http.Request with signature headers
func (msg *SignedMessage) Request() (request *http.Request, err error) {
	// Check for mandatorys
	if msg.Key == nil {
		err = fmt.Errorf("Missing Private Key")
		return
	}
	if msg.HTTPMethod == "" {
		err = fmt.Errorf("Missing HTTPMethod")
		return
	}
	if msg.BaseURL == "" {
		err = fmt.Errorf("Missing BaseURL")
		return
	}
	if msg.Endpoint == "" {
		err = fmt.Errorf("Missing Endpoint")
		return
	}

	// Mangle BaseURL to allow for optional port number
	baseURL := msg.BaseURL
	if msg.Port != 0 {
		parts := strings.Split(baseURL, "/")
		parts[2] = fmt.Sprintf("%s:%d", parts[2], msg.Port)
		baseURL = strings.Join(parts, "/")
	}

	// Add Timestamp/Nonce to Endpoint
	endpoint := msg.Endpoint
	if !strings.Contains(endpoint, "?") {
		endpoint = endpoint + "?"
	} else {
		endpoint = endpoint + "&"
	}

	nonce, err := getNonce()
	if err != nil {
		return
	}
	endpoint = fmt.Sprintf(
		"%stimestamp=%s&nonce=%s",
		endpoint,
		getTimestamp(),
		nonce,
	)

	// Generate and Sign the message digest
	var digest string
	if msg.Body != nil {
		digest = fmt.Sprintf(
			"%s&%s&%s",
			msg.HTTPMethod,
			endpoint,
			base64.StdEncoding.EncodeToString(msg.Body),
		)
	} else {
		digest = fmt.Sprintf("%s&%s",
			msg.HTTPMethod,
			endpoint,
		)
	}
	signedDigest, err := msg.signDigest([]byte(digest))
	if err != nil {
		return
	}

	// Construct the HTTP Request
	request, err = http.NewRequest(
		msg.HTTPMethod,
		baseURL+endpoint,
		bytes.NewReader(msg.Body),
	)
	if err != nil {
		return
	}
	request.Header.Add("X-Yoti-Auth-Digest", signedDigest)
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&msg.Key.PublicKey)
	request.Header.Add("X-Yoti-Auth-Key", base64.StdEncoding.EncodeToString(publicKeyBytes))
	for key, values := range msg.Headers {
		for _, value := range values {
			request.Header.Add(key, value)
		}
	}
	return request, err
}
