// Copyright 2012 Marko Kungla. All rights reserved.
// Use of this source code is governed by a The Apache-style
// license that can be found in the LICENSE file.

package slices

import (
	"testing"
)

func TestStringSlice(t *testing.T) {
	tests := []struct {
		name     string
		defaults []string
	}{
		{
			"string set", []string{
				"994834176",
				"1000000001",
				"1000000000",
				string('⌘'),
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
			if len(got.All()) != len(tt.defaults) {
				t.Fatalf("Len should be same: len(got.All())  = %d want: %d", len(got.All()), len(tt.defaults))
			}
			i := 0
			for el := got.First(); el != nil; el = el.Next() {
				str := el.Value.String()
				if tt.defaults[i] != str {
					t.Errorf("el.Value.String() want: %s got %s", tt.defaults[i], str)
				}
				i++
			}
		})
	}
}
