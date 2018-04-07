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
		for _, contributor := range w.MetaData.Contributors {
			w.Log.Line(contributor)
		}
		return
	}
	buildDate, _ := w.Flag("build-date")
	if show, _ := buildDate.Value().Bool(); show {
		fmt.Print(w.MetaData.BuildDate)
		return
	}
	version, _ := w.Flag("version")
	if show, _ := version.Value().Bool(); show {
		fmt.Print(w.MetaData.Version)
		return
	}

	w.Log.Line("ABOUT")
	w.Log.Line("------------------------------------------------------------------------")
	if w.MetaData.LongDescription != "" {
		w.Log.Line(w.MetaData.LongDescription)
	} else {
		w.Log.Line(w.MetaData.ShortDescription)
	}
	w.Log.Line("------------------------------------------------------------------------")
	w.Log.Line(tableRow("Version:", w.MetaData.Version))
	w.Log.Line(tableRow("Build date:", w.MetaData.BuildDate))
	w.Log.Line(tableRow("Total contributors:", len(w.MetaData.Contributors)))
	w.Log.Line("------------------------------------------------------------------------")
	w.Log.Line("Project Contributors\n")
	for _, contributor := range w.MetaData.Contributors {
		w.Log.Line(contributor)
	}
	w.Log.Line("------------------------------------------------------------------------")
	w.Log.Line("for flags printing additional info use --help")

}

func tableRow(key string, val interface{}) string {
	return fmt.Sprintf("%-30s %v", key, val)
}
