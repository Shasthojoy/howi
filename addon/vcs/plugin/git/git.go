// Copyright 2005-2017 Marko Kungla. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// license that can be found in the LICENSE file.

package git

import "os/exec"

// STATIC API

// GlobalConfig returns global git config
func GlobalConfig() ([]string, error) {
	gitconfig, err := cmdgit("config", "--global", "--list")
	return gitconfig.Lines(), err
}

// LookPath searches for an executable git binary
// in the directories named by the PATH environment variable.
// If file contains a slash, it is tried directly and the PATH is not consulted.
// The result may be an absolute path or a path relative to the current directory.
func LookPath() (string, error) {
	var err error
	if executable == "" {
		executable, err = exec.LookPath("git")
	}
	return executable, err
}
