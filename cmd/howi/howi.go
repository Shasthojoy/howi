package main

import (
	"time"

	"github.com/okramlabs/howi"
)

func main() {
	howi := howi.New()

	// parse .howi.yaml or set defaults
	appMeta := howi.Meta()
	appMeta.SetName("howi")
	appMeta.SetNamespace("okramlabs")
	appMeta.SetTitle("HOWI")
	appMeta.SetShortDesc("The extreme simplicity of HOWICLI makes the building of CLI applications in go super fun and easy. Includes collection of extended Go standard libraries, replacements, helpers and additional packages to transform HOWI API from it's other language bindings into Go.")
	appMeta.SetKeywords("golang-tools", "go", "golang", "golang-library", "howi")
	appMeta.SetCopyRight(2005, "Marko Kungla")
	appMeta.SetVersion("5.0.0-alpha.a.3")
	appMeta.SetURL("https://github.com/howi-ce/howi")
	appMeta.AddContributor("Marko Kungla <marko.kungla@gmail.com>")
	buildDate, err := time.Parse(time.RFC3339, "2018-03-06T03:06:34+02:00")
	if err != nil {
		buildDate = time.Now().UTC()
	}
	appMeta.SetBuildDate(buildDate)

	// Command-line interface
	howicli := howi.CLI()
	howicli.Log.Colors()
	howicli.Log.SetPrimaryColor("yellow")
	// Application header
	howicli.Header.SetTemplate(`
################################################################################
# {{ .Title }}{{ if .CopyRight}}
#  {{ .CopyRight }}{{end}}
# {{if .Version}}
#   Version:    {{ .Version }}{{end}}{{if .BuildDate}}
#   Build date: {{ .BuildDate | funcDate }}{{end}}
################################################################################
`)
	// Application footer
	howicli.Footer.SetTemplate(`
################################################################################
# elapsed: {{ funcElapsed }}
################################################################################`)
	// start the Application
	howi.Start()
}
