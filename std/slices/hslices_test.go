// Copyright 2005-2017 Marko Kungla. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// license that can be found in the LICENSE file.

package slices

import (
	"reflect"
	"testing"
)

func TestMakeFloat32Slice(t *testing.T) {
	tests := []struct {
		name     string
		defaults []float32
		want     *Float32Slice
	}{
		{"floats 1", []float32{
			1.000000059604644775390625,
			1.000000059604644775390624,
			1.000000059604644775390626,
			340282346638528859811704183484516925440,
			-340282346638528859811704183484516925440,
			3.402823567e38,
			1e-38,
			4951760157141521099596496896,
		},
			&Float32Slice{
				raw: []float32{1, 1, 1.0000001, 3.4028235e+38, -3.4028235e+38, 3.4028235e+38, 1e-38, 4.9517602e+27},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeFloat32Slice(tt.defaults...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeFloat32Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32Slice_Add(t *testing.T) {
	tests := []struct {
		name    string
		raw     []float32
		value   string
		wantErr bool
	}{
		{"floats 1", []float32{1}, "1", false},
		{"floats 2", []float32{1}, "3.4028236e38", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := MakeFloat32Slice(tt.raw...)
			if err := s.Add(tt.value); (err != nil) != tt.wantErr {
				t.Errorf("Float32Slice.Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

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
			if got := MakeFloat64Slice(tt.defaults...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeFloat64Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}
