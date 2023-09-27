package yoti

import (
	"crypto/rsa"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/getyoti/yoti-go-sdk/v3/digitalidentity"
	"github.com/getyoti/yoti-go-sdk/v3/test"
	"gotest.tools/v3/assert"
)

func TestDigitalIDClient(t *testing.T) {
	key, err := os.ReadFile("./test/test-key.pem")
	assert.NilError(t, err)

	_, err = NewDigitalIdentityClient("some-sdk-id", key)
	assert.NilError(t, err)
}

func TestDigitalIDClient_KeyLoad_Failure(t *testing.T) {
	key, err := os.ReadFile("test/test-key-invalid-format.pem")
	assert.NilError(t, err)

	_, err = NewDigitalIdentityClient("", key)

	assert.ErrorContains(t, err, "invalid key: not PEM-encoded")

	tempError, temporary := err.(interface {
		Temporary() bool
	})
	assert.Check(t, !temporary || !tempError.Temporary())
}

func TestYotiClient_CreateShareSession(t *testing.T) {
	key, err := os.ReadFile("./test/test-key.pem")
	assert.NilError(t, err)

	client, err := NewDigitalIdentityClient("some-sdk-id", key)
	assert.NilError(t, err)

	client.HTTPClient = &mockHTTPClient{
		do: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 201,
				Body:       io.NopCloser(strings.NewReader(`{"id":"SOME_ID","status":"SOME_STATUS","expiry":"SOME_EXPIRY","created":"SOME_CREATED","updated":"SOME_UPDATED","qrCode":{"id":"SOME_QRCODE_ID"},"receipt":{"id":"SOME_RECEIPT_ID"}}`)),
			}, nil
		},
	}

	policy, err := (&digitalidentity.PolicyBuilder{}).WithFullName().WithWantedRememberMe().Build()
	assert.NilError(t, err)

	session, err := (&digitalidentity.ShareSessionRequestBuilder{}).WithPolicy(policy).Build()
	assert.NilError(t, err)

	result, err := client.CreateShareSession(&session)

	assert.NilError(t, err)
	assert.Equal(t, result.Status, "SOME_STATUS")
}

func TestDigitalIDClient_HttpFailure_ReturnsApplicationNotFound(t *testing.T) {
	key := getDigitalValidKey()
	client := DigitalIdentityClient{
		HTTPClient: &mockHTTPClient{
			do: func(*http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 404,
				}, nil
			},
		},
		Key: key,
	}

	_, err := client.GetSession("SOME ID")

	assert.ErrorContains(t, err, "Application was not found")
	tempError, temporary := err.(interface {
		Temporary() bool
	})
	assert.Check(t, !temporary || !tempError.Temporary())
}

func TestDigitalIDClient_GetSession(t *testing.T) {
	key, err := os.ReadFile("./test/test-key.pem")

	mockSessionID := "SOME_SESSION_ID"
	client, err := NewDigitalIdentityClient("some-sdk-id", key)

	client.HTTPClient = &mockHTTPClient{
		do: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(strings.NewReader(`{"id":"SOME_ID","status":"SOME_STATUS","expiry":"SOME_EXPIRY","created":"SOME_CREATED","updated":"SOME_UPDATED","qrCode":{"id":"SOME_QRCODE_ID"},"receipt":{"id":"SOME_RECEIPT_ID"}}`)),
			}, nil
		},
	}

	result, err := client.GetSession(mockSessionID)
	if err != nil {
		return
	}

	assert.Equal(t, result.Id, "SOME_ID")
	assert.Equal(t, result.Status, "SOME_STATUS")
	assert.Equal(t, result.Created, "SOME_CREATED")

}

func TestDigitalIDClient_OverrideAPIURL_ShouldSetAPIURL(t *testing.T) {
	client := &DigitalIdentityClient{}
	expectedURL := "expectedurl.com"
	client.OverrideAPIURL(expectedURL)
	assert.Equal(t, client.getAPIURL(), expectedURL)
}

func TestDigitalIDClient_GetAPIURLUsesOverriddenBaseUrlOverEnvVariable(t *testing.T) {
	client := DigitalIdentityClient{}
	client.OverrideAPIURL("overridenBaseUrl")

	os.Setenv("YOTI_API_URL", "envBaseUrl")

	result := client.getAPIURL()

	assert.Equal(t, "overridenBaseUrl", result)
}

func TestDigitalIDClient_GetAPIURLUsesEnvVariable(t *testing.T) {
	client := DigitalIdentityClient{}

	os.Setenv("YOTI_API_URL", "envBaseUrl")

	result := client.getAPIURL()

	assert.Equal(t, "envBaseUrl", result)
}

func TestDigitalIDClient_GetAPIURLUsesDefaultUrlAsFallbackWithEmptyEnvValue(t *testing.T) {
	client := DigitalIdentityClient{}

	os.Setenv("YOTI_API_URL", "")

	result := client.getAPIURL()

	assert.Equal(t, "https://api.yoti.com/share/", result)
}

func TestDigitalIDClient_GetAPIURLUsesDefaultUrlAsFallbackWithNoEnvValue(t *testing.T) {
	client := DigitalIdentityClient{}

	os.Unsetenv("YOTI_API_URL")

	result := client.getAPIURL()

	assert.Equal(t, "https://api.yoti.com/share/", result)
}

func getDigitalValidKey() *rsa.PrivateKey {
	return test.GetValidKey("test/test-key.pem")
}
