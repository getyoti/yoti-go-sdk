package requests

import (
	"net/http"
)

// HttpClient is a mockable HTTP Client Interface
type HttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

// ClientInterface defines the interface required to Mock the YotiClient for
// testing
type ClientInterface interface {
	MakeRequest(string, string, []byte, bool, ...map[int]string) (string, error)
	GetSdkID() string
}
