// Copyright 2012 Marko Kungla. All rights reserved.
// Use of this source code is governed by a The Apache-style
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
			got := MakeFloat32Slice(tt.defaults...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeFloat32Slice() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got.raw, got.Raw()) {
				t.Errorf("MakeFloat32Slice() raw %v and .Raw %v should equal", got.raw, got.Raw())
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
		{"floats 2", []float32{3.4028236}, "3.4028236e38", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := MakeFloat32Slice(tt.raw...)
			if err := s.Add(tt.value); (err != nil) != tt.wantErr {
				t.Errorf("MakeFloat32Slice.Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
