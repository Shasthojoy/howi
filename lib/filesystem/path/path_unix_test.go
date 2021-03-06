// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !windows

package path

import (
	"runtime"
	"testing"
)

var cleantests = []PathTest{
	// Already clean

	{"/.", "/"},
	{"/abc", "/abc"},
	{"/abc/def", "/abc/def"},
	{"/a/b/c", "/a/b/c"},
	{"/..", "/"},
	{"/../..", "/"},
	{"/./", "/"},
	{"/../", "/"},
	{"/../../", "/"},
	{"/../../abc", "/abc"},
	{"/abc", "/abc"},
	{"/", "/"},

	// Remove trailing slash
	{"/abc/", "/abc"},
	{"/abc/def/", "/abc/def"},
	{"/a/b/c/", "/a/b/c"},

	{"/abc/", "/abc"},

	// Remove doubled slash
	{"/abc//def//ghi", "/abc/def/ghi"},
	{"//abc", "/abc"},
	{"///abc", "/abc"},
	{"//abc//", "/abc"},
	{"/abc//", "/abc"},

	// Remove . elements
	{"/abc/./def", "/abc/def"},
	{"/./abc/def", "/abc/def"},
	{"/abc/.", "/abc"},

	// Remove .. elements
	{"/abc/def/ghi/../jkl", "/abc/def/jkl"},
	{"/abc/def/../ghi/../jkl", "/abc/jkl"},
	{"/abc/def/..", "/abc"},
	{"/abc/def/../..", "/"},
	{"/abc/def/../..", "/"},
	{"/abc/def/../..", "/"},
	{"/abc/def/../../..", "/"},
	{"/abc/def/../../../ghi/jkl/../../mno", "/mno"},

	// Combinations
	{"/abc/./../def", "/def"},
	{"/abc//./../def", "/def"},
	{"/abc/../../././../def", "/def"},
}

func TestClean(t *testing.T) {
	for _, test := range cleantests {
		p, _ := New(test.path)
		if s := p.Clean(); s != test.result {
			t.Errorf("Clean(%q) = %q, want %q", test.path, s, test.result)
		}
		r, _ := New(test.result)
		if s := r.Clean(); s != test.result {
			t.Errorf("Clean(%q) = %q, want %q", test.result, s, test.result)
		}
	}
}

func TestCleanMallocs(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping malloc count in short mode")
	}
	if runtime.GOMAXPROCS(0) > 1 {
		t.Log("skipping AllocsPerRun checks; GOMAXPROCS>1")
		return
	}

	for _, test := range cleantests {
		p, _ := New(test.result)
		allocs := testing.AllocsPerRun(100, func() { p.Clean() })
		if allocs > 0 {
			t.Errorf("Clean(%q): %v allocs, want zero", test.result, allocs)
		}
	}
}

type SplitTest struct {
	path, dir, file string
}

var splittests = []SplitTest{
	{"/a/b", "/a/", "b"},
	{"/a/", "/", "a"},
	{"/a", "/", "a"},
	{"/", "/", ""},
}

func TestSplit(t *testing.T) {
	for _, test := range splittests {
		p, _ := New(test.path)
		if d, f := p.Split(); d != test.dir || f != test.file {
			t.Errorf("Split(%q) = %q, %q, want %q, %q", test.path, d, f, test.dir, test.file)
		}
	}
}

type JoinTest struct {
	elem []string
	path string
}

var jointests = []JoinTest{
	// zero parameters
	{[]string{}, ""},

	// one parameter
	{[]string{""}, ""},
	{[]string{"a"}, "a"},

	// two parameters
	{[]string{"a", "b"}, "a/b"},
	{[]string{"a", ""}, "a"},
	{[]string{"", "b"}, "b"},
	{[]string{"/", "a"}, "/a"},
	{[]string{"/", ""}, "/"},
	{[]string{"a/", "b"}, "a/b"},
	{[]string{"a/", ""}, "a"},
	{[]string{"", ""}, ""},
}

func TestJoin(t *testing.T) {
	for _, test := range jointests {
		tt, _ := New("")
		if p := tt.Join(test.elem...); p != test.path {
			t.Errorf("join(%q) = %q, want %q", test.elem, p, test.path)
		}
	}
}

type ExtTest struct {
	path, ext string
}

var exttests = []ExtTest{
	{"path.go", ".go"},
	{"path.pb.go", ".go"},
	{"a.dir/b", ""},
	{"a.dir/b.go", ".go"},
	{"a.dir/", ""},
}

func TestExt(t *testing.T) {
	for _, test := range exttests {
		p, _ := New(test.path)
		if x := p.Ext(); x != test.ext {
			t.Errorf("Ext(%q) = %q, want %q", test.path, x, test.ext)
		}
	}
}

var basetests = []PathTest{
	// Already clean
	{"", ""},
	{"/.", "/"},
	{"/", "/"},
	{"////", "/"},
	{"x/", "x"},
	{"abc", "abc"},
	{"abc/def", "def"},
	{"a/b/.x", ".x"},
	{"a/b/c.", "c."},
	{"a/b/c.x", "c.x"},
}

func TestBase(t *testing.T) {
	for _, test := range basetests {
		p, _ := New(test.path)
		if s := p.Base(); s != test.result {
			t.Errorf("Base(%q) = %q, want %q", test.path, s, test.result)
		}
	}
}

var dirtests = []PathTest{
	{"", ""},
	{"/.", "/"},
	{"/", "/"},
	{"////", "/"},
	{"/foo", "/"},
	{"/x/", "/"},
	{"/abc", "/"},
	{"/abc/def", "/abc"},
	{"/abc////def", "/abc"},
	{"/a/b/.x", "/a/b"},
	{"/a/b/c.", "/a/b"},
	{"/a/b/c.x", "/a/b"},
}

func TestDir(t *testing.T) {
	for _, test := range dirtests {
		p, _ := New(test.path)
		if s := p.Dir(); s != test.result {
			t.Errorf("Dir(%q) = %q, want %q", test.path, s, test.result)
		}
	}
}
