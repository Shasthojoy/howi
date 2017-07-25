// Copyright 2005-2017 Marko Kungla. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// license that can be found in the LICENSE file.

package hlog

import "time"

type timestamp struct {
	raw []byte
}

func (ts *timestamp) now(r uint8) (f []byte) {
	t := time.Now()
	ts.raw, _ = t.MarshalText()
	if r == t0 {
		return nil
	}
	switch r {
	case t2:
		f = ts.raw[11:19]
		break
	default:
		f = ts.raw[:19]
		f[10] = ' '
		break
	}
	return
}
