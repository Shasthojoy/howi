// Copyright 2012 Marko Kungla. All rights reserved.
// Use of this source code is governed by a The Apache-style
// license that can be found in the LICENSE file.

package errors

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		err  string
		want error
	}{
		{"", fmt.Errorf("")},
		{"foo", fmt.Errorf("foo")},
		{"foo", New("foo")},
		{"string with format spec: %v", errors.New("string with format spec: %v")},
	}

	for _, tt := range tests {
		got := New(tt.err)
		if got.Error() != tt.want.Error() {
			t.Errorf("New.Error(): got: %q, want %q", got, tt.want)
		}
	}
}

func TestErrorf(t *testing.T) {
	tests := []struct {
		err  error
		want string
	}{
		{Newf("error without format spec"), "error without format spec"},
		{Newf("error with %d format spec", 1), "error with 1 format spec"},
		{Newf("error with %q format spec", "var"), "error with \"var\" format spec"},
		{Newf("error with %t format spec", true), "error with true format spec"},
		{Newf("error with %t format spec", false), "error with false format spec"},
	}

	for _, tt := range tests {
		got := tt.err.Error()
		if got != tt.want {
			t.Errorf("Errorf(%v): got: %q, want %q", tt.err, got, tt.want)
		}
	}
}

func TestNewDeprecatedf(t *testing.T) {
	tests := []struct {
		err  error
		want string
	}{
		{NewDeprecatedf("error without format spec"), "error without format spec"},
		{NewDeprecatedf("error with %d format spec", 1), "error with 1 format spec"},
		{NewDeprecatedf("error with %q format spec", "var"), "error with \"var\" format spec"},
		{NewDeprecatedf("error with %t format spec", true), "error with true format spec"},
		{NewDeprecatedf("error with %t format spec", false), "error with false format spec"},
	}

	for _, tt := range tests {
		got := tt.err.Error()
		if got != tt.want {
			t.Errorf("Errorf(%v): got: %q, want %q", tt.err, got, tt.want)
		}
	}
}

func TestNewNotImplementedf(t *testing.T) {
	tests := []struct {
		err  error
		want string
	}{
		{NewNotImplementedf("error without format spec"), "error without format spec"},
		{NewNotImplementedf("error with %d format spec", 1), "error with 1 format spec"},
		{NewNotImplementedf("error with %q format spec", "var"), "error with \"var\" format spec"},
		{NewNotImplementedf("error with %t format spec", true), "error with true format spec"},
		{NewNotImplementedf("error with %t format spec", false), "error with false format spec"},
	}

	for _, tt := range tests {
		got := tt.err.Error()
		if got != tt.want {
			t.Errorf("Errorf(%v): got: %q, want %q", tt.err, got, tt.want)
		}
	}
}

func testErrorFormat(t *testing.T, n int, arg interface{}, format, want string) {
	got := fmt.Sprintf(format, arg)
	gotLines := strings.SplitN(got, "\n", -1)
	wantLines := strings.SplitN(want, "\n", -1)

	if len(wantLines) > len(gotLines) {
		t.Errorf("test %d: wantLines(%d) > gotLines(%d):\n got: %q\nwant: %q", n+1, len(wantLines), len(gotLines), got, want)
		return
	}

	for i, w := range wantLines {
		match, err := regexp.MatchString(w, gotLines[i])
		if err != nil {
			t.Fatal(err)
		}
		if !match {
			t.Errorf("test %d: line %d: fmt.Sprintf(%q, err):\n got: %q\nwant: %q", n+1, i+1, format, got, want)
		}
	}
}
