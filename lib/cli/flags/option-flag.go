// Copyright 2016 Marko Kungla. All rights reserved.
// Use of this source code is governed by a The Apache-style
// license that can be found in the LICENSE file.

package flags

import (
	"strings"

	"github.com/digaverse/howi/pkg/vars"
)

// NewOptionFlag returns new string flag. Argument "opts" is string slice
// of options this flag accepts
func NewOptionFlag(name string, opts []string, a ...string) *OptionFlag {
	f := &OptionFlag{}
	f.name = strings.TrimLeft(name, "-")
	f.opts = make(map[string]bool)
	for _, o := range opts {
		f.opts[o] = true
	}
	for _, alias := range a {
		f.aliases = append(f.aliases, strings.TrimLeft(alias, "-"))
	}
	f.value = vars.Value("")
	return f
}

// OptionFlag is string flag type which can have value of one of the options
type OptionFlag struct {
	opts map[string]bool
	FlagCommon
}

// Parse the OptionFlag
func (f *OptionFlag) Parse(args *[]string) (bool, error) {
	_, err := f.parser(args, func(v *vars.Value) {
		if v.Empty() {
			*v = vars.Value("")
		}
	})
	if _, isSet := f.opts[f.value.String()]; !isSet {
		f.isPresent = false
		return false, err
	}

	return true, nil
}
