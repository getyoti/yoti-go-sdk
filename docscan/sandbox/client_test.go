package sandbox

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/getyoti/yoti-go-sdk/v3/cryptoutil"
	"github.com/getyoti/yoti-go-sdk/v3/docscan/sandbox/request"
	"github.com/getyoti/yoti-go-sdk/v3/yotierror"
	"gotest.tools/v3/assert"
)

func TestClient_httpClient_ShouldReturnDefaultClient(t *testing.T) {
	client := Client{}
	assert.Check(t, client.getHTTPClient() != nil)
}

func TestClient_ConfigureSessionResponse_ShouldReturnErrorIfNotCreated(t *testing.T) {
	key, _ := rsa.GenerateKey(rand.Reader, 1024)
	client := Client{
		Key: key,
		HTTPClient: &mockHTTPClient{
			do: func(*http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 400,
					Body:       ioutil.NopCloser(strings.NewReader("")),
				}, nil
			},
		},
	}
	err := client.ConfigureSessionResponse("some_session_id", request.ResponseConfig{})
	assert.ErrorContains(t, err, "400: unknown HTTP error")
}

func TestClient_ConfigureSessionResponse_ShouldReturnFormattedErrorWithResponse(t *testing.T) {
	key, _ := rsa.GenerateKey(rand.Reader, 1024)
	jsonBytes, _ := json.Marshal(yotierror.DataObject{
		Code:    "SOME_CODE",
		Message: "some message",
	})
	response := &http.Response{
		StatusCode: 400,
		Body:       ioutil.NopCloser(bytes.NewReader(jsonBytes)),
	}

	client := Client{
		Key: key,
		HTTPClient: &mockHTTPClient{
			do: func(*http.Request) (*http.Response, error) {
				return response, nil
			},
		},
	}
	err := client.ConfigureSessionResponse("some_session_id", request.ResponseConfig{})
	assert.ErrorContains(t, err, "400: SOME_CODE - some message")

	errorResponse := err.(*yotierror.Error).Response
	assert.Equal(t, response, errorResponse)

	body, _ := ioutil.ReadAll(errorResponse.Body)
	assert.Equal(t, string(body), string(jsonBytes))
}

func TestClient_ConfigureSessionResponse_ShouldReturnMissingKeyError(t *testing.T) {
	client := Client{}
	err := client.ConfigureSessionResponse("some_session_id", request.ResponseConfig{})
	assert.ErrorContains(t, err, "missing private key")
}

func TestClient_ConfigureSessionResponse_ShouldReturnJsonError(t *testing.T) {
	client := Client{
		jsonMarshaler: &mockJSONMarshaler{
			marshal: func(v interface{}) ([]byte, error) {
				return []byte{}, errors.New("some json error")
			},
		},
	}
	err := client.ConfigureSessionResponse("some_session_id", request.ResponseConfig{})
	assert.ErrorContains(t, err, "some json error")
}

func TestNewClient_ConfigureSessionResponse_Success(t *testing.T) {
	key, _ := ioutil.ReadFile("../../test/test-key.pem")
	client, err := NewClient("ClientSDKID", key)
	assert.NilError(t, err)

	client.HTTPClient = &mockHTTPClient{
		do: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 201,
			}, nil
		},
	}

	responseErr := client.ConfigureSessionResponse("some_session_id", request.ResponseConfig{})
	assert.NilError(t, responseErr)
}

func TestNewClient_KeyLoad_Failure(t *testing.T) {
	key, _ := ioutil.ReadFile("../../test/test-key-invalid-format.pem")
	_, err := NewClient("", key)

	assert.ErrorContains(t, err, "invalid Key: not PEM-encoded")

	tempError, temporary := err.(interface {
		Temporary() bool
	})
	assert.Check(t, !temporary || !tempError.Temporary())
}

func TestClient_ConfigureSessionResponse_Success(t *testing.T) {
	key, _ := rsa.GenerateKey(rand.Reader, 1024)
	client := Client{
		Key: key,
		HTTPClient: &mockHTTPClient{
			do: func(*http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 201,
				}, nil
			},
		},
	}
	err := client.ConfigureSessionResponse("some_session_id", request.ResponseConfig{})
	assert.NilError(t, err)
}

func TestClient_ConfigureSessionResponse_ShouldReturnHttpClientError(t *testing.T) {
	key, _ := rsa.GenerateKey(rand.Reader, 1024)
	client := Client{
		Key: key,
		HTTPClient: &mockHTTPClient{
			do: func(*http.Request) (*http.Response, error) {
				return &http.Response{}, errors.New("some error")
			},
		},
	}
	err := client.ConfigureSessionResponse("some_session_id", request.ResponseConfig{})
	assert.ErrorContains(t, err, "some error")
}

func TestClient_ConfigureApplicationResponse_ShouldReturnErrorIfNotCreated(t *testing.T) {
	key, _ := rsa.GenerateKey(rand.Reader, 1024)
	client := Client{
		Key: key,
		HTTPClient: &mockHTTPClient{
			do: func(*http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 401,
					Body:       ioutil.NopCloser(strings.NewReader("")),
				}, nil
			},
		},
	}
	err := client.ConfigureApplicationResponse(request.ResponseConfig{})
	assert.ErrorContains(t, err, "401: unknown HTTP error")
}

func TestClient_ConfigureApplicationResponse_ShouldReturnMissingKeyError(t *testing.T) {
	client := Client{}
	err := client.ConfigureApplicationResponse(request.ResponseConfig{})
	assert.ErrorContains(t, err, "missing private key")
}

func TestClient_ConfigureApplicationResponse_ShouldReturnJsonError(t *testing.T) {
	client := Client{
		jsonMarshaler: &mockJSONMarshaler{
			marshal: func(v interface{}) ([]byte, error) {
				return []byte{}, errors.New("some json error")
			},
		},
	}
	err := client.ConfigureApplicationResponse(request.ResponseConfig{})
	assert.ErrorContains(t, err, "some json error")
}

func TestClient_ConfigureApplicationResponse_Success(t *testing.T) {
	key, _ := rsa.GenerateKey(rand.Reader, 1024)
	client := Client{
		Key: key,
		HTTPClient: &mockHTTPClient{
			do: func(*http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 201,
				}, nil
			},
		},
	}
	err := client.ConfigureApplicationResponse(request.ResponseConfig{})
	assert.NilError(t, err)
}

func TestClient_ConfigureApplicationResponse_ShouldReturnHttpClientError(t *testing.T) {
	key, _ := rsa.GenerateKey(rand.Reader, 1024)
	client := Client{
		Key: key,
		HTTPClient: &mockHTTPClient{
			do: func(*http.Request) (*http.Response, error) {
				return &http.Response{}, errors.New("some error")
			},
		},
	}
	err := client.ConfigureApplicationResponse(request.ResponseConfig{})
	assert.ErrorContains(t, err, "some error")
}

func TestClient_ConfigureSessionResponseUsesConstructorApiUrlOverEnvVariable(t *testing.T) {
	client := createSandboxClient(t, "constuctorApiURL")
	os.Setenv("YOTI_DOC_SCAN_API_URL", "envBaseUrl")

	err := client.ConfigureSessionResponse("some_session_id", request.ResponseConfig{})
	assert.NilError(t, err)

	assert.Equal(t, "constuctorApiURL", client.getAPIURL())
}

func TestClient_ConfigureSessionResponseUsesOverrideApiUrlOverEnvVariable(t *testing.T) {
	client := createSandboxClient(t, "")
	client.OverrideAPIURL("overrideApiURL")
	os.Setenv("YOTI_DOC_SCAN_API_URL", "envBaseUrl")

	err := client.ConfigureSessionResponse("some_session_id", request.ResponseConfig{})
	assert.NilError(t, err)

	assert.Equal(t, "overrideApiURL", client.getAPIURL())
}

func TestClient_ConfigureSessionResponseUsesEnvVariable(t *testing.T) {
	client := createSandboxClient(t, "")

	os.Setenv("YOTI_DOC_SCAN_API_URL", "envApiURL")

	err := client.ConfigureSessionResponse("some_session_id", request.ResponseConfig{})
	assert.NilError(t, err)

	assert.Equal(t, "envApiURL", client.getAPIURL())
}

func TestClient_ConfigureSessionResponseUsesDefaultUrlAsFallbackWithEmptyEnvValue(t *testing.T) {
	os.Setenv("YOTI_DOC_SCAN_API_URL", "")

	client := createSandboxClient(t, "")

	err := client.ConfigureSessionResponse("some_session_id", request.ResponseConfig{})
	assert.NilError(t, err)

	assert.Equal(t, "https://api.yoti.com/sandbox/idverify/v1", client.getAPIURL())
}

func TestClient_ConfigureSessionResponseUsesDefaultUrlAsFallbackWithNoEnvValue(t *testing.T) {
	os.Unsetenv("YOTI_DOC_SCAN_API_URL")

	client := createSandboxClient(t, "")

	err := client.ConfigureSessionResponse("some_session_id", request.ResponseConfig{})
	assert.NilError(t, err)

	assert.Equal(t, "https://api.yoti.com/sandbox/idverify/v1", client.getAPIURL())
}

func createSandboxClient(t *testing.T, constructorApiURL string) (client Client) {
	keyBytes, fileErr := ioutil.ReadFile("../../test/test-key.pem")
	assert.NilError(t, fileErr)

	pemFile, parseErr := cryptoutil.ParseRSAKey(keyBytes)
	assert.NilError(t, parseErr)

	if constructorApiURL == "" {
		return Client{
			Key:        pemFile,
			SdkID:      "ClientSDKID",
			HTTPClient: mockHTTPClientCreatedResponse(),
		}
	}

	return Client{
		Key:        pemFile,
		SdkID:      "ClientSDKID",
		HTTPClient: mockHTTPClientCreatedResponse(),
		apiURL:     constructorApiURL,
	}

}

type mockHTTPClient struct {
	do func(*http.Request) (*http.Response, error)
}

func (mock *mockHTTPClient) Do(request *http.Request) (*http.Response, error) {
	if mock.do != nil {
		return mock.do(request)
	}
	return nil, nil
}

func mockHTTPClientCreatedResponse() *mockHTTPClient {
	return &mockHTTPClient{
		do: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 201,
			}, nil
		},
	}
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
