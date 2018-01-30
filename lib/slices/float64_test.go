// Copyright 2012 Marko Kungla. All rights reserved.
// Use of this source code is governed by a The Apache-style
// license that can be found in the LICENSE file.

package slices

import (
	"testing"
)

func TestNewFloat64Slice(t *testing.T) {
	tests := []struct {
		name     string
		defaults []float64
	}{
		{
			"floats 64", []float64{
				99999999999999974834176,
				100000000000000000000001,
				1000008388608,
				100000000000000016777215,
				100000000000000016777216,
				1.0000000000000003e+23,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New()
			for _, el := range tt.defaults {
				got.Append(el)
			}
			if got.Len() != len(tt.defaults) {
				t.Fatalf("Len should be same: got.Len() = %d want: %d", got.Len(), len(tt.defaults))
			}
			i := 0
			for el := got.First(); el != nil; el = el.Next() {
				f64, _ := el.Value.Float(64)
				if tt.defaults[i] != f64 {
					t.Errorf("el.Value.Float(32) want: %f got %f", tt.defaults[i], f64)
				}
				i++
			}
		})
	}
}
