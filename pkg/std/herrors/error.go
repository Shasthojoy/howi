package herrors

import (
	"errors"
	"fmt"
)

// New returns new standard error
func New(msg string) error {
	return errors.New(msg)
}

// Newf returns new standard error.  Arguments are handled in the manner of
// fmt.Esprintf followed by \n.
func Newf(format string, v ...interface{}) error {
	return fmt.Errorf(format, v...)
}
