// Copyright 2005-2017 Marko Kungla.
// Use of this source code is governed by a The MIT License
// license that can be found in the LICENSE file.

package cli

import (
	"os"
	"time"

	"github.com/howi-ce/howi/addon/application/plugin/cli/flags"
	"github.com/howi-ce/howi/std/errors"
)

// add builtin flags
func (cli *Plugin) addInternalFlags() {
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

// verifyConfig verifies that configuration is correct
func (cli *Plugin) verifyConfig() error {
	lenc := len(cli.commands)
	lenf := len(cli.flags)
	cli.Log.Debugf("Application:verifyConfig - has total %d command(s)", cli.MetaData.Name(), lenc)
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
func (cli *Plugin) exit(code int) {
	os.Exit(code)
}

// Elapsed returns time.Duration since application was started
func (cli *Plugin) elapsed() time.Duration {
	return time.Now().Sub(cli.started)
}

// Flag looks up flag with given name and returns flags.Interface. If no flag
// was found empty bool flag will be returned. Instead of returning error you
// can check returned flags .IsPresent
func (cli *Plugin) flag(name string) flags.Interface {
	if id, exists := cli.flagAliases[name]; exists {
		return cli.flags[id]
	}
	return flags.NewBoolFlag(name)
}

// prepare runtime
func (cli *Plugin) prepare() error {
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
func (cli *Plugin) checkRuntimeErrors() {
	hasErrors := !cli.errs.Nil()
	cli.Log.Debugf("Application:checkRuntimeErrors - has errors (%t)", hasErrors)
	// log errors and exit if present
	if hasErrors {
		elapsed := cli.elapsed()

		cli.Header.Print(cli.Log, cli.MetaData.GetInfo(), elapsed)
		cli.Log.Error(cli.errs.Error())
		cli.Footer.Print(cli.Log, cli.MetaData.GetInfo(), elapsed)

		cli.exit(2)
	}
}

// handleBashCompletion handles bash completion calls
func (cli *Plugin) handleBashCompletion() {
	cli.Log.Debugf("Application:handleBashCompletion - is bash completion call (%t)",
		cli.flag("show-bash-completion").Present())

	if cli.flag("show-bash-completion").Present() {
		// TODO(mkungla): https://github.com/howi-ce/howi/issues/15
		cli.Log.Error("bash completion not implemented")
		os.Exit(0)
	}
}

// handleHelp prints help menu depending on request
func (cli *Plugin) handleHelp() {
	cli.Log.Debugf("Application:handleHelp - was it help call (%t)",
		cli.flag("help").Present())
	if cli.flag("help").Present() {
		elapsed := cli.elapsed()
		cli.Header.Print(cli.Log, cli.MetaData.GetInfo(), elapsed)
		if cli.flag("help").IsGlobal() {
			help := HelpGlobal{
				Info:     cli.MetaData.GetInfo(),
				Commands: cli.commands,
				Flags:    cli.flags,
			}
			help.Print(cli.Log)
		} else {
			help := HelpCommand{
				Info:    cli.MetaData.GetInfo(),
				Command: *cli.currentCmd,
			}
			help.Print(cli.Log)
		}
		cli.Footer.Print(cli.Log, cli.MetaData.GetInfo(), elapsed)
		cli.exit(0)
	}
}
