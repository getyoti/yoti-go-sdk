package requests

import (
	"errors"
	"net/http"
	"testing"
	"time"

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

func TestExecute_Success(t *testing.T) {
	client := &mockHTTPClient{
		do: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
			}, nil
		},
	}

	request := &http.Request{
		Method: http.MethodGet,
	}

	response, err := Execute(client, request)

	assert.NilError(t, err)
	assert.Equal(t, response.StatusCode, 200)
}

func TestExecute_Failure(t *testing.T) {
	client := &mockHTTPClient{
		do: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 400,
			}, nil
		},
	}

	request := &http.Request{
		Method: http.MethodGet,
	}

	response, err := Execute(client, request)

	assert.ErrorContains(t, err, "400: unknown HTTP error")
	assert.Equal(t, response.StatusCode, 400)
}

func TestExecute_ClientError(t *testing.T) {
	client := &mockHTTPClient{
		do: func(*http.Request) (*http.Response, error) {
			return nil, errors.New("some error")
		},
	}

	request := &http.Request{
		Method: http.MethodGet,
	}

	_, err := Execute(client, request)

	assert.ErrorContains(t, err, "some error")
}

func TestEnsureHttpClientTimeout_NilHTTPClientShouldUse10sTimeout(t *testing.T) {
	result := ensureHttpClientTimeout(nil).(*http.Client)

	assert.Equal(t, 10*time.Second, result.Timeout)
}

func TestEnsureHttpClientTimeout(t *testing.T) {
	httpClient := &http.Client{
		Timeout: time.Minute * 12,
	}
	result := ensureHttpClientTimeout(httpClient).(*http.Client)

	assert.Equal(t, 12*time.Minute, result.Timeout)
}
