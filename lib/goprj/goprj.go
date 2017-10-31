// Copyright 2005-2017 Marko Kungla.
// Use of this source code is governed by a The MIT License
// license that can be found in the LICENSE file.

package goprj

import (
	"os"
	"os/exec"
	"os/user"
	"strings"

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
	// Get environment vars and merge these with govars
	v := append(os.Environ(), govars...)

	// Get git variables
	gitconf, err := git.GlobalConfig()
	if err != nil {
		return nil, err
	}

	v = append(v, gitconf...)
	prj.Vars = vars.ParseFromStrings(v)

	// Create new empty contibutor intsance to hold current users info
	prj.Contributor = contributors.NewContributor()

	// Load Git library
	prj.Git, err = git.New(prjPath)
	prj.errs.AppendError(err)

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
	prj.errs.AppendError(err)

	// Check is working directory actually directory
	if !prj.Git.FS.IsDir() && err == nil {
		prj.errs.AppendString("provided path can not be used as working directory")
	}

	// Check is working directory in GOPATH
	if !prj.Git.FS.InGOPATH() {
		prj.errs.AppendString("provided path is not in GOPATH")
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
