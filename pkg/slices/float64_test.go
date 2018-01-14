// Copyright 2012 Marko Kungla. All rights reserved.
// Use of this source code is governed by a The Apache-style
// license that can be found in the LICENSE file.

package slices

import (
	"reflect"
	"testing"
)

func TestMakeFloat64Slice(t *testing.T) {
	tests := []struct {
		name     string
		defaults []float64
		want     *Float64Slice
	}{
		{"floats 1", []float64{
			99999999999999974834176,
			100000000000000000000001,
			100000000000000008388608,
			100000000000000016777215,
			100000000000000016777216,
		},
			&Float64Slice{
				raw: []float64{
					9.999999999999997e+22,
					1.0000000000000001e+23,
					1.0000000000000001e+23,
					1.0000000000000001e+23,
					1.0000000000000003e+23,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MakeFloat64Slice(tt.defaults...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeFloat64Slice() = %v, want %v", got, tt.want)
			}

			if err := got.Add("0.1"); err != nil {
				t.Errorf("Float64Slice.Add() error = %v", err)
			}
			if !reflect.DeepEqual(got.raw, got.Raw()) {
				t.Errorf("Float64Slice raw %v and .Raw %v should equal", got.raw, got.Raw())
			}
		})
	}
}

func TestFloat64Slice_Add(t *testing.T) {
	tests := []struct {
		name    string
		raw     []float64
		value   string
		wantErr bool
	}{
		{"floats 1", []float64{1}, "1", false},
		{"floats 2", []float64{1.0000000000000003e+23}, "x", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := MakeFloat64Slice(tt.raw...)
			if err := s.Add(tt.value); (err != nil) != tt.wantErr {
				t.Errorf("MakeFloat32Slice.Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
