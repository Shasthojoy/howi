package herrors

import "fmt"

// NewMultiError returns new multierror
func NewMultiError() MultiError {
	return MultiError{}
}

// MultiError is returned by batch operations when there are errors in
// particular situations where returning single error is not possible
// or type of error is not blocking further execution until MultiError
// is evaluated.
type MultiError struct {
	Errors []error
}

// Error returns first occurred error string with additional suffix with
// total count of errors.
func (merr *MultiError) Error() (str string) {
	l := len(merr.Errors)
	if l > 0 {
		str = merr.Errors[0].Error()
	}
	return fmt.Sprintf("%s (total errors: %d)", str, l)
}

// AppendError appends error to multierror if provided error is not nil
func (merr *MultiError) AppendError(err error) error {
	if err != nil {
		n := len(merr.Errors)
		if n == cap(merr.Errors) {
			errs := make([]error, len(merr.Errors), 2*len(merr.Errors)+1)
			copy(errs, merr.Errors)
			merr.Errors = errs
		}
		merr.Errors = merr.Errors[0 : n+1]
		merr.Errors[n] = err
	}
	return err
}

// AppendString creates error from provided string and appends it to multierror.
// It also returns same created error if any case you would need it.
// It returns nil if provided string is empty.
func (merr *MultiError) AppendString(str string) error {
	if str == "" {
		return nil
	}
	err := fmt.Errorf("%s", str)
	merr.AppendError(err)
	return err
}

// AppendSprintf creates error from provided format, interface and appends it
// to multierror. It also returns same created error if any case you would need it.
// It returns nil if provided values are empty.
func (merr *MultiError) AppendSprintf(format string, v ...interface{}) error {
	if len(v) == 0 && format == "" {
		return nil
	}
	err := fmt.Errorf(format, v...)
	merr.AppendError(err)
	return err
}

// Len returns total count of errors
func (merr *MultiError) Len() int {
	return len(merr.Errors)
}

// Nil returns true if there are no errors
func (merr *MultiError) Nil() bool {
	return len(merr.Errors) == 0
}
