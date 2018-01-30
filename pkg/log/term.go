// Copyright 2016 Marko Kungla. All rights reserved.
// Use of this source code is governed by a The Apache-style
// license that can be found in the LICENSE file.

package log

import (
	"golang.org/x/crypto/ssh/terminal"
)

var t *term

// Width returns cuurent line with
func Width() int {
	if t != nil {
		return t.size.w
	}
	return 80
}

type tsize struct {
	w int
	h int
}

// Terminal instance
type term struct {
	fd    int
	size  tsize
	sch   chan struct{}
	evch  chan tsize
	state *terminal.State
}
