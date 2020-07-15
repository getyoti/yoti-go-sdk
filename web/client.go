package web

import (
	"net/http"
)

var (
	DefaultUnknownErrorMessageConst = "Unknown HTTP Error: %[1]d: %[2]s"

	// DefaultHTTPErrorMessages maps HTTP error status codes to format strings
	// to create useful error messages. -1 is used to specify a default message
	// that can be used if an error code is not explicitly defined
	DefaultHTTPErrorMessages = map[int]string{
		-1: DefaultUnknownErrorMessageConst,
	}
)

//HttpClient is a mockable HTTP Client Interface
type HttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

// ClientInterface defines the interface required to Mock the YotiClient for
// testing
type ClientInterface interface {
	MakeRequest(string, string, []byte, bool, ...map[int]string) (string, error)
	GetSdkID() string
}
