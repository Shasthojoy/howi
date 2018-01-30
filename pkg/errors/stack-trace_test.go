// Copyright 2012 Marko Kungla. All rights reserved.
// Use of this source code is governed by a The Apache-style
// license that can be found in the LICENSE file.

package errors

import (
	"fmt"
	"runtime"
	"testing"
)

func TestFrameLine(t *testing.T) {
	var tests = []struct {
		Frame
		want int
	}{{
		func() Frame {
			var pc, _, _, _ = runtime.Caller(0)
			return newFrame(pc)
		}(),
		19,
	}, {
		func() Frame {
			var pc, _, _, _ = runtime.Caller(1)
			return newFrame(pc)
		}(),
		27,
	}, {
		newFrame(0), // invalid PC
		0,
	}}

	for _, tt := range tests {
		got := tt.Frame.Line()
		want := tt.want
		if want != got {
			t.Errorf("Frame(%v): want: %v, got: %v", uintptr(tt.Frame.pc), want, got)
		}
	}
}

func TestPackageInfo(t *testing.T) {
	tests := []struct {
		path, wantPkgName, wantFnName string
	}{
		{"", "", ""},
		{"errors.New", "errors", "New"},
		{"errors.Newf", "errors", "Newf"},
		{"errors.NewWithContext", "errors", "NewWithContext"},
		{"errors.NewDeprecated", "errors", "NewDeprecated"},
		{"errors.NewDeprecatedf", "errors", "NewDeprecatedf"},
		{"errors.NewNotImplemented", "errors", "NewNotImplemented"},
		{"errors.NewNotImplementedf", "errors", "NewNotImplementedf"},
		{"errors.GetTypeOf", "errors", "GetTypeOf"},
		{"errors.NewMultiError", "errors", "NewMultiError"},
		{"errors.WithStackTrace", "errors", "WithStackTrace"},
		{"runtime.main", "runtime", "main"},
		{"packageInfo", "", "packageInfo"},
		{"github.com/okramlabs/howicli/pkg/errors.packageInfo", "errors", "packageInfo"},
	}

	for _, tt := range tests {
		pkgName, fnName := packageInfo(tt.path)
		if pkgName != tt.wantPkgName {
			t.Errorf("packageInfo(%q): want: %q, got %q", tt.path, tt.wantPkgName, pkgName)
		}
		if fnName != tt.wantFnName {
			t.Errorf("packageInfo(%q): want: %q, got %q", tt.path, tt.wantFnName, fnName)
		}
	}
}

func TestWithStackTrace(t *testing.T) {
	err := WithStackTrace("your error msg")
	if err == nil {
		t.Fatal("err should not be nil")
	}
	if err.Error() != "your error msg" {
		t.Errorf("err.Error() wanr: your error msg got: %s", err.Error())
	}
}

func TestStackTraceFrames(t *testing.T) {
	err := WithStackTrace("your error msg")
	// 	github.com/okramlabs/howicli/pkg/errors/stack-trace_test.go:85 errors.TestStackTraceFrames
	// testing/testing.go:746 testing.tRunner
	// runtime/asm_amd64.s:2337 runtime.goexit
	tests := []struct {
		wantFile          string
		wantLine          int
		wantPkg, wantFunc string
	}{
		{"github.com/okramlabs/howicli/pkg/errors/stack-trace_test.go", 85, "errors", "TestStackTraceFrames"},
	}

	st := err.GetStackTrace()
	for i, tt := range tests {
		if i >= len(st) {
			t.Fatal("Length of tests and StackTrace mismatch")
		}
		if tt.wantFile != st[i].File() {
			t.Errorf("File(): want: %q, got %q", tt.wantFile, st[i].File())
		}
		if tt.wantLine != st[i].Line() {
			t.Errorf("Line(): want: %d, got %d", tt.wantLine, st[i].Line())
		}
		if tt.wantPkg != st[i].Package() {
			t.Errorf("Package(): want: %q, got %q", tt.wantPkg, st[i].Package())
		}
		if tt.wantFunc != st[i].Func() {
			t.Errorf("Func(): want: %q, got %q", tt.wantFunc, st[i].Func())
		}

		want := fmt.Sprintf("%s:%d %s.%s", tt.wantFile, tt.wantLine, tt.wantPkg, tt.wantFunc)
		if want != st[i].String() {
			t.Errorf("String(): want: %q, got %q", want, st[i].String())
		}
	}
}
