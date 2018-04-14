// Copyright 2016 Marko Kungla. All rights reserved.
// Use of this source code is governed by a The Apache-style
// license that can be found in the LICENSE file.

package cli

import (
	"fmt"

	"github.com/digaverse/howi/lib/cli/flags"
)

func cmdAbout() Command {
	cmd := NewCommand("about-howi")
	cmd.SetShortDesc("Display information about this application")
	cmd.SetCategory("internal")

	contributors := flags.NewBoolFlag("contributors")
	contributors.SetUsage("print project contributors list")
	cmd.AddFlag(contributors)

	version := flags.NewBoolFlag("version")
	version.SetUsage("print version string")
	cmd.AddFlag(version)

	buildDate := flags.NewBoolFlag("build-date")
	buildDate.SetUsage("print build date")
	cmd.AddFlag(buildDate)

	cmd.Before(func(w *Worker) {
		version, _ := w.Flag("version")
		showVersion, _ := version.Value().Bool()
		buildDate, _ := w.Flag("build-date")
		showBuildDate, _ := buildDate.Value().Bool()
		if showVersion || showBuildDate {
			w.Config.ShowHeader = false
			w.Config.ShowFooter = false
		}
	})

	cmd.Do(aboutCLIdo)

	return cmd
}

func aboutCLIdo(w *Worker) {
	contributors, _ := w.Flag("contributors")
	if show, _ := contributors.Value().Bool(); show {
		w.Log.Line("Project Contributors\n")
		for _, contributor := range w.Project.Contributors {
			w.Log.Line(contributor.String())
		}
		return
	}
	buildDate, _ := w.Flag("build-date")
	if show, _ := buildDate.Value().Bool(); show {
		fmt.Print(w.Project.BuildDate)
		return
	}
	version, _ := w.Flag("version")
	if show, _ := version.Value().Bool(); show {
		fmt.Print(w.Project.Version)
		return
	}

	w.Log.Line("ABOUT")
	w.Log.Line("------------------------------------------------------------------------")
	w.Log.Line(w.Project.Description)
	w.Log.Line("------------------------------------------------------------------------")
	w.Log.Line(tableRow("Version:", w.Project.Version))
	w.Log.Line(tableRow("Build date:", w.Project.BuildDate))
	w.Log.Line(tableRow("Total contributors:", len(w.Project.Contributors)))
	w.Log.Line("------------------------------------------------------------------------")
	w.Log.Line("Project Contributors\n")
	for _, contributor := range w.Project.Contributors {
		w.Log.Line(contributor.String())
	}
	w.Log.Line("------------------------------------------------------------------------")
	w.Log.Line("for flags printing additional info use --help")

}

func tableRow(key string, val interface{}) string {
	return fmt.Sprintf("%-30s %v", key, val)
}
