// Copyright 2005-2017 Marko Kungla. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// license that can be found in the LICENSE file.

package cli

import (
	"os"
	"time"

	"github.com/howi-ce/howi/addon/application/plugin/cli/clitmpl"
	"github.com/howi-ce/howi/addon/application/plugin/cli/flags"
	"github.com/howi-ce/howi/std/herrors"
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

// NewPlugin constructs new CLI Application Plugin and returns it's instance for
// configuration. Returned Application has basic initialization done and
// all defaults set.
func NewPlugin() *Plugin {
	cli := &Plugin{
		Log:         hlog.NewStdout(hlog.NOTICE),
		commands:    make(map[string]Command),
		flags:       make(map[int]flags.Interface),
		flagAliases: make(map[string]int),
		osArgs:      os.Args[1:],
	}
	// set initial startup time
	cli.started = time.Now()
	cli.Log.TsDisabled()

	cli.addInternalFlags()

	// Set log level to debug and lock the log level, but only if --debug
	// flag was found before any command. If --debug flag was found later
	// then we want to set debugging later for that command only.
	if cli.flag("debug").IsGlobal() && cli.flag("debug").Present() {
		cli.Log.SetLogLevel(hlog.DEBUG)
		cli.Log.LockLevel()
		cli.flag("verbose").Unset()
	}

	// Only lock log level to verbose if no --debug flag was set
	if !cli.flag("debug").Present() && cli.flag("verbose").Present() {
		cli.Log.LockLevel()
	}

	cli.Log.Debugf("Application:Create - accepting configuration changes debugging(%t)",
		cli.flag("debug").Present())
	return cli
}

// Plugin for CLI Application instance
type Plugin struct {
	started     time.Time               // when application was started
	Log         *hlog.Logger            // logger
	MetaData    MetaData                // Application MetaData
	Header      clitmpl.Header          // header if used
	Footer      clitmpl.Footer          // footer if used
	errs        herrors.MultiError      // internal errors
	commands    map[string]Command      // commands
	flags       map[int]flags.Interface // global flags
	flagAliases map[string]int          // global flag aliases
	isLoaded    bool                    // is application loaded
	osArgs      []string                // raw os args from beginning of the execution
	currentCmd  *Command
}

// AddCommand to application. Commands and command flags will be verified upon
// application startup and will prevent application to start if command was
// invalid or command introduces any flag shadowing.
func (cli *Plugin) AddCommand(c Command) {
	cli.Log.Debugf("Application:AddCommand - %q", c.Name())
	// Can only check command name here since nothing stops you to add possible
	// shadow flags after this command was added.
	if _, exists := cli.commands[c.Name()]; exists {
		cli.errs.AppendError(herrors.Newf(FmtErrCommandNameInUse, c.Name()))
		return
	}
	cli.commands[c.Name()] = c
}

// AddFlag to application. Invalid flag will add error to multierror and
// prevents application to start.
func (cli *Plugin) AddFlag(f flags.Interface) {
	cli.Log.Debugf("Application:AddFlag - %q", f.Name())
	// Verify flag
	cli.errs.AppendError(f.Verify())

	// add flag if there was no errors
	if cli.errs.Nil() {
		nextFlagID := len(cli.flags) + 1
		// assign the flag and set error if there is flag alias shadowing.
		cli.flags[nextFlagID] = f
		// Check that flag or it's aliases do not shadow other global flag
		for _, alias := range f.GetAliases() {
			if flagID, exists := cli.flagAliases[alias]; exists {
				cli.errs.AppendStringf(FmtErrFlagShadowing, f.Name(), alias, cli.flags[flagID].Name())
			} else {
				cli.flagAliases[alias] = nextFlagID
			}
		}
	}
}

// Start the application
func (cli *Plugin) Start() {
	cli.Log.Debug("Application:Start - preparing runtime")
	// Setup internals if not setup already
	if !cli.isLoaded {
		// Check for application configuration and validity of flags and commands
		cli.errs.AppendError(cli.verifyConfig())

		// parse request flags and arguments
		cli.errs.AppendError(cli.prepare())

		// Will exit if there are any errors adding some command()
		cli.checkRuntimeErrors()
	}
	cli.isLoaded = true

	// Check was it bash completion request and respond to it if so.
	// exits with 0 if request was for bash completion
	cli.handleBashCompletion()

	// Shall we display default help if so print it and exit with 0
	cli.handleHelp()

	if cli.currentCmd == nil {
		cli.Log.Errorf(FmtErrCommandNotProvided, cli.MetaData.name)
		cli.exit(2)
	}
	// If debug flag was present. but not as global flag then set the level now
	llvl := cli.Log.GetCurrentLevel()
	if llvl != hlog.DEBUG && cli.flag("debug").Present() && !cli.flag("debug").IsGlobal() {
		cli.Log.SetLogLevel(hlog.DEBUG)
	}

	worker := newWorker(cli.Log)
	worker.Info = cli.MetaData.GetInfo()
	worker.args = cli.currentCmd.getArgs()
	// add global flags to worker
	for _, flag := range cli.flags {
		worker.attachFlag(flag)
	}
	for _, flag := range cli.currentCmd.getFlags() {
		worker.attachFlag(flag)
	}

	// Start the application and reset the start time
	now := time.Now()
	cli.Log.Debugf("Application:Start - startup took %f seconds (excluding before function)",
		cli.elapsed().Seconds())
	cli.started = now
	// show header if command has not disabled it
	if worker.Config.ShowHeader {
		cli.Header.Print(cli.Log, cli.MetaData.GetInfo(), cli.elapsed())
	}
	cli.currentCmd.executeBeforeFn(worker)
	if worker.Phase().status <= StatusSuccess {
		cli.currentCmd.executeDoFn(worker)
	}
	// Do funxtion must exits
	if worker.Phase().status == StatusSuccess {
		cli.currentCmd.executeAfterSuccessFn(worker)
		cli.currentCmd.executeAfterAlwaysFn(worker)
		// show footer if command has not disabled it
		if worker.Config.ShowFooter {
			cli.Footer.Print(cli.Log, cli.MetaData.GetInfo(), cli.elapsed())
		}
		cli.exit(0)
	}
	// failure
	cli.Log.Errorf(FmtErrPhaseFailed, worker.Phase().Name(), worker.Phase().msg)
	cli.currentCmd.executeAfterFailureFn(worker)
	cli.currentCmd.executeAfterAlwaysFn(worker)
	// restore loglevel
	if llvl != hlog.DEBUG && cli.flag("debug").Present() && !cli.flag("debug").IsGlobal() {
		cli.Log.SetLogLevel(llvl)
	}
	// show footer if command has not disabled it
	if worker.Config.ShowFooter {
		cli.Footer.Print(cli.Log, cli.MetaData.GetInfo(), cli.elapsed())
	}
	cli.exit(1)
}
