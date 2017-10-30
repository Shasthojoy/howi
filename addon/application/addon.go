// Copyright 2005-2017 Marko Kungla.
// Use of this source code is governed by a The MIT License
// license that can be found in the LICENSE file.

package application

import (
	"github.com/howi-ce/howi/addon/application/plugin/cli"
	"github.com/howi-ce/howi/lib/appinfo"
)

// Addon instance
type Addon struct {
	info *appinfo.Metadata
	cli  *cli.Plugin
}

// Info returns application info
func (a *Addon) Info() *appinfo.Metadata {
	return a.info
}

// CLI for the application
func (a *Addon) CLI() *cli.Plugin {
	if a.cli == nil {
		// Should not modify application info on runtime
		a.cli = cli.NewPlugin(*a.info)
	}
	return a.cli
}

// Start the application
func (a *Addon) Start() {
	if a.cli != nil {
		a.cli.Start()
	}
}
