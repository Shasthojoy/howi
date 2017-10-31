// Copyright 2005-2017 Marko Kungla.
// Use of this source code is governed by a The MIT License
// license that can be found in the LICENSE file.

package goprj

import (
	"github.com/howi-ce/howi/lib/goprj/contributors"
	"github.com/howi-ce/howi/lib/vcs/git"
	"github.com/howi-ce/howi/std/errors"
	"github.com/howi-ce/howi/std/vars"
)

// Project holds information about the golang project
type Project struct {
	errs errors.MultiError // internal errors

	Contributor *contributors.Contributor // Current Contributor
	Git         *git.Git                  // Git instance for project
	Vars        vars.Collection
}
