package yoti

import (
	"crypto/rsa"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/getyoti/yoti-go-sdk/v3/test"
	"gotest.tools/v3/assert"
)

type mockHTTPClient struct {
	do func(*http.Request) (*http.Response, error)
}

func (mock *mockHTTPClient) Do(request *http.Request) (*http.Response, error) {
	if mock.do != nil {
		return mock.do(request)
	}
	return nil, nil
}

func TestYotiClient_KeyLoad_Failure(t *testing.T) {
	key, _ := ioutil.ReadFile("test/test-key-invalid-format.pem")
	_, err := NewClient("", key)

	assert.ErrorContains(t, err, "Invalid Key: not PEM-encoded")

	tempError, temporary := err.(interface {
		Temporary() bool
	})
	assert.Check(t, !temporary || !tempError.Temporary())
}

func TestYotiClient_HttpFailure_ReturnsFailure(t *testing.T) {
	key := getValidKey()

	client := Client{
		HTTPClient: &mockHTTPClient{
			do: func(*http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 500,
				}, nil
			},
		},
		Key: key,
	}

	_, err := client.GetActivityDetails(test.EncryptedToken)

	assert.Check(t, err != nil)
	assert.ErrorContains(t, err, "unknown HTTP error")
	tempError, temporary := err.(interface {
		Temporary() bool
		Unwrap() error
	})
	assert.Check(t, temporary)
	assert.Check(t, tempError.Temporary())
	assert.ErrorContains(t, tempError.Unwrap(), "unknown HTTP error")
}

func TestYotiClient_HttpFailure_ReturnsProfileNotFound(t *testing.T) {
	key := getValidKey()

	client := Client{
		HTTPClient: &mockHTTPClient{
			do: func(*http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 404,
				}, nil
			},
		},
		Key: key,
	}

	_, err := client.GetActivityDetails(test.EncryptedToken)

	assert.ErrorContains(t, err, "Profile Not Found")
	tempError, temporary := err.(interface {
		Temporary() bool
	})
	assert.Check(t, !temporary || !tempError.Temporary())
}

func TestClient_OverrideAPIURL_ShouldSetAPIURL(t *testing.T) {
	client := &Client{}
	expectedURL := "expectedurl.com"
	client.OverrideAPIURL(expectedURL)
	assert.Equal(t, client.getAPIURL(), expectedURL)
}

func TestYotiClient_GetAPIURLUsesOverriddenBaseUrlOverEnvVariable(t *testing.T) {
	client := Client{}
	client.OverrideAPIURL("overridenBaseUrl")

	os.Setenv("YOTI_API_URL", "envBaseUrl")

	result := client.getAPIURL()

	assert.Equal(t, "overridenBaseUrl", result)
}

func TestYotiClient_GetAPIURLUsesEnvVariable(t *testing.T) {
	client := Client{}

	os.Setenv("YOTI_API_URL", "envBaseUrl")

	result := client.getAPIURL()

	assert.Equal(t, "envBaseUrl", result)
}

func TestYotiClient_GetAPIURLUsesDefaultUrlAsFallbackWithEmptyEnvValue(t *testing.T) {
	client := Client{}

	os.Setenv("YOTI_API_URL", "")

	result := client.getAPIURL()

	assert.Equal(t, "https://api.yoti.com/api/v1", result)
}

func TestYotiClient_GetAPIURLUsesDefaultUrlAsFallbackWithNoEnvValue(t *testing.T) {
	client := Client{}

	os.Unsetenv("YOTI_API_URL")

	result := client.getAPIURL()

	assert.Equal(t, "https://api.yoti.com/api/v1", result)
}

func getValidKey() *rsa.PrivateKey {
	return test.GetValidKey("test/test-key.pem")
}
