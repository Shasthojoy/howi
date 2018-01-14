// Copyright 2012 Marko Kungla. All rights reserved.
// Use of this source code is governed by a The Apache-style
// license that can be found in the LICENSE file.

package slices

import (
	"reflect"
	"testing"
)

func TestMakeRuneSlice(t *testing.T) {
	tests := []struct {
		name     string
		defaults []rune
		want     *RuneSlice
	}{
		{"rune set", []rune{
			994834176,
			1000000001,
			1000000000,
			'⌘',
		},
			&RuneSlice{
				raw: []rune{
					994834176,
					1000000001,
					1000000000,
					'⌘',
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MakeRuneSlice(tt.defaults...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeRuneSlice() = %v, want %v", got, tt.want)
			}

			if err := got.Add("1"); err != nil {
				t.Errorf("RuneSlice(.Add() error = %v", err)
			}

			if !reflect.DeepEqual(got.raw, got.Raw()) {
				t.Errorf("RuneSlice( raw %v and .Raw %v should equal", got.raw, got.Raw())
			}
		})
	}
}

func TestRuneSlice_Add(t *testing.T) {
	tests := []struct {
		name    string
		raw     []rune
		value   string
		wantErr bool
	}{
		{"int 1", []rune{1}, "1", false},
		{"int 2", []rune{23}, "x", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := MakeRuneSlice(tt.raw...)
			if err := s.Add(tt.value); (err != nil) != tt.wantErr {
				t.Errorf("RuneSlice(.Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
