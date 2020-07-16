package sandbox

import (
	"crypto/rand"
	"crypto/rsa"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox/request"

	"github.com/getyoti/yoti-go-sdk/v3/cryptoutil"
	"gotest.tools/v3/assert"
)

func TestClient_ConfigureSessionResponse_ShouldReturnErrorIfNotCreated(t *testing.T) {
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
	err := client.ConfigureSessionResponse("some_session_id", request.ResponseConfig{})
	assert.ErrorContains(t, err, "Response config not created")
}

func TestClient_ConfigureSessionResponse_Success(t *testing.T) {
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
	err := client.ConfigureSessionResponse("some_session_id", request.ResponseConfig{})
	assert.NilError(t, err)
}

func TestClient_ConfigureApplicationResponse_ShouldReturnErrorIfNotCreated(t *testing.T) {
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
	err := client.ConfigureApplicationResponse(request.ResponseConfig{})
	assert.ErrorContains(t, err, "Response config not created")
}

func TestClient_ConfigureApplicationResponse_Success(t *testing.T) {
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
	err := client.ConfigureApplicationResponse(request.ResponseConfig{})
	assert.NilError(t, err)
}

func TestClient_ConfigureSessionResponseUsesConstructorBaseUrlOverEnvVariable(t *testing.T) {
	client := createSandboxClient(t, "constuctorBaseUrl")
	os.Setenv("YOTI_DOC_SCAN_API_URL", "envBaseUrl")

	err := client.ConfigureSessionResponse("some_session_id", request.ResponseConfig{})
	assert.NilError(t, err)

	assert.Equal(t, "constuctorBaseUrl", client.BaseURL)
}

func TestClient_ConfigureSessionResponseUsesEnvVariable(t *testing.T) {
	client := createSandboxClient(t, "")

	os.Setenv("YOTI_DOC_SCAN_API_URL", "envBaseUrl")

	err := client.ConfigureSessionResponse("some_session_id", request.ResponseConfig{})
	assert.NilError(t, err)

	assert.Equal(t, "envBaseUrl", client.BaseURL)
}

func TestClient_ConfigureSessionResponseUsesDefaultUrlAsFallbackWithEmptyEnvValue(t *testing.T) {
	os.Setenv("YOTI_DOC_SCAN_API_URL", "")

	client := createSandboxClient(t, "")

	err := client.ConfigureSessionResponse("some_session_id", request.ResponseConfig{})
	assert.NilError(t, err)

	assert.Equal(t, "https://api.yoti.com/sandbox/idverify/v1", client.BaseURL)
}

func TestClient_ConfigureSessionResponseUsesDefaultUrlAsFallbackWithNoEnvValue(t *testing.T) {
	os.Unsetenv("YOTI_DOC_SCAN_API_URL")

	client := createSandboxClient(t, "")

	err := client.ConfigureSessionResponse("some_session_id", request.ResponseConfig{})
	assert.NilError(t, err)

	assert.Equal(t, "https://api.yoti.com/sandbox/idverify/v1", client.BaseURL)
}

func createSandboxClient(t *testing.T, constructorBaseURL string) (client Client) {
	keyBytes, fileErr := ioutil.ReadFile("../../test-key.pem")
	assert.NilError(t, fileErr)

	pemFile, parseErr := cryptoutil.ParseRSAKey(keyBytes)
	assert.NilError(t, parseErr)

	if constructorBaseURL == "" {
		return Client{
			Key:         pemFile,
			ClientSdkID: "ClientSDKID",
			HTTPClient:  mockHTTPClientCreatedResponse(),
		}
	}

	return Client{
		Key:         pemFile,
		BaseURL:     constructorBaseURL,
		ClientSdkID: "ClientSDKID",
		HTTPClient:  mockHTTPClientCreatedResponse(),
	}

}

type mockHTTPClient struct {
	do func(*http.Request) (*http.Response, error)
}

func (mock *mockHTTPClient) Do(request *http.Request) (*http.Response, error) {
	if mock.do != nil {
		return mock.do(request)
	}
	return nil, nil
}

func mockHTTPClientCreatedResponse() *mockHTTPClient {
	return &mockHTTPClient{
		do: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 201,
				Body:       ioutil.NopCloser(strings.NewReader(`{"token":"tokenValue"}`)),
			}, nil
		},
	}
}
