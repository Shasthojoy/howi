// Copyright 2016 Marko Kungla. All rights reserved.
// Use of this source code is governed by a The Apache-style
// license that can be found in the LICENSE file.

package howi

import (
	"github.com/digaverse/howi/lib/cli"
	"github.com/digaverse/howi/lib/metadata"
)

// New creates new howi application instance
func New() *HOWI {
	return &HOWI{}
}

// HOWI Application wrapper
type HOWI struct {
	metadata *metadata.Basic
	cli      *cli.Application
}

// Meta [creates] returns metadata pointer
func (h *HOWI) Meta() *metadata.Basic {
	if h.metadata == nil {
		h.metadata = &metadata.Basic{}
	}
	return h.metadata
}

// CLI [creates] returns application command-line interface instance
func (h *HOWI) CLI() *cli.Application {
	h.cli = cli.New(h.metadata)
	return h.cli
}

// Start the application
func (h *HOWI) Start() {
	if h.cli != nil {
		h.cli.Start()
	}
}
