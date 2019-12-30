package yoti

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

func TestYotiClient_PerformAmlCheck_Success(t *testing.T) {
	key, _ := ioutil.ReadFile("test-key.pem")

	client := Client{
		HTTPClient: &mockHTTPClient{
			do: func(*http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(strings.NewReader(`{"on_fraud_list":true,"on_pep_list":true,"on_watch_list":true}`)),
				}, nil
			},
		},
	}
	var err error
	client.Key, err = loadRsaKey(key)
	assert.NilError(t, err)

	result, err := client.PerformAmlCheck(createStandardAmlProfile())

	assert.Assert(t, is.Nil(err))

	assert.Check(t, result.OnFraudList)
	assert.Check(t, result.OnPEPList)
	assert.Check(t, result.OnWatchList)

}

func TestYotiClient_PerformAmlCheck_Unsuccessful(t *testing.T) {
	key, _ := ioutil.ReadFile("test-key.pem")
	client := Client{
		HTTPClient: &mockHTTPClient{
			do: func(*http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 503,
					Body:       ioutil.NopCloser(strings.NewReader(`SERVICE UNAVAILABLE - Unable to reach the Integrity Service`)),
				}, nil
			},
		},
	}
	var err error
	client.Key, err = loadRsaKey(key)
	assert.NilError(t, err)

	_, err = client.PerformAmlCheck(createStandardAmlProfile())

	var expectedErrString = "AML Check was unsuccessful"

	assert.Assert(t, err != nil)
	assert.Check(t, strings.HasPrefix(err.Error(), expectedErrString))
	tempError, temporary := err.(interface {
		Temporary() bool
	})
	assert.Check(t, temporary && tempError.Temporary())
}
