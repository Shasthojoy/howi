// Copyright 2012 Marko Kungla. All rights reserved.
// Use of this source code is governed by a The Apache-style
// license that can be found in the LICENSE file.

package slices

import (
	"testing"
)

func checkSliceLen(t *testing.T, s *Slice, len int) bool {
	if n := s.Len(); n != len {
		t.Errorf("Slice.Len() = %d, want %d", n, len)
		return false
	}
	return true
}

func checkSlicePtr(t *testing.T, s *Slice, es []*Element) {
	root := &s.root

	if !checkSliceLen(t, s, len(es)) {
		return
	}

	// zero length lists must be the zero value or properly initialized (sentinel circle)
	if len(es) == 0 {
		if s.root.next != nil && s.root.next != root || s.root.prev != nil && s.root.prev != root {
			t.Errorf("Slice.root.next = %p, Slice.root.prev = %p; both should both be nil or %p", s.root.next, s.root.prev, root)
		}
		return
	}
	// len(es) > 0

	// check internal and external prev/next connections
	for i, e := range es {
		prev := root
		Prev := (*Element)(nil)
		if i > 0 {
			prev = es[i-1]
			Prev = prev
		}
		if p := e.prev; p != prev {
			t.Errorf("elt[%d](%p).prev = %p, want %p", i, e, p, prev)
		}
		if p := e.Prev(); p != Prev {
			t.Errorf("elt[%d](%p).Prev() = %p, want %p", i, e, p, Prev)
		}

		next := root
		Next := (*Element)(nil)
		if i < len(es)-1 {
			next = es[i+1]
			Next = next
		}
		if n := e.next; n != next {
			t.Errorf("elt[%d](%p).next = %p, want %p", i, e, n, next)
		}
		if n := e.Next(); n != Next {
			t.Errorf("elt[%d](%p).Next() = %p, want %p", i, e, n, Next)
		}
	}
}

func checkSlice(t *testing.T, s *Slice, es []interface{}) {
	if !checkSliceLen(t, s, len(es)) {
		return
	}

	i := 0
	for e := s.First(); e != nil; e = e.Next() {
		le, _ := e.Value.AsInt()
		if le != es[i] {
			t.Errorf("elt[%d].Value = %v, want %v", i, le, es[i])
		}
		i++
	}
}
