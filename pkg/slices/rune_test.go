// Copyright 2012 Marko Kungla. All rights reserved.
// Use of this source code is governed by a The Apache-style
// license that can be found in the LICENSE file.

package slices

import (
	"testing"
)

func TestRuneSlice(t *testing.T) {
	tests := []struct {
		name     string
		defaults []rune
	}{
		{
			"rune set", []rune{
				994834176,
				1000000001,
				1000000000,
				'⌘',
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
			rv := New()
			rv.Append(string(tt.defaults))
			if string(tt.defaults) != rv.First().Value.String() {
				t.Error("failed to get []rune")
			}
		})
	}
}
