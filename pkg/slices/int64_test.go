// Copyright 2012 Marko Kungla. All rights reserved.
// Use of this source code is governed by a The Apache-style
// license that can be found in the LICENSE file.

package slices

import (
	"testing"
)

func TestInt64Slice(t *testing.T) {
	tests := []struct {
		name     string
		defaults []int64
	}{
		{
			"int set", []int64{
				999974834176,
				1000000000000000001,
				100000000008388608,
				1000000000016777215,
				100000000016777216,
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
				d64, _ := el.Value.Int(10, 64)
				if tt.defaults[i] != d64 {
					t.Errorf("el.Value.Float(32) want: %d got %d", tt.defaults[i], d64)
				}
				i++
			}
		})
	}
}
