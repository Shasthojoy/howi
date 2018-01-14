// Copyright 2012 Marko Kungla. All rights reserved.
// Use of this source code is governed by a The Apache-style
// license that can be found in the LICENSE file.

package vars

import "testing"

// These test focus on basic type tests
// https://github.com/okramlabs/howicli/issues/15
func TestType_string(t *testing.T) {
	// string
	val := ValueFromString("string var")
	val2 := Value("string var")
	if val != val2 {
		t.Errorf("want: ValueFromString == Value got: ValueFromString = %q, val2 = %q", val, val2)
	}
}

func TestType_bool(t *testing.T) {
	tests := []struct {
		val  Value
		want string
	}{
		{ValueFromBool(true), "true"},
		{ValueFromBool(false), "false"},
	}
	for _, tt := range tests {
		if tt.val.String() != tt.want {
			t.Errorf("want: %q got %q", tt.val.String(), tt.want)
		}
	}
}
