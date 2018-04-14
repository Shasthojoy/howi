// Copyright 2016 Marko Kungla. All rights reserved.
// Use of this source code is governed by a The Apache-style
// license that can be found in the LICENSE file.

package howi

import (
	"github.com/digaverse/howi/lib/cli"
	"github.com/digaverse/howi/lib/ui"
	"github.com/digaverse/howi/pkg/project"
)

// New creates new howi application instance
func New(project *project.Project) (h *HOWI, err error) {
	h = &HOWI{Project: project}
	return h, nil
}

// HOWI Application wrapper
type HOWI struct {
	Project *project.Project
	cli     *cli.Application
	ui      *ui.Application
}

// CLI [creates] returns application command-line interface instance
//
// This instance enables you to build all your CLI functionality for your app
func (h *HOWI) CLI() *cli.Application {
	if h.cli == nil {
		h.cli = cli.New(h.Project)
	}
	return h.cli
}

// UI [creates] returns application user interface instance
//
// This instance enables you to build all your user interface functionality for your app
func (h *HOWI) UI() *ui.Application {
	if h.ui == nil {
		h.ui = ui.New(h.Project)
		// update logger
		if h.cli != nil {
			h.ui.Log = h.cli.Log
		}
	}
	return h.ui
}

// Start the application
func (h *HOWI) Start() {
	if h.ui != nil {
		h.ui.Start()
	}
	if h.cli != nil {
		h.cli.Start()
	}
}
