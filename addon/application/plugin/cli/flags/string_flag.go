package flags

import (
	"strings"

	"github.com/howi-ce/howi/std/hvars"
)

// NewStringFlag returns new string flag. Argument "a" can be any nr of aliases
func NewStringFlag(name string, a ...string) *StringFlag {
	f := &StringFlag{}
	f.name = strings.TrimLeft(name, "-")
	f.aliases = append(f.aliases, f.name)
	for _, alias := range a {
		f.aliases = append(f.aliases, strings.TrimLeft(alias, "-"))
	}
	f.value = hvars.ValueFromString("")
	return f
}

// StringFlag is string flag type with default value ""
type StringFlag struct {
	FlagCommon
}

// Parse the StringFlag
func (f *StringFlag) Parse(args *[]string) (bool, error) {
	return f.parser(args, func(v *hvars.Value) {
		if v.Empty() {
			*v = hvars.ValueFromString("")
		}
	})
}
