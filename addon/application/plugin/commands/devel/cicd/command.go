// Copyright 2005-2017 Marko Kungla.
// Use of this source code is governed by a The MIT License
// license that can be found in the LICENSE file.

package cicd

import (
	"github.com/howi-ce/howi/addon/application/plugin/cli"
	"github.com/howi-ce/howi/addon/application/plugin/cli/flags"
	"github.com/howi-ce/howi/lib/alm/pipeline"
	"github.com/howi-ce/howi/lib/goprj"
)

// Command adds CI/CD commad to your application
// it requires valid path to the project root
func Command(path string) (cli.Command, error) {
	cmd := cli.NewCommand("cicd")

	cmd.SetShortDesc(pipeline.ShortDesc)

	// Flags
	nFlag := flags.NewBoolFlag("n")
	nFlag.SetUsage("prints commands that would be executed.")
	cmd.AddFlag(nFlag)

	varsFlag := flags.NewStringFlag("vars")
	varsFlag.SetUsage(`prints all variables available within CI/CD pipeline and exits. Use grep to filter (howi-ce devel cicd | grep HOWI_)!`)
	cmd.AddFlag(varsFlag)

	// Set pipline stage
	stageFlag := flags.NewOptionFlag("stage", pipeline.Stages)
	stageFlag.SetUsagef("define CI or CD pipeline stage %q", pipeline.Stages)
	cmd.AddFlag(stageFlag)

	project, err := goprj.Open(path)
	if err != nil {
		return cmd, err
	}
	cmd.Before(func(w *cli.Worker) {
		if varsFlag.Present() {
			w.Config.ShowFooter = false
			w.Config.ShowHeader = false
		}
	})

	cmd.Do(func(w *cli.Worker) {
		// We exit if path is not valid project
		if !project.Exists() {
			w.Log.Errorf(".howi.yaml is missing on project root %s", path)
			w.Fail("Can not execute CI / CD commands")
			return
		}

		diplayOnly, err := nFlag.Value().ParseBool()
		if err != nil {
			w.Failf("%q %q failed to parse bool", nFlag.Name(), nFlag.Usage())
			return
		}

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

		if diplayOnly {
			w.Log.Warning("Only printing commands that would be executed.")
		}
		w.Log.Okf("loaded project %q", project.Config.Info.Name)

		// stageFlag.Value().String()
		jobs, err := cicd.GetJobs()
		if err != nil {
			w.Fail(err.Error())
			return
		}

		// TEST
		w.Log.Okf("loading phase %q", cicd.GetPhaseName())
		if jobs.Test.CanRun() {
			w.Log.Line()
			w.Log.Ok("starting test job")
			w.Task("test", func(task *cli.Task) {
				if diplayOnly {
					jobs.Test.PrintCommands()
					return
				}
				err := jobs.Test.Run(project.Path, w.Log)
				if err != nil && !jobs.Test.AllowedToFail {
					task.Fail(err.Error())
				}

				w.Log.ColoredLine("job test elapsed: ", task.Elapsed())
			})
			w.Wait()
			w.Log.Ok("test job DONE")
		} else {
			w.Log.Notice("skipping test job")
		}

		// BUILD
		if jobs.Build.CanRun() {
			w.Log.Line()
			w.Log.Ok("starting build job")
			w.Task("build", func(task *cli.Task) {
				if diplayOnly {
					jobs.Build.PrintCommands()
					return
				}
				err := jobs.Build.Run(project.Path, w.Log)
				if err != nil && !jobs.Build.AllowedToFail {
					task.Fail(err.Error())
				}
				w.Log.ColoredLine("job build elapsed: ", task.Elapsed())
			})
			w.Wait()
			w.Log.Ok("build job DONE")
		} else {
			w.Log.Notice("skipping build job")
		}

		// DEPLOY
		if jobs.Deploy.CanRun() {
			w.Log.Line()
			w.Log.Ok("starting deploy job")
			w.Task("deploy", func(task *cli.Task) {
				if diplayOnly {
					jobs.Deploy.PrintCommands()
					return
				}
				err := jobs.Deploy.Run(project.Path, w.Log)
				if err != nil && !jobs.Deploy.AllowedToFail {
					task.Fail(err.Error())
				}
				w.Log.ColoredLine("job deploy elapsed: ", task.Elapsed())
			})
			w.Wait()
			w.Log.Ok("deploy job DONE")
		} else {
			w.Log.Notice("skipping deploy job")
		}

		// ALWAYS
		if jobs.Always.CanRun() {
			w.Log.Line()
			w.Log.Ok("starting always job")
			if diplayOnly {
				jobs.Always.PrintCommands()
			} else {
				err := jobs.Always.Run(project.Path, w.Log)
				if err != nil && !jobs.Always.AllowedToFail {
					w.Fail(err.Error())
				}
			}
			w.Log.Ok("always job DONE")
		} else {
			w.Log.Notice("skipping always job")
		}

		// FAILURE
		if jobs.Failure.CanRun() && w.Failed() {
			w.Log.Line()
			w.Log.Ok("starting failure job")
			if diplayOnly {
				jobs.Failure.PrintCommands()
			} else {
				err := jobs.Failure.Run(project.Path, w.Log)
				if err != nil && !jobs.Failure.AllowedToFail {
					w.Fail(err.Error())
				}
			}
			w.Log.Ok("failure job DONE")
		} else {
			w.Log.Notice("skipping failure job")
		}

		// SUCCESS
		if jobs.Success.CanRun() && !w.Failed() {
			w.Log.Line()
			w.Log.Ok("starting success job")
			w.Task("success", func(task *cli.Task) {
				if diplayOnly {
					jobs.Success.PrintCommands()
					return
				}
				err := jobs.Success.Run(project.Path, w.Log)
				if err != nil && !jobs.Success.AllowedToFail {
					task.Fail(err.Error())
				}
				w.Log.ColoredLine("job success elapsed: ", task.Elapsed())
			})
			w.Wait()
			w.Log.Ok("success job DONE")
		} else {
			w.Log.Notice("skipping success job")
		}
	})

	return cmd, nil
}
