// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package path

import (
	"testing"
)

type PathTest struct {
	path, result string
}

type IsAbsTest struct {
	path  string
	isAbs bool
}

var isAbsTests = []IsAbsTest{
	{"", false},
	{"/", true},
	{"/usr/bin/gcc", true},
	{"..", true},
	{"/a/../bb", true},
	{".", true},
	{"./", true},
	{"lala", true},
}

func TestIsAbs(t *testing.T) {
	for _, test := range isAbsTests {
		p, _ := New(test.path)
		if r := p.IsAbs(); r != test.isAbs {
			t.Errorf("IsAbs(%q) = %v, want %v", test.path, r, test.isAbs)
		}
	}
}
