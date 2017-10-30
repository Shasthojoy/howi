// Copyright 2005-2017 Marko Kungla.
// Use of this source code is governed by a The MIT License
// license that can be found in the LICENSE file.

package filesystem

import (
	"strings"

	"github.com/howi-ce/howi/addon/filesystem/plugin/path"
)

// NewAddon creates new file system Addon instance for given path as wprking directory
// It returns error if current user does not have read access to that path.
func NewAddon(root string) (*Addon, error) {
	wdobj, err := path.NewPlugin(root)
	return &Addon{root: wdobj}, err
}

// Addon for file system abstraction
type Addon struct {
	root path.Plugin
}

// RealAbs returns real abs path of that files system root in system
func (a *Addon) RealAbs() string {
	return a.root.Abs()
}

// IsGitRepository checks if the root is git repository.
func (a *Addon) IsGitRepository() bool {
	return a.root.IsGitRepository()
}

// LoadPath joins provided path to root current fs and returns new path plugin.
func (a *Addon) LoadPath(p string) (path.Plugin, error) {
	if strings.HasPrefix(p, a.root.Abs()) {
		return path.NewPlugin(p)
	}
	return path.NewPlugin(a.root.Join(p))
}
