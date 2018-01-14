// Copyright 2012 Marko Kungla. All rights reserved.
// Use of this source code is governed by a The Apache-style
// license that can be found in the LICENSE file.

package slices

import (
	"reflect"
	"testing"
)

func TestMakeStringSlice(t *testing.T) {
	tests := []struct {
		name     string
		defaults []string
		want     *StringSlice
	}{
		{"string set", []string{
			"994834176",
			"1000000001",
			"1000000000",
			string('⌘'),
		},
			&StringSlice{
				raw: []string{
					"994834176",
					"1000000001",
					"1000000000",
					string('⌘'),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MakeStringSlice(tt.defaults...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeStringSlice() = %v, want %v", got, tt.want)
			}

			if err := got.Add("1"); err != nil {
				t.Errorf("StringSlice(.Add() error = %v", err)
			}

			if !reflect.DeepEqual(got.raw, got.Raw()) {
				t.Errorf("StringSlice( raw %v and .Raw %v should equal", got.raw, got.Raw())
			}
		})
	}
}

// func TestStringSlice_String(t *testing.T) {
// 	tests := []struct {
// 		name  string
// 		raw   []string
// 		value string
// 	}{
// 		{"string 1", []string{"xxx"}, "xxx"},
// 		{"string 2", []string{"yyy"}, "yyy"},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			s := MakeStringSlice(tt.raw...)
// 			if s.Raw() != tt.value {
// 				t.Errorf("StringSlice.String() %q != %q", s.String(), tt.value)
// 			}
// 		})
// 	}
// }
