// Copyright 2005-2017 Marko Kungla.
// Use of this source code is governed by a The MIT License
// license that can be found in the LICENSE file.

package appinfo

// Info can be consumed by header and footer templates.
// It is used by MetaData.GetInfo()
type Info struct {
	Name             string   `json:"name"`
	Title            string   `json:"title"`
	ShortDescription string   `json:"short-description"`
	LongDescription  string   `json:"long-description"`
	CopyRight        string   `json:"copyright"`
	CopySince        int      `json:"copy-since"`
	CopyMsg          string   `json:"copy-msg"`
	URL              string   `json:"url"`
	Version          string   `json:"version"`
	BuildDate        string   `json:"build-date"`
	Contributors     []string `json:"authors"`
}
