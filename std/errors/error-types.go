// Copyright 2005-2017 Marko Kungla.
// Use of this source code is governed by a The MIT License
// license that can be found in the LICENSE file.

package errors

// ErrDeprecated is error returned when git command with matching name is deprecated.
type ErrDeprecated error

// NewErrDeprecated constructs ErrDeprecated
func NewErrDeprecated(s ...string) ErrDeprecated {
	return ErrDeprecated(New(s...))
}

// NewErrDeprecatedf constructs ErrDeprecated
func NewErrDeprecatedf(format string, v ...interface{}) ErrDeprecated {
	return ErrDeprecated(Newf(format, v...))
}

// ErrNotImplemented is error returned when git command with matching name is
// not implemented at this point and may be removed in next release.
type ErrNotImplemented error

// NewErrNotImplemented constructs ErrNotImplemented
func NewErrNotImplemented(s ...string) ErrNotImplemented {
	return ErrNotImplemented(New(s...))
}

// NewErrNotImplementedf constructs ErrNotImplemented
func NewErrNotImplementedf(format string, v ...interface{}) ErrNotImplemented {
	return ErrNotImplemented(Newf(format, v...))
}
