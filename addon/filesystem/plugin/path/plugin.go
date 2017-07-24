// Copyright 2005-2017 Marko Kungla. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// license that can be found in the LICENSE file.

package path

import "path/filepath"

// Plugin returns Path for given string, It tries to set absolute representation of path,
// but sets provided string if that fails.
func Plugin(path string) (Path, error) {
	abs, err := filepath.Abs(filepath.FromSlash(path))
	p := Path{}
	if err != nil {
		p.abs = path
	} else {
		p.abs = abs
	}
	// first load
	p.Stat()
	p.base = filepath.Base(abs)
	p.clean = filepath.Clean(abs)
	p.dir = filepath.Dir(abs)
	p.ext = filepath.Ext(abs)
	p.isAbs = filepath.IsAbs(abs)
	return p, err
}
