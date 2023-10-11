package digitalidentity

import (
	"fmt"
	"io"
	"net/http"
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
				Body:       io.NopCloser(strings.NewReader(`{"id":"0","status":"success","expiry": ""}`)),
			}, nil
		},
	}

	policy, err := (&PolicyBuilder{}).WithFullName().WithWantedRememberMe().Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	session, err := (&ShareSessionRequestBuilder{}).WithPolicy(policy).Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	result, err := CreateShareSession(client, &session, "sdkId", "https://apiurl", key)

	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	fmt.Printf("Status code: %s", result.Status)
	// Output: Status code: success
}

func TestCreateShareURL_Unsuccessful_401(t *testing.T) {
	_, err := createShareSessionWithErrorResponse(401, `{"id":"8f6a9dfe72128de20909af0d476769b6","status":401,"error":"INVALID_REQUEST_SIGNATURE","message":"Invalid request signature"}`)

	assert.ErrorContains(t, err, "INVALID_REQUEST_SIGNATURE")

	tempError, temporary := err.(interface {
		Temporary() bool
	})
	assert.Check(t, !temporary || !tempError.Temporary())
}

func createShareSessionWithErrorResponse(statusCode int, responseBody string) (*ShareSession, error) {
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
		return nil, err
	}
	session, err := (&ShareSessionRequestBuilder{}).WithPolicy(policy).Build()
	if err != nil {
		return nil, err
	}

	return CreateShareSession(client, &session, "sdkId", "https://apiurl", key)
}

func ExampleGetShareSession() {
	key := test.GetValidKey("../test/test-key.pem")
	mockSessionID := "SOME_SESSION_ID"
	mockClientSdkId := "SOME_CLIENT_SDK_ID"
	mockApiUrl := "https://example.com/api"
	client := &mockHTTPClient{
		do: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 201,
				Body:       io.NopCloser(strings.NewReader(`{"id":"SOME_ID","status":"SOME_STATUS","expiry":"SOME_EXPIRY","created":"SOME_CREATED","updated":"SOME_UPDATED","qrCode":{"id":"SOME_QRCODE_ID"},"receipt":{"id":"SOME_RECEIPT_ID"}}`)),
			}, nil
		},
	}

	result, err := GetShareSession(client, mockSessionID, mockClientSdkId, mockApiUrl, key)
	if err != nil {
		return
	}
	fmt.Printf("Status code: %s", result.Status)
	// Output:Status code: SOME_STATUS
}
