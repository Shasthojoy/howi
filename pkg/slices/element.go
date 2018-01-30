// Copyright 2012 Marko Kungla. All rights reserved.
// Use of this source code is governed by a The Apache-style
// license that can be found in the LICENSE file.

package slices

import "github.com/okramlabs/howicli/pkg/vars"

// Element is an element of a slice
type Element struct {
	// Next and previous pointers in the slice of elements.
	next, prev *Element

	// The slice to which this element belongs to.
	parent *Slice

	// The value stored with this element.
	Value vars.Value
}

// Next returns the next slice element or nil.
func (e *Element) Next() *Element {
	if p := e.next; e.parent != nil && p != &e.parent.root {
		return p
	}
	return nil
}

// Prev returns the previous slice element or nil.
func (e *Element) Prev() *Element {
	if p := e.prev; e.parent != nil && p != &e.parent.root {
		return p
	}
	return nil
}
