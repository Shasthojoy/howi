// Copyright 2005-2017 Marko Kungla.
// Use of this source code is governed by a The MIT License
// license that can be found in the LICENSE file.

package devel

import (
	"github.com/howi-ce/howi/addon/application/plugin/cli"
	"github.com/howi-ce/howi/addon/application/plugin/commands/devel/cicd"
	"github.com/howi-ce/howi/lib/filesystem/path"
	"github.com/howi-ce/howi/lib/goprj"
	"github.com/howi-ce/howi/std/vars"
)

// CMD adds subcommands for development and contributing to HOWI CE.
func CMD() cli.Command {
	cmd := cli.NewCommand("devel")
	cmd.SetShortDesc("Subcommands for development and contributing to HOWI CE.")
	cmd.SetCategory("howi-ce")

	cmd.Do(func(w *cli.Worker) {
		w.Fail("See howi-ce devel --help for more info.")
	})
	cmd.AfterFailure(func(w *cli.Worker) {
		w.Log.Warning("Get involved! Check out the Contributing Guide for how to get started.")
		w.Log.Line("https://github.com/howi-ce/howi/blob/master/CONTRIBUTING.md")
	})

	cmd.AddSubcommand(issues())

	// cicd
	// Get go variables
	govars, err := goprj.GetGoVars()
	if err != nil {
		return cmd
	}
	vars := vars.ParseFromStrings(govars)
	gopath, _ := path.New(vars.Getvar("GOPATH").String())
	repopath := gopath.Join("src", "github.com/howi-ce/howi")
	cicdCmd, err := cicd.Command(repopath)
	if err != nil {
		return cmd
	}
	cmd.AddSubcommand(cicdCmd)
	return cmd
}
