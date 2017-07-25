package flags

import (
	"strings"

	"github.com/howi-ce/howi/std/hvars"
)

// NewNumFlag returns new numeric flag. Argument "a" can be any nr of aliases
func NewNumFlag(name string, a ...string) *NumFlag {
	f := &NumFlag{}
	f.name = strings.TrimLeft(name, "-")
	f.aliases = append(f.aliases, f.name)
	for _, alias := range a {
		f.aliases = append(f.aliases, strings.TrimLeft(alias, "-"))
	}
	f.value = hvars.ValueFromString("")
	return f
}

// NumFlag is numeric flag type with default value 0
type NumFlag struct {
	FlagCommon
}

// Parse the NumFlag
func (f *NumFlag) Parse(args *[]string) (bool, error) {
	return f.parser(args, func(v *hvars.Value) {
		if v.Empty() {
			*v = hvars.ValueFromString("0")
		}
	})
}
