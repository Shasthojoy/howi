// Copyright 2016 Marko Kungla. All rights reserved.
// Use of this source code is governed by a The Apache-style
// license that can be found in the LICENSE file.

package cli

import (
	"fmt"
	"strings"

	"github.com/okramlabs/howi/lib/cli/flags"
	"github.com/okramlabs/howi/pkg/errors"
	"github.com/okramlabs/howi/pkg/namespace"
	"github.com/okramlabs/howi/pkg/vars"
)

// Command for application
type Command struct {
	name           string
	hidden         bool
	category       string
	usage          string
	shortDesc      string
	longDesc       string
	errs           errors.MultiError // internal errors
	beforeFn       func(w *Worker)
	doFn           func(w *Worker)
	afterFailureFn func(w *Worker)
	afterSuccessFn func(w *Worker)
	afterAlwaysFn  func(w *Worker)
	subCommands    map[string]Command      // subcommands
	flags          map[int]flags.Interface // command flags
	flagAliases    map[string]int          // command flag aliases
	acceptArgs     int
	args           []vars.Value
	subCmd         *Command // if subcommand was called
	parents        []string
}

// Hide the command from help menu and bash completion.
func (c *Command) Hide() {
	c.hidden = true
}

// ArgsAllowed sets nr of arguments this command accepts. If provided argument
// is not a subcommand and AllowNumArgs is 0 (default) then application fails to
// start with error "invalid subcommand"
func (c *Command) ArgsAllowed(n int) {
	c.acceptArgs = n
}

// AddSubcommand to application which are verified in application startup
func (c *Command) AddSubcommand(cmd Command) {
	if c.subCommands == nil {
		c.subCommands = make(map[string]Command)
	}
	cmd.parents = (append(c.parents, c.name))
	c.subCommands[cmd.name] = cmd
}

// AddFlag adds provided flag to command or subcommand.
// Invalid flag will add error to multierror and prevents application to start.
func (c *Command) AddFlag(f flags.Interface) {
	// Verify flag
	c.errs.Add(f.Verify())
	// Add flag if there was no errors while verifying the flag
	if c.errs.Nil() {
		c.attachFlag(f)
	}
}

// SetCategory sets help category to categorize commands in help output
func (c *Command) SetCategory(category string) {
	c.category = strings.TrimSpace(category)
}

// SetUsage adds extra usage string following the (app-name cmd-name ...)
// auto generated usage string would be something like:
//   `app-name cmd-name [flags] [subcommand] [flags] [args]`
// this text would appear next line
func (c *Command) SetUsage(usage string) {
	c.usage = strings.TrimSpace(usage)
}

// SetUsagef is same as SetUsage, but enables format the usage string
// Arguments are handled in the manner of fmt.Srintf
func (c *Command) SetUsagef(format string, v ...interface{}) {
	c.usage = fmt.Sprintf(format, v...)
}

// SetShortDesc sets commands short description used when describing command within list.
func (c *Command) SetShortDesc(desc string) {
	c.shortDesc = desc
}

// SetLongDesc sets commands long description which is used when displaying commands help
func (c *Command) SetLongDesc(desc string) {
	c.longDesc = desc
}

// Name returns name of the command
func (c *Command) Name() string {
	if c.subCmd != nil {
		return c.subCmd.Name()
	}
	return c.name
}

// Usage returns commands usage string
func (c *Command) Usage() string {
	if c.subCmd != nil {
		return c.subCmd.Usage()
	}
	if c.usage != "" {
		return c.usage
	}
	return ""
}

// ShortDesc returns commands short description
func (c *Command) ShortDesc() string {
	if c.subCmd != nil {
		return c.subCmd.ShortDesc()
	}
	return c.shortDesc
}

// LongDesc returns commands long description
func (c *Command) LongDesc() string {
	if c.subCmd != nil {
		return c.subCmd.LongDesc()
	}
	return c.longDesc
}

// AcceptsFlags returns true if command accepts any flags
func (c *Command) AcceptsFlags() bool {
	if c.subCmd != nil {
		return c.subCmd.AcceptsFlags()
	}
	return c.flags != nil && len(c.flags) > 0
}

// AcceptsArgs returns true if command accepts any arguments
func (c *Command) AcceptsArgs() bool {
	if c.subCmd != nil {
		return c.subCmd.AcceptsArgs()
	}
	return c.acceptArgs > 0
}

// HasSubcommands returns true if command has any subcommands
func (c *Command) HasSubcommands() bool {
	if c.subCmd != nil {
		return c.subCmd.HasSubcommands()
	}
	return c.subCommands != nil && len(c.subCommands) > 0
}

// GetSubcommands returns slice with all subcommands for the command
func (c *Command) GetSubcommands() []Command {
	if c.subCmd != nil {
		return c.subCmd.GetSubcommands()
	}
	var scmd []Command
	for _, cmd := range c.subCommands {
		scmd = append(scmd, cmd)
	}
	return scmd
}

// Before is first function called if it is set.
// It will continue executing worker queue set within this function until first
// failure occurs which is not allowed to continue.
func (c *Command) Before(f func(w *Worker)) {
	c.beforeFn = f
}

// Do should contain main of this command
// This function is called when:
//   - BeforeFunc is not set
//   - BeforeFunc succeeds
//   - BeforeFunc fails but failed tasks have status "allow failure"
func (c *Command) Do(f func(w *Worker)) {
	c.doFn = f
}

// AfterFailure is called when DoFunc fails.
// This function is called when:
//   - DoFunc is not set (this case default AfterFailure function is called)
//   - DoFunc task fails which has no mark "allow failure"
func (c *Command) AfterFailure(f func(w *Worker)) {
	c.afterFailureFn = f
}

// AfterSuccess is called when AfterFailure states that there has been no failures.
// This function is called when:
//   - AfterFailure states that there has been no fatal errors
func (c *Command) AfterSuccess(f func(w *Worker)) {
	c.afterSuccessFn = f
}

// AfterAlways is final function called and is waiting until all tasks whithin
// AfterFailure and/or AfterSuccess functions are done executing.
// If this function if set then it is called always regardless what was the final state of
// any previous phase.
func (c *Command) AfterAlways(f func(w *Worker)) {
	c.afterAlwaysFn = f
}

// Parse command
func (c *Command) parse(args *[]string) error {

	if len(*args) == 0 || (*args)[0] != c.name {
		return errors.Newf(FmtErrInvalidCommandArgs, c.name)
	}

	// remove name of this command
	*args = (*args)[1:]
	// command flags
	for i := 1; i <= len(c.flags); i++ {
		if _, err := c.flags[i].Parse(args); err != nil {
			return err
		}
	}

	// If we still have arg 0 which is not a argument or subcommand
	if len(*args) > 0 && (*args)[0][0] == '-' {
		return errors.Newf(FmtErrUnknownFlag, (*args)[0], c.name)
	}

	for _, arg := range *args {
		// parse subcommand
		if scmd, isSubcommand := c.subCommands[arg]; isSubcommand {
			c.subCmd = &scmd
			return c.subCmd.parse(args)
		}
		// can parse args
		if c.acceptArgs == 0 {
			return errors.Newf(FmtErrUnknownSubcommand, arg, c.name)
		}
		// too many arguments
		if len(c.args) >= c.acceptArgs {
			return errors.Newf(FmtErrTooManyArgs, c.name, c.acceptArgs)
		}
		// add this arg
		c.appendArg(arg)
	}
	return nil
}

// get all already parsed flags for the worker
func (c *Command) getFlags() []flags.Interface {
	var flags []flags.Interface
	for _, flag := range c.flags {
		flags = append(flags, flag)
	}
	if c.subCmd != nil {
		subFlags := c.subCmd.getFlags()
		for _, flag := range subFlags {
			flags = append(flags, flag)
		}
	}
	return flags
}

// appendArg is used by parser to attach provided args to command
func (c *Command) appendArg(arg string) {
	n := len(c.args)
	if n == cap(c.args) {
		args := make([]vars.Value, n, 2*n+1)
		copy(args, c.args)
		c.args = args
	}
	c.args = c.args[0 : n+1]
	c.args[n] = vars.Value(arg)
}

// return slice with names of parent commands.
func (c *Command) getParents() []string {
	if c.subCmd != nil {
		return c.subCmd.getParents()
	}
	return c.parents
}

// getArgs returns args passed for command or subcommand
func (c *Command) getArgs() []vars.Value {
	if c.subCmd != nil {
		return c.subCmd.getArgs()
	}
	return c.args
}

// Verify ranges over command flags and the sub commands
//   - verify that commands are valid and have atleast Do function
//   - verify that subcommand do not shadow flags of any parent command
func (c *Command) verify(reservedFlags map[string]int) error {
	if !c.errs.Nil() {
		return c.errs.AsError()
	}
	// Chck commands name
	if !namespace.IsValid(c.name) {
		return errors.Newf(FmtErrCommandNameInvalid, c.name, namespace.NamespaceMustCompile)
	}
	// must have Do function
	if c.doFn == nil {
		if c.subCommands != nil {
			goto SubCommands
		}
		return errors.Newf(FmtErrCommandMissingDoFn, c.name)
	}
	// Check command flags
	if c.flagAliases != nil && c.flags != nil {
		for flagAlias, flagID := range c.flagAliases {
			if _, isReserved := reservedFlags[flagAlias]; isReserved {
				return errors.Newf(FmtErrCommandFlagShadowing, c.name, c.flags[flagID].Name(), flagAlias)
			}
			// Add it as reserved for subcommands
			reservedFlags[flagAlias] = flagID
		}
	}
SubCommands:
	// Check subcommand flags if any
	if c.subCommands != nil {
		for _, cmd := range c.subCommands {
			err := cmd.verify(reservedFlags)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (c *Command) attachFlag(f flags.Interface) {
	if c.flags == nil {
		c.flags = make(map[int]flags.Interface)
		c.flagAliases = make(map[string]int)
	}
	nextFlagID := len(c.flags) + 1
	c.flags[nextFlagID] = f
	// Chack that flag or it's aliases do not shadow other flag for current command
	for _, alias := range f.GetAliases() {
		if flagID, exists := c.flagAliases[alias]; exists {
			c.errs.Appendf(FmtErrFlagShadowing, f.Name(), alias, c.flags[flagID].Name())
		} else {
			c.flagAliases[alias] = nextFlagID
		}
	}
}

// execute Before function if it was set.
func (c *Command) executeBeforeFn(w *Worker) {
	w.phase = "before"
	if c.subCmd != nil {
		c.subCmd.executeBeforeFn(w)
		return
	}

	if c.beforeFn == nil {
		w.phases[w.phase].status = StatusSkipped
		w.Log.Debug(w.phase, " skipped")
		return
	}
	w.Phase().start()
	c.beforeFn(w)
	// wait
	w.phasewait()
}

// execute Do function.
func (c *Command) executeDoFn(w *Worker) {
	w.phase = "do"
	if c.subCmd != nil {
		c.subCmd.executeDoFn(w)
		return
	}
	w.Phase().start()
	if c.doFn == nil {
		w.Log.Line(c.ShortDesc())
		w.Failf(FmtErrCommandNotProvided, c.Name())
		return
	}
	c.doFn(w)
	// wait
	w.phasewait()
}

// execute AfterFailure function.
func (c *Command) executeAfterFailureFn(w *Worker) {
	w.phase = "after-failure"
	if c.subCmd != nil {
		c.subCmd.executeAfterFailureFn(w)
		return
	}
	if c.afterFailureFn == nil {
		w.phases[w.phase].status = StatusSkipped
		w.Log.Debug(w.phase, " skipped")
		return
	}
	w.Phase().start()
	c.afterFailureFn(w)
	// wait
	w.phasewait()
}

// execute AfterSuccess function.
func (c *Command) executeAfterSuccessFn(w *Worker) {
	w.phase = "after-success"
	if c.subCmd != nil {
		c.subCmd.executeAfterSuccessFn(w)
		return
	}
	if c.afterSuccessFn == nil {
		w.phases[w.phase].status = StatusSkipped
		w.Log.Debug(w.phase, " skipped")
		return
	}
	w.Phase().start()
	c.afterSuccessFn(w)
	// wait
	w.phasewait()
}

// execute AfterAlways function.
func (c *Command) executeAfterAlwaysFn(w *Worker) {
	w.phase = "after-always"
	if c.subCmd != nil {
		c.subCmd.executeAfterAlwaysFn(w)
		return
	}

	if c.afterAlwaysFn == nil {
		w.phases[w.phase].status = StatusSkipped
		w.Log.Debug(w.phase, " skipped")
		return
	}
	w.Phase().start()
	c.afterAlwaysFn(w)
	// wait
	w.phasewait()
}
