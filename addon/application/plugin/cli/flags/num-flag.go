// Copyright 2005-2017 Marko Kungla.
// Use of this source code is governed by a The MIT License
// license that can be found in the LICENSE file.

package flags

import (
	"strings"

	"github.com/howi-ce/howi/std/vars"
)

// NewNumFlag returns new numeric flag. Argument "a" can be any nr of aliases
func NewNumFlag(name string, a ...string) *NumFlag {
	f := &NumFlag{}
	f.name = strings.TrimLeft(name, "-")
	f.aliases = append(f.aliases, f.name)
	for _, alias := range a {
		f.aliases = append(f.aliases, strings.TrimLeft(alias, "-"))
	}
	f.value = vars.ValueFromString("")
	return f
}

// NumFlag is numeric flag type with default value 0
type NumFlag struct {
	FlagCommon
}

// Parse the NumFlag
func (f *NumFlag) Parse(args *[]string) (bool, error) {
	return f.parser(args, func(v *vars.Value) {
		if v.Empty() {
			*v = vars.ValueFromString("0")
		}
	})
}
