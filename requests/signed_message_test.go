package requests

import (
	"crypto/rsa"
	"fmt"
	"math/rand"
	"regexp"
	"testing"

	"gotest.tools/v3/assert"
)

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
