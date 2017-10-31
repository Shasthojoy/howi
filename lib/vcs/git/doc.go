// Copyright 2005-2017 Marko Kungla.
// Use of this source code is governed by a The MIT License
// license that can be found in the LICENSE file.

/*
Package git provides the API to work with git repositories.
It attends to be pure Go implementation of Git and to have no dependencies
like libgit2. Until that library will spawn a shell processes and use the
Git command-line tool to do the work for methods which are not yet in pure go
therefore currently target system has to have git installed.
*/
package git

// TODO(mkungla): Complete the docs
