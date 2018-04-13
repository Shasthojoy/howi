// Copyright 2016 Marko Kungla. All rights reserved.
// Use of this source code is governed by a The Apache-style
// license that can be found in the LICENSE file.

package howi

import (
	"github.com/digaverse/howi/lib/cli"
	"github.com/digaverse/howi/lib/metadata"
)

// New creates new howi application instance
func New(name string) (h *HOWI, err error) {
	h = &HOWI{}
	h.name = name
	h.meta, err = metadata.New(h.name)
	if err != nil {
		return nil, err
	}
	return h, nil
}

// HOWI Application wrapper
type HOWI struct {
	name string
	cli  *cli.Application
	meta *metadata.Basic
}

// CLI [creates] returns application command-line interface instance
//
// This instance enables you to build all your CLI functionality for your app
func (h *HOWI) CLI() *cli.Application {
	if h.cli == nil {
		h.cli = cli.New(h.meta)
	}
	return h.cli
}

// Meta for application
func (h *HOWI) Meta() *metadata.Basic {
	return h.meta
}

// Start the application
func (h *HOWI) Start() {
	if h.cli != nil {
		h.cli.Start()
	}
}
