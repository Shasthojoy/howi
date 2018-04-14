// Copyright 2018 DIGAVERSE. All rights reserved.
// Use of this source code is governed by a The Apache-style
// license that can be found in the LICENSE file.

package project

// Config of the application
type Config struct {
	LogLevel int    `json:"loglevel,omitempty"`
	Color    string `json:"color,omitempty"`
	InitTerm bool   `json:"initterm,omitempty"`
}
