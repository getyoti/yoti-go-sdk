package yoti

import (
	"crypto/rsa"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/getyoti/yoti-go-sdk/v3/aml"
	"github.com/getyoti/yoti-go-sdk/v3/dynamic"
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

func TestNewClient(t *testing.T) {
	key, readErr := ioutil.ReadFile("./test/test-key.pem")
	assert.NilError(t, readErr)

	_, err := NewClient("some-sdk-id", key)
	assert.NilError(t, err)
}

func TestNewClient_KeyLoad_Failure(t *testing.T) {
	key, _ := ioutil.ReadFile("test/test-key-invalid-format.pem")
	_, err := NewClient("", key)

	assert.ErrorContains(t, err, "invalid key: not PEM-encoded")

	tempError, temporary := err.(interface {
		Temporary() bool
	})
	assert.Check(t, !temporary || !tempError.Temporary())
}

func TestYotiClient_PerformAmlCheck(t *testing.T) {
	key, readErr := ioutil.ReadFile("./test/test-key.pem")
	assert.NilError(t, readErr)

	client, clientErr := NewClient("some-sdk-id", key)
	assert.NilError(t, clientErr)

	client.HTTPClient = &mockHTTPClient{
		do: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       ioutil.NopCloser(strings.NewReader(`{"on_fraud_list":true}`)),
			}, nil
		},
	}

	var amlAddress = aml.Address{
		Country: "GBR"}

	var amlProfile = aml.Profile{
		GivenNames: "Edward Richard George",
		FamilyName: "Heath",
		Address:    amlAddress}

	result, err := client.PerformAmlCheck(amlProfile)
	assert.NilError(t, err)

	assert.Check(t, result.OnFraudList)
}

func TestYotiClient_CreateShareURL(t *testing.T) {
	key, readErr := ioutil.ReadFile("./test/test-key.pem")
	assert.NilError(t, readErr)

	client, clientErr := NewClient("some-sdk-id", key)
	assert.NilError(t, clientErr)

	client.HTTPClient = &mockHTTPClient{
		do: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 201,
				Body:       ioutil.NopCloser(strings.NewReader(`{"qrcode":"https://code.yoti.com/some-qr","ref_id":"0"}`)),
			}, nil
		},
	}

	policy, policyErr := (&dynamic.PolicyBuilder{}).WithFullName().WithWantedRememberMe().Build()
	assert.NilError(t, policyErr)

	scenario, scenarioErr := (&dynamic.ScenarioBuilder{}).WithPolicy(policy).Build()
	assert.NilError(t, scenarioErr)

	result, err := client.CreateShareURL(&scenario)
	assert.NilError(t, err)
	assert.Equal(t, result.ShareURL, "https://code.yoti.com/some-qr")
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
	})
	assert.Check(t, temporary)
	assert.Check(t, tempError.Temporary())
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

	assert.ErrorContains(t, err, "Profile not found")
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
