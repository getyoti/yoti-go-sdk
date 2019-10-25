package requests

import (
	"crypto/rsa"
	"fmt"
	"math/rand"
	"regexp"
	"testing"

	"gotest.tools/assert"
)

func ExampleMergeHeaders() {
	left := map[string][]string{"A": []string{"Value"}}
	right := map[string][]string{"B": []string{"Value"}}

	fmt.Println(MergeHeaders(left, right))
	// Output: map[A:[Value] B:[Value]]
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
