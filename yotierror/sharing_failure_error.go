package yotierror

type DetailedSharingFailureError struct {
	Code        string
	Description string
}

func (d DetailedSharingFailureError) Error() string {
	return "sharing failure"
}
