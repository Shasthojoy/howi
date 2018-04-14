// Copyright 2018 Marko Kungla. All rights reserved.
// Use of this source code is governed by a The Apache-style
// license that can be found in the LICENSE file.

package ui

import (
	"os"
	"time"

	"github.com/digaverse/howi/pkg/errors"
	"github.com/digaverse/howi/pkg/log"
	"github.com/digaverse/howi/pkg/project"
)

const (
	// FmtErrAppUnnamed formats error for unnamed application.
	FmtErrAppUnnamed = "application must have a name"
)

// Application for CLI Application instance
type Application struct {
	started  time.Time         // when application was started
	Log      *log.Logger       // logger
	Project  *project.Project  // Application MetaData
	errs     errors.MultiError // internal errors
	isLoaded bool              // is application loaded
}

// New constructs new CLI Application Plugin and returns it's instance for
// configuration. Returned Application has basic initialization done and
// all defaults set.
func New(project *project.Project) *Application {
	ui := &Application{
		Log:     log.NewStdout(log.NOTICE),
		Project: project,
	}
	ui.Log.TsDisabled()

	// set initial startup time
	ui.started = time.Now()
	ui.Log.TsDisabled()
	return ui
}

// Start the user interface background job
func (ui *Application) Start() {
	ui.Log.Debug("UI:Start - preparing runtime")

	// Setup internals if not setup already
	if !ui.isLoaded {

		// Check for application configuration and validity
		ui.errs.Add(ui.verifyConfig())

		// preapre ui
		ui.errs.Add(ui.prepare())

		// Will exit if there are any errors adding some ui features
		ui.checkRuntimeErrors()
	}
	ui.isLoaded = true

	// Start the application and reset the start time
	now := time.Now()
	ui.Log.Debugf("UI:Start - startup took %f seconds (excluding before function)",
		ui.elapsed().Seconds())
	ui.started = now
}

// Launch the user interface
func (ui *Application) Launch() {
	ui.Log.Debug("UI:Launch called")
}

// Elapsed returns time.Duration since application was started
func (ui *Application) elapsed() time.Duration {
	return time.Now().Sub(ui.started)
}

// verifyConfig verifies that configuration is correct
func (ui *Application) verifyConfig() error {
	ui.Log.Debug("UI:verifyConfig")
	if ui.Project.Name == "" {
		return errors.New(FmtErrAppUnnamed)
	}
	return nil
}

// prepare runtime
func (ui *Application) prepare() error {
	ui.Log.Debug("UI:prepare")
	return nil
}

// checkRuntimeErrors checks if any errors have been added to application
// level multierror if so then calls immediately .Log.Fatal which exits after
// printing the error
func (ui *Application) checkRuntimeErrors() {
	hasErrors := !ui.errs.Nil()
	ui.Log.Debugf("UI:checkRuntimeErrors - has errors (%t)", hasErrors)
	// log errors and exit if present
	if hasErrors {
		elapsed := ui.elapsed()

		ui.Log.Error(ui.errs.Error())
		ui.Log.Infof("elapsed: %s", elapsed.String())

		ui.exit(2)
	}
}

// Exit application
// This is called in the end of the execution and takes care of cleaning up runtime before exiting.
func (ui *Application) exit(code int) {
	os.Exit(code)
}
