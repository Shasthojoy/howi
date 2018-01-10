// Copyright 2012 Marko Kungla. All rights reserved.
// Use of this source code is governed by a The Apache-style
// license that can be found in the LICENSE file.

package errors

import (
	"fmt"
	"testing"
)

func TestGetStackTrace(t *testing.T) {
	err := New("your error msg")
	st := GetStackTrace(err)
	fmt.Println(st)
}
