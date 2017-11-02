// Copyright 2005-2017 Marko Kungla.
// Use of this source code is governed by a The MIT License
// license that can be found in the LICENSE file.

package cicd

import (
	"github.com/howi-ce/howi/addon/application/plugin/cli"
	"github.com/howi-ce/howi/addon/application/plugin/cli/flags"
	"github.com/howi-ce/howi/lib/alm/pipeline"
	"github.com/howi-ce/howi/lib/goprj"
	"github.com/howi-ce/howi/std/errors"
	"github.com/howi-ce/howi/std/log"
)

// Command adds CI/CD commad to your application
// it requires valid path to the project root
func Command(path string) (cli.Command, error) {
	cmd := cli.NewCommand("cicd")

	cmd.SetShortDesc(pipeline.ShortDesc)

	// Flags
	varsFlag := flags.NewStringFlag("vars")
	varsFlag.SetUsage(`prints all variables available within CI/CD pipeline and exits. Use grep to filter (howi-ce devel cicd | grep HOWI_)!`)
	cmd.AddFlag(varsFlag)

	// Set pipline stage
	stageFlag := flags.NewOptionFlag("stage", pipeline.Stages)
	stageFlag.SetUsagef("define CI or CD pipeline stage %q", pipeline.Stages)
	cmd.AddFlag(stageFlag)

	nFlag := flags.NewBoolFlag("n")
	nFlag.SetUsage("prints commands that would be executed.")
	cmd.AddFlag(nFlag)

	cmd.Before(addBeforeFn(varsFlag))

	cmd.Do(addDoFn(path, varsFlag, stageFlag, nFlag))

	return cmd, nil
}

func addBeforeFn(varsFlag *flags.StringFlag) func(w *cli.Worker) {
	return func(w *cli.Worker) {
		if varsFlag.Present() {
			w.Config.ShowFooter = false
			w.Config.ShowHeader = false
		}
	}
}

func addDoFn(path string, varsFlag, stageFlag, nFlag flags.Interface) func(w *cli.Worker) {
	return func(w *cli.Worker) {
		n, err := nFlag.Value().ParseBool()
		if err != nil {
			w.Failf("%q %q failed to parse bool", nFlag.Name(), nFlag.Usage())
			return
		}

		// Load the project
		project, err := projectLoad(path)
		if err != nil {
			w.Fail(err.Error())
			return
		}
		w.Log.Okf("loaded project %q", project.Config.Info.Name)

		// Load pipeline
		cicd, err := project.Pipeline()
		if err != nil {
			w.Fail(err.Error())
			return
		}
		if varsFlag.Present() {
			for k, v := range cicd.Vars {
				w.Log.Linef("%s=%q", k, v.String())
			}
			return
		}

		// stageFlag.Value().String()
		jobs, err := cicd.GetJobs()
		if err != nil {
			w.Fail(err.Error())
			return
		}

		// TEST
		w.Log.Okf("loading phase %q", cicd.GetPhaseName())

		if jobs.Test.CanRun() {
			w.Log.Line("starting test job\n")
			w.Task("test", taskTest(n, jobs.Test, project, w.Log))
			w.Wait()
			w.Log.Ok("test job DONE")
		} else {
			w.Log.Notice("skipping test job")
		}

		// BUILD
		if jobs.Build.CanRun() {
			w.Log.Line()
			w.Log.Line("starting build job\n")
			w.Task("build", taskBuild(n, jobs.Build, project, w.Log))
			w.Wait()
			w.Log.Ok("build job DONE")
		} else {
			w.Log.Notice("skipping build job")
		}

		// DEPLOY
		if jobs.Deploy.CanRun() {
			w.Log.Line("starting deploy job\n")
			w.Task("deploy", taskDeploy(n, jobs.Deploy, project, w.Log))
			w.Wait()
			w.Log.Ok("deploy job DONE")
		} else {
			w.Log.Notice("skipping deploy job")
		}

		// ALWAYS
		if jobs.Always.CanRun() {
			w.Log.Line("starting always job\n")
			taskAlways(n, jobs.Always, project, w)
			w.Log.Ok("always job DONE")
		} else {
			w.Log.Notice("skipping always job")
		}

		// FAILURE
		if jobs.Failure.CanRun() && w.Failed() {
			w.Log.Line("starting failure job\n")
			taskFailure(n, jobs.Failure, project, w)
			w.Log.Ok("failure job DONE")
		} else {
			w.Log.Notice("skipping failure job")
		}

		// SUCCESS
		if jobs.Success.CanRun() && !w.Failed() {
			w.Log.Line("starting success job\n")
			w.Task("success", taskSuccess(n, jobs.Success, project, w.Log))
			w.Wait()
			w.Log.Ok("Success job DONE")
		} else {
			w.Log.Notice("skipping success job")
		}
	}
}

func projectLoad(path string) (prj *goprj.Project, err error) {
	prj, err = goprj.Open(path)
	if err != nil {
		return nil, err
	}
	// We exit if path is not valid project
	if !prj.Exists() {
		return nil, errors.Newf(".howi.yaml is missing on project root %s", path)
	}
	return prj, nil
}

// Test job
func taskTest(n bool, testJob pipeline.Job, prj *goprj.Project, wlog *log.Logger) func(task *cli.Task) {
	return func(task *cli.Task) {
		if n {
			testJob.PrintCommands()
			return
		}
		err := testJob.Run(prj.Path, wlog)
		if err != nil && !testJob.AllowedToFail {
			task.Fail(err.Error())
		}
		wlog.ColoredLine("job test elapsed: ", task.Elapsed())
	}
}

// Build job
func taskBuild(n bool, buildJob pipeline.Job, prj *goprj.Project, wlog *log.Logger) func(task *cli.Task) {
	return func(task *cli.Task) {
		if n {
			buildJob.PrintCommands()
			return
		}
		err := buildJob.Run(prj.Path, wlog)
		if err != nil && !buildJob.AllowedToFail {
			task.Fail(err.Error())
		}
		wlog.ColoredLine("job build elapsed: ", task.Elapsed())
	}
}

// Deploy job
func taskDeploy(n bool, deployJob pipeline.Job, prj *goprj.Project, wlog *log.Logger) func(task *cli.Task) {
	return func(task *cli.Task) {
		if n {
			deployJob.PrintCommands()
			return
		}
		err := deployJob.Run(prj.Path, wlog)
		if err != nil && !deployJob.AllowedToFail {
			task.Fail(err.Error())
		}
		wlog.ColoredLine("job deploy elapsed: ", task.Elapsed())
	}
}

// Always job
func taskAlways(n bool, alwaysJob pipeline.Job, prj *goprj.Project, w *cli.Worker) {
	if n {
		alwaysJob.PrintCommands()
	} else {
		err := alwaysJob.Run(prj.Path, w.Log)
		if err != nil && !alwaysJob.AllowedToFail {
			w.Fail(err.Error())
		}
	}
}

// Failure Job
func taskFailure(n bool, failureJob pipeline.Job, prj *goprj.Project, w *cli.Worker) {
	if n {
		failureJob.PrintCommands()
	} else {
		err := failureJob.Run(prj.Path, w.Log)
		if err != nil && !failureJob.AllowedToFail {
			w.Fail(err.Error())
		}
	}
}

// Success Job
func taskSuccess(n bool, successJob pipeline.Job, prj *goprj.Project, wlog *log.Logger) func(task *cli.Task) {
	return func(task *cli.Task) {
		if n {
			successJob.PrintCommands()
			return
		}
		err := successJob.Run(prj.Path, wlog)
		if err != nil && !successJob.AllowedToFail {
			task.Fail(err.Error())
		}
		wlog.ColoredLine("job success elapsed: ", task.Elapsed())
	}
}
