package flags

import (
	"strings"

	"github.com/howi-ce/howi/pkg/vars"
)

// NewBoolFlag returns new bool flag. Argument "a" can be any nr of aliases
func NewBoolFlag(name string, a ...string) *BoolFlag {
	f := &BoolFlag{}
	f.name = strings.TrimLeft(name, "-")
	f.aliases = append(f.aliases, f.name)
	for _, alias := range a {
		f.aliases = append(f.aliases, strings.TrimLeft(alias, "-"))
	}
	f.value = vars.ValueFromString("false")
	return f
}

// BoolFlag is boolean flag type with default value "false"
type BoolFlag struct {
	FlagCommon
}

// Parse the BoolFlag
func (f *BoolFlag) Parse(args *[]string) (bool, error) {
	return f.parser(args, func(v *vars.Value) {
		if v.Empty() {
			*v = vars.ValueFromString("true")
		}
	})
}
