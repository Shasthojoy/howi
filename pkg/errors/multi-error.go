// Copyright 2012 Marko Kungla. All rights reserved.
// Use of this source code is governed by a The Apache-style
// license that can be found in the LICENSE file.

package errors

import "fmt"

// MultiError is returned by batch operations when there are errors in
// particular situations where returning single error is not possible
// or type of error is not blocking further execution until MultiError
// is evaluated.
type MultiError struct {
	Errors []error
}

// Nil returns true if there are no errors
func (merr *MultiError) Nil() bool {
	return len(merr.Errors) == 0
}

// Len returns total count of errors
func (merr *MultiError) Len() int {
	return len(merr.Errors)
}

// Error returns first occurred error string with additional suffix with
// total count of errors.
func (merr *MultiError) Error() (str string) {
	l := len(merr.Errors)
	if l > 0 {
		str = merr.Errors[len(merr.Errors)-1].Error() + " "
	}
	return fmt.Sprintf("%s(total errors: %d)", str, l)
}

// AsError returns first occurred error with additional suffix with
// total count of errors.
func (merr *MultiError) AsError() error {
	if merr.Nil() {
		return nil
	}
	return New(merr.Error())
}

// Add appends error to multierror if provided error is not nil
func (merr *MultiError) Add(err error) error {
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
	merr.Add(err)
	return err
}

// AppendStringf creates error from provided format, interface and appends it
// to multierror. It also returns same created error if any case you would need it.
// It returns nil if provided values are empty.
func (merr *MultiError) AppendStringf(format string, v ...interface{}) (err error) {
	if len(v) == 0 {
		return nil
	} else if format == "" {
		err = Newf("%s", v...)
	} else {
		err = Newf(format, v...)
	}
	merr.Add(err)
	return err
}
