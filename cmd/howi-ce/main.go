// Copyright 2005-2017 Marko Kungla.
// Use of this source code is governed by a The MIT License
// license that can be found in the LICENSE file.

package main

import (
	"time"

	"github.com/howi-ce/howi/addon/application"
	"github.com/howi-ce/howi/cmd/howi-ce/commands/devel"
	"github.com/howi-ce/howi/cmd/howi-ce/internal"
	"github.com/howi-ce/howi/std/log"
)

func main() {
	app := application.NewAddon()

	buildDate, err := time.Parse(time.RFC3339, internal.BuildDate)
	if err != nil {
		buildDate = time.Now().UTC()
	}
	// Set application info
	info := app.Info()
	info.SetName("howi-ce")
	info.SetTitle("HOWI CE 5")
	info.SetShortDesc("Collection of extended Go standard libraries, replacements, helpers and additional packages to transform HOWI API from it's other language bindings into Go")
	// info.SetLongDesc("")
	info.SetCopyRightInfo(2005, "Marko Kungla")
	info.SetURL("https://github.com/howi-ce/howi")
	info.SetVersion(internal.Version)
	info.SetBuildDate(buildDate)
	for _, contributor := range internal.Contributors {
		info.AddContributor(contributor)
	}

	appcli := app.CLI()
	appcli.Log.Colors()

	appcli.Log.SetPrimaryColor("yellow")
	appcli.Log.SetLogLevel(log.NOTICE)

	// Application header
	appcli.Header.SetTemplate(internal.CLIheader)
	// Application footer
	appcli.Footer.SetTemplate(internal.CLIfooter)

	// Attach Commands
	appcli.AddCommand(devel.CMD())

	app.Start()
}
