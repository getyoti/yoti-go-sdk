package sandbox

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/getyoti/yoti-go-sdk/v2/cryptoutil"
	"gotest.tools/v3/assert"
)

func TestClient_SetupSharingProfileUsesConstructorBaseUrlOverEnvVariable(t *testing.T) {
	client := createSandboxClient(t, "constuctorBaseUrl")
	os.Setenv("YOTI_API_URL", "envBaseUrl")

	_, err := client.SetupSharingProfile(TokenRequest{})
	assert.NilError(t, err)

	assert.Equal(t, "constuctorBaseUrl", client.BaseURL)
}

func TestClient_SetupSharingProfileUsesEnvVariable(t *testing.T) {
	client := createSandboxClient(t, "")

	os.Setenv("YOTI_API_URL", "envBaseUrl")

	_, err := client.SetupSharingProfile(TokenRequest{})
	assert.NilError(t, err)

	assert.Equal(t, "envBaseUrl", client.BaseURL)
}

func TestClient_SetupSharingProfileUsesDefaultUrlAsFallbackWithEmptyEnvValue(t *testing.T) {
	os.Setenv("YOTI_API_URL", "")

	client := createSandboxClient(t, "")

	_, err := client.SetupSharingProfile(TokenRequest{})
	assert.NilError(t, err)

	assert.Equal(t, "https://api.yoti.com/sandbox/v1", client.BaseURL)
}

func TestClient_SetupSharingProfileUsesDefaultUrlAsFallbackWithNoEnvValue(t *testing.T) {
	os.Unsetenv("YOTI_API_URL")

	client := createSandboxClient(t, "")

	_, err := client.SetupSharingProfile(TokenRequest{})
	assert.NilError(t, err)

	assert.Equal(t, "https://api.yoti.com/sandbox/v1", client.BaseURL)
}

func createSandboxClient(t *testing.T, constructorBaseUrl string) (client Client) {
	keyBytes, fileErr := ioutil.ReadFile("../../test-key.pem")
	assert.NilError(t, fileErr)

	pemFile, parseErr := cryptoutil.ParseRSAKey(keyBytes)
	assert.NilError(t, parseErr)

	if constructorBaseUrl == "" {
		return Client{
			Key:         pemFile,
			ClientSdkID: "ClientSDKID",
			httpClient:  mockHttpClientCreatedResponse(),
		}
	}

	return Client{
		Key:         pemFile,
		BaseURL:     constructorBaseUrl,
		ClientSdkID: "ClientSDKID",
		httpClient:  mockHttpClientCreatedResponse(),
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

func mockHttpClientCreatedResponse() *mockHTTPClient {
	return &mockHTTPClient{
		do: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 201,
				Body:       ioutil.NopCloser(strings.NewReader(`{"token":"tokenValue"}`)),
			}, nil
		},
	}
}
