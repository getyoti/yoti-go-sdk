package digitalidentity

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
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

func ExampleCreateShareSession() {
	key := test.GetValidKey("../test/test-key.pem")

	client := &mockHTTPClient{
		do: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 201,
				Body:       io.NopCloser(strings.NewReader(`{"status":"success","ref_id":"0"}`)),
			}, nil
		},
	}

	policy, err := (&PolicyBuilder{}).WithFullName().WithWantedRememberMe().Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	session, err := (&ShareSessionBuilder{}).WithPolicy(policy).Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	result, err := CreateShareSession(client, session, "sdkId", "https://apiurl", key)

	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	fmt.Printf("Status code: %s", result.Status)
	// Output: Status code: success
}

func TestCreateShareURL_Unsuccessful_503(t *testing.T) {
	_, err := createShareUrlWithErrorResponse(503, "some service unavailable response")

	assert.ErrorContains(t, err, "503: unknown HTTP error - some service unavailable response")

	tempError, temporary := err.(interface {
		Temporary() bool
	})
	assert.Check(t, temporary && tempError.Temporary())
}

func TestCreateShareURL_Unsuccessful_404(t *testing.T) {
	_, err := createShareUrlWithErrorResponse(404, "some not found response")

	assert.ErrorContains(t, err, "404: Application was not found - some not found response")

	tempError, temporary := err.(interface {
		Temporary() bool
	})
	assert.Check(t, !temporary || !tempError.Temporary())
}

func TestCreateShareURL_Unsuccessful_400(t *testing.T) {
	_, err := createShareUrlWithErrorResponse(400, "some invalid JSON response")

	assert.ErrorContains(t, err, "400: JSON is incorrect, contains invalid data - some invalid JSON response")

	tempError, temporary := err.(interface {
		Temporary() bool
	})
	assert.Check(t, !temporary || !tempError.Temporary())
}

func createShareUrlWithErrorResponse(statusCode int, responseBody string) (share ShareSession, err error) {
	key := test.GetValidKey("../test/test-key.pem")

	client := &mockHTTPClient{
		do: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: statusCode,
				Body:       io.NopCloser(strings.NewReader(responseBody)),
			}, nil
		},
	}

	policy, err := (&PolicyBuilder{}).WithFullName().WithWantedRememberMe().Build()
	if err != nil {
		return
	}
	session, err := (&ShareSessionBuilder{}).WithPolicy(policy).Build()
	if err != nil {
		return
	}

	return CreateShareSession(client, session, "sdkId", "https://apiurl", key)
}

type MockHttpClient struct{}

func (m *MockHttpClient) Do(req *http.Request) (*http.Response, error) {
	// Simulate a successful response
	mockResponse := httptest.NewRecorder()
	mockResponse.WriteString(`{
		"id": "SOME_ID",
		"status": "SOME_STATUS",
		"expiry": "SOME_EXPIRY",
		"created": "SOME_CREATED",
		"updated": "SOME_UPDATED",
		"qrCode": {"id": "SOME_QRCODE_ID"},
		"receipt": {"id": "SOME_RECEIPT_ID"}
	}`)
	return mockResponse.Result(), nil
}

func TestGetSession(t *testing.T) {
	mockPrivateKey := test.GetValidKey("../test/test-key.pem")

	mockHttpClient := &MockHttpClient{}

	mockSessionID := "SOME_SESSION_ID"
	mockClientSdkId := "SOME_CLIENT_SDK_ID"
	mockApiUrl := "https://example.com/api"

	result, err := GetSession(mockHttpClient, mockSessionID, mockClientSdkId, mockApiUrl, mockPrivateKey)

	assert.NilError(t, err)

	assert.Equal(t, "SOME_ID", result.Id)
	assert.Equal(t, "SOME_STATUS", result.Status)
	assert.Equal(t, "SOME_EXPIRY", result.Expiry)
	assert.Equal(t, "SOME_CREATED", result.Created)
	assert.Equal(t, "SOME_UPDATED", result.Updated)
	assert.Equal(t, "SOME_QRCODE_ID", result.QrCode.Id)
	assert.Equal(t, "SOME_RECEIPT_ID", result.Receipt.Id)
}
