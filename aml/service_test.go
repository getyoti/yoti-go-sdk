package aml

import (
	"crypto/rsa"
	"errors"
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

type mockReadCloser struct {
	read  func(p []byte) (n int, err error)
	close func() error
}

func (mock *mockReadCloser) Read(p []byte) (n int, err error) {
	if mock.read != nil {
		return mock.read(p)
	}
	return 0, nil
}

func (mock *mockReadCloser) Close() error {
	if mock.close != nil {
		return mock.close()
	}
	return nil
}

func TestPerformCheck_WithInvalidJSON(t *testing.T) {
	key := getValidKey()

	client := &mockHTTPClient{
		do: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       ioutil.NopCloser(strings.NewReader("Not a JSON document")),
			}, nil
		},
	}

	_, err := PerformCheck(client, createStandardProfile(), "clientSdkId", "https://apiUrl", key)
	assert.Check(t, strings.Contains(err.Error(), "invalid character"))
}

func TestPerformCheck_Success(t *testing.T) {
	key := getValidKey()

	client := &mockHTTPClient{
		do: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       ioutil.NopCloser(strings.NewReader(`{"on_fraud_list":true,"on_pep_list":true,"on_watch_list":true}`)),
			}, nil
		},
	}

	result, err := PerformCheck(client, createStandardProfile(), "clientSdkId", "https://apiUrl", key)
	assert.NilError(t, err)

	assert.Check(t, result.OnFraudList)
	assert.Check(t, result.OnPEPList)
	assert.Check(t, result.OnWatchList)
}

func TestPerformCheck_Unsuccessful(t *testing.T) {
	key := getValidKey()
	responseBody := "some service unavailable response"

	client := &mockHTTPClient{
		do: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 503,
				Body:       ioutil.NopCloser(strings.NewReader(responseBody)),
			}, nil
		},
	}

	_, err := PerformCheck(client, createStandardProfile(), "clientSdkId", "https://apiUrl", key)
	assert.ErrorContains(t, err, fmt.Sprintf("%d: AML Check was unsuccessful - %s", 503, responseBody))

	tempError, temporary := err.(interface {
		Temporary() bool
	})
	assert.Check(t, temporary && tempError.Temporary())
}

func TestPerformCheck_ShouldReturnMissingBaseURLError(t *testing.T) {
	key := getValidKey()

	client := &mockHTTPClient{
		do: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
			}, nil
		},
	}

	_, err := PerformCheck(client, createStandardProfile(), "clientSdkId", "", key)
	assert.ErrorContains(t, err, "missing BaseURL")
}

func TestPerformCheck_ShouldReturnBodyError(t *testing.T) {
	key := getValidKey()

	body := &mockReadCloser{
		read: func(p []byte) (n int, err error) {
			return 0, errors.New("some read error")
		},
	}

	client := &mockHTTPClient{
		do: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       body,
			}, nil
		},
	}

	_, err := PerformCheck(client, createStandardProfile(), "clientSdkId", "https://apiUrl", key)
	assert.ErrorContains(t, err, "some read error")
}

func getValidKey() *rsa.PrivateKey {
	return test.GetValidKey("../test/test-key.pem")
}

func createStandardProfile() Profile {
	var amlAddress = Address{
		Country: "GBR"}

	var amlProfile = Profile{
		GivenNames: "Edward Richard George",
		FamilyName: "Heath",
		Address:    amlAddress}

	return amlProfile
}
