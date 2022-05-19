package yotierror

import "errors"

var (
	// InvalidTokenError means that the token used to call GetActivityDetails is invalid. Make sure you are retrieving this token correctly.
	InvalidTokenError = errors.New("invalid token")

	// TokenDecryptError means that the token could not be decrypted. Ensure you are using the correct .pem file.
	TokenDecryptError = errors.New("unable to decrypt token")

	// SharingFailureError means that the share between a user and an application was not successful.
	SharingFailureError = DetailedSharingFailureError{}
)
