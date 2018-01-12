// Copyright 2012 Marko Kungla. All rights reserved.
// Use of this source code is governed by a The Apache-style
// license that can be found in the LICENSE file.

package errors

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNewMultiError(t *testing.T) {
	merr := NewMultiError()
	tests := []struct {
		errstr string
	}{
		{"new error"},
		{"another error"},
		{http.ErrShortBody.Error()},
		// errors which should not be added
		{""},
		{"nil"},
	}
	for i := 1; i < len(tests); i++ {
		merr.Append(tests[i].errstr)
		if i != merr.Len() && tests[i].errstr != "" && tests[i].errstr != "nil" {
			t.Errorf("multierror.AppendError() = %d, want %d", i, merr.Len())
		}
	}
}

func TestMultiErrorNil(t *testing.T) {
	merr := NewMultiError()
	if !merr.Nil() {
		t.Error("multierror.Nil() = true, want false")
	}
	if merr.AsError() != nil {
		t.Error("multierror.Nil() = true, want false")
	}
}

func TestMultiErrorError(t *testing.T) {
	merr := NewMultiError()
	if merr.Error() != "(total errors: 0)" {
		t.Errorf("multierror.Error() = %q, want (total errors: 0)", merr.Error())
	}
}

func TestAppendString(t *testing.T) {
	merr := NewMultiError()
	tests := []struct {
		errstr string
	}{
		{"new error"},
		{"another error"},
		{http.ErrShortBody.Error()},
		{"your errors go on"},
	}
	for i := 0; i < len(tests); i++ {
		merr.Append(tests[i].errstr)
		// tests[0].errst shows always last error
		want := fmt.Sprintf("%s (total errors: %d)", tests[i].errstr, merr.Len())
		if merr.Error() != want {
			t.Errorf("multierror.Error() = %q, want %q", merr.Error(), want)
		}
	}
}

func TestAppendStringf(t *testing.T) {
	merr := NewMultiError()
	tests := []struct {
		errstr []string
	}{
		{[]string{"new", "error"}},
		{[]string{"another", "error"}},
		{[]string{http.ErrShortBody.Error()}},
		{[]string{"your", "errors", "go", "on"}},
		{},
	}
	for i := 0; i < len(tests); i++ {
		merr.Appendf("%s", tests[i].errstr)
		// tests[0].errst shows always last error
		want := fmt.Sprintf("%s (total errors: %d)", tests[i].errstr, merr.Len())
		if merr.Error() != want {
			t.Errorf("multierror.Error() = %q, want %q", merr.Error(), want)
		}
	}
}

func TestAppendStringf_nilFormat(t *testing.T) {
	merr := NewMultiError()
	tests := []struct {
		errstr []string
	}{
		{[]string{"new", "error"}},
		{[]string{"another", "error"}},
		{[]string{http.ErrShortBody.Error()}},
		{[]string{"your", "errors", "go", "on"}},
		{},
	}
	for i := 0; i < len(tests); i++ {
		merr.Appendf("", tests[i].errstr)
		// tests[0].errst shows always last error
		want := fmt.Sprintf("%s (total errors: %d)", tests[i].errstr, merr.Len())
		if merr.Error() != want {
			t.Errorf("multierror.Error() = %q, want %q", merr.Error(), want)
		}
	}
}

func TestAppendStringf_nilValue(t *testing.T) {
	merr := NewMultiError()
	merr.Appendf("")
	// tests[0].errst shows always last error
	want := fmt.Sprintf("(total errors: %d)", merr.Len())
	if merr.Error() != want {
		t.Errorf("multierror.Error() = %q, want %q", merr.Error(), want)
	}
}

func TestAsError(t *testing.T) {
	merr := NewMultiError()
	tests := []struct {
		errstr string
	}{
		{"new error"},
		{"another error"},
		{http.ErrShortBody.Error()},
		{"your errors go on"},
	}
	for i := 0; i < len(tests); i++ {
		merr.Append(tests[i].errstr)
		// tests[0].errst shows always last error
		want := fmt.Sprintf("%s (total errors: %d)", tests[i].errstr, merr.Len())
		if merr.AsError().Error() != want {
			t.Errorf("multierror.Error() = %q, want %q", merr.AsError().Error(), want)
		}
	}
}
