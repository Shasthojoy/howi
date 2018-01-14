// Copyright 2012 Marko Kungla. All rights reserved.
// Use of this source code is governed by a The Apache-style
// license that can be found in the LICENSE file.

package vars

import (
	"strconv"
	"strings"
)

// Value describes the variable value
type Value string

func (v Value) String() string {
	return string(v)
}

// ParseBool calls strconv.ParseBool on Value
func (v Value) ParseBool() (bool, error) {
	return strconv.ParseBool(v.String())
}

// ParseFloat calls strconv.ParseFloat on Value
func (v Value) ParseFloat(bitSize int) (float64, error) {
	return strconv.ParseFloat(v.String(), bitSize)
}

// ParseInt calls strconv.ParseInt on Value
func (v Value) ParseInt(base int, bitSize int) (int64, error) {
	return strconv.ParseInt(v.String(), base, bitSize)
}

// ParseUint calls strconv.ParseUint on Value
func (v Value) ParseUint(base int, bitSize int) (uint64, error) {
	return strconv.ParseUint(v.String(), base, bitSize)
}

// ParseFields calls strings.Fields on Value string
func (v Value) ParseFields() []string {
	return strings.Fields(v.String())
}

// ParseComplex64 tries to split Value to strings.Fields and
// use 2 first fields to return complex64
func (v Value) ParseComplex64() (complex64, error) {
	var err error
	fields := v.ParseFields()
	if len(fields) != 2 {
		return complex64(0), strconv.ErrSyntax
	}

	var f1 float64
	var f2 float64
	if f1, err = strconv.ParseFloat(fields[0], 32); err != nil {
		return complex64(0), err
	}
	if f2, err = strconv.ParseFloat(fields[1], 32); err != nil {
		return complex64(0), err
	}
	return complex64(complex(f1, f2)), nil
}

// ParseComplex128 tries to split Value to strings.Fields and
// use 2 first fields to return complex128
func (v Value) ParseComplex128() (complex128, error) {
	var err error
	fields := v.ParseFields()
	if len(fields) != 2 {
		return complex128(0), strconv.ErrSyntax
	}
	var f1 float64
	var f2 float64
	if f1, err = strconv.ParseFloat(fields[0], 64); err != nil {
		return complex128(0), err
	}
	if f2, err = strconv.ParseFloat(fields[1], 64); err != nil {
		return complex128(0), err
	}
	return complex128(complex(f1, f2)), nil
}

// Len returns the length of the string representation of the Value
func (v Value) Len() int {
	return len(v.String())
}

// Empty returns true if this Value is empty
func (v Value) Empty() bool {
	return v.Len() == 0
}

// ValueFromString trims spaces and returns Value
func ValueFromString(val string) Value {
	return Value(strings.TrimSpace(val))
}

// ValueFromBool returns Value pared from bool
func ValueFromBool(val bool) Value {
	return Value(strconv.FormatBool(val))
}
