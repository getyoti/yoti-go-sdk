package docscan

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/getyoti/yoti-go-sdk/v3/docscan/session/create"
	"github.com/getyoti/yoti-go-sdk/v3/docscan/supported"
	"github.com/getyoti/yoti-go-sdk/v3/media"
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

func TestClient_CreateSession(t *testing.T) {
	key, _ := rsa.GenerateKey(rand.Reader, 1024)

	var clientSessionTokenTTL int = 100
	var clientSessionToken string = "8c91671a-7194-4ad7-8483-32703b965cfc"
	var sessionID string = "c87c4f2a-13fd-4cc8-a0e4-f1637cf32f71"

	jsonResponse := fmt.Sprintf(`{"client_session_token_ttl":%d,"client_session_token":"%s","session_id":"%s"}`, clientSessionTokenTTL, clientSessionToken, sessionID)

	HTTPClient := &mockHTTPClient{
		do: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusCreated,
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
	assert.Equal(t, sessionID, createSessionResult.SessionID)
}

func TestClient_GetSession(t *testing.T) {
	key, _ := rsa.GenerateKey(rand.Reader, 1024)

	var clientSessionTokenTTL int = 100
	var clientSessionToken string = "8c91671a-7194-4ad7-8483-32703b965cfc"
	var sessionID string = "c87c4f2a-13fd-4cc8-a0e4-f1637cf32f71"
	var userTrackingID string = "user-tracking-id"
	var state = "COMPLETED"
	jsonResponse := fmt.Sprintf(`{"client_session_token_ttl":%d,"client_session_token":"%s","session_id":"%s","user_tracking_id":"%s","state":"%s"}`, clientSessionTokenTTL, clientSessionToken, sessionID, userTrackingID, state)

	HTTPClient := &mockHTTPClient{
		do: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(strings.NewReader(jsonResponse)),
			}, nil
		},
	}

	client := Client{
		SdkID:      "sdkId",
		Key:        key,
		HTTPClient: HTTPClient,
		apiURL:     "https://apiurl.com",
	}

	getSessionResult, err := client.GetSession(sessionID)
	assert.NilError(t, err)

	assert.Equal(t, clientSessionTokenTTL, getSessionResult.ClientSessionTokenTTL)
	assert.Equal(t, clientSessionToken, getSessionResult.ClientSessionToken)
	assert.Equal(t, sessionID, getSessionResult.SessionID)
	assert.Equal(t, userTrackingID, getSessionResult.UserTrackingID)
	assert.Equal(t, state, getSessionResult.State)
}

func TestClient_DeleteSession(t *testing.T) {
	key, _ := rsa.GenerateKey(rand.Reader, 1024)
	HTTPClient := &mockHTTPClient{
		do: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
			}, nil
		},
	}

	client := Client{
		SdkID:      "sdkId",
		Key:        key,
		HTTPClient: HTTPClient,
		apiURL:     "https://apiurl.com",
	}

	err := client.DeleteSession("some-session-id")
	assert.NilError(t, err)
}

func TestClient_GetMediaContent(t *testing.T) {
	key, _ := rsa.GenerateKey(rand.Reader, 1024)

	jpegImage := []byte("value")
	HTTPClient := &mockHTTPClient{
		do: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewReader(jpegImage)),
				Header:     map[string][]string{"Content-Type": {media.ImageTypeJPEG}},
			}, nil
		},
	}

	client := Client{
		SdkID:      "sdkId",
		Key:        key,
		HTTPClient: HTTPClient,
		apiURL:     "https://apiurl.com",
	}

	result, err := client.GetMediaContent("some-sessionID", "some-mediaID")
	assert.NilError(t, err)

	assert.DeepEqual(t, jpegImage, result.Data())
	assert.Equal(t, media.ImageTypeJPEG, result.MIME())
}

func TestClient_DeleteMediaContent(t *testing.T) {
	key, _ := rsa.GenerateKey(rand.Reader, 1024)
	HTTPClient := &mockHTTPClient{
		do: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
			}, nil
		},
	}

	client := Client{
		SdkID:      "sdkId",
		Key:        key,
		HTTPClient: HTTPClient,
		apiURL:     "https://apiurl.com",
	}

	err := client.DeleteMediaContent("some-session-id", "some-media-id")
	assert.NilError(t, err)
}

func TestClient_GetSupportedDocuments(t *testing.T) {
	key, _ := rsa.GenerateKey(rand.Reader, 1024)
	countryCodeUSA := "USA"
	documentTypePassport := "PASSPORT"
	documentsResponse := supported.DocumentsResponse{
		SupportedCountries: []*supported.Country{
			{
				Code: countryCodeUSA,
				SupportedDocuments: []*supported.Document{
					{
						Type: documentTypePassport,
					},
				},
			},
		},
	}

	jsonBytes, err := json.Marshal(documentsResponse)
	assert.NilError(t, err)

	HTTPClient := &mockHTTPClient{
		do: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewReader(jsonBytes)),
			}, nil
		},
	}

	client := Client{
		SdkID:      "sdkId",
		Key:        key,
		HTTPClient: HTTPClient,
		apiURL:     "https://apiurl.com",
	}

	result, err := client.GetSupportedDocuments()
	assert.NilError(t, err)

	assert.Equal(t, result.SupportedCountries[0].Code, countryCodeUSA)
	assert.Equal(t, result.SupportedCountries[0].SupportedDocuments[0].Type, documentTypePassport)
}

func TestNewClient(t *testing.T) {
	key, err := ioutil.ReadFile("../test/test-key.pem")
	assert.NilError(t, err)

	client, err := NewClient("sdkID", key)
	assert.NilError(t, err)

	assert.Equal(t, "sdkID", client.SdkID)
}

func TestNewClient_EmptySdkID(t *testing.T) {
	_, err := NewClient("", []byte("someKey"))

	assert.ErrorContains(t, err, "SdkID cannot be an empty string")
}

func TestClient_GetSession_EmptySessionID(t *testing.T) {
	key, err := ioutil.ReadFile("../test/test-key.pem")
	assert.NilError(t, err)

	client, err := NewClient("sdkID", key)
	assert.NilError(t, err)

	_, err = client.GetSession("")
	assert.ErrorContains(t, err, "sessionID cannot be an empty string")
}

func TestClient_DeleteSession_EmptySessionID(t *testing.T) {
	key, err := ioutil.ReadFile("../test/test-key.pem")
	assert.NilError(t, err)

	client, err := NewClient("sdkID", key)
	assert.NilError(t, err)

	err = client.DeleteSession("")
	assert.ErrorContains(t, err, "sessionID cannot be an empty string")
}

func TestClient_GetMediaContent_EmptySessionID(t *testing.T) {
	key, err := ioutil.ReadFile("../test/test-key.pem")
	assert.NilError(t, err)

	client, err := NewClient("sdkID", key)
	assert.NilError(t, err)

	_, err = client.GetMediaContent("", "someMediaID")
	assert.ErrorContains(t, err, "sessionID cannot be an empty string")
}

func TestClient_GetMediaContent_EmptyMediaID(t *testing.T) {
	key, err := ioutil.ReadFile("../test/test-key.pem")
	assert.NilError(t, err)

	client, err := NewClient("sdkID", key)
	assert.NilError(t, err)

	_, err = client.GetMediaContent("someSessionID", "")
	assert.ErrorContains(t, err, "mediaID cannot be an empty string")
}

func TestClient_DeleteMediaContent_EmptySessionID(t *testing.T) {
	key, err := ioutil.ReadFile("../test/test-key.pem")
	assert.NilError(t, err)

	client, err := NewClient("sdkID", key)
	assert.NilError(t, err)

	err = client.DeleteMediaContent("", "someMediaID")
	assert.ErrorContains(t, err, "sessionID cannot be an empty string")
}

func TestClient_DeleteMediaContent_EmptyMediaID(t *testing.T) {
	key, err := ioutil.ReadFile("../test/test-key.pem")
	assert.NilError(t, err)

	client, err := NewClient("sdkID", key)
	assert.NilError(t, err)

	err = client.DeleteMediaContent("someSessionID", "")
	assert.ErrorContains(t, err, "mediaID cannot be an empty string")
}

func Test_EmptySdkID(t *testing.T) {
	key, err := ioutil.ReadFile("../test/test-key.pem")
	assert.NilError(t, err)

	client, err := NewClient("sdkID", key)
	assert.NilError(t, err)

	_, err = client.GetSession("")
	assert.ErrorContains(t, err, "sessionID cannot be an empty string")
}
