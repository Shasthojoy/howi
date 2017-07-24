// Copyright 2005-2017 Marko Kungla. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// license that can be found in the LICENSE file.

package hlog

import (
	"os"
	"os/signal"
	"syscall"

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
	fd         int
	size       tsize
	sch        chan struct{}
	evch       chan tsize
	winch      chan os.Signal
	state      *terminal.State
	monitoring bool
}

func (t *term) monitor() {
	t.monitoring = true
	signal.Notify(t.winch, syscall.SIGWINCH)
	defer func() {
		signal.Stop(t.winch)
		t.monitoring = false
	}()
	for {
		select {
		case <-t.winch:
			w, h, err := terminal.GetSize(t.fd)
			t.size = tsize{w: w, h: h}
			if err != nil {
				return
			}
		case <-t.sch:
			if t != nil && t.state != nil {
				terminal.Restore(t.fd, t.state)
			}
			return
		}
	}
}
