package yotierror

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var (
	defaultUnknownErrorCodeConst    = "UNKNOWN_ERROR"
	defaultUnknownErrorMessageConst = "unknown HTTP error"
)

// Error indicates errors related to the Yoti API.
type Error struct {
	Id        string `json:"id"`
	Status    int    `json:"status"`
	ErrorCode string `json:"error"`
	Message   string `json:"message"`
}

func (e Error) Error() string {
	return e.ErrorCode + " - " + e.Message
}

// NewResponseError creates a new Error
func NewResponseError(response *http.Response) *Error {
	err := &Error{
		ErrorCode: defaultUnknownErrorCodeConst,
		Message:   defaultUnknownErrorMessageConst,
	}
	if response == nil {
		return err
	}
	err.Status = response.StatusCode
	if response.Body == nil {
		return err
	}
	defer response.Body.Close()
	b, e := io.ReadAll(response.Body)
	if e != nil {
		err.Message = fmt.Sprintf(defaultUnknownErrorMessageConst+": %q", e)
		return err
	}
	e = json.Unmarshal(b, err)
	if e != nil {
		err.Message = fmt.Sprintf(defaultUnknownErrorMessageConst+": %q", e)
	}
	return err
}

// Temporary indicates this ErrorCode is a temporary ErrorCode
func (e Error) Temporary() bool {
	return e.Status >= 500
}
