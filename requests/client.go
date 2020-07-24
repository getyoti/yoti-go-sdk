package requests

import (
	"net/http"
)

var (
	defaultUnknownErrorMessageConst = "unknown HTTP error"

	// DefaultHTTPErrorMessages maps HTTP error status codes to format strings
	// to create useful error messages. -1 is used to specify a default message
	// that can be used if an error code is not explicitly defined
	DefaultHTTPErrorMessages = map[int]string{
		-1: defaultUnknownErrorMessageConst,
	}
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
