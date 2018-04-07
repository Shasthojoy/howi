// Copyright 2012 Marko Kungla. All rights reserved.
// Use of this source code is governed by a The Apache-style
// license that can be found in the LICENSE file.

package slices_test

import (
	"fmt"

	"github.com/digaverse/howi/lib/slices"
)

func Example() {
	// Create a new slice and put some numbers in it.
	slice := slices.New()
	e4 := slice.Append(4)
	e1 := slice.Prepend(1)
	slice.InsertBefore(3, e4)
	slice.InsertAfter(2, e1)

	// Iterate through list and print its contents.
	for e := slice.First(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	// Output:
	// 1
	// 2
	// 3
	// 4
}
