// Copyright 2005-2017 Marko Kungla. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// license that can be found in the LICENSE file.

package git

import (
	"os"
	"os/exec"

	"github.com/howi-ce/howi/std/herrors"
	"github.com/howi-ce/howi/std/hos/hexec"
)

// PRIVATE STATIC API

func cmdgit(v ...string) (hexec.Output, error) {
	b, err := exec.Command("git", v...).Output()
	return hexec.Output(b), err
}

func cmdgitInPath(p string, v ...string) (hexec.Output, error) {
	cur, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	if err := os.Chdir(p); err != nil {
		return nil, err
	}
	b, rderr := exec.Command("git", v...).Output()
	if err := os.Chdir(cur); err != nil {
		return nil, err
	}
	return hexec.Output(b), rderr
}

// newErrDeprecated constructs herrors.ErrDeprecated
func newErrDeprecated(method string, alternative string, docs string) herrors.ErrDeprecated {
	return herrors.NewErrDeprecatedf("method %s is deprecated - alternatives (%s) - docs (%s)",
		method, alternative, docs)
}

// newErrNotImplemented constructs herrors.ErrNotImplemented
func newErrNotImplemented(method string, reason ...string) herrors.ErrNotImplemented {
	return herrors.NewErrNotImplementedf(
		"method %s is not implemented - it may be removed in next version or is in development (%s)",
		method, reason)
}
