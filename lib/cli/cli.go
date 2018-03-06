// Copyright 2016 Marko Kungla. All rights reserved.
// Use of this source code is governed by a The Apache-style
// license that can be found in the LICENSE file.

package cli

import (
	"os"
	"time"

	"github.com/okramlabs/howi/lib/cli/flags"
	"github.com/okramlabs/howi/lib/metadata"
	"github.com/okramlabs/howi/pkg/errors"
	"github.com/okramlabs/howi/pkg/log"
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
	// FmtErrRequiredFlag formats error if required flag is missing
	FmtErrRequiredFlag = "%q requires flag %q %q"
	// FmtErrUnknownSubcommand formats error for unknown subcommand request.
	FmtErrUnknownSubcommand = "unknown subcommand %q for command %q"
	// FmtErrTooManyArgs formats error when too many arguments are passed.
	FmtErrTooManyArgs = "too many arguments for command %q which accepts max (%d) args"
	// FmtErrInvalidCommandArgs is returned when invalid args are received by
	// command parser.
	FmtErrInvalidCommandArgs = "invalid arguments passed for (%s).Parse"
	// FmtErrCommandNotProvided when no command is provided calling the application
	FmtErrCommandNotProvided = "no command, see (%s --help) for available commands"
)

// Application for CLI Application instance
type Application struct {
	started     time.Time               // when application was started
	Log         *log.Logger             // logger
	MetaData    *metadata.Basic         // Application MetaData
	Header      Header                  // header if used
	Footer      Footer                  // footer if used
	errs        errors.MultiError       // internal errors
	commands    map[string]Command      // commands
	flags       map[int]flags.Interface // global flags
	flagAliases map[string]int          // global flag aliases
	isLoaded    bool                    // is application loaded
	osArgs      []string                // raw os args from beginning of the execution
	currentCmd  *Command
}

// New constructs new CLI Application Plugin and returns it's instance for
// configuration. Returned Application has basic initialization done and
// all defaults set.
func New(m *metadata.Basic) *Application {
	cli := &Application{
		Log:         log.NewStdout(log.NOTICE),
		MetaData:    m,
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
		cli.Log.SetLogLevel(log.DEBUG)
		cli.Log.LockLevel()
		cli.flag("verbose").Unset()
	}

	// Only lock log level to verbose if no --debug flag was set
	if !cli.flag("debug").Present() && cli.flag("verbose").Present() {
		cli.Log.SetLogLevel(log.INFO)
		cli.Log.LockLevel()
	}

	cli.Log.Debugf("Application:Create - accepting configuration changes debugging(%t)",
		cli.flag("debug").Present())

	// Add internal commands besides help
	cli.AddCommand(cmdAbout())
	return cli
}

// AddCommand to application. Commands and command flags will be verified upon
// application startup and will prevent application to start if command was
// invalid or command introduces any flag shadowing.
func (cli *Application) AddCommand(c Command) {
	cli.Log.Debugf("Application:AddCommand - %q", c.Name())
	// Can only check command name here since nothing stops you to add possible
	// shadow flags after this command was added.
	if _, exists := cli.commands[c.Name()]; exists {
		cli.errs.Add(errors.Newf(FmtErrCommandNameInUse, c.Name()))
		return
	}
	cli.commands[c.Name()] = c
}

// AddFlag to application. Invalid flag will add error to multierror and
// prevents application to start.
func (cli *Application) AddFlag(f flags.Interface) {
	cli.Log.Debugf("Application:AddFlag - %q", f.Name())
	// Verify flag
	cli.errs.Add(f.Verify())

	// add flag if there was no errors
	if cli.errs.Nil() {
		nextFlagID := len(cli.flags) + 1
		// assign the flag and set error if there is flag alias shadowing.
		cli.flags[nextFlagID] = f
		// Check that flag or it's aliases do not shadow other global flag
		for _, alias := range f.GetAliases() {
			if flagID, exists := cli.flagAliases[alias]; exists {
				cli.errs.Appendf(FmtErrFlagShadowing, f.Name(), alias, cli.flags[flagID].Name())
			} else {
				cli.flagAliases[alias] = nextFlagID
			}
		}
	}
}

// Start the application
func (cli *Application) Start() {
	cli.Log.Debug("Application:Start - preparing runtime")
	// Setup internals if not setup already
	if !cli.isLoaded {
		// Check for application configuration and validity of flags and commands
		cli.errs.Add(cli.verifyConfig())

		// parse request flags and arguments
		cli.errs.Add(cli.prepare())

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
		cli.Log.Errorf(FmtErrCommandNotProvided, cli.MetaData.Name())
		cli.exit(2)
	}

	// If debug flag was present. but not as global flag then set the level now
	llvl := cli.Log.GetCurrentLevel()
	if llvl != log.DEBUG && cli.flag("debug").Present() && !cli.flag("debug").IsGlobal() {
		cli.Log.SetLogLevel(log.DEBUG)
	}

	worker := newWorker(cli.Log)
	worker.MetaData = cli.MetaData.JSON()
	worker.args = cli.currentCmd.getArgs()

	// Add flags
	cli.processFlags(worker)

	// Start the application and reset the start time
	now := time.Now()
	cli.Log.Debugf("Application:Start - startup took %f seconds (excluding before function)",
		cli.elapsed().Seconds())
	cli.started = now
	// show header if command has not disabled it
	cli.currentCmd.executeBeforeFn(worker)
	if worker.Config.ShowHeader {
		cli.Header.Print(cli.Log, cli.MetaData.JSON(), cli.elapsed())
	}
	if worker.Phase().status <= StatusSuccess {
		cli.currentCmd.executeDoFn(worker)
	}
	// Do funxtion must exits
	if worker.Phase().status == StatusSuccess {
		cli.currentCmd.executeAfterSuccessFn(worker)
		cli.currentCmd.executeAfterAlwaysFn(worker)
		// show footer if command has not disabled it
		if worker.Config.ShowFooter {
			cli.Footer.Print(cli.Log, cli.MetaData.JSON(), cli.elapsed())
		}
		cli.exit(0)
	}
	// failure
	cli.Log.Debugf(FmtErrPhaseFailed, worker.Phase().Name(), worker.Phase().msg)
	cli.Log.Error(worker.Phase().msg)
	cli.currentCmd.executeAfterFailureFn(worker)
	cli.currentCmd.executeAfterAlwaysFn(worker)
	// restore loglevel
	if llvl != log.DEBUG && cli.flag("debug").Present() && !cli.flag("debug").IsGlobal() {
		cli.Log.SetLogLevel(llvl)
	}
	// show footer if command has not disabled it
	if worker.Config.ShowFooter {
		cli.Footer.Print(cli.Log, cli.MetaData.JSON(), cli.elapsed())
	}
	cli.exit(1)
}

// verifyConfig verifies that configuration is correct
func (cli *Application) verifyConfig() error {
	lenc := len(cli.commands)
	lenf := len(cli.flags)
	cli.Log.Debugf("Application:verifyConfig - %q has total %d command(s)", cli.MetaData.Name(), lenc)
	if (cli.commands == nil || lenc == 0) || (cli.flags == nil || lenf == 0) {
		return errors.New(FmtErrAppWithNoCommandsOrFlags)
	}
	if cli.MetaData.Name() == "" {
		return errors.New(FmtErrAppUnnamed)
	}
	return nil
}

// Exit application
// This is called in the end of the execution and takes care of cleaning up runtime before exiting.
func (cli *Application) exit(code int) {
	os.Exit(code)
}

// Elapsed returns time.Duration since application was started
func (cli *Application) elapsed() time.Duration {
	return time.Now().Sub(cli.started)
}

// Flag looks up flag with given name and returns flags.Interface. If no flag
// was found empty bool flag will be returned. Instead of returning error you
// can check returned flags .IsPresent
func (cli *Application) flag(name string) flags.Interface {
	if id, exists := cli.flagAliases[name]; exists {
		return cli.flags[id]
	}
	return flags.NewBoolFlag(name)
}

// prepare runtime
func (cli *Application) prepare() error {
	// global flags
	for i := 1; i <= len(cli.flags); i++ {
		// ignore already parsed error since it is valid for predefined global flags
		if ok, err := cli.flags[i].Parse(&cli.osArgs); err != nil && !ok {
			return err
		}
	}

	// If we still have global flags left
	if len(cli.osArgs) > 0 && cli.osArgs[0][0] == '-' {
		return errors.Newf(FmtErrUnknownGlobalFlag, cli.osArgs[0])
	}

	// verify configuration of commands
	for _, cmd := range cli.commands {
		if err := cmd.verify(cli.flagAliases); err != nil {
			return err
		}
	}
	// parse requested command
	if len(cli.osArgs) > 0 {
		cmd, exists := cli.commands[cli.osArgs[0]]
		if !exists {
			return errors.Newf(FmtErrUnknownCommand, cli.osArgs[0])
		}
		cli.currentCmd = &cmd
		if err := cli.currentCmd.parse(&cli.osArgs); err != nil {
			return err
		}
	}
	if cli.currentCmd != nil {
		return cli.currentCmd.errs.AsError()
	}
	return nil
}

// checkRuntimeErrors checks if any errors have been added to application
// level multierror if so then calls immediately .Log.Fatal which exits after
// printing the error
func (cli *Application) checkRuntimeErrors() {
	hasErrors := !cli.errs.Nil()
	cli.Log.Debugf("Application:checkRuntimeErrors - has errors (%t)", hasErrors)
	// log errors and exit if present
	if hasErrors {
		elapsed := cli.elapsed()

		cli.Header.Print(cli.Log, cli.MetaData.JSON(), elapsed)
		cli.Log.Error(cli.errs.Error())
		cli.Footer.Print(cli.Log, cli.MetaData.JSON(), elapsed)

		cli.exit(2)
	}
}

// handleBashCompletion handles bash completion calls
func (cli *Application) handleBashCompletion() {
	cli.Log.Debugf("Application:handleBashCompletion - is bash completion call (%t)",
		cli.flag("show-bash-completion").Present())

	if cli.flag("show-bash-completion").Present() {
		// TODO(mkungla): https://github.com/howi-ce/howi/issues/15
		cli.Log.Error("bash completion not implemented")
		os.Exit(0)
	}
}

// handleHelp prints help menu depending on request
func (cli *Application) handleHelp() {
	cli.Log.Debugf("Application:handleHelp - was it help call (%t)",
		cli.flag("help").Present())
	if cli.flag("help").Present() {
		elapsed := cli.elapsed()
		cli.Header.Print(cli.Log, cli.MetaData.JSON(), elapsed)
		if cli.flag("help").IsGlobal() {
			help := HelpGlobal{
				Info:     cli.MetaData.JSON(),
				Commands: cli.commands,
				Flags:    cli.flags,
			}
			help.Print(cli.Log)
		} else {
			help := HelpCommand{
				Info:    cli.MetaData.JSON(),
				Command: *cli.currentCmd,
			}
			help.Print(cli.Log)
		}
		cli.Footer.Print(cli.Log, cli.MetaData.JSON(), elapsed)
		cli.exit(0)
	}
}

// Add flags
func (cli *Application) processFlags(worker *Worker) {
	// add global flags to worker
	for _, flag := range cli.flags {
		worker.attachFlag(flag)
		if flag.IsRequired() && !flag.Present() {
			// show footer if command has not disabled it
			if worker.Config.ShowHeader {
				cli.Header.Print(cli.Log, cli.MetaData.JSON(), cli.elapsed())
			}
			worker.Log.Errorf(FmtErrRequiredFlag, "global", flag.Name(), flag.Usage())
			// show footer if command has not disabled it
			if worker.Config.ShowFooter {
				cli.Footer.Print(cli.Log, cli.MetaData.JSON(), cli.elapsed())
			}
			cli.exit(1)
		}
	}
	// Add flags from current command
	for _, flag := range cli.currentCmd.getFlags() {
		worker.attachFlag(flag)
		// check did we have any required flags missing
		if flag.IsRequired() && !flag.Present() {
			// show footer if command has not disabled it
			if worker.Config.ShowHeader {
				cli.Header.Print(cli.Log, cli.MetaData.JSON(), cli.elapsed())
			}
			worker.Log.Errorf(FmtErrRequiredFlag, cli.currentCmd.Name(), flag.Name(), flag.Usage())
			// show footer if command has not disabled it
			if worker.Config.ShowFooter {
				cli.Footer.Print(cli.Log, cli.MetaData.JSON(), cli.elapsed())
			}
			cli.exit(1)
		}
	}
}

// add builtin flags
func (cli *Application) addInternalFlags() {
	debug := flags.NewBoolFlag("debug")
	debug.SetUsage("enable debug log level. when debug flag is after the command then debugging will be enabled only for that command")
	debug.Parse(&cli.osArgs)
	cli.AddFlag(debug)

	verbose := flags.NewBoolFlag("verbose", "v")
	verbose.SetUsage("enable verbose log level")
	verbose.Parse(&cli.osArgs)
	cli.AddFlag(verbose)

	help := flags.NewBoolFlag("help", "h")
	help.SetUsage("display help or help for the command. [...command --help]")
	help.Parse(&cli.osArgs)
	cli.AddFlag(help)

	bashCompletion := flags.NewBoolFlag("show-bash-completion")
	bashCompletion.Parse(&cli.osArgs)
	bashCompletion.Hide()
	cli.AddFlag(bashCompletion)
}

// NewCommand returns new command constructor.
func NewCommand(name string) Command {
	return Command{name: name}
}
