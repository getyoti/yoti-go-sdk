package yotierror

const (
	// InvalidRequestSignature can be returned by any endpoint that requires a signed request.
	InvalidRequestSignature = "INVALID_REQUEST_SIGNATURE"
	// InvalidAuthHeader can be returned by any endpoint that requires a signed request.
	InvalidAuthHeader = "INVALID_AUTH_HEADER"
)
