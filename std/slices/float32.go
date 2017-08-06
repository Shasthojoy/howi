// Copyright 2005-2017 Marko Kungla. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// license that can be found in the LICENSE file.

package slices

import "strconv"

// MakeFloat32Slice returns new *Float32Slice with optionally default values
func MakeFloat32Slice(defaults ...float32) *Float32Slice {
	return &Float32Slice{raw: append([]float32{}, defaults...)}
}

// Float32Slice wraps a []float32
type Float32Slice struct {
	slice
	raw []float32
}

// Add parses the value into a float32 and appends it to the list of values
func (s *Float32Slice) Add(value string) error {
	if !s.notEmpty {
		s.raw = []float32{}
		s.notEmpty = true
	}

	tmp, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return err
	}

	s.raw = append(s.raw, float32(tmp))
	return nil
}

// Raw returns []float32 slice
func (s *Float32Slice) Raw() []float32 {
	return s.raw
}
