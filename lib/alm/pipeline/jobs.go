// Copyright 2005-2017 Marko Kungla.
// Use of this source code is governed by a The MIT License
// license that can be found in the LICENSE file.

package pipeline

import (
	"container/list"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/howi-ce/howi/std/log"

	"github.com/howi-ce/howi/lib/filesystem/path"
)

// Jobs of the pipeline
type Jobs struct {
	Test    Job
	Build   Job
	Deploy  Job
	Always  Job
	Failure Job
	Success Job
}

// AddTestJob add new job
func (j *Jobs) AddTestJob(commonConf ConfigStageJob, jobConf ConfigStageJob) {
	if !jobConf.Present() {
		return
	}
	j.Test = NewJob(commonConf, jobConf)
}

// AddBuildJob add new job
func (j *Jobs) AddBuildJob(commonConf ConfigStageJob, jobConf ConfigStageJob) {
	if !jobConf.Present() {
		return
	}
	j.Build = NewJob(commonConf, jobConf)
}

// AddDeployJob add new job
func (j *Jobs) AddDeployJob(commonConf ConfigStageJob, jobConf ConfigStageJob) {
	if !jobConf.Present() {
		return
	}
	j.Deploy = NewJob(commonConf, jobConf)
}

// AddAlwaysJob add new job
func (j *Jobs) AddAlwaysJob(commonConf ConfigStageJob, jobConf ConfigStageJob) {
	if !jobConf.Present() {
		return
	}
	j.Always = NewJob(commonConf, jobConf)
}

// AddFailureJob add new job
func (j *Jobs) AddFailureJob(commonConf ConfigStageJob, jobConf ConfigStageJob) {
	if !jobConf.Present() {
		return
	}
	j.Failure = NewJob(commonConf, jobConf)
}

// AddSuccessJob add new job
func (j *Jobs) AddSuccessJob(commonConf ConfigStageJob, jobConf ConfigStageJob) {
	if !jobConf.Present() {
		return
	}
	j.Success = NewJob(commonConf, jobConf)
}

// Job is single job
type Job struct {
	AllowedToFail bool
	Commands      *list.List
}

// CanRun report
func (j *Job) CanRun() bool {
	if j.Commands == nil {
		return false
	}
	return j.Commands.Len() > 0
}

// Run the job
func (j *Job) Run(path path.Obj, log *log.Logger) error {
	if j.Commands == nil {
		return errors.New("no commands to run")
	}
	if !path.IsDir() {
		return errors.New("path must be directory")
	}
	err := os.Chdir(path.Abs())
	if err != nil {
		return err
	}
	for c := j.Commands.Front(); c != nil; c = c.Next() {
		log.ColoredLine(c.Value)
		args := strings.Fields(fmt.Sprintf("%s", c.Value))
		cmd := args[0]
		args = args[1:len(args)]
		command := exec.Command(cmd, args...)
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr
		command.Start()

		if status := command.Wait(); status != nil {
			return status
		}
	}
	return nil
}

// PrintCommands prints commands that would be executed.
func (j *Job) PrintCommands() {
	if j.Commands == nil {
		return
	}
	for c := j.Commands.Front(); c != nil; c = c.Next() {
		fmt.Println(c.Value)
	}
}

// NewJob Create new job
func NewJob(commonConf ConfigStageJob, jobConf ConfigStageJob) Job {
	job := Job{
		Commands:      list.New(),
		AllowedToFail: jobConf.AllowFailure || commonConf.AllowFailure,
	}
	for _, cmd := range commonConf.Script {
		job.Commands.PushBack(cmd)
	}
	for _, cmd := range jobConf.Script {
		job.Commands.PushBack(cmd)
	}
	return job
}
