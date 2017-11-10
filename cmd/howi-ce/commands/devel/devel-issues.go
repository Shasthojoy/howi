// Copyright 2005-2017 Marko Kungla.
// Use of this source code is governed by a The MIT License
// license that can be found in the LICENSE file.

package devel

import "github.com/howi-ce/howi/addon/application/plugin/cli"

func issues() cli.Command {
	cmd := cli.NewCommand("issues")
	cmd.SetShortDesc("HOWI CE issue tracker. See howi-ce devel issues --help for more info.")
	cmd.Do(func(w *cli.Worker) {
		w.Log.Error("not implemented")
	})
	return cmd
}
