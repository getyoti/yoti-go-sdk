package dynamic

import (
	"fmt"
	"io/ioutil"
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

func ExampleCreateShareURL() {
	key := test.GetValidKey("../test/test-key.pem")

	client := &mockHTTPClient{
		do: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 201,
				Body:       ioutil.NopCloser(strings.NewReader(`{"qrcode":"https://code.yoti.com/CAEaJDQzNzllZDc0LTU0YjItNDkxMy04OTE4LTExYzM2ZDU2OTU3ZDAC","ref_id":"0"}`)),
			}, nil
		},
	}

	policy, err := (&DynamicPolicyBuilder{}).WithFullName().WithWantedRememberMe().Build()
	if err != nil {
		return
	}
	scenario, err := (&DynamicScenarioBuilder{}).WithPolicy(policy).Build()
	if err != nil {
		return
	}

	result, err := CreateShareURL(client, &scenario, "sdkId", "https://apiurl", key)

	if err != nil {
		return
	}
	fmt.Printf("QR code: %s", result.ShareURL)
	// Output: QR code: https://code.yoti.com/CAEaJDQzNzllZDc0LTU0YjItNDkxMy04OTE4LTExYzM2ZDU2OTU3ZDAC
}

func TestCreateShareURL_Unsuccessful_503(t *testing.T) {
	_, err := createShareUrlWithErrorResponse(503, "SERVICE UNAVAILABLE")

	assert.ErrorContains(t, err, "503: unknown HTTP error: SERVICE UNAVAILABLE")

	tempError, temporary := err.(interface {
		Temporary() bool
	})
	assert.Check(t, temporary && tempError.Temporary())
}

func TestCreateShareURL_Unsuccessful_404(t *testing.T) {
	_, err := createShareUrlWithErrorResponse(404, "NOT FOUND")

	assert.ErrorContains(t, err, "404: Application was not found: NOT FOUND")

	tempError, temporary := err.(interface {
		Temporary() bool
	})
	assert.Check(t, !temporary || !tempError.Temporary())
}

func TestCreateShareURL_Unsuccessful_400(t *testing.T) {
	_, err := createShareUrlWithErrorResponse(400, "INVALID JSON")

	assert.ErrorContains(t, err, "400: JSON is incorrect, contains invalid data: INVALID JSON")

	tempError, temporary := err.(interface {
		Temporary() bool
	})
	assert.Check(t, !temporary || !tempError.Temporary())
}

func createShareUrlWithErrorResponse(statusCode int, responseBody string) (share ShareURL, err error) {
	key := test.GetValidKey("../test/test-key.pem")

	client := &mockHTTPClient{
		do: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: statusCode,
				Body:       ioutil.NopCloser(strings.NewReader(responseBody)),
			}, nil
		},
	}

	policy, err := (&DynamicPolicyBuilder{}).WithFullName().WithWantedRememberMe().Build()
	if err != nil {
		return
	}
	scenario, err := (&DynamicScenarioBuilder{}).WithPolicy(policy).Build()
	if err != nil {
		return
	}

	return CreateShareURL(client, &scenario, "sdkId", "https://apiurl", key)
}
