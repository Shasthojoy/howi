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

// CLI Application instance
type CLI struct {
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
func (app *CLI) AddCommand(c Command) {
	app.Log.Debugf("Application:AddCommand - %q", c.Name())
	// Can only check command name here since nothing stops you to add possible
	// shadow flags after this command was added.
	if _, exists := app.commands[c.Name()]; exists {
		app.errs.AppendError(herrors.Newf(FmtErrCommandNameInUse, c.Name()))
		return
	}
	app.commands[c.Name()] = c
}

// AddFlag to application. Invalid flag will add error to multierror and
// prevents application to start.
func (app *CLI) AddFlag(f flags.Interface) {
	app.Log.Debugf("Application:AddFlag - %q", f.Name())
	// Verify flag
	app.errs.AppendError(f.Verify())

	// add flag if there was no errors
	if app.errs.Nil() {
		nextFlagID := len(app.flags) + 1
		// assign the flag and set error if there is flag alias shadowing.
		app.flags[nextFlagID] = f
		// Check that flag or it's aliases do not shadow other global flag
		for _, alias := range f.GetAliases() {
			if flagID, exists := app.flagAliases[alias]; exists {
				app.errs.AppendStringf(FmtErrFlagShadowing, f.Name(), alias, app.flags[flagID].Name())
			} else {
				app.flagAliases[alias] = nextFlagID
			}
		}
	}
}

// Start the application
func (app *CLI) Start() {
	app.Log.Debug("Application:Start - preparing runtime")
	// Setup internals if not setup already
	if !app.isLoaded {
		// Check for application configuration and validity of flags and commands
		app.errs.AppendError(app.verifyConfig())

		// parse request flags and arguments
		app.errs.AppendError(app.prepare())

		// Will exit if there are any errors adding some command()
		app.checkRuntimeErrors()
	}
	app.isLoaded = true

	// Check was it bash completion request and respond to it if so.
	// exits with 0 if request was for bash completion
	app.handleBashCompletion()

	// Shall we display default help if so print it and exit with 0
	app.handleHelp()

	if app.currentCmd == nil {
		app.Log.Errorf(FmtErrCommandNotProvided, app.MetaData.name)
		app.exit(2)
	}
	// If debug flag was present. but not as global flag then set the level now
	llvl := app.Log.GetCurrentLevel()
	if llvl != hlog.DEBUG && app.flag("debug").Present() && !app.flag("debug").IsGlobal() {
		app.Log.SetLogLevel(hlog.DEBUG)
	}

	worker := newWorker(app.Log)
	worker.Info = app.MetaData.GetInfo()
	worker.args = app.currentCmd.getArgs()
	// add global flags to worker
	for _, flag := range app.flags {
		worker.attachFlag(flag)
	}
	for _, flag := range app.currentCmd.getFlags() {
		worker.attachFlag(flag)
	}

	// Start the application and reset the start time
	now := time.Now()
	app.Log.Debugf("Application:Start - startup took %f seconds (excluding before function)",
		app.elapsed().Seconds())
	app.started = now
	// show header if command has not disabled it
	if worker.Config.ShowHeader {
		app.Header.Print(app.Log, app.MetaData.GetInfo(), app.elapsed())
	}
	app.currentCmd.executeBeforeFn(worker)
	if worker.Phase().status <= StatusSuccess {
		app.currentCmd.executeDoFn(worker)
	}
	// Do funxtion must exits
	if worker.Phase().status == StatusSuccess {
		app.currentCmd.executeAfterSuccessFn(worker)
		app.currentCmd.executeAfterAlwaysFn(worker)
		// show footer if command has not disabled it
		if worker.Config.ShowFooter {
			app.Footer.Print(app.Log, app.MetaData.GetInfo(), app.elapsed())
		}
		app.exit(0)
	}
	// failure
	app.Log.Errorf(FmtErrPhaseFailed, worker.Phase().Name(), worker.Phase().msg)
	app.currentCmd.executeAfterFailureFn(worker)
	app.currentCmd.executeAfterAlwaysFn(worker)
	// restore loglevel
	if llvl != hlog.DEBUG && app.flag("debug").Present() && !app.flag("debug").IsGlobal() {
		app.Log.SetLogLevel(llvl)
	}
	// show footer if command has not disabled it
	if worker.Config.ShowFooter {
		app.Footer.Print(app.Log, app.MetaData.GetInfo(), app.elapsed())
	}
	app.exit(1)
}

// add builtin flags
func (app *CLI) addInternalFlags() {
	debug := flags.NewBoolFlag("debug")
	debug.SetUsage("enable debug log level. when debug flag is after the command then debugging will be enabled only for that command")
	debug.Parse(&app.osArgs)
	app.AddFlag(debug)

	verbose := flags.NewBoolFlag("verbose", "v")
	verbose.SetUsage("enable verbose log level")
	verbose.Parse(&app.osArgs)
	app.AddFlag(verbose)

	help := flags.NewBoolFlag("help", "h")
	help.SetUsage("display help or help for the command. [...command --help]")
	help.Parse(&app.osArgs)
	app.AddFlag(help)

	bashCompletion := flags.NewBoolFlag("show-bash-completion")
	bashCompletion.Parse(&app.osArgs)
	bashCompletion.Hide()
	app.AddFlag(bashCompletion)
}

// verifyConfig verifies that configuration is correct
func (app *CLI) verifyConfig() error {
	lenc := len(app.commands)
	lenf := len(app.flags)
	app.Log.Debugf("Application:verifyConfig - %q has total %d command(s)", app.MetaData.name, lenc)
	if (app.commands == nil || lenc == 0) || (app.flags == nil || lenf == 0) {
		return herrors.New(FmtErrAppWithNoCommandsOrFlags)
	}
	if app.MetaData.name == "" {
		return herrors.New(FmtErrAppUnnamed)
	}
	return nil
}

// Exit application
// This is called in the end of the execution and takes care of cleaning up runtime before exiting.
func (app *CLI) exit(code int) {
	os.Exit(code)
}

// Elapsed returns time.Duration since application was started
func (app *CLI) elapsed() time.Duration {
	return time.Now().Sub(app.started)
}

// Flag looks up flag with given name and returns flags.Interface. If no flag
// was found empty bool flag will be returned. Instead of returning error you
// can check returned flags .IsPresent
func (app *CLI) flag(name string) flags.Interface {
	if id, exists := app.flagAliases[name]; exists {
		return app.flags[id]
	}
	return flags.NewBoolFlag(name)
}

// prepare runtime
func (app *CLI) prepare() error {
	// global flags
	for i := 1; i <= len(app.flags); i++ {
		// ignore already parsed error since it is valid for predefined global flags
		if ok, err := app.flags[i].Parse(&app.osArgs); err != nil && !ok {
			return err
		}
	}

	// If we still have global flags left
	if len(app.osArgs) > 0 && app.osArgs[0][0] == '-' {
		return herrors.Newf(FmtErrUnknownGlobalFlag, app.osArgs[0])
	}

	// verify configuration of commands
	for _, cmd := range app.commands {
		if err := cmd.verify(app.flagAliases); err != nil {
			return err
		}
	}
	// parse requested command
	if len(app.osArgs) > 0 {
		cmd, exists := app.commands[app.osArgs[0]]
		if !exists {
			return herrors.Newf(FmtErrUnknownCommand, app.osArgs[0])
		}
		app.currentCmd = &cmd
		if err := app.currentCmd.parse(&app.osArgs); err != nil {
			return err
		}
	}
	if app.currentCmd != nil {
		return app.currentCmd.errs.AsError()
	}
	return nil
}

// checkRuntimeErrors checks if any errors have been added to application
// level multierror if so then calls immediately .Log.Fatal which exits after
// printing the error
func (app *CLI) checkRuntimeErrors() {
	hasErrors := !app.errs.Nil()
	app.Log.Debugf("Application:checkRuntimeErrors - has errors (%t)", hasErrors)
	// log errors and exit if present
	if hasErrors {
		elapsed := app.elapsed()

		app.Header.Print(app.Log, app.MetaData.GetInfo(), elapsed)
		app.Log.Error(app.errs.Error())
		app.Footer.Print(app.Log, app.MetaData.GetInfo(), elapsed)

		app.exit(2)
	}
}

// handleBashCompletion handles bash completion calls
func (app *CLI) handleBashCompletion() {
	app.Log.Debugf("Application:handleBashCompletion - is bash completion call (%t)",
		app.flag("show-bash-completion").Present())

	if app.flag("show-bash-completion").Present() {
		// TODO(mkungla): https://github.com/howi-ce/howi/issues/15
		app.Log.Error("bash completion not implemented")
		os.Exit(0)
	}
}

// handleHelp prints help menu depending on request
func (app *CLI) handleHelp() {
	app.Log.Debugf("Application:handleHelp - was it help call (%t)",
		app.flag("help").Present())
	if app.flag("help").Present() {
		elapsed := app.elapsed()
		app.Header.Print(app.Log, app.MetaData.GetInfo(), elapsed)
		if app.flag("help").IsGlobal() {
			help := HelpGlobal{
				Info:     app.MetaData.GetInfo(),
				Commands: app.commands,
				Flags:    app.flags,
			}
			help.Print(app.Log)
		} else {
			help := HelpCommand{
				Info:    app.MetaData.GetInfo(),
				Command: *app.currentCmd,
			}
			help.Print(app.Log)
		}
		app.Footer.Print(app.Log, app.MetaData.GetInfo(), elapsed)
		app.exit(0)
	}
}
