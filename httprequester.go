package yoti

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	//HTTPMethodPost Post HTTP method
	HTTPMethodPost = "POST"
	//HTTPMethodGet Get HTTP method
	HTTPMethodGet = "GET"
	//HTTPMethodPut Put HTTP method
	HTTPMethodPut = "PUT"
	//HTTPMethodPatch Patch HTTP method
	HTTPMethodPatch = "PATCH"
)

type httpResponse struct {
	Success    bool
	StatusCode int
	Content    string
}

type httpRequester func(uri string, headers map[string]string, httpRequestMethod string, contentBytes []byte) (result *httpResponse, err error)

func doRequest(uri string, headers map[string]string, httpRequestMethod string, contentBytes []byte) (result *httpResponse, err error) {
	client := &http.Client{}

	supportedHTTPMethods := map[string]bool{"GET": true, "POST": true, "PUT": true, "PATCH": true}

	if !supportedHTTPMethods[httpRequestMethod] {
		err = fmt.Errorf("HTTP Method: '%s' is unsupported", httpRequestMethod)
		return
	}

	var req *http.Request
	if req, err = http.NewRequest(
		httpRequestMethod,
		uri,
		bytes.NewBuffer(contentBytes)); err != nil {
		return
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	var resp *http.Response
	resp, err = client.Do(req)

	if err != nil {
		return
	}

	defer resp.Body.Close()

	var responseBody []byte
	if responseBody, err = ioutil.ReadAll(resp.Body); err != nil {
		log.Printf("Unable to read the HTTP response, error: %s", err)
	}

	result = &httpResponse{
		Success:    resp.StatusCode < 300,
		StatusCode: resp.StatusCode,
		Content:    string(responseBody)}

	return
}
