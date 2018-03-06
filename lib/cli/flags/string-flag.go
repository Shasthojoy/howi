// Copyright 2016 Marko Kungla. All rights reserved.
// Use of this source code is governed by a The Apache-style
// license that can be found in the LICENSE file.

package flags

import (
	"strings"

	"github.com/okramlabs/howi/pkg/vars"
)

// NewStringFlag returns new string flag. Argument "a" can be any nr of aliases
func NewStringFlag(name string, a ...string) *StringFlag {
	f := &StringFlag{}
	f.name = strings.TrimLeft(name, "-")
	f.aliases = append(f.aliases, f.name)
	for _, alias := range a {
		f.aliases = append(f.aliases, strings.TrimLeft(alias, "-"))
	}
	f.value = vars.Value("")
	return f
}

// StringFlag is string flag type with default value ""
type StringFlag struct {
	FlagCommon
}

// Parse the StringFlag
func (f *StringFlag) Parse(args *[]string) (bool, error) {
	return f.parser(args, func(v *vars.Value) {
		if v.Empty() {
			*v = vars.Value("")
		}
	})
}
