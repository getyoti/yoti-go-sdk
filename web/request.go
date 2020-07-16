package web

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/getyoti/yoti-go-sdk/v3/yotierror"
)

// MakeRequest makes a request to the specified endpoint, with an optional payload
func MakeRequest(httpClient HttpClient, request *http.Request, httpErrorMessages ...map[int]string) (responseData string, err error) {
	var response *http.Response
	if response, err = doRequest(request, httpClient); err != nil {
		return
	}

	if response.StatusCode < 300 && response.StatusCode >= 200 {
		var tmp []byte
		if response.Body != nil {
			tmp, err = ioutil.ReadAll(response.Body)
		} else {
			tmp = make([]byte, 0)
		}
		responseData = string(tmp)
		return
	}
	err = handleHTTPError(response, httpErrorMessages...)
	if response.StatusCode >= 500 {
		err = yotierror.NewTemporary(err)
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
		DefaultUnknownErrorMessageConst,
		response.StatusCode,
		body,
	)
}
