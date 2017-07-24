package herrors

import (
	"errors"
	"fmt"
	"strings"
)

// New returns new standard error msg argument is handled in manner of print
func New(msg ...string) error {
	return errors.New(strings.Join(msg, " "))
}

// Newf returns new standard error.  Arguments are handled in the manner of fmt.Errorf
func Newf(format string, v ...interface{}) error {
	return fmt.Errorf(format, v...)
}
