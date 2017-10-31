// Copyright 2005-2017 Marko Kungla.
// Use of this source code is governed by a The MIT License
// license that can be found in the LICENSE file.

package flags

import (
	"fmt"
	"strings"

	"github.com/howi-ce/howi/std/errors"
	"github.com/howi-ce/howi/std/vars"
)

// Interface for the flags
type Interface interface {
	// Parse value for the flag from given string. It returns true if flag was found in
	// provided args string and false if not.
	// error is returned when flag was set but had invalid value.
	Parse(*[]string) (bool, error)
	// Get primary name for the flag. Usually that is long option
	Name() string
	// Usage returns a usage description for that flag
	Usage() string
	// HelpName returns string for help menu
	HelpName() string
	// HelpAliases returns string for help menu
	HelpAliases() string
	// IsHidden reports whether to show that flag in help menu or not.
	IsHidden() bool
	// IsGlobal reports whether this flag was global and was set before any command or arg
	IsGlobal() bool
	// Pos returns flags position after command. Case of global since app name
	// min value 1 which means first global flag or first flag after command
	Pos() int
	// Return flag aliases
	GetAliases() []string
	// Verify the flag
	Verify() error
	// Unset unsets the value for the flag if it was parsed, handy for cases where
	// one flag cancels another like --debug cancels --verbose
	Unset()
	// Present reports whether flag was set in commandline
	Present() bool
	// Value returns vars.Value for given flag
	Value() vars.Value
	// Required sets this flag as required
	Required()
	// IsRequired returns true if this flag is required
	IsRequired() bool
}

// FlagCommon shares private fields and some function with flags
type FlagCommon struct {
	// name of this flag
	name string
	// aliases for this flag
	aliases []string
	// hide from help menu
	hidden bool
	// global is set to true if value was parsed before any command or arg occurred
	global bool
	// position in os args how many commands where before that flag
	pos int
	// usage string
	usage string
	// isPresent enables to mock removal and .Unset the flag it reports whether flag was "present"
	isPresent bool
	// value for this flag
	value vars.Value
	// is this flag required
	required bool
}

// Name returns primary name for the flag usually that is long option
func (f *FlagCommon) Name() string {
	return f.name
}

// Usage returns a usage description for that flag
func (f *FlagCommon) Usage() string {
	return f.usage
}

// SetUsage sets flag description
func (f *FlagCommon) SetUsage(desc string) {
	f.usage = desc
}

// HelpName returns string for help menu
func (f *FlagCommon) HelpName() string {
	if len(f.name) == 1 {
		return fmt.Sprintf("-%s", f.name)
	}
	return fmt.Sprintf("--%s", f.name)
}

// HelpAliases returns string for help menu
func (f *FlagCommon) HelpAliases() string {
	if len(f.aliases) == 0 {
		return ""
	}
	var aliases []string
	for _, a := range f.aliases {
		if a == f.name {
			continue
		}
		if len(a) < 2 {
			aliases = append(aliases, fmt.Sprintf("-%s", a))
			continue
		}
		aliases = append(aliases, fmt.Sprintf("--%s", a))
	}
	return strings.Join(aliases, ",")
}

// IsHidden reports whether flag should be visible in help menu
func (f *FlagCommon) IsHidden() bool {
	return f.hidden
}

// Hide flag from help menu
func (f *FlagCommon) Hide() {
	f.hidden = true
}

// IsGlobal reports whether this flag was global and was set before any command or arg
func (f *FlagCommon) IsGlobal() bool {
	return f.global
}

// Pos returns flags position after command and case of global since app name
// min value is 1 which means first global flag or first flag after command
func (f *FlagCommon) Pos() int {
	return f.pos
}

// GetAliases Returns all aliases for the flag together with primary "name"
func (f *FlagCommon) GetAliases() []string {
	return f.aliases
}

// Verify the flag
func (f *FlagCommon) Verify() error {
	if f.name == "" {
		return errors.Newf("flag name %q is not valid", f.name)
	}
	return nil
}

// Unset the value
func (f *FlagCommon) Unset() {
	f.value = vars.Value("")
	f.isPresent = false
}

// Present reports whether flag was set in commandline
func (f *FlagCommon) Present() bool {
	return f.isPresent
}

// Value returns vars.Value for this flag
func (f *FlagCommon) Value() vars.Value {
	return f.value
}

// Required setts this flag as required
func (f *FlagCommon) Required() {
	f.required = true
}

// IsRequired returns true if this flag is required
func (f *FlagCommon) IsRequired() bool {
	return f.required
}

// Parse value for the flag from given string. It returns true if flag has been parsed
// and errro if flag has been already parsed.
func (f *FlagCommon) parser(args *[]string, read func(*vars.Value)) (bool, error) {

	if f.isPresent {
		return f.isPresent, errors.Newf("flag %q is already parsed", f.name)
	}

	for i, arg := range *args {
		if arg[0] != '-' {
			f.pos++
			continue
		}
		cur := strings.TrimLeft(arg, "-")
		flag, value := vars.ParseFromString(cur)
		if flag == f.name {
			f.isPresent = true
			read(&value)
			f.value = value
			*args = append((*args)[:i], (*args)[i+1:]...)
			goto checkIsItGlobal
		}
		// if we got so far lets search alias then
		for _, alias := range f.aliases {
			if flag == alias {
				f.isPresent = true
				read(&value)
				f.value = value
				*args = append((*args)[:i], (*args)[i+1:]...)
				goto checkIsItGlobal
			}
		}
	}
checkIsItGlobal:
	// was it global
	if !f.value.Empty() && f.pos == 0 {
		f.global = true
	}
	return f.isPresent, nil
}
