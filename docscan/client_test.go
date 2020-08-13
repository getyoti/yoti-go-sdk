package docscan

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/getyoti/yoti-go-sdk/v3/docscan/session/create"
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

func ExampleClient_CreateSession() {
	// session := CreateSession(nil, )
}

func TestCreateSession(t *testing.T) {
	key, _ := rsa.GenerateKey(rand.Reader, 1024)

	var clientSessionTokenTTL int = 100
	var clientSessionToken string = "8c91671a-7194-4ad7-8483-32703b965cfc"
	var sessionId string = "c87c4f2a-13fd-4cc8-a0e4-f1637cf32f71"

	jsonResponse := fmt.Sprintf(`{"client_session_token_ttl":%d,"client_session_token":"%s","session_id":"%s"}`, clientSessionTokenTTL, clientSessionToken, sessionId)

	HTTPClient := &mockHTTPClient{
		do: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 201,
				Body:       ioutil.NopCloser(strings.NewReader(jsonResponse)),
			}, nil
		},
	}

	var sessionSpec *create.SessionSpecification

	client := Client{
		SdkID:      "sdkId",
		Key:        key,
		HTTPClient: HTTPClient,
		apiURL:     "https://apiurl.com",
	}
	createSessionResult, err := client.CreateSession(sessionSpec)
	assert.NilError(t, err)

	assert.Equal(t, clientSessionTokenTTL, createSessionResult.ClientSessionTokenTTL)
	assert.Equal(t, clientSessionToken, createSessionResult.ClientSessionToken)
	assert.Equal(t, sessionId, createSessionResult.SessionID)
}
