// Copyright 2005-2017 Marko Kungla.
// Use of this source code is governed by a The MIT License
// license that can be found in the LICENSE file.

package internal

var (
	// Version number
	Version = "5.0.0-alpha1+96.89b773c"
	// BuildDate of the application
	BuildDate = "2017-07-29T21:16:34+02:00"
	// Contributors of the application
	Contributors = []string{
		"Marko Kungla <marko.kungla@gmail.com>",
	}
	// CLIheader template
	CLIheader = `
################################################################################
# {{ .Title }}{{ if .CopyRight}}
#  {{ .CopyRight }}{{end}}
# {{if .Version}}
#   Version:    {{ .Version }}{{end}}{{if .BuildDate}}
#   Build date: {{ .BuildDate | funcDate }}{{end}}
################################################################################
`
	// CLIfooter template
	CLIfooter = `
################################################################################
# elapsed: {{ funcElapsed }}
################################################################################`
)
