// Copyright 2018 DIGAVERSE. All rights reserved.
// Use of this source code is governed by a The Apache-style
// license that can be found in the LICENSE file.

package project

import "github.com/digaverse/howi/pkg/emailaddr"

// Bugs tracker info
type Bugs struct {
	URL   string            `json:"url,omitempty"`
	Email emailaddr.Address `json:"email,omitempty"`
}
