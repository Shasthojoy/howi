// Copyright 2012 Marko Kungla. All rights reserved.
// Use of this source code is governed by a The Apache-style
// license that can be found in the LICENSE file.

package metadata

// JSON can be consumed by header and footer templates.
// It is used by MetaData.JSON()
type JSON struct {
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
