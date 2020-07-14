package requests

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/getyoti/yoti-go-sdk/v3/consts"
)

// MergeHeaders merges two or more header prototypes together from left to right
func MergeHeaders(headers ...map[string][]string) map[string][]string {
	if len(headers) == 0 {
		return make(map[string][]string)
	}
	out := headers[0]
	for _, element := range headers[1:] {
		for k, v := range element {
			out[k] = v
		}
	}
	return out
}

// JSONHeaders is a header prototype for JSON based requests
func JSONHeaders() map[string][]string {
	return map[string][]string{
		"Content-Type": {"application/json"},
		"Accept":       {"application/json"},
	}
}

// AuthKeyHeader is a header prototype including an encoded RSA PublicKey
func AuthKeyHeader(key *rsa.PublicKey) map[string][]string {
	return map[string][]string{
		"X-Yoti-Auth-Key": {
			base64.StdEncoding.EncodeToString(
				func(a []byte, _ error) []byte {
					return a
				}(x509.MarshalPKIXPublicKey(key)),
			),
		},
	}
}

// SignedRequest is a builder for constructing a http.Request with Yoti signing
type SignedRequest struct {
	Key        *rsa.PrivateKey
	HTTPMethod string
	BaseURL    string
	Endpoint   string
	Headers    map[string][]string
	Params     map[string]string
	Body       []byte
	Error      error
}

func (msg *SignedRequest) signDigest(digest []byte) (string, error) {
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

// WithPemFile loads the private key from a PEM file reader
func (msg SignedRequest) WithPemFile(in []byte) SignedRequest {
	block, _ := pem.Decode(in)
	if block == nil {
		msg.Error = errors.New("Not PEM-encoded")
		return msg
	}
	if block.Type != "RSA PRIVATE KEY" {
		msg.Error = errors.New("Not an RSA Private Key")
		return msg
	}

	msg.Key, msg.Error = x509.ParsePKCS1PrivateKey(block.Bytes)
	return msg
}

func (msg *SignedRequest) addParametersToEndpoint() (endpoint string, err error) {
	if msg.Params == nil {
		msg.Params = make(map[string]string)
	}
	// Add Timestamp/Nonce
	if _, ok := msg.Params["nonce"]; !ok {
		nonce, err := getNonce()
		if err != nil {
			return "", err
		}
		msg.Params["nonce"] = nonce
	}
	if _, ok := msg.Params["timestamp"]; !ok {
		msg.Params["timestamp"] = getTimestamp()
	}

	endpoint = msg.Endpoint
	if !strings.Contains(endpoint, "?") {
		endpoint = endpoint + "?"
	} else {
		endpoint = endpoint + "&"
	}

	var firstParam = true
	for param, value := range msg.Params {
		var formatString = "%s&%s=%s"
		if firstParam {
			formatString = "%s%s=%s"
		}
		endpoint = fmt.Sprintf(formatString, endpoint, param, value)
		firstParam = false
	}
	return
}

func (msg *SignedRequest) generateDigest(endpoint string) (digest string) {
	// Generate the message digest
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
	return
}

func (msg *SignedRequest) checkMandatories() error {
	if msg.Error != nil {
		return msg.Error
	}
	if msg.Key == nil {
		return fmt.Errorf("Missing Private Key")
	}
	if msg.HTTPMethod == "" {
		return fmt.Errorf("Missing HTTPMethod")
	}
	if msg.BaseURL == "" {
		return fmt.Errorf("Missing BaseURL")
	}
	if msg.Endpoint == "" {
		return fmt.Errorf("Missing Endpoint")
	}
	return nil
}

// Request builds a http.Request with signature headers
func (msg SignedRequest) Request() (request *http.Request, err error) {
	err = msg.checkMandatories()
	if err != nil {
		return
	}

	endpoint, err := msg.addParametersToEndpoint()
	if err != nil {
		return
	}
	signedDigest, err := msg.signDigest([]byte(msg.generateDigest(endpoint)))
	if err != nil {
		return
	}

	// Construct the HTTP Request
	request, err = http.NewRequest(
		msg.HTTPMethod,
		msg.BaseURL+endpoint,
		bytes.NewReader(msg.Body),
	)
	if err != nil {
		return
	}
	request.Header.Add("X-Yoti-Auth-Digest", signedDigest)
	request.Header.Add("X-Yoti-SDK", consts.SDKIdentifier)
	request.Header.Add("X-Yoti-SDK-Verson", consts.SDKVersionIdentifier)

	for key, values := range msg.Headers {
		for _, value := range values {
			request.Header.Add(key, value)
		}
	}
	return request, err
}
