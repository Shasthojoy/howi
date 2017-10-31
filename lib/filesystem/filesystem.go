// Copyright 2005-2017 Marko Kungla.
// Use of this source code is governed by a The MIT License
// license that can be found in the LICENSE file.

package filesystem

import (
	"strings"

	"github.com/howi-ce/howi/lib/filesystem/path"
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
func (fs *Container) LoadPath(p string) (path.Obj, error) {
	if strings.HasPrefix(p, fs.root.Abs()) {
		return path.New(p)
	}
	return path.New(fs.root.Join(p))
}
