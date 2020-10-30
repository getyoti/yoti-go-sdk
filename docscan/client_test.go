package docscan

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
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

func TestClient_CreateSession_ShouldReturnJsonMarshalError(t *testing.T) {
	client := Client{
		jsonMarshaler: &mockJSONMarshaler{
			marshal: func(v interface{}) ([]byte, error) {
				return []byte{}, errors.New("some json error")
			},
		},
	}
	_, err := client.CreateSession(&create.SessionSpecification{})
	assert.ErrorContains(t, err, "some json error")
}

func TestClient_CreateSession_ShouldReturnMissingKeyError(t *testing.T) {
	client := Client{}
	_, err := client.CreateSession(&create.SessionSpecification{})
	assert.ErrorContains(t, err, "missing private key")
}

func TestClient_CreateSession_ShouldReturnResponseError(t *testing.T) {
	key, _ := rsa.GenerateKey(rand.Reader, 1024)

	HTTPClient := &mockHTTPClient{
		do: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusBadRequest,
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
	_, err := client.CreateSession(sessionSpec)
	assert.ErrorContains(t, err, "400: unknown HTTP error")
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

func TestClient_GetSession_ShouldReturnMissingKeyError(t *testing.T) {
	client := Client{}
	_, err := client.GetSession("some-id")
	assert.ErrorContains(t, err, "missing private key")
}

func TestClient_GetSession_ShouldReturnResponseError(t *testing.T) {
	key, _ := rsa.GenerateKey(rand.Reader, 1024)

	HTTPClient := &mockHTTPClient{
		do: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusBadRequest,
			}, nil
		},
	}

	client := Client{
		SdkID:      "sdkId",
		Key:        key,
		HTTPClient: HTTPClient,
		apiURL:     "https://apiurl.com",
	}

	_, err := client.GetSession("some-id")
	assert.ErrorContains(t, err, "400: unknown HTTP error")
}

func TestClient_GetSession_ShouldReturnJsonError(t *testing.T) {
	key, _ := rsa.GenerateKey(rand.Reader, 1024)

	HTTPClient := &mockHTTPClient{
		do: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(strings.NewReader("some-invalid-json")),
			}, nil
		},
	}

	client := Client{
		SdkID:      "sdkId",
		Key:        key,
		HTTPClient: HTTPClient,
		apiURL:     "https://apiurl.com",
	}

	_, err := client.GetSession("some-id")
	assert.ErrorContains(t, err, "invalid character")
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

func TestClient_DeleteSession_ShouldReturnMissingKeyError(t *testing.T) {
	client := Client{}
	err := client.DeleteSession("some-id")
	assert.ErrorContains(t, err, "missing private key")
}

func TestClient_DeleteSession_ShouldReturnResponseError(t *testing.T) {
	key, _ := rsa.GenerateKey(rand.Reader, 1024)

	HTTPClient := &mockHTTPClient{
		do: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusBadRequest,
			}, nil
		},
	}

	client := Client{
		SdkID:      "sdkId",
		Key:        key,
		HTTPClient: HTTPClient,
		apiURL:     "https://apiurl.com",
	}

	err := client.DeleteSession("some-id")
	assert.ErrorContains(t, err, "400: unknown HTTP error")
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

func TestClient_GetMediaContent_NoContent(t *testing.T) {
	key, _ := rsa.GenerateKey(rand.Reader, 1024)

	HTTPClient := &mockHTTPClient{
		do: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusNoContent,
				Header:     map[string][]string{},
			}, nil
		},
	}

	client := Client{
		SdkID:      "sdkId",
		Key:        key,
		HTTPClient: HTTPClient,
		apiURL:     "https://apiurl.com",
	}

	media, err := client.GetMediaContent("some-sessionID", "some-mediaID")
	assert.Equal(t, media, nil)
	assert.NilError(t, err)
}

func TestClient_GetMediaContent_NoContentType(t *testing.T) {
	key, _ := rsa.GenerateKey(rand.Reader, 1024)

	jpegImage := []byte("value")
	HTTPClient := &mockHTTPClient{
		do: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewReader(jpegImage)),
				Header:     map[string][]string{},
			}, nil
		},
	}

	client := Client{
		SdkID:      "sdkId",
		Key:        key,
		HTTPClient: HTTPClient,
		apiURL:     "https://apiurl.com",
	}

	_, err := client.GetMediaContent("some-sessionID", "some-mediaID")
	assert.ErrorContains(t, err, "unable to parse content type from response")
}

func TestClient_GetMediaContent_ShouldReturnMissingKeyError(t *testing.T) {
	client := Client{}
	_, err := client.GetMediaContent("some-id", "some-media-id")
	assert.ErrorContains(t, err, "missing private key")
}

func TestClient_GetMediaContent_ShouldReturnResponseError(t *testing.T) {
	key, _ := rsa.GenerateKey(rand.Reader, 1024)

	HTTPClient := &mockHTTPClient{
		do: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusBadRequest,
			}, nil
		},
	}

	client := Client{
		SdkID:      "sdkId",
		Key:        key,
		HTTPClient: HTTPClient,
		apiURL:     "https://apiurl.com",
	}

	_, err := client.GetMediaContent("some-id", "some-media-id")
	assert.ErrorContains(t, err, "400: unknown HTTP error")
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

func TestClient_DeleteMediaContent_ShouldReturnMissingKeyError(t *testing.T) {
	client := Client{}
	err := client.DeleteMediaContent("some-id", "some-media-id")
	assert.ErrorContains(t, err, "missing private key")
}

func TestClient_DeleteMediaContent_ShouldReturnResponseError(t *testing.T) {
	key, _ := rsa.GenerateKey(rand.Reader, 1024)

	HTTPClient := &mockHTTPClient{
		do: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusBadRequest,
			}, nil
		},
	}

	client := Client{
		SdkID:      "sdkId",
		Key:        key,
		HTTPClient: HTTPClient,
		apiURL:     "https://apiurl.com",
	}

	err := client.DeleteMediaContent("some-id", "some-media-id")
	assert.ErrorContains(t, err, "400: unknown HTTP error")
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

func TestClient_GetSupportedDocuments_ShouldReturnMissingKeyError(t *testing.T) {
	client := Client{}
	_, err := client.GetSupportedDocuments()
	assert.ErrorContains(t, err, "missing private key")
}

func TestClient_GetSupportedDocuments_ShouldReturnResponseError(t *testing.T) {
	key, _ := rsa.GenerateKey(rand.Reader, 1024)

	HTTPClient := &mockHTTPClient{
		do: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusBadRequest,
			}, nil
		},
	}

	client := Client{
		SdkID:      "sdkId",
		Key:        key,
		HTTPClient: HTTPClient,
		apiURL:     "https://apiurl.com",
	}

	_, err := client.GetSupportedDocuments()
	assert.ErrorContains(t, err, "400: unknown HTTP error")
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

func TestNewClient_KeyLoad_Failure(t *testing.T) {
	key, _ := ioutil.ReadFile("../test/test-key-invalid-format.pem")
	_, err := NewClient("sdkID", key)

	assert.ErrorContains(t, err, "invalid key: not PEM-encoded")

	tempError, temporary := err.(interface {
		Temporary() bool
	})
	assert.Check(t, !temporary || !tempError.Temporary())
}

func TestClient_UsesDefaultApiUrl(t *testing.T) {
	key, err := ioutil.ReadFile("../test/test-key.pem")
	assert.NilError(t, err)

	client, err := NewClient("sdkID", key)
	assert.NilError(t, err)

	assert.Equal(t, "https://api.yoti.com/idverify/v1", client.apiURL)
}

func TestClient_UsesEnvVariable(t *testing.T) {
	key, err := ioutil.ReadFile("../test/test-key.pem")
	assert.NilError(t, err)

	os.Setenv("YOTI_DOC_SCAN_API_URL", "envBaseUrl")

	client, err := NewClient("sdkID", key)
	assert.NilError(t, err)

	assert.Equal(t, "envBaseUrl", client.apiURL)
}

func TestClient_UsesOverrideApiUrlOverEnvVariable(t *testing.T) {
	key, err := ioutil.ReadFile("../test/test-key.pem")
	assert.NilError(t, err)

	os.Setenv("YOTI_DOC_SCAN_API_URL", "envBaseUrl")

	client, err := NewClient("sdkID", key)
	assert.NilError(t, err)

	client.OverrideAPIURL("overrideApiURL")

	assert.Equal(t, "overrideApiURL", client.apiURL)
}

type mockJSONMarshaler struct {
	marshal func(v interface{}) ([]byte, error)
}

func (mock *mockJSONMarshaler) Marshal(v interface{}) ([]byte, error) {
	if mock.marshal != nil {
		return mock.marshal(v)
	}
	return nil, nil
}
