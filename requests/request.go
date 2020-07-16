package requests

import (
	"fmt"
	"io/ioutil"
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
		err = handleHTTPError(response, httpErrorMessages...)
		if response.StatusCode >= 500 {
			err = yotierror.NewTemporary(err)
		}
	}

	return
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

func handleHTTPError(response *http.Response, errorMessages ...map[int]string) error {
	var body []byte
	if response.Body != nil {
		body, _ = ioutil.ReadAll(response.Body)
	} else {
		body = make([]byte, 0)
	}
	for _, handler := range errorMessages {
		for code, message := range handler {
			if code == response.StatusCode {
				return fmt.Errorf(
					message,
					response.StatusCode,
					body,
				)
			}

		}
		if defaultMessage, ok := handler[-1]; ok {
			return fmt.Errorf(
				defaultMessage,
				response.StatusCode,
				body,
			)
		}

	}
	return fmt.Errorf(
		defaultUnknownErrorMessageConst,
		response.StatusCode,
		body,
	)
}
