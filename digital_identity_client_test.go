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

func TestDigitalIDClient_HttpFailure_ReturnsUnKnownHttpError(t *testing.T) {
	key := getDigitalValidKey()
	client := DigitalIdentityClient{
		HTTPClient: &mockHTTPClient{
			do: func(*http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 401,
				}, nil
			},
		},
		SdkID: "some-sdk-id",
		Key:   key,
	}

	_, err := client.GetShareSession("SOME ID")

	assert.ErrorContains(t, err, "unknown HTTP error")
	tempError, temporary := err.(interface {
		Temporary() bool
	})
	assert.Check(t, !temporary || !tempError.Temporary())
}

func TestDigitalIDClient_GetSession(t *testing.T) {
	key, err := os.ReadFile("./test/test-key.pem")
	if err != nil {
		t.Fatalf("failed to read pem file :: %v", err)
	}

	mockSessionID := "SOME_SESSION_ID"
	client, err := NewDigitalIdentityClient("some-sdk-id", key)
	if err != nil {
		t.Fatalf("failed to build the DigitalIdClient :: %v", err)
	}

	client.HTTPClient = &mockHTTPClient{
		do: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(strings.NewReader(`{"id":"SOME_ID","status":"SOME_STATUS","expiry":"SOME_EXPIRY","created":"SOME_CREATED","updated":"SOME_UPDATED","qrCode":{"id":"SOME_QRCODE_ID"},"receipt":{"id":"SOME_RECEIPT_ID"}}`)),
			}, nil
		},
	}

	result, err := client.GetShareSession(mockSessionID)
	if err != nil {
		t.Fatalf("failed to GetShareSesssion :: %v", err)
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

	assert.Equal(t, "https://api.yoti.com/share", result)
}

func TestDigitalIDClient_GetAPIURLUsesDefaultUrlAsFallbackWithNoEnvValue(t *testing.T) {
	client := DigitalIdentityClient{}

	os.Unsetenv("YOTI_API_URL")
	result := client.getAPIURL()

	assert.Equal(t, "https://api.yoti.com/share", result)
}

func getDigitalValidKey() *rsa.PrivateKey {
	return test.GetValidKey("test/test-key.pem")
}

func TestNewDigitalIdentityClientWithToken(t *testing.T) {
	client, err := NewDigitalIdentityClientWithToken("my-auth-token")
	assert.NilError(t, err)
	assert.Equal(t, client.AuthToken, "my-auth-token")
}

func TestNewDigitalIdentityClientWithToken_EmptyToken(t *testing.T) {
	_, err := NewDigitalIdentityClientWithToken("")
	assert.ErrorContains(t, err, "authentication token must not be empty")
}

func TestDigitalIDClient_GetAuthStrategy_BearerToken(t *testing.T) {
	client := DigitalIdentityClient{
		AuthToken: "my-bearer-token",
	}

	strategy, err := client.GetAuthStrategy()
	assert.NilError(t, err)

	headers, err := strategy.CreateAuthHeaders("GET", "/v2/sessions", nil)
	assert.NilError(t, err)
	assert.Equal(t, headers["Authorization"][0], "Bearer my-bearer-token")
}

func TestDigitalIDClient_GetAuthStrategy_SignedRequest(t *testing.T) {
	key := getDigitalValidKey()
	client := DigitalIdentityClient{
		SdkID: "test-sdk-id",
		Key:   key,
	}

	strategy, err := client.GetAuthStrategy()
	assert.NilError(t, err)

	params, err := strategy.CreateQueryParams()
	assert.NilError(t, err)
	assert.Equal(t, params["sdkID"], "test-sdk-id")
}

func TestDigitalIDClient_GetAuthStrategy_MutualExclusion(t *testing.T) {
	key := getDigitalValidKey()
	client := DigitalIdentityClient{
		AuthToken: "my-token",
		SdkID:     "test-sdk-id",
		Key:       key,
	}

	_, err := client.GetAuthStrategy()
	assert.ErrorContains(t, err, "must not supply both authentication token and SDK credentials")
}

func TestDigitalIDClient_GetAuthStrategy_MissingKey(t *testing.T) {
	client := DigitalIdentityClient{
		SdkID: "test-sdk-id",
	}

	_, err := client.GetAuthStrategy()
	assert.ErrorContains(t, err, "missing private key")
}

func TestDigitalIDClient_GetAuthStrategy_MissingSdkId(t *testing.T) {
	key := getDigitalValidKey()
	client := DigitalIdentityClient{
		Key: key,
	}

	_, err := client.GetAuthStrategy()
	assert.ErrorContains(t, err, "missing SDK ID")
}

func TestDigitalIDClient_WithToken_CreateShareSession(t *testing.T) {
	client, err := NewDigitalIdentityClientWithToken("test-token")
	assert.NilError(t, err)

	client.HTTPClient = &mockHTTPClient{
		do: func(req *http.Request) (*http.Response, error) {
			// Verify bearer token is sent
			assert.Equal(t, req.Header.Get("Authorization"), "Bearer test-token")
			// Verify no signed-request artifacts
			assert.Equal(t, req.Header.Get("X-Yoti-Auth-Digest"), "")
			assert.Equal(t, req.Header.Get("X-Yoti-Auth-Id"), "")
			// Verify no nonce/timestamp query params
			assert.Equal(t, req.URL.Query().Get("nonce"), "")
			assert.Equal(t, req.URL.Query().Get("timestamp"), "")

			return &http.Response{
				StatusCode: 201,
				Body:       io.NopCloser(strings.NewReader(`{"id":"SOME_ID","status":"SOME_STATUS","expiry":"SOME_EXPIRY"}`)),
			}, nil
		},
	}

	policy, err := (&digitalidentity.PolicyBuilder{}).WithFullName().Build()
	assert.NilError(t, err)

	session, err := (&digitalidentity.ShareSessionRequestBuilder{}).WithPolicy(policy).Build()
	assert.NilError(t, err)

	result, err := client.CreateShareSession(&session)
	assert.NilError(t, err)
	assert.Equal(t, result.Status, "SOME_STATUS")
}

func TestDigitalIDClient_WithToken_GetShareReceiptReturnsErrorNotPanic(t *testing.T) {
	client, err := NewDigitalIdentityClientWithToken("test-token")
	assert.NilError(t, err)

	_, err = client.GetShareReceipt("some-receipt-id")
	assert.ErrorContains(t, err, "GetShareReceipt requires a private key")
}
