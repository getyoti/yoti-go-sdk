package validate

import (
	"errors"
)

// NotEmpty checks if the string validate is "" and returns an error value if it
// is. It returns nil if it is not empty-string
func NotEmpty(validate, info string) error {
	if validate == "" {
		return errors.New(info)
	}
	return nil
}
