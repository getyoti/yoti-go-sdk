package yoti

// TemporaryError indicates that a temporary outage has occured and the
// previous request can be reattempted without modification.
type TemporaryError struct {
	Err error
}

func (e TemporaryError) Error() string {
	return e.Err.Error()
}

// Unwrap returns the internal error for debugging
func (e TemporaryError) Unwrap() error {
	return e.Err
}
