package aml

import (
	"crypto/rsa"
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

func TestYotiClient_PerformAmlCheck_WithInvalidJSON(t *testing.T) {
	key := getValidKey()

	client := &mockHTTPClient{
		do: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       ioutil.NopCloser(strings.NewReader("Not a JSON document")),
			}, nil
		},
	}

	_, err := PerformAmlCheck(client, createStandardAmlProfile(), "clientSdkId", "https://apiUrl", key)
	assert.Check(t, strings.Contains(err.Error(), "invalid character"))
}

func TestYotiClient_PerformAmlCheck_Success(t *testing.T) {
	key := getValidKey()

	client := &mockHTTPClient{
		do: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       ioutil.NopCloser(strings.NewReader(`{"on_fraud_list":true,"on_pep_list":true,"on_watch_list":true}`)),
			}, nil
		},
	}

	result, err := PerformAmlCheck(client, createStandardAmlProfile(), "clientSdkId", "https://apiUrl", key)
	assert.NilError(t, err)

	assert.Check(t, result.OnFraudList)
	assert.Check(t, result.OnPEPList)
	assert.Check(t, result.OnWatchList)
}

func TestYotiClient_PerformAmlCheck_Unsuccessful(t *testing.T) {
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

	_, err := PerformAmlCheck(client, createStandardAmlProfile(), "clientSdkId", "https://apiUrl", key)
	assert.ErrorContains(t, err, fmt.Sprintf("%d: AML Check was unsuccessful - %s", 503, responseBody))

	tempError, temporary := err.(interface {
		Temporary() bool
	})
	assert.Check(t, temporary && tempError.Temporary())
}

func getValidKey() *rsa.PrivateKey {
	return test.GetValidKey("../test/test-key.pem")
}
func createStandardAmlProfile() AmlProfile {
	var amlAddress = AmlAddress{
		Country: "GBR"}

	var amlProfile = AmlProfile{
		GivenNames: "Edward Richard George",
		FamilyName: "Heath",
		Address:    amlAddress}

	return amlProfile
}
