package dynamic

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/getyoti/yoti-go-sdk/v3/test"
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
