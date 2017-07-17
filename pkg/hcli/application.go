package hcli

import (
	"os"
	"time"

	"github.com/howi-ce/howi/pkg/hcli/clitmpl"
	"github.com/howi-ce/howi/pkg/hcli/flags"
	"github.com/howi-ce/howi/pkg/std/herrors"
	"github.com/howi-ce/howi/pkg/std/hlog"
)

// Application instance
type Application struct {
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
func (a *Application) AddCommand(c Command) {
	a.Log.Debugf("Application:AddCommand - %q", c.Name())
	// Can only check command name here since nothing stops you to add possible
	// shadow flags after this command was added.
	if _, exits := a.commands[c.Name()]; exits {
		a.errs.AppendError(herrors.Newf(FmtErrCommandNameInUse, c.Name()))
		return
	}
	a.commands[c.Name()] = c
}

// AddFlag to application. Invalid flag will add error to multierror and
// prevents application to start.
func (a *Application) AddFlag(f flags.Interface) {
	a.Log.Debugf("Application:AddFlag - %q", f.Name())
	// Verify flag
	a.errs.AppendError(f.Verify())

	// add flag if there was no errors
	if a.errs.Nil() {
		nextFlagID := len(a.flags) + 1
		// assign the flag and set error if there is flag alias shadowing.
		a.flags[nextFlagID] = f
		// Check that flag or it's aliases do not shadow other global flag
		for _, alias := range f.GetAliases() {
			if flagID, exists := a.flagAliases[alias]; exists {
				a.errs.AppendStringf(FmtErrFlagShadowing, f.Name(), alias, a.flags[flagID].Name())
			} else {
				a.flagAliases[alias] = nextFlagID
			}
		}
	}
}

// Start the application
func (a *Application) Start() {
	a.Log.Debug("Application:Start - preparing runtime")
	// Setup internals if not setup already
	if !a.isLoaded {
		// Check for application configuration and validity of flags and commands
		a.errs.AppendError(a.verifyConfig())

		// parse request flags and arguments
		a.errs.AppendError(a.prepare())

		// Will exit if there are any errors adding some command()
		a.checkRuntimeErrors()
	}

	// Check was it bash completion request and respond to it if so.
	// exits with 0 if request was for bash completion
	a.handleBashCompletion()

	// Shall we display default help if so print it and exit with 0
	a.handleHelp()

	// If debug flag was present. but not as global flag then set the level now
	ll := a.Log.GetCurrentLevel()
	if ll != hlog.DEBUG && a.flag("debug").Present() && !a.flag("debug").IsGlobal() {
		a.Log.SetLogLevel(hlog.DEBUG)
	}

	worker := newWorker(a.Log)
	worker.args = a.currentCmd.getArgs()
	// add global flags to worker
	for _, flag := range a.flags {
		worker.attachFlag(flag)
	}
	for _, flag := range a.currentCmd.getFlags() {
		worker.attachFlag(flag)
	}
	a.currentCmd.executeBeforeFn(worker)
	if worker.Phase().status <= PhaseSuccess {
		// Start the application and reset the start time
		now := time.Now()
		a.Log.Debugf("Application:Start - startup took %f seconds (including before function)",
			a.elapsed().Seconds())
		a.started = now
		// show header if command has not disabled it
		if worker.Config.ShowHeader {
			a.Header.Print(a.Log, a.MetaData.GetInfo(), a.elapsed())
		}
		a.currentCmd.executeDoFn(worker)
		// show footer if command has not disabled it
		if worker.Config.ShowFooter {
			a.Footer.Print(a.Log, a.MetaData.GetInfo(), a.elapsed())
		}
	}
	// Do funxtion must exits
	if worker.Phase().status == PhaseSuccess {
		a.currentCmd.executeAfterSuccessFn(worker)
		a.currentCmd.executeAfterAlwaysFn(worker)
		a.exit(0)
	}
	// failure
	a.Log.Errorf(FmtErrPhaseFailed, worker.Phase().Name(), worker.Phase().msg)
	a.currentCmd.executeAfterFailureFn(worker)
	a.currentCmd.executeAfterAlwaysFn(worker)
	// restore loglevel
	if ll != hlog.DEBUG && a.flag("debug").Present() && !a.flag("debug").IsGlobal() {
		a.Log.SetLogLevel(ll)
	}
	a.exit(1)
}

// add builtin flags
func (a *Application) addInternalFlags() {
	debug := flags.NewBoolFlag("debug")
	debug.SetUsage("enable debug log level. when debug flag is after the command then debugging will be enabled only for that command")
	debug.Parse(&a.osArgs)
	a.AddFlag(debug)

	verbose := flags.NewBoolFlag("verbose", "v")
	verbose.SetUsage("enable verbose log level")
	verbose.Parse(&a.osArgs)
	a.AddFlag(verbose)

	help := flags.NewBoolFlag("help", "h")
	help.SetUsage("display help or help for the command. [...command --help]")
	help.Parse(&a.osArgs)
	a.AddFlag(help)

	bashCompletion := flags.NewBoolFlag("show-bash-completion")
	bashCompletion.Parse(&a.osArgs)
	bashCompletion.Hide()
	a.AddFlag(bashCompletion)
}

// verifyConfig verifies that configuration is correct
func (a *Application) verifyConfig() error {
	lenc := len(a.commands)
	lenf := len(a.flags)
	a.Log.Debugf("Application:verifyConfig - %q has total %d command(s)", a.MetaData.name, lenc)
	if (a.commands == nil || lenc == 0) || (a.flags == nil || lenf == 0) {
		return herrors.New(FmtErrAppWithNoCommandsOrFlags)
	}
	if a.MetaData.name == "" {
		return herrors.New(FmtErrAppUnnamed)
	}
	return nil
}

// Exit application
// This is called in the end of the execution and takes care of cleaning up runtime before exiting.
func (a *Application) exit(code int) {
	os.Exit(code)
}

// Elapsed returns time.Duration since application was started
func (a *Application) elapsed() time.Duration {
	return time.Now().Sub(a.started)
}

// Flag looks up flag with given name and returns flags.Interface. If no flag
// was found empty bool flag will be returned. Instead of returning error you
// can check returned flags .IsPresent
func (a *Application) flag(name string) flags.Interface {
	if id, exists := a.flagAliases[name]; exists {
		return a.flags[id]
	}
	return flags.NewBoolFlag(name)
}

// prepare runtime
func (a *Application) prepare() error {
	// global flags
	for i := 1; i <= len(a.flags); i++ {
		// ignore already parsed error since it is valid for predefined global flags
		if ok, err := a.flags[i].Parse(&a.osArgs); err != nil && !ok {
			return err
		}
	}

	// If we still have global flags left
	if len(a.osArgs) > 0 && a.osArgs[0][0] == '-' {
		return herrors.Newf(FmtErrUnknownGlobalFlag, a.osArgs[0])
	}

	// verify configuration of commands
	for _, cmd := range a.commands {
		if err := cmd.verify(a.flagAliases); err != nil {
			return err
		}
	}
	// parse requested command
	if len(a.osArgs) > 0 {
		cmd, exists := a.commands[a.osArgs[0]]
		if !exists {
			return herrors.Newf(FmtErrUnknownCommand, a.osArgs[0])
		}
		a.currentCmd = &cmd
		if err := a.currentCmd.parse(&a.osArgs); err != nil {
			return err
		}
	}
	if a.currentCmd != nil {
		return a.currentCmd.errs.AsError()
	}
	return nil
}

// checkRuntimeErrors checks if any errors have been added to application
// level multierror if so then calls immediately .Log.Fatal which exits after
// printing the error
func (a *Application) checkRuntimeErrors() {
	hasErrors := !a.errs.Nil()
	a.Log.Debugf("Application:checkRuntimeErrors - has errors (%t)", hasErrors)
	// log errors and exit if present
	if hasErrors {
		elapsed := a.elapsed()

		a.Header.Print(a.Log, a.MetaData.GetInfo(), elapsed)
		a.Log.Error(a.errs.Error())
		a.Footer.Print(a.Log, a.MetaData.GetInfo(), elapsed)

		a.exit(2)
	}
}

// handleBashCompletion handles bash completion calls
func (a *Application) handleBashCompletion() {
	a.Log.Debugf("Application:handleBashCompletion - is bash completion call (%t)",
		a.flag("show-bash-completion").Present())

	if a.flag("show-bash-completion").Present() {
		// TODO(mkungla): https://github.com/howi-ce/howi/issues/15
		a.Log.Error("bash completion not implemented")
		os.Exit(0)
	}
}

// handleHelp prints help menu depending on request
func (a *Application) handleHelp() {
	a.Log.Debugf("Application:handleHelp - was it help call (%t)",
		a.flag("help").Present())
	if a.flag("help").Present() {
		elapsed := a.elapsed()
		a.Header.Print(a.Log, a.MetaData.GetInfo(), elapsed)
		if a.flag("help").IsGlobal() {
			help := HelpGlobal{
				Info:     a.MetaData.GetInfo(),
				Commands: a.commands,
				Flags:    a.flags,
			}
			help.Print(a.Log)
		} else {
			help := HelpCommand{
				Info:    a.MetaData.GetInfo(),
				Command: *a.currentCmd,
			}
			help.Print(a.Log)
		}
		a.Footer.Print(a.Log, a.MetaData.GetInfo(), elapsed)
		a.exit(0)
	}
}
