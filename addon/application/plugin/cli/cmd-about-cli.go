// Copyright 2005-2017 Marko Kungla.
// Use of this source code is governed by a The MIT License
// license that can be found in the LICENSE file.

package cli

import (
	"fmt"

	"github.com/howi-ce/howi/addon/application/plugin/cli/flags"
)

func aboutCLI() Command {
	cmd := NewCommand("about-cli")
	cmd.SetShortDesc("Display information about this CLI app")
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
		showVersion, _ := version.Value().ParseBool()
		buildDate, _ := w.Flag("build-date")
		showBuildDate, _ := buildDate.Value().ParseBool()
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
	if show, _ := contributors.Value().ParseBool(); show {
		w.Log.Line("Project Contributors\n")
		for _, contributor := range w.Info.Contributors {
			w.Log.Line(contributor)
		}
		return
	}
	buildDate, _ := w.Flag("build-date")
	if show, _ := buildDate.Value().ParseBool(); show {
		fmt.Println(w.Info.BuildDate)
		return
	}
	version, _ := w.Flag("version")
	if show, _ := version.Value().ParseBool(); show {
		fmt.Println(w.Info.Version)
		return
	}

	w.Log.Line("ABOUT")
	w.Log.Line("------------------------------------------------------------------------")
	if w.Info.LongDescription != "" {
		w.Log.Line(w.Info.LongDescription)
	} else {
		w.Log.Line(w.Info.ShortDescription)
	}
	w.Log.Line("------------------------------------------------------------------------")
	w.Log.Line(tableRow("Version:", w.Info.Version))
	w.Log.Line(tableRow("Build date:", w.Info.BuildDate))
	w.Log.Line(tableRow("Total contributors:", len(w.Info.Contributors)))
	w.Log.Line("------------------------------------------------------------------------")
	w.Log.Line("Project Contributors\n")
	for _, contributor := range w.Info.Contributors {
		w.Log.Line(contributor)
	}
	w.Log.Line("------------------------------------------------------------------------")
	w.Log.Line("for flags printing additional info use --help")

}

func tableRow(key string, val interface{}) string {
	return fmt.Sprintf("%-30s %v", key, val)
}
