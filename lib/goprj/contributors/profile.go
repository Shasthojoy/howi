// Copyright 2005-2017 Marko Kungla.
// Use of this source code is governed by a The MIT License
// license that can be found in the LICENSE file.

package contributors

// ContributorProfile exposes public info about contributor
type ContributorProfile struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	TotalCommits string `json:"total-commits"`
	Additions    string `json:"total-additions"`
	Deletions    string `json:"total-deletions"`
	LastCommit   string `json:"last-commit"`
}
