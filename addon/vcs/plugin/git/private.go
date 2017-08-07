// Copyright 2005-2017 Marko Kungla. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// license that can be found in the LICENSE file.

package git

import (
	"os"
	goOsExec "os/exec"

	"github.com/howi-ce/howi/std/errors"
	"github.com/howi-ce/howi/std/os/exec"
)

// PRIVATE STATIC API

func cmdgit(v ...string) (exec.Output, error) {
	b, err := goOsExec.Command("git", v...).Output()
	return exec.Output(b), err
}

func cmdgitInPath(p string, v ...string) (exec.Output, error) {
	cur, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	if err := os.Chdir(p); err != nil {
		return nil, err
	}
	b, rderr := goOsExec.Command("git", v...).Output()
	if err := os.Chdir(cur); err != nil {
		return nil, err
	}
	return exec.Output(b), rderr
}

// newErrDeprecated constructs errors.ErrDeprecated
func newErrDeprecated(method string, alternative string, docs string) errors.ErrDeprecated {
	return errors.NewErrDeprecatedf("method %s is deprecated - alternatives (%s) - docs (%s)",
		method, alternative, docs)
}

// newErrNotImplemented constructs errors.ErrNotImplemented
func newErrNotImplemented(method string, reason ...string) errors.ErrNotImplemented {
	return errors.NewErrNotImplementedf(
		"method %s is not implemented - it may be removed in next version or is in development (%s)",
		method, reason)
}
