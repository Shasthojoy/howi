// Copyright 2016 Marko Kungla. All rights reserved.
// Use of this source code is governed by a The Apache-style
// license that can be found in the LICENSE file.

package filesystem

import (
	"path/filepath"
	"strings"

	"github.com/okramlabs/howi/lib/filesystem/path"
	"github.com/okramlabs/howi/pkg/errors"
)

// Open creates new file system instance for given path as wprking directory
// It returns error if current user does not have read access to that path.
func Open(root string) (*Container, error) {
	wdobj, err := path.New(root)
	return &Container{root: wdobj}, err
}

// Container for file system abstraction
type Container struct {
	root path.Obj
}

// RealAbs returns real abs path of that files system root in system
func (fs *Container) RealAbs() string {
	return fs.root.Abs()
}

// IsDir checks if a given root path is a diirectory
func (fs *Container) IsDir() bool {
	return fs.root.IsDir()
}

// InGOPATH reports whether root path is in GOPATH
func (fs *Container) InGOPATH() bool {
	return fs.root.InGOPATH()
}

// IsGitRepository checks if the root is git repository.
func (fs *Container) IsGitRepository() bool {
	return fs.root.IsGitRepository()
}

// LoadPath joins provided path to root current fs and returns new path plugin.
func (fs *Container) Resolve(p ...string) (path.Obj, error) {
	if len(p) == 0 {
		return fs.root, errors.New("no path segment passed")
	}
	s := filepath.Join(p...)
	if strings.HasPrefix(s, fs.root.Abs()) {
		return path.New(s)
	}
	return path.New(fs.root.Join(s))
}
