// Copyright 2012 Marko Kungla. All rights reserved.
// Use of this source code is governed by a The Apache-style
// license that can be found in the LICENSE file.

package slices

import (
	"testing"
)

func TestFloat32Slice(t *testing.T) {
	tests := []struct {
		name     string
		defaults []float32
	}{
		{
			"floats 32", []float32{
				3.44,
				1.000000059604644775390625,
				1.000000059604644775390624,
				1.000000059604644775390626,
				340282346638528859811704183484516925440,
				-340282346638528859811704183484516925440,
				3.402823567e38,
				1e-38,
				4951760157141521099596496896,
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
				f32, _ := el.Value.Float(32)
				if tt.defaults[i] != float32(f32) {
					t.Errorf("el.Value.Float(32) want: %f got %f", tt.defaults[i], float32(f32))
				}
				i++
			}
		})
	}
}
