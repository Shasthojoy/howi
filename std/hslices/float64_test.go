// Copyright 2005-2017 Marko Kungla. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// license that can be found in the LICENSE file.

package hslices

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
			if got := MakeFloat64Slice(tt.defaults...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeFloat64Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}
