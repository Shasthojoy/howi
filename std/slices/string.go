// Copyright 2005-2017 Marko Kungla.
// Use of this source code is governed by a The MIT License
// license that can be found in the LICENSE file.

package slices

// MakeStringSlice returns new *StringSlice with optionally default values
func MakeStringSlice(defaults ...string) *StringSlice {
	return &StringSlice{raw: append([]string{}, defaults...)}
}

// StringSlice wraps a []string
type StringSlice struct {
	slice
	raw []string
}

// Add appends the string value to the list of values
func (s *StringSlice) Add(value string) error {
	if !s.notEmpty {
		s.raw = []string{}
		s.notEmpty = true
	}

	s.raw = append(s.raw, value)
	return nil
}

// Raw returns the slice of strings
func (s *StringSlice) Raw() []string {
	return s.raw
}
