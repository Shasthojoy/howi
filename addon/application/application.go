// Copyright 2005-2017 Marko Kungla.
// Use of this source code is governed by a The MIT License
// license that can be found in the LICENSE file.

package application

import "github.com/howi-ce/howi/lib/app"

// NewAddon provides Application addon instance
func NewAddon() *Addon {
	addon := &Addon{}
	addon.info = &app.Metadata{}
	return addon
}
