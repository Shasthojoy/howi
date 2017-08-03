// Copyright 2005-2017 Marko Kungla. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// license that can be found in the LICENSE file.

package hexec

import "strings"

// Output of git commands
type Output []byte

// Lines as string slice
func (o Output) Lines() []string {
	return strings.Split(string(o), "\n")
}

// String retruns string representation of output
func (o Output) String() string {
	return strings.Trim(string(o), "\n")
}
