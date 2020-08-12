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

	policy, err := (&PolicyBuilder{}).WithFullName().WithWantedRememberMe().Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	scenario, err := (&ScenarioBuilder{}).WithPolicy(policy).Build()
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	result, err := CreateShareURL(client, &scenario, "sdkId", "https://apiurl", key)

	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	fmt.Printf("QR code: %s", result.ShareURL)
	// Output: QR code: https://code.yoti.com/CAEaJDQzNzllZDc0LTU0YjItNDkxMy04OTE4LTExYzM2ZDU2OTU3ZDAC
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

	policy, err := (&PolicyBuilder{}).WithFullName().WithWantedRememberMe().Build()
	if err != nil {
		return
	}
	scenario, err := (&ScenarioBuilder{}).WithPolicy(policy).Build()
	if err != nil {
		return
	}

	return CreateShareURL(client, &scenario, "sdkId", "https://apiurl", key)
}
