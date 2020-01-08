package sandbox

import (
	"crypto/rand"
	"crypto/rsa"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"gotest.tools/assert"
)

type mockHTTPClient struct {
	do func(*http.Request) (*http.Response, error)
}

func (mock *mockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	return mock.do(req)
}

func TestClient_SetupSharingProfile_ShouldReturnErrorIfProfileNotCreated(t *testing.T) {
	key, _ := rsa.GenerateKey(rand.Reader, 1024)
	client := Client{
		Key:     key,
		BaseURL: "example.com",
		HTTPClient: &mockHTTPClient{
			do: func(*http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 401,
					Body:       ioutil.NopCloser(strings.NewReader("")),
				}, nil
			},
		},
	}
	_, err := client.SetupSharingProfile(Profile{})
	assert.ErrorContains(t, err, "Sharing Profile not created")
}

func TestClient_SetupSharingProfile_Success(t *testing.T) {
	expectedToken := "shareToken"
	key, _ := rsa.GenerateKey(rand.Reader, 1024)
	client := Client{
		Key:     key,
		BaseURL: "example.com",
		HTTPClient: &mockHTTPClient{
			do: func(*http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 201,
					Body:       ioutil.NopCloser(strings.NewReader(`{"token":"` + expectedToken + `"}`)),
				}, nil
			},
		},
	}
	token, err := client.SetupSharingProfile(Profile{})
	assert.NilError(t, err)

	assert.Equal(t, token, expectedToken)
}
