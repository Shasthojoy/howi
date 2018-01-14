// Copyright 2012 Marko Kungla. All rights reserved.
// Use of this source code is governed by a The Apache-style
// license that can be found in the LICENSE file.

package slices

import (
	"reflect"
	"testing"
)

func TestMakeInt64Slice(t *testing.T) {
	tests := []struct {
		name     string
		defaults []int64
		want     *Int64Slice
	}{
		{"int set", []int64{
			999974834176,
			1000000000000000001,
			100000000008388608,
			1000000000016777215,
			100000000016777216,
		},
			&Int64Slice{
				raw: []int64{
					999974834176,
					1000000000000000001,
					100000000008388608,
					1000000000016777215,
					100000000016777216,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MakeInt64Slice(tt.defaults...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeIntSlice() = %v, want %v", got, tt.want)
			}

			if err := got.Add("1"); err != nil {
				t.Errorf("IntSlice.Add() error = %v", err)
			}

			if !reflect.DeepEqual(got.raw, got.Raw()) {
				t.Errorf("IntSlice raw %v and .Raw %v should equal", got.raw, got.Raw())
			}
		})
	}
}

func TestInt64Slice_Add(t *testing.T) {
	tests := []struct {
		name    string
		raw     []int64
		value   string
		wantErr bool
	}{
		{"int 1", []int64{1}, "1", false},
		{"int 2", []int64{23}, "x", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := MakeInt64Slice(tt.raw...)
			if err := s.Add(tt.value); (err != nil) != tt.wantErr {
				t.Errorf("Int64Slice.Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
