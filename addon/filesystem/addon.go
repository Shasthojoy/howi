// Copyright 2005-2017 Marko Kungla. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// license that can be found in the LICENSE file.

package filesystem

import "github.com/howi-ce/howi/addon/filesystem/plugin/path"

// Addon creates new FileSystem instance for given path as root
// It returns error if current user does not have read access to that path.
func Addon(root string) (*FileSystem, error) {
	wdobj, err := path.Plugin(root)
	return &FileSystem{root: wdobj}, err
}
