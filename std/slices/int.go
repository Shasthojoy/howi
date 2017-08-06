// Copyright 2005-2017 Marko Kungla. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// license that can be found in the LICENSE file.

package slices

import "strconv"

// MakeIntSlice returns new *IntSlice with optionally default values
func MakeIntSlice(defaults ...int) *IntSlice {
	return &IntSlice{raw: append([]int{}, defaults...)}
}

// IntSlice wraps an []int
type IntSlice struct {
	slice
	raw []int
}

// Add parses the value into an integer and appends it to the list of values
func (s *IntSlice) Add(value string) error {
	if !s.notEmpty {
		s.raw = []int{}
		s.notEmpty = true
	}

	tmp, err := strconv.ParseInt(value, 0, 64)
	if err != nil {
		return err
	}

	s.raw = append(s.raw, int(tmp))
	return nil
}

// Raw returns []int slice
func (s *IntSlice) Raw() []int {
	return s.raw
}
