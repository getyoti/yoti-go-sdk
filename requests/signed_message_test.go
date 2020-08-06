package requests

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"regexp"
	"testing"

	"gotest.tools/v3/assert"
)

const exampleKey = "MIICXgIBAAKBgQCpTiICtL+ujx8D0FquVWIaXg+ajJadN5hsTlGUXymiFAunSZjLjTsoGfSPz8PJm6pG9ax1Qb+R5UsSgTRTcpZTps2RLRWr5oPfD66bz4l38QXPSvfg5o+5kNxyCb8QANitF7Ht/DcpsGpL7anruHg/RgCLCBFRaGAodfuJCCM9zwIDAQABAoGBAIJL7GbSvjZUVVU1E6TZd0+9lhqmGf/S2o5309bxSfQ/oxxSyrHU9nMNTqcjCZXuJCTKS7hOKmXY5mbOYvvZ0xA7DXfOc+A4LGXQl0r3ZMzhHZTPKboUSh16E4WI4pr98KagFdkeB/0KBURM3x5d/6dSKip8ZpEyqVpuc9d1xtvhAkEAxabfsqfb4fgBsrhZ/qt133yB0FBHs1alRxvUXZWbVPTOegKi5KBdPptf2QfCy8WK3An/lg8cFQG78PyNll/P0QJBANtJBUHTuRDCoYLhqZLdSTQ52qOWRNutZ2fho9ZcLquokB4SFFeC2I4T+s3oSJ8SNh9vW1nNeXW6Zipx+zz8O58CQQCjV9qNGf40zDITEhmFxwt967aYgpAO3O9wScaCpM4fMsWkvaMDEKiewec/RBOvNY0hdb3ctJX/olRAv2b/vCTRAkAuLmCnDlnJR9QP5kp6HZRPJWgAT6NMyGYgoIqKmHtTt3oyewhBrdLBiT+moaa5qXIwiJkqfnV377uYcMzCeTRtAkEAwHdhM3v01GprmHqE2kvlKOXNq9CB1Z4j/vXSQxBYoSrFWLv5nW9e69ngX+n7qhvO3Gs9CBoy/oqOLatFZOuFEw=="

var keyBytes, _ = base64.StdEncoding.DecodeString(exampleKey)
var privateKey, _ = x509.ParsePKCS1PrivateKey(keyBytes)

func ExampleMergeHeaders() {
	left := map[string][]string{"A": {"Value Of A"}}
	right := map[string][]string{"B": {"Value Of B"}}

	merged := MergeHeaders(left, right)
	fmt.Println(merged["A"])
	fmt.Println(merged["B"])
	// Output:
	// [Value Of A]
	// [Value Of B]
}

func TestMergeHeaders_HandleNullCaseGracefully(t *testing.T) {
	assert.Equal(t, len(MergeHeaders()), 0)
}

func ExampleJSONHeaders() {
	jsonHeaders, _ := json.Marshal(JSONHeaders())
	fmt.Println(string(jsonHeaders))
	// Output: {"Accept":["application/json"],"Content-Type":["application/json"]}
}

func ExampleAuthKeyHeader() {
	headers, _ := json.Marshal(AuthKeyHeader(&privateKey.PublicKey))
	fmt.Println(string(headers))
	// Output: {"X-Yoti-Auth-Key":["MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCpTiICtL+ujx8D0FquVWIaXg+ajJadN5hsTlGUXymiFAunSZjLjTsoGfSPz8PJm6pG9ax1Qb+R5UsSgTRTcpZTps2RLRWr5oPfD66bz4l38QXPSvfg5o+5kNxyCb8QANitF7Ht/DcpsGpL7anruHg/RgCLCBFRaGAodfuJCCM9zwIDAQAB"]}
}

func TestRequestShouldBuildForValid(t *testing.T) {
	random := rand.New(rand.NewSource(25))
	key, err := rsa.GenerateKey(random, 1024)

	assert.NilError(t, err)
	httpMethod := "GET"
	baseURL := "example.com"
	endpoint := "/"

	request := SignedRequest{
		Key:        key,
		HTTPMethod: httpMethod,
		BaseURL:    baseURL,
		Endpoint:   endpoint,
	}
	signed, err := request.Request()
	assert.NilError(t, err)
	assert.Equal(t, httpMethod, signed.Method)
	urlCheck, err := regexp.Match(baseURL+endpoint, []byte(signed.URL.String()))
	assert.NilError(t, err)
	assert.Check(t, urlCheck)
	assert.Check(t, signed.Header["X-Yoti-Auth-Digest"][0] != "")
}

func TestRequestShouldAddHeaders(t *testing.T) {
	random := rand.New(rand.NewSource(25))
	key, err := rsa.GenerateKey(random, 1024)

	assert.NilError(t, err)
	httpMethod := "GET"
	baseURL := "example.com"
	endpoint := "/"

	request := SignedRequest{
		Key:        key,
		HTTPMethod: httpMethod,
		BaseURL:    baseURL,
		Endpoint:   endpoint,
		Headers:    JSONHeaders(),
	}
	signed, err := request.Request()
	assert.NilError(t, err)
	assert.Check(t, signed.Header["X-Yoti-Auth-Digest"][0] != "")
	assert.Equal(t, signed.Header["Accept"][0], "application/json")
}

func TestSignedRequest_checkMandatories_WhenErrorIsSetReturnIt(t *testing.T) {
	msg := &SignedRequest{Error: fmt.Errorf("exampleError")}
	assert.Error(t, msg.checkMandatories(), "exampleError")
}

func TestSignedRequest_checkMandatories_WhenKeyMissing(t *testing.T) {
	msg := &SignedRequest{}
	assert.Error(t, msg.checkMandatories(), "missing private key")
}

func TestSignedRequest_checkMandatories_WhenHTTPMethodMissing(t *testing.T) {
	msg := &SignedRequest{Key: privateKey}
	assert.Error(t, msg.checkMandatories(), "missing HTTPMethod")
}

func TestSignedRequest_checkMandatories_WhenBaseURLMissing(t *testing.T) {
	msg := &SignedRequest{
		Key:        privateKey,
		HTTPMethod: http.MethodPost,
	}
	assert.Error(t, msg.checkMandatories(), "missing BaseURL")
}

func TestSignedRequest_checkMandatories_WhenEndpointMissing(t *testing.T) {
	msg := &SignedRequest{
		Key:        privateKey,
		HTTPMethod: http.MethodPost,
		BaseURL:    "example.com",
	}
	assert.Error(t, msg.checkMandatories(), "missing Endpoint")
}

func ExampleSignedRequest_generateDigest() {
	msg := &SignedRequest{
		HTTPMethod: http.MethodPost,
		Body:       []byte("simple message body"),
	}
	fmt.Println(msg.generateDigest("endpoint"))
	// Output: POST&endpoint&c2ltcGxlIG1lc3NhZ2UgYm9keQ==

}

func ExampleSignedRequest_WithPemFile() {
	msg := SignedRequest{}.WithPemFile([]byte(`
-----BEGIN RSA PRIVATE KEY-----
` + exampleKey + `
-----END RSA PRIVATE KEY-----`))
	fmt.Println(AuthKeyHeader(&msg.Key.PublicKey))
	// Output: map[X-Yoti-Auth-Key:[MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCpTiICtL+ujx8D0FquVWIaXg+ajJadN5hsTlGUXymiFAunSZjLjTsoGfSPz8PJm6pG9ax1Qb+R5UsSgTRTcpZTps2RLRWr5oPfD66bz4l38QXPSvfg5o+5kNxyCb8QANitF7Ht/DcpsGpL7anruHg/RgCLCBFRaGAodfuJCCM9zwIDAQAB]]
}

func TestSignedRequest_WithPemFile_NotPemEncodedShouldError(t *testing.T) {
	msg := SignedRequest{}.WithPemFile([]byte("not pem encoded"))
	assert.ErrorContains(t, msg.Error, "not PEM-encoded")
}

func TestSignedRequest_WithPemFile_NotRSAKeyShouldError(t *testing.T) {
	msg := SignedRequest{}.WithPemFile([]byte(`-----BEGIN RSA PUBLIC KEY-----
` + exampleKey + `
-----END RSA PUBLIC KEY-----`))
	assert.ErrorContains(t, msg.Error, "not an RSA Private Key")
}
