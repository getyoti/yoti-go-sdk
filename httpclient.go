package yoti

import (
	"net/http"
)

//HttpClient is a mockable HTTP Client Interface
type HttpClient interface {
	Do(*http.Request) (*http.Response, error)
}
