// Copyright 2005-2017 Marko Kungla.
// Use of this source code is governed by a The MIT License
// license that can be found in the LICENSE file.

package slices

import "strconv"

// MakeFloat64Slice returns new *Float64Slice with optionally default values
func MakeFloat64Slice(defaults ...float64) *Float64Slice {
	return &Float64Slice{raw: append([]float64{}, defaults...)}
}

// Float64Slice wraps a []float64
type Float64Slice struct {
	slice
	raw []float64
}

// Add parses the value into a float64 and appends it to the list of values
func (s *Float64Slice) Add(value string) error {
	if !s.notEmpty {
		s.raw = []float64{}
		s.notEmpty = true
	}

	tmp, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return err
	}

	s.raw = append(s.raw, tmp)
	return nil
}

// Raw returns []float64 slice
func (s *Float64Slice) Raw() []float64 {
	return s.raw
}
