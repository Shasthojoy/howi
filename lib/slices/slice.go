// Copyright 2012 Marko Kungla. All rights reserved.
// Use of this source code is governed by a The Apache-style
// license that can be found in the LICENSE file.

package slices

import (
	"fmt"

	"github.com/digaverse/howi/pkg/vars"
)

// Slice represents a doubly linked list like go's container/list
type Slice struct {
	root Element // sentinel slice element, only &root, root.prev, and root.next are used
	len  int     // current slice length excluding (this) sentinel element
}

// Len returns the number of elements of slice
func (s *Slice) Len() int { return s.len }

// First returns the first element of slice or nil if there are no elements.
func (s *Slice) First() *Element {
	if s.len == 0 {
		return nil
	}
	return s.root.next
}

// Last returns the last element of slice or nil if there are no elements.
func (s *Slice) Last() *Element {
	if s.len == 0 {
		return nil
	}
	return s.root.prev
}

// Remove removes e from slice if e is an element of list slice.
// It returns the element value e.Value.
func (s *Slice) Remove(e *Element) interface{} {
	// if e.parent == s, s must have been initialized when e was inserted
	// in s or s == nil (e is a zero Element) and s.delete will crash
	if e.parent == s {
		s.delete(e)
	}
	return e.Value
}

// Prepend inserts a new element e with value v at the front of slice and returns e.
func (s *Slice) Prepend(v interface{}) *Element {
	s.lazyLoad()
	return s.insertValue(v, &s.root)
}

// Append inserts a new element e with value v at the end of the slice and returns e.
func (s *Slice) Append(v interface{}) *Element {
	s.lazyLoad()
	return s.insertValue(v, s.root.prev)
}

// InsertBefore inserts a new element e with value v immediately before mark and returns e.
// If mark is not an element of the s, slice is not modified.
func (s *Slice) InsertBefore(v interface{}, mark *Element) *Element {
	if mark.parent != s {
		return nil
	}
	// see comment in List.Remove about initialization of l
	return s.insertValue(v, mark.prev)
}

// InsertAfter inserts a new element e with value v immediately after mark and returns e.
// If mark is not an element of s, the slice is not modified.
func (s *Slice) InsertAfter(v interface{}, mark *Element) *Element {
	if mark.parent != s {
		return nil
	}
	// see comment in List.Remove about initialization of l
	return s.insertValue(v, mark)
}

// MoveToBeginning moves element e to the front of slice s.
// If e is not an element of s, the slice is not modified.
func (s *Slice) MoveToBeginning(e *Element) {
	if e.parent != s || s.root.next == e {
		return
	}
	s.insert(s.delete(e), &s.root)
}

// MoveToEnd moves element e to the end of slice s.
// If e is not an element of s, the slice is not modified.
func (s *Slice) MoveToEnd(e *Element) {
	if e.parent != s || s.root.prev == e {
		return
	}
	s.insert(s.delete(e), s.root.prev)
}

// MoveBefore moves element e to its new position before mark.
// If e or mark is not an element of s, or e == mark, the slice is not modified.
func (s *Slice) MoveBefore(e, mark *Element) {
	if e.parent != s || e == mark || mark.parent != s {
		return
	}
	s.insert(s.delete(e), mark.prev)
}

// MoveAfter moves element e to its new position after mark.
// If e or mark is not an element of s, or e == mark, the slice is not modified.
func (s *Slice) MoveAfter(e, mark *Element) {
	if e.parent != s || e == mark || mark.parent != s {
		return
	}
	s.insert(s.delete(e), mark)
}

// AppendSlice inserts a copy of an other slice at the end of the slice s.
// The slices s and other may be the same.
func (s *Slice) AppendSlice(other *Slice) {
	s.lazyLoad()
	for i, e := other.Len(), other.First(); i > 0; i, e = i-1, e.Next() {
		s.insertValue(e.Value, s.root.prev)
	}
}

// PrependSlice inserts a copy of an other slice at the front of slice s.
// The slices s and other may be the same.
func (s *Slice) PrependSlice(other *Slice) {
	s.lazyLoad()
	for i, e := other.Len(), other.Last(); i > 0; i, e = i-1, e.Prev() {
		s.insertValue(e.Value, &s.root)
	}
}

// All returns slice []vars.Value
func (s *Slice) All() []vars.Value {
	s.lazyLoad()
	var res []vars.Value
	for e := s.First(); e != nil; e = e.Next() {
		res = append(res, e.Value)
	}
	return res
}

////////////////////////////////////////////////////////////////////////////////
// PRIVATE API

// Init initializes or clears slice l.
func (s *Slice) init() *Slice {
	s.root.next = &s.root
	s.root.prev = &s.root
	s.len = 0
	return s
}

// lazyLoad lazily initializes a Slice.
// Useful when Slice struct is constucted and method is called which needs
// basics being setup
func (s *Slice) lazyLoad() {
	if s.root.next == nil {
		s.init()
	}
}

// insert inserts e after at, increments slice.len, and returns element.
func (s *Slice) insert(e, at *Element) *Element {
	n := at.next
	at.next = e
	e.prev = at
	e.next = n
	n.prev = e
	e.parent = s
	s.len++
	return e
}

// insertValue is a convenience wrapper for insert(&Element{Value: v}, at).
func (s *Slice) insertValue(v interface{}, at *Element) *Element {
	val := vars.NewValue(fmt.Sprintf("%v", v))
	return s.insert(&Element{Value: val}, at)
}

// delete removes e from its slice, decrements s.len, and returns e.
func (s *Slice) delete(e *Element) *Element {
	e.prev.next = e.next
	e.next.prev = e.prev
	e.next = nil // avoid memory leaks
	e.prev = nil // avoid memory leaks
	e.parent = nil
	s.len--
	return e
}
