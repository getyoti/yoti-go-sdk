package yoti

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/getyoti/yoti-go-sdk/v3/digitalidentity"
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

	session, err := (&digitalidentity.ShareSessionBuilder{}).WithPolicy(policy).Build()
	assert.NilError(t, err)

	result, err := client.CreateShareSession(session)

	b, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

	assert.NilError(t, err)
	assert.Equal(t, result.Status, "SOME_STATUS")
}

func TestDigitalIDClient_HttpFailure_ReturnsApplicationNotFound(t *testing.T) {
	key := getValidKey()

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
