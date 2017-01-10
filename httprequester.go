package yoti

import (
	"io/ioutil"
	"net/http"
)

type httpResponse struct {
	Success    bool
	StatusCode int
	Content    string
}

type httpRequester func(uri string, headers map[string]string) (result *httpResponse, err error)

func doRequest(uri string, headers map[string]string) (result *httpResponse, err error) {
	client := &http.Client{}

	var req *http.Request
	if req, err = http.NewRequest("GET", uri, nil); err != nil {
		return
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	var resp *http.Response
	resp, err = client.Do(req)
	defer resp.Body.Close()

	if err != nil {
		return
	}

	var body []byte
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		return
	}

	result = &httpResponse{
		Success:    resp.StatusCode < 300,
		StatusCode: resp.StatusCode,
		Content:    string(body)}

	return
}
