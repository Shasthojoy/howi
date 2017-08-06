// Copyright 2005-2017 Marko Kungla. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// license that can be found in the LICENSE file.

package slices

import "strconv"

// MakeRuneSlice returns new *RuneSlice with optionally default values
func MakeRuneSlice(defaults ...rune) *RuneSlice {
	return &RuneSlice{raw: append([]rune{}, defaults...)}
}

// RuneSlice wraps an []rune
type RuneSlice struct {
	slice
	raw []rune
}

// Add parses the value into an rune and appends it to the list of values
func (s *RuneSlice) Add(value string) error {
	if !s.notEmpty {
		s.raw = []rune{}
		s.notEmpty = true
	}

	tmp, err := strconv.ParseInt(value, 0, 64)
	if err != nil {
		return err
	}

	s.raw = append(s.raw, rune(tmp))
	return nil
}

// Raw returns rune slice
func (s *RuneSlice) Raw() []rune {
	return s.raw
}
