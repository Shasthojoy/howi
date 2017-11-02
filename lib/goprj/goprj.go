// Copyright 2005-2017 Marko Kungla.
// Use of this source code is governed by a The MIT License
// license that can be found in the LICENSE file.

package goprj

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"

	"github.com/blang/semver"
	"github.com/howi-ce/howi/lib/goprj/contributors"
	"github.com/howi-ce/howi/lib/vcs/git"
	"github.com/howi-ce/howi/std/vars"
)

// Open creates new Project instance with provided path as working directory
func Open(prjPath string) (*Project, error) {
	prj := &Project{}
	var err error

	// Get go variables
	govars, err := GetGoVars()
	if err != nil {
		return nil, err
	}

	// Create new empty contibutor intsance to hold current users info
	prj.Contributor = contributors.NewContributor()

	// Load Git library
	prj.Git, err = git.New(prjPath)
	prj.errs.AppendError(err)

	// Get environment vars and merge these with govars
	v := append(os.Environ(), govars...)
	// Get git variables
	gitconf, err := git.GlobalConfig()
	if err != nil {
		return nil, err
	}
	v = append(v, gitconf...)
	prj.Vars = vars.ParseFromStrings(v)

	// Set some gontributor information from git gonfig or from os.User
	if err == nil {
		name := prj.Vars.Getvar("user.name")
		email := prj.Vars.Getvar("user.email")
		if email != "" {
			prj.Contributor.AddEmail(name.String()+"<"+email.String()+">", true)
			prj.Contributor.SetName(name.String())
		}
		if name == "" {
			u, oerr := user.Current()
			if oerr == nil {
				prj.Contributor.SetName(u.Name)
			}
			prj.errs.AppendError(oerr)
		}
	}

	// Check is working directory actually directory
	if !prj.Git.FS.IsDir() && err == nil {
		prj.errs.AppendString("provided path can not be used as working directory")
	}

	// Check is working directory in GOPATH
	if !prj.Git.FS.InGOPATH() {
		prj.errs.AppendString("provided path is not in GOPATH")
	}

	// If project exists
	if prj.Exists() {
		prj.Config.Reload()
		prj.createVarsGit()
		prj.Version, err = semver.Make(func() string {
			commits, _ := prj.Vars.Getvar("HOWI_GIT_NUM_COMMITS_SINCE_LAST_TAG").ParseInt(10, 0)
			tag := prj.Vars.Getvar("HOWI_GIT_LAST_TAG").String()
			if commits > 0 {
				return fmt.Sprintf("%s+%d.%s", tag, commits,
					prj.Vars.Getvar("HOWI_GIT_COMMIT_SHA_ABBREV").String())
			}
			return tag
		}())
		if err == nil {
			// We need Git vars to greate project vars and that version is parsed
			prj.createVarsPrj()
		}
		prj.errs.AppendError(err)
	}

	return prj, prj.errs.AsError()
}

// GetGoVars returns values from output of go env as slice
func GetGoVars() (govars []string, err error) {
	goenv, err := exec.Command("go", "env").Output()
	if err == nil {
		govars = strings.Split(string(goenv), "\n")
	}
	return
}
