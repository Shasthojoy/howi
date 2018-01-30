// Copyright 2012 Marko Kungla. All rights reserved.
// Use of this source code is governed by a The Apache-style
// license that can be found in the LICENSE file.

package slices

import "testing"

func TestSlice(t *testing.T) {
	l := New()
	checkSlicePtr(t, l, []*Element{})

	// Single element list
	e := l.Prepend("a")
	checkSlicePtr(t, l, []*Element{e})
	l.MoveToBeginning(e)
	checkSlicePtr(t, l, []*Element{e})
	l.MoveToEnd(e)
	checkSlicePtr(t, l, []*Element{e})
	l.Remove(e)
	checkSlicePtr(t, l, []*Element{})

	// Bigger list
	e2 := l.Prepend(2)
	e1 := l.Prepend(1)
	e3 := l.Append(3)
	e4 := l.Append("banana")
	checkSlicePtr(t, l, []*Element{e1, e2, e3, e4})

	l.Remove(e2)
	checkSlicePtr(t, l, []*Element{e1, e3, e4})

	l.MoveToBeginning(e3) // move from middle
	checkSlicePtr(t, l, []*Element{e3, e1, e4})

	l.MoveToBeginning(e1)
	l.MoveToEnd(e3) // move from middle
	checkSlicePtr(t, l, []*Element{e1, e4, e3})

	l.MoveToBeginning(e3) // move from back
	checkSlicePtr(t, l, []*Element{e3, e1, e4})
	l.MoveToBeginning(e3) // should be no-op
	checkSlicePtr(t, l, []*Element{e3, e1, e4})

	l.MoveToEnd(e3) // move from front
	checkSlicePtr(t, l, []*Element{e1, e4, e3})
	l.MoveToEnd(e3) // should be no-op
	checkSlicePtr(t, l, []*Element{e1, e4, e3})

	e2 = l.InsertBefore(2, e1) // insert before front
	checkSlicePtr(t, l, []*Element{e2, e1, e4, e3})
	l.Remove(e2)
	e2 = l.InsertBefore(2, e4) // insert before middle
	checkSlicePtr(t, l, []*Element{e1, e2, e4, e3})
	l.Remove(e2)
	e2 = l.InsertBefore(2, e3) // insert before back
	checkSlicePtr(t, l, []*Element{e1, e4, e2, e3})
	l.Remove(e2)

	e2 = l.InsertAfter(2, e1) // insert after front
	checkSlicePtr(t, l, []*Element{e1, e2, e4, e3})
	l.Remove(e2)
	e2 = l.InsertAfter(2, e4) // insert after middle
	checkSlicePtr(t, l, []*Element{e1, e4, e2, e3})
	l.Remove(e2)
	e2 = l.InsertAfter(2, e3) // insert after back
	checkSlicePtr(t, l, []*Element{e1, e4, e3, e2})
	l.Remove(e2)

	// Check standard iteration.
	sum := 0
	for e := l.First(); e != nil; e = e.Next() {
		i, err := e.Value.AsInt()
		if err == nil {
			sum += i
		}
	}
	if sum != 4 {
		t.Errorf("sum over l = %d, want 4", sum)
	}

	// Clear all elements by iterating
	var next *Element
	for e := l.First(); e != nil; e = next {
		next = e.Next()
		l.Remove(e)
	}
	checkSlicePtr(t, l, []*Element{})
}

func TestExtending(t *testing.T) {
	l1 := New()
	l2 := New()

	l1.Append(1)
	l1.Append(2)
	l1.Append(3)

	l2.Append(4)
	l2.Append(5)

	l3 := New()
	l3.AppendSlice(l1)
	checkSlice(t, l3, []interface{}{1, 2, 3})
	l3.AppendSlice(l2)
	checkSlice(t, l3, []interface{}{1, 2, 3, 4, 5})

	l3 = New()
	l3.PrependSlice(l2)
	checkSlice(t, l3, []interface{}{4, 5})
	l3.PrependSlice(l1)
	checkSlice(t, l3, []interface{}{1, 2, 3, 4, 5})

	checkSlice(t, l1, []interface{}{1, 2, 3})
	checkSlice(t, l2, []interface{}{4, 5})

	l3 = New()
	l3.AppendSlice(l1)
	checkSlice(t, l3, []interface{}{1, 2, 3})
	l3.AppendSlice(l3)
	checkSlice(t, l3, []interface{}{1, 2, 3, 1, 2, 3})

	l3 = New()
	l3.PrependSlice(l1)
	checkSlice(t, l3, []interface{}{1, 2, 3})
	l3.PrependSlice(l3)
	checkSlice(t, l3, []interface{}{1, 2, 3, 1, 2, 3})

	l3 = New()
	l1.AppendSlice(l3)
	checkSlice(t, l1, []interface{}{1, 2, 3})
	l1.PrependSlice(l3)
	checkSlice(t, l1, []interface{}{1, 2, 3})
}

func TestRemove(t *testing.T) {
	l := New()
	e1 := l.Append(1)
	e2 := l.Append(2)
	checkSlicePtr(t, l, []*Element{e1, e2})
	e := l.First()
	l.Remove(e)
	checkSlicePtr(t, l, []*Element{e2})
	l.Remove(e)
	checkSlicePtr(t, l, []*Element{e2})
}

func TestMove(t *testing.T) {
	l := New()
	e1 := l.Append(1)
	e2 := l.Append(2)
	e3 := l.Append(3)
	e4 := l.Append(4)

	l.MoveAfter(e3, e3)
	checkSlicePtr(t, l, []*Element{e1, e2, e3, e4})
	l.MoveBefore(e2, e2)
	checkSlicePtr(t, l, []*Element{e1, e2, e3, e4})

	l.MoveAfter(e3, e2)
	checkSlicePtr(t, l, []*Element{e1, e2, e3, e4})
	l.MoveBefore(e2, e3)
	checkSlicePtr(t, l, []*Element{e1, e2, e3, e4})

	l.MoveBefore(e2, e4)
	checkSlicePtr(t, l, []*Element{e1, e3, e2, e4})
	e2, e3 = e3, e2

	l.MoveBefore(e4, e1)
	checkSlicePtr(t, l, []*Element{e4, e1, e2, e3})
	e1, e2, e3, e4 = e4, e1, e2, e3

	l.MoveAfter(e4, e1)
	checkSlicePtr(t, l, []*Element{e1, e4, e2, e3})
	e2, e3, e4 = e4, e2, e3

	l.MoveAfter(e2, e3)
	checkSlicePtr(t, l, []*Element{e1, e3, e2, e4})
}

// Test PushFront, PushBack, PushFrontList, PushBackSlice with uninitialized List
func TestZeroSlice(t *testing.T) {
	var l1 = new(Slice)
	l1.Prepend(1)
	checkSlice(t, l1, []interface{}{1})

	var l2 = new(Slice)
	l2.Append(1)
	checkSlice(t, l2, []interface{}{1})

	var l3 = new(Slice)
	l3.PrependSlice(l1)
	checkSlice(t, l3, []interface{}{1})

	var l4 = new(Slice)
	l4.AppendSlice(l2)
	checkSlice(t, l4, []interface{}{1})
}

// Test that a slice is not modified when calling InsertBefore with a mark that is not an element of s.
func TestInsertBeforeUnknownMark(t *testing.T) {
	var s Slice
	s.Append(1)
	s.Append(2)
	s.Append(3)
	s.InsertBefore(1, new(Element))
	checkSlice(t, &s, []interface{}{1, 2, 3})
}

// Test that a slice is not modified when calling InsertAfter with a mark that is not an element of s.
func TestInsertAfterUnknownMark(t *testing.T) {
	var s Slice
	s.Append(1)
	s.Append(2)
	s.Append(3)
	s.InsertAfter(1, new(Element))
	checkSlice(t, &s, []interface{}{1, 2, 3})
}

// Test that a slice is not modified when calling MoveAfter or MoveBefore with a mark that is not an element of s.
func TestMoveUnknownMark(t *testing.T) {
	var l1 Slice
	e1 := l1.Append(1)

	var l2 Slice
	e2 := l2.Append(2)

	l1.MoveAfter(e1, e2)
	checkSlice(t, &l1, []interface{}{1})
	checkSlice(t, &l2, []interface{}{2})

	l1.MoveBefore(e1, e2)
	checkSlice(t, &l1, []interface{}{1})
	checkSlice(t, &l2, []interface{}{2})
}
