// Copyright 2012 Marko Kungla. All rights reserved.
// Use of this source code is governed by a The Apache-style
// license that can be found in the LICENSE file.

package errors

import (
	"context"
	"errors"
	"fmt"
	"strings"
)

// DeprecatedErr to mark something as deprecated and may be removed in next release.
type DeprecatedErr error

// NotImplementedErr is error to be used when your application API func name is
// not implemented at this point e.g reserved or in dev.
type NotImplementedErr error

// New returns new standard error msg argument is handled in manner of print
func New(msg ...string) error {
	return errors.New(strings.Join(msg, " "))
}

// Newf returns new standard error.  Arguments are handled in the manner of fmt.Errorf
func Newf(format string, v ...interface{}) error {
	return fmt.Errorf(format, v...)
}

// NewWithContext returns error only if "key" is found in context Value
func NewWithContext(ctx context.Context, ctxkey string, msg ...string) error {
	if v := ctx.Value("err-ctx"); v != nil && ctxkey == v {
		msg = append([]string{"ctx:", ctxkey, "msg:"}, msg...)
		return New(msg...)
	}
	return nil
}

// NewDeprecated constructs ErrDeprecated
func NewDeprecated(s ...string) DeprecatedErr {
	return DeprecatedErr(New(s...))
}

// NewDeprecatedf constructs ErrDeprecated
func NewDeprecatedf(format string, v ...interface{}) DeprecatedErr {
	return DeprecatedErr(Newf(format, v...))
}

// NewNotImplemented constructs ErrNotImplemented
func NewNotImplemented(s ...string) NotImplementedErr {
	return NotImplementedErr(New(s...))
}

// NewNotImplementedf constructs ErrNotImplemented
func NewNotImplementedf(format string, v ...interface{}) NotImplementedErr {
	return NotImplementedErr(Newf(format, v...))
}

// GetTypeOf provided error
func GetTypeOf(err interface{}) string {
	// remove pointer
	typ := "<nil>"
	if err != nil {
		typ = fmt.Sprintf("%T", err)[1:]
	}
	if strings.HasPrefix(typ, "errors.") {
		typ = typ[7:]
	}
	return typ
}

// NewMultiError returns new multierror
func NewMultiError() MultiError {
	return MultiError{}
}

// GetStackTrace returns StackTrace for that error
func GetStackTrace(err error) StackTrace {
	return StackTrace{}
}
