package yoti

import (
	"io"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/getyoti/yoti-go-sdk/v3/dynamic"
	"gotest.tools/v3/assert"
)

func TestDigitalIDClient(t *testing.T) {
	key, readErr := os.ReadFile("./test/test-key.pem")
	assert.NilError(t, readErr)

	_, err := NewDigitalIdentityClient("some-sdk-id", key)
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

func TestDigitalIDClient_CreateShareURL(t *testing.T) {
	key, readErr := os.ReadFile("./test/test-key.pem")
	assert.NilError(t, readErr)

	client, clientErr := NewDigitalIdentityClient("some-sdk-id", key)
	assert.NilError(t, clientErr)

	client.HTTPClient = &mockHTTPClient{
		do: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 201,
				Body:       io.NopCloser(strings.NewReader(`{"qrcode":"https://code.yoti.com/some-qr","ref_id":"0"}`)),
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
