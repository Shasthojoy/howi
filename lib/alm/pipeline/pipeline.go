// Copyright 2005-2017 Marko Kungla.
// Use of this source code is governed by a The MIT License
// license that can be found in the LICENSE file.

package pipeline

import (
	"github.com/howi-ce/howi/lib/filesystem/path"
	"github.com/howi-ce/howi/std/errors"
	"github.com/howi-ce/howi/std/vars"
)

const (
	// ShortDesc public short description
	ShortDesc = "Continuous integration and Continuous delivery."
)

var (
	// Stages of the CI/CD pipeline
	Stages = []string{"test", "build", "deploy", "always", "failure", "success"}
)

// Library provides pipeline instance for CI/CD
type Library struct {
	phaseName string
	configDir path.Obj
	jobs      *Jobs
	Vars      vars.Collection
	Config    Config
}

// GetJobs return pipline jobs
func (l *Library) GetJobs() (*Jobs, error) {
	if l.jobs == nil {
		err := l.loadPhase()
		if err != nil {
			return nil, err
		}
		l.jobs = &Jobs{}
		switch l.phaseName {
		case "development":
			break
		case "prereleases":
			break
		case "staging":
			l.jobs.AddTestJob(l.Config.TestCommon, l.Config.TestStaging)
			l.jobs.AddBuildJob(l.Config.BuildCommon, l.Config.BuildStaging)
			l.jobs.AddDeployJob(l.Config.DeployCommon, l.Config.DeployStaging)
			l.jobs.AddAlwaysJob(l.Config.AlwaysCommon, l.Config.AlwaysStaging)
			l.jobs.AddFailureJob(l.Config.FailureCommon, l.Config.FailureStaging)
			l.jobs.AddSuccessJob(l.Config.SuccessCommon, l.Config.SuccessStaging)
			break
		case "stable":
			break
		}
	}
	return l.jobs, nil
}

// SetConfigDir where pipeline configuration files can be found
func (l *Library) SetConfigDir(configDir string) (err error) {
	l.configDir, err = path.New(configDir)
	if err == nil && !l.configDir.Exists() || !l.configDir.IsDir() {
		err = errors.New("pipeline configuration directory does not exist")
	}
	if err == nil {
		err = l.Config.Load(l.configDir.Join("pipeline-config.yaml"))
	}
	return
}

// GetPhaseName returns the name of the phase
func (l *Library) GetPhaseName() string {
	return l.phaseName
}

func (l *Library) loadPhase(stage ...string) error {
	isTag, err := l.Vars.Getvar("HOWI_GIT_IS_TAG").ParseBool()
	if err != nil {
		return err
	}
	refName := l.Vars.Getvar("HOWI_GIT_REF_NAME").String()

	devSr, err := l.Config.Phases.Development.ShouldRun(refName, isTag)
	if err != nil {
		return err
	}
	if devSr {
		l.phaseName = "development"
	}

	preSr, err := l.Config.Phases.Prereleases.ShouldRun(refName, isTag)
	if err != nil {
		return err
	}
	if preSr {
		if len(l.phaseName) > 0 {
			return errors.Newf("phases %s and prereleases have conflicting configurations for ref-name %q and is-tag %q",
				l.phaseName, refName, isTag)
		}
		l.phaseName = "prereleases"
	}

	stagingSr, err := l.Config.Phases.Staging.ShouldRun(refName, isTag)
	if err != nil {
		return err
	}
	if stagingSr {
		if len(l.phaseName) > 0 {
			return errors.Newf("phases %s and staging have conflicting configurations for ref-name %q and is-tag %q",
				l.phaseName, refName, isTag)
		}
		l.phaseName = "staging"
	}

	stableSr, err := l.Config.Phases.Stable.ShouldRun(refName, isTag)
	if err != nil {
		return err
	}
	if stableSr {
		if len(l.phaseName) > 0 {
			return errors.Newf("phases %s and stable have conflicting configurations for ref-name %q and is-tag %q",
				l.phaseName, refName, isTag)
		}
		l.phaseName = "stable"
	}
	return nil
}
