package yoti

import "errors"

var (
	//ErrProfileNotFound profile was not found during activity retrieval for the provided one time use token
	ErrProfileNotFound = errors.New("ProfileNotFound")
	//ErrFailure there was a failure during activity retrieval
	ErrFailure = errors.New("Failure")
	//ErrSharingFailure there was a failure when sharing
	ErrSharingFailure = errors.New("SharingFailure")
)
