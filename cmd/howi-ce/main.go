// Copyright 2005-2017 Marko Kungla. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// license that can be found in the LICENSE file.

package main

import "github.com/howi-ce/howi/std/hlog"

func main() {
	hlog.Colors()
	hlog.Fatal("howi-ce is not yet available as command, use howi instead - see howi --help")
}
