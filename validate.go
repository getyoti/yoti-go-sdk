package yoti

import (
	"errors"
)

func notEmpty(validate, info string) error {
	if validate == "" {
		return errors.New(info)
	}
	return nil
}
