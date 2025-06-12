package docscan

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/getyoti/yoti-go-sdk/v3/docscan/session/create"
	"github.com/getyoti/yoti-go-sdk/v3/docscan/session/create/facecapture"
	"github.com/getyoti/yoti-go-sdk/v3/docscan/session/retrieve"
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
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	assert.NilError(t, err)

	var clientSessionTokenTTL int = 100
	var clientSessionToken string = "8c91671a-7194-4ad7-8483-32703b965cfc"
	var sessionID string = "c87c4f2a-13fd-4cc8-a0e4-f1637cf32f71"

	jsonResponse := fmt.Sprintf(`{"client_session_token_ttl":%d,"client_session_token":"%s","session_id":"%s"}`, clientSessionTokenTTL, clientSessionToken, sessionID)

	HTTPClient := &mockHTTPClient{
		do: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusCreated,
				Body:       io.NopCloser(strings.NewReader(jsonResponse)),
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
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	assert.NilError(t, err)

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
	_, err = client.CreateSession(sessionSpec)
	assert.ErrorContains(t, err, "400: unknown HTTP error")
}

func TestClient_GetSession(t *testing.T) {
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	assert.NilError(t, err)

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
				Body:       io.NopCloser(strings.NewReader(jsonResponse)),
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
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	assert.NilError(t, err)

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

	_, err = client.GetSession("some-id")
	assert.ErrorContains(t, err, "400: unknown HTTP error")
}

func TestClient_GetSession_ShouldReturnJsonError(t *testing.T) {
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	assert.NilError(t, err)

	HTTPClient := &mockHTTPClient{
		do: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader("some-invalid-json")),
			}, nil
		},
	}

	client := Client{
		SdkID:      "sdkId",
		Key:        key,
		HTTPClient: HTTPClient,
		apiURL:     "https://apiurl.com",
	}

	_, err = client.GetSession("some-id")
	assert.ErrorContains(t, err, "invalid character")
}

func TestClient_DeleteSession(t *testing.T) {
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	assert.NilError(t, err)

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

	err = client.DeleteSession("some-session-id")
	assert.NilError(t, err)
}

func TestClient_DeleteSession_ShouldReturnMissingKeyError(t *testing.T) {
	client := Client{}
	err := client.DeleteSession("some-id")
	assert.ErrorContains(t, err, "missing private key")
}

func TestClient_DeleteSession_ShouldReturnResponseError(t *testing.T) {
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	assert.NilError(t, err)

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

	err = client.DeleteSession("some-id")
	assert.ErrorContains(t, err, "400: unknown HTTP error")
}

func TestClient_GetMediaContent(t *testing.T) {
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	assert.NilError(t, err)

	jpegImage := []byte("value")
	HTTPClient := &mockHTTPClient{
		do: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader(jpegImage)),
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
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	assert.NilError(t, err)

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
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	assert.NilError(t, err)

	jpegImage := []byte("value")
	HTTPClient := &mockHTTPClient{
		do: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader(jpegImage)),
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

	_, err = client.GetMediaContent("some-sessionID", "some-mediaID")
	assert.ErrorContains(t, err, "unable to parse content type from response")
}

func TestClient_GetMediaContent_ShouldReturnMissingKeyError(t *testing.T) {
	client := Client{}
	_, err := client.GetMediaContent("some-id", "some-media-id")
	assert.ErrorContains(t, err, "missing private key")
}

func TestClient_GetMediaContent_ShouldReturnResponseError(t *testing.T) {
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	assert.NilError(t, err)

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

	_, err = client.GetMediaContent("some-id", "some-media-id")
	assert.ErrorContains(t, err, "400: unknown HTTP error")
}

func TestClient_DeleteMediaContent(t *testing.T) {
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	assert.NilError(t, err)

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

	err = client.DeleteMediaContent("some-session-id", "some-media-id")
	assert.NilError(t, err)
}

func TestClient_DeleteMediaContent_ShouldReturnMissingKeyError(t *testing.T) {
	client := Client{}
	err := client.DeleteMediaContent("some-id", "some-media-id")
	assert.ErrorContains(t, err, "missing private key")
}

func TestClient_DeleteMediaContent_ShouldReturnResponseError(t *testing.T) {
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	assert.NilError(t, err)

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

	err = client.DeleteMediaContent("some-id", "some-media-id")
	assert.ErrorContains(t, err, "400: unknown HTTP error")
}

func TestClient_GetSupportedDocuments(t *testing.T) {
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	assert.NilError(t, err)

	countryCodeUSA := "USA"
	documentTypePassport := "PASSPORT"
	isStrictlyLatin := true
	documentsResponse := supported.DocumentsResponse{
		SupportedCountries: []*supported.Country{
			{
				Code: countryCodeUSA,
				SupportedDocuments: []*supported.Document{
					{
						Type:            documentTypePassport,
						IsStrictlyLatin: isStrictlyLatin,
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
				Body:       io.NopCloser(bytes.NewReader(jsonBytes)),
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
	assert.Equal(t, result.SupportedCountries[0].SupportedDocuments[0].IsStrictlyLatin, isStrictlyLatin)
}

func TestClient_GetSupportedDocuments_ShouldReturnMissingKeyError(t *testing.T) {
	client := Client{}
	_, err := client.GetSupportedDocuments()
	assert.ErrorContains(t, err, "missing private key")
}

func TestClient_GetSupportedDocuments_ShouldReturnResponseError(t *testing.T) {
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	assert.NilError(t, err)

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

	_, err = client.GetSupportedDocuments()
	assert.ErrorContains(t, err, "400: unknown HTTP error")
}

func TestNewClient(t *testing.T) {
	key, err := os.ReadFile("../test/test-key.pem")
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
	key, err := os.ReadFile("../test/test-key.pem")
	assert.NilError(t, err)

	client, err := NewClient("sdkID", key)
	assert.NilError(t, err)

	_, err = client.GetSession("")
	assert.ErrorContains(t, err, "sessionID cannot be an empty string")
}

func TestClient_DeleteSession_EmptySessionID(t *testing.T) {
	key, err := os.ReadFile("../test/test-key.pem")
	assert.NilError(t, err)

	client, err := NewClient("sdkID", key)
	assert.NilError(t, err)

	err = client.DeleteSession("")
	assert.ErrorContains(t, err, "sessionID cannot be an empty string")
}

func TestClient_GetMediaContent_EmptySessionID(t *testing.T) {
	key, err := os.ReadFile("../test/test-key.pem")
	assert.NilError(t, err)

	client, err := NewClient("sdkID", key)
	assert.NilError(t, err)

	_, err = client.GetMediaContent("", "someMediaID")
	assert.ErrorContains(t, err, "sessionID cannot be an empty string")
}

func TestClient_GetMediaContent_EmptyMediaID(t *testing.T) {
	key, err := os.ReadFile("../test/test-key.pem")
	assert.NilError(t, err)

	client, err := NewClient("sdkID", key)
	assert.NilError(t, err)

	_, err = client.GetMediaContent("someSessionID", "")
	assert.ErrorContains(t, err, "mediaID cannot be an empty string")
}

func TestClient_DeleteMediaContent_EmptySessionID(t *testing.T) {
	key, err := os.ReadFile("../test/test-key.pem")
	assert.NilError(t, err)

	client, err := NewClient("sdkID", key)
	assert.NilError(t, err)

	err = client.DeleteMediaContent("", "someMediaID")
	assert.ErrorContains(t, err, "sessionID cannot be an empty string")
}

func TestClient_DeleteMediaContent_EmptyMediaID(t *testing.T) {
	key, err := os.ReadFile("../test/test-key.pem")
	assert.NilError(t, err)

	client, err := NewClient("sdkID", key)
	assert.NilError(t, err)

	err = client.DeleteMediaContent("someSessionID", "")
	assert.ErrorContains(t, err, "mediaID cannot be an empty string")
}

func Test_EmptySdkID(t *testing.T) {
	key, err := os.ReadFile("../test/test-key.pem")
	assert.NilError(t, err)

	client, err := NewClient("sdkID", key)
	assert.NilError(t, err)

	_, err = client.GetSession("")
	assert.ErrorContains(t, err, "sessionID cannot be an empty string")
}

func TestNewClient_KeyLoad_Failure(t *testing.T) {
	key, err := os.ReadFile("../test/test-key-invalid-format.pem")
	assert.NilError(t, err)

	_, err = NewClient("sdkID", key)

	assert.ErrorContains(t, err, "invalid key: not PEM-encoded")

	tempError, temporary := err.(interface {
		Temporary() bool
	})
	assert.Check(t, !temporary || !tempError.Temporary())
}

func TestClient_UsesDefaultApiUrl(t *testing.T) {
	key, err := os.ReadFile("../test/test-key.pem")
	assert.NilError(t, err)

	client, err := NewClient("sdkID", key)
	assert.NilError(t, err)

	assert.Equal(t, "https://api.yoti.com/idverify/v1", client.apiURL)
}

func TestClient_UsesEnvVariable(t *testing.T) {
	key, err := os.ReadFile("../test/test-key.pem")
	assert.NilError(t, err)

	os.Setenv("YOTI_DOC_SCAN_API_URL", "envBaseUrl")

	client, err := NewClient("sdkID", key)
	assert.NilError(t, err)

	assert.Equal(t, "envBaseUrl", client.apiURL)
}

func TestClient_UsesOverrideApiUrlOverEnvVariable(t *testing.T) {
	key, err := os.ReadFile("../test/test-key.pem")
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

type testJSONMarshaler struct{}

func (m testJSONMarshaler) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func TestClient_CreateFaceCaptureResource(t *testing.T) {
	key, _ := rsa.GenerateKey(rand.Reader, 1024)

	expected := &retrieve.FaceCaptureResourceResponse{ID: "resource-id"}
	expectedBytes, _ := json.Marshal(expected)

	client := Client{
		Key: key,
		HTTPClient: &mockHTTPClient{
			do: func(r *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body:       io.NopCloser(bytes.NewReader(expectedBytes)),
				}, nil
			},
		},
		apiURL:        "https://example.com",
		SdkID:         "sdk-id",
		jsonMarshaler: testJSONMarshaler{},
	}

	payload := facecapture.NewCreateFaceCaptureResourcePayload("requirement-id")
	result, err := client.CreateFaceCaptureResource("session-id", payload)

	assert.NilError(t, err)
	assert.Equal(t, result.ID, expected.ID)
}

func TestClient_UploadFaceCaptureImage(t *testing.T) {
	key, _ := rsa.GenerateKey(rand.Reader, 1024)

	image := []byte("test-image")
	payload := facecapture.NewUploadFaceCaptureImagePayload("image/png", image)

	client := Client{
		Key: key,
		HTTPClient: &mockHTTPClient{
			do: func(r *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 204,
					Body:       io.NopCloser(bytes.NewReader([]byte{})),
				}, nil
			},
		},
		apiURL: "https://example.com",
		SdkID:  "sdk-id",
	}

	err := client.UploadFaceCaptureImage("session-id", "resource-id", payload)
	assert.NilError(t, err)
}

func TestClient_GetSessionConfiguration(t *testing.T) {
	key, _ := rsa.GenerateKey(rand.Reader, 1024)

	expected := &retrieve.SessionConfigurationResponse{}
	expectedBytes, _ := json.Marshal(expected)

	client := Client{
		Key: key,
		HTTPClient: &mockHTTPClient{
			do: func(r *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body:       io.NopCloser(bytes.NewReader(expectedBytes)),
				}, nil
			},
		},
		apiURL: "https://example.com",
		SdkID:  "sdk-id",
	}

	result, err := client.GetSessionConfiguration("session-id")
	assert.NilError(t, err)
	assert.DeepEqual(t, result, expected)
}

func TestClient_GetSessionConfiguration_FailsOnEmptySessionID(t *testing.T) {
	key, _ := rsa.GenerateKey(rand.Reader, 1024)
	client := Client{Key: key}

	_, err := client.GetSessionConfiguration("")
	assert.Error(t, err, "sessionID cannot be an empty string")
}

func TestClient_CreateFaceCaptureResource_FailsOnEmptySessionID(t *testing.T) {
	key, _ := rsa.GenerateKey(rand.Reader, 1024)
	client := Client{Key: key}

	payload := facecapture.NewCreateFaceCaptureResourcePayload("requirement-id")
	_, err := client.CreateFaceCaptureResource("", payload)
	assert.Error(t, err, "sessionID cannot be an empty string")
}

func TestClient_UploadFaceCaptureImage_FailsOnEmptyIDs(t *testing.T) {
	key, _ := rsa.GenerateKey(rand.Reader, 1024)
	client := Client{Key: key}

	payload := facecapture.NewUploadFaceCaptureImagePayload("image/png", []byte("img"))

	err := client.UploadFaceCaptureImage("", "resource-id", payload)
	assert.Error(t, err, "sessionID and resourceID must not be empty")

	err = client.UploadFaceCaptureImage("session-id", "", payload)
	assert.Error(t, err, "sessionID and resourceID must not be empty")
}

func TestClient_AddFaceCaptureResourceToSession_HappyPath(t *testing.T) {
	key, _ := rsa.GenerateKey(rand.Reader, 1024)
	mockBase64Image := "aW1hZ2U=" // "image"

	mockConfig := &retrieve.SessionConfigurationResponse{
		Capture: &retrieve.CaptureResponse{
			RequiredResources: []retrieve.RequiredResourceResponse{
				&retrieve.RequiredFaceCaptureResourceResponse{
					BaseRequiredResource: retrieve.BaseRequiredResource{
						Type: "FACE_CAPTURE",
						ID:   "requirement-id-123",
					},
				},
			},
		},
	}
	configBytes, _ := json.Marshal(mockConfig)

	mockResource := &retrieve.FaceCaptureResourceResponse{ID: "resource-id-456"}
	resourceBytes, _ := json.Marshal(mockResource)

	var getConfigCalled, createResourceCalled, uploadImageCalled bool

	mockClient := &mockHTTPClient{
		do: func(req *http.Request) (*http.Response, error) {
			if strings.Contains(req.URL.Path, "/configuration") {
				getConfigCalled = true
				return &http.Response{StatusCode: 200, Body: io.NopCloser(bytes.NewReader(configBytes))}, nil
			}
			if req.Method == http.MethodPost && strings.Contains(req.URL.Path, "/face-capture") {
				createResourceCalled = true
				return &http.Response{StatusCode: 200, Body: io.NopCloser(bytes.NewReader(resourceBytes))}, nil
			}
			if req.Method == http.MethodPut && strings.Contains(req.URL.Path, "/image") {
				uploadImageCalled = true
				return &http.Response{StatusCode: 200, Body: io.NopCloser(bytes.NewReader([]byte{}))}, nil
			}
			return nil, fmt.Errorf("unexpected request: %s %s", req.Method, req.URL.Path)
		},
	}

	client := Client{
		Key:           key,
		HTTPClient:    mockClient,
		apiURL:        "https://example.com",
		SdkID:         "sdk-id",
		jsonMarshaler: testJSONMarshaler{},
	}

	err := client.AddFaceCaptureResourceToSession("session-id", mockBase64Image)

	assert.NilError(t, err)
	assert.Assert(t, getConfigCalled)
	assert.Assert(t, createResourceCalled)
	assert.Assert(t, uploadImageCalled)
}

func TestClient_AddFaceCaptureResourceToSession_GetConfigFails(t *testing.T) {
	key, _ := rsa.GenerateKey(rand.Reader, 1024)
	expectedErr := errors.New("failed to get config")

	mockClient := &mockHTTPClient{
		do: func(req *http.Request) (*http.Response, error) {
			return nil, expectedErr
		},
	}

	client := Client{Key: key, HTTPClient: mockClient, apiURL: "https://example.com", SdkID: "sdk-id"}
	err := client.AddFaceCaptureResourceToSession("session-id", "aW1hZ2U=")

	assert.ErrorIs(t, err, expectedErr)
}

func TestClient_AddFaceCaptureResourceToSession_NoRequirements(t *testing.T) {
	key, _ := rsa.GenerateKey(rand.Reader, 1024)
	mockConfig := &retrieve.SessionConfigurationResponse{
		Capture: &retrieve.CaptureResponse{
			RequiredResources: []retrieve.RequiredResourceResponse{}, // No requirements
		},
	}
	configBytes, _ := json.Marshal(mockConfig)

	mockClient := &mockHTTPClient{
		do: func(req *http.Request) (*http.Response, error) {
			// Only expect a call to the configuration endpoint
			if strings.Contains(req.URL.Path, "/configuration") {
				return &http.Response{StatusCode: 200, Body: io.NopCloser(bytes.NewReader(configBytes))}, nil
			}
			return nil, fmt.Errorf("unexpected request: %s %s", req.Method, req.URL.Path)
		},
	}
	client := Client{Key: key, HTTPClient: mockClient, apiURL: "https://example.com", SdkID: "sdk-id"}

	err := client.AddFaceCaptureResourceToSession("session-id", "aW1hZ2U=")
	assert.NilError(t, err)
}

func TestClient_AddFaceCaptureResourceToSession_InvalidBase64(t *testing.T) {
	key, _ := rsa.GenerateKey(rand.Reader, 1024)
	mockConfig := &retrieve.SessionConfigurationResponse{
		Capture: &retrieve.CaptureResponse{
			RequiredResources: []retrieve.RequiredResourceResponse{
				&retrieve.RequiredFaceCaptureResourceResponse{
					BaseRequiredResource: retrieve.BaseRequiredResource{
						Type: "FACE_CAPTURE",
						ID:   "requirement-id-123",
					},
				},
			},
		},
	}
	configBytes, _ := json.Marshal(mockConfig)
	mockResource := &retrieve.FaceCaptureResourceResponse{ID: "resource-id-456"}
	resourceBytes, _ := json.Marshal(mockResource)

	mockClient := &mockHTTPClient{
		do: func(req *http.Request) (*http.Response, error) {
			if strings.Contains(req.URL.Path, "/configuration") {
				return &http.Response{StatusCode: 200, Body: io.NopCloser(bytes.NewReader(configBytes))}, nil
			}
			if strings.Contains(req.URL.Path, "/face-capture") {
				return &http.Response{StatusCode: 200, Body: io.NopCloser(bytes.NewReader(resourceBytes))}, nil
			}
			return nil, fmt.Errorf("unexpected request: %s", req.URL.Path)
		},
	}
	client := Client{Key: key, HTTPClient: mockClient, apiURL: "https://example.com", SdkID: "sdk-id"}

	err := client.AddFaceCaptureResourceToSession("session-id", "this is not base64")
	assert.ErrorContains(t, err, "illegal base64 data")
}
