// Copyright 2005-2017 Marko Kungla. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// license that can be found in the LICENSE file.

package cli

import (
	"os"
	"time"

	"github.com/howi-ce/howi/addon/application/plugin/cli/flags"
	"github.com/howi-ce/howi/std/hlog"
)

const (
	// StatusPending marks that phase is pending (800)
	StatusPending = iota + 800
	// StatusSkipped marks that phase was skipped (801)
	StatusSkipped
	// StatusSuccess marks that phase was successful (802)
	StatusSuccess
	// StatusRunning marks that phase is running (803)
	StatusRunning
	// StatusFailed marks that phase failed (804)
	StatusFailed
	// FmtErrFlagShadowing formats shadowed flag error.
	FmtErrFlagShadowing = "flag(%s) alias %q shadows existing flag for %q"
	// FmtErrCommandFlagShadowing formats command shaddowed flag error.
	FmtErrCommandFlagShadowing = "command (%s) flag %q alias %q shadows existing flag"
	// FmtErrCommandNameInvalid formats invalid command name error.
	FmtErrCommandNameInvalid = "command name %q is invalid - must match following regex %v"
	// FmtErrCommandNameInUse formats command name in use error.
	FmtErrCommandNameInUse = "command %q is already in use"
	// FmtErrCommandMissingDoFn formats command missing Do function error.
	FmtErrCommandMissingDoFn = "command (%s) must have DoFn"
	// FmtErrPhaseFailed formats phase failure error.
	FmtErrPhaseFailed = "phase: %q failed (%s)"
	// FmtErrAppWithNoCommandsOrFlags formats error when application is started
	// without any commands or flags configured.
	FmtErrAppWithNoCommandsOrFlags = "application has no flags or commands"
	// FmtErrAppUnnamed formats error for unnamed application.
	FmtErrAppUnnamed = "application must have a name"
	// FmtErrUnknownGlobalFlag formats invalid global flag error.
	FmtErrUnknownGlobalFlag = "unknown global flag %q"
	// FmtErrUnknownCommand formats unknown command error.
	FmtErrUnknownCommand = "unknown command %q"
	// FmtErrUnknownFlag formats error for any request looking non existing flag.
	FmtErrUnknownFlag = "unknown flag %q for command %q"
	// FmtErrUnknownSubcommand formats error for unknown subcommand request.
	FmtErrUnknownSubcommand = "unknown subcommand %q for command %q"
	// FmtErrTooManyArgs formats error when too many arguments are passed.
	FmtErrTooManyArgs = "too many arguments for command %q which accepts max (%d) args"
	// FmtErrInvalidCommandArgs is returned when invalid args are recieved by
	// command parser.
	FmtErrInvalidCommandArgs = "invalid arguments passed for (%s).Parse"
	// FmtErrCommandNotProvided when no command is provided calling the application
	FmtErrCommandNotProvided = "no command, see (%s --help) for available commands"
)

// Plugin constructs new CLI Application and returns it's instance for
// configuration. Returned Application has basic initialization done and
// all defaults set.
func Plugin() *CLI {
	c := &CLI{
		Log:         hlog.NewStdout(hlog.NOTICE),
		commands:    make(map[string]Command),
		flags:       make(map[int]flags.Interface),
		flagAliases: make(map[string]int),
		osArgs:      os.Args[1:],
	}
	// set initial startup time
	c.started = time.Now()
	c.Log.TsDisabled()

	c.addInternalFlags()

	// Set log level to debug and lock the log level, but only if --debug
	// flag was found before any command. If --debug flag was found later
	// then we want to set debugging later for that command only.
	if c.flag("debug").IsGlobal() && c.flag("debug").Present() {
		c.Log.SetLogLevel(hlog.DEBUG)
		c.Log.LockLevel()
		c.flag("verbose").Unset()
	}

	// Only lock log level to verbose if no --debug flag was set
	if !c.flag("debug").Present() && c.flag("verbose").Present() {
		c.Log.LockLevel()
	}

	c.Log.Debugf("Application:Create - accepting configuration changes debugging(%t)",
		c.flag("debug").Present())
	return c
}

// NewCommand returns new command constructor.
func NewCommand(name string) Command {
	return Command{name: name}
}
