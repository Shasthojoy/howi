// Copyright 2005-2017 Marko Kungla. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// license that can be found in the LICENSE file.

package filesystem

import "github.com/howi-ce/howi/addon/filesystem/plugin/path"

// FileSystem is wdfs locked for path
type FileSystem struct {
	root path.Path
}
