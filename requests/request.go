package requests

import (
	"net/http"
	"time"

	"github.com/getyoti/yoti-go-sdk/v3/yotierror"
)

// Execute makes a request to the specified endpoint, with an optional payload
func Execute(httpClient HttpClient, request *http.Request, httpErrorMessages ...map[int]string) (response *http.Response, err error) {
	if response, err = doRequest(request, httpClient); err != nil {
		return
	}

	statusCodeIsFailure := response.StatusCode >= 300 || response.StatusCode < 200

	if statusCodeIsFailure {
		return response, yotierror.New(response, httpErrorMessages...)
	}

	return response, nil
}

func doRequest(request *http.Request, httpClient HttpClient) (*http.Response, error) {
	httpClient = ensureHttpClientTimeout(httpClient)
	return httpClient.Do(request)
}

func ensureHttpClientTimeout(httpClient HttpClient) HttpClient {
	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: time.Second * 10,
		}
	}

	return httpClient
}
