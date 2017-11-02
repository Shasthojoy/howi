// Copyright 2005-2017 Marko Kungla.
// Use of this source code is governed by a The MIT License
// license that can be found in the LICENSE file.

package pipeline

import (
	"io/ioutil"
	"os"
	"regexp"

	"github.com/howi-ce/howi/lib/filesystem/path"
	"github.com/howi-ce/howi/std/errors"
	yaml "gopkg.in/yaml.v2"
)

// ConfigDevPhase configuration
type ConfigDevPhase struct {
	Only   []string `yaml:"only"`
	Except []string `yaml:"except"`
}

// ShouldRun returns true if this Development Phase is allowing
// execution of the pipeline.
func (d *ConfigDevPhase) ShouldRun(refName string, isTag bool) (bool, error) {
	var run bool
	for _, only := range d.Only {
		if only == "branches" && isTag || only == "tags" && !isTag {
			return false, nil
		}
		m, err := regexp.MatchString(refName, only)
		if err != nil {
			return false, err
		}
		if m {
			run = true
			break
		}
	}
	for _, except := range d.Except {
		if except == "branches" && !isTag || except == "tags" && isTag {
			return false, nil
		}
		m, err := regexp.MatchString(refName, except)
		if err != nil {
			return false, err
		}
		if m {
			run = false
			break
		}
	}
	return run, nil
}

// ConfigStageJob contains individual job
type ConfigStageJob struct {
	AllowFailure bool     `yaml:"allow-failure"`
	Script       []string `yaml:"script"`
}

// Present returns true if job should be used
func (s *ConfigStageJob) Present() bool {
	return len(s.Script) > 0
}

// Config of the pipeline
type Config struct {
	Phases struct {
		Development ConfigDevPhase `yaml:"development"`
		Prereleases ConfigDevPhase `yaml:"prereleases"`
		Staging     ConfigDevPhase `yaml:"staging"`
		Stable      ConfigDevPhase `yaml:"stable"`
	} `yaml:"phases"`
	TestCommon      ConfigStageJob `yaml:"test-common"`
	TestDevelopment ConfigStageJob `yaml:"test-development"`
	TestPrereleases ConfigStageJob `yaml:"test-prereleases"`
	TestStaging     ConfigStageJob `yaml:"test-staging"`
	TestStable      ConfigStageJob `yaml:"test-stable"`

	BuildCommon      ConfigStageJob `yaml:"build-common"`
	BuildDevelopment ConfigStageJob `yaml:"build-development"`
	BuildPrereleases ConfigStageJob `yaml:"build-prereleases"`
	BuildStaging     ConfigStageJob `yaml:"build-staging"`
	BuildStable      ConfigStageJob `yaml:"build-stable"`

	DeployCommon      ConfigStageJob `yaml:"deploy-common"`
	DeployDevelopment ConfigStageJob `yaml:"deploy-development"`
	DeployPrereleases ConfigStageJob `yaml:"deploy-prereleases"`
	DeployStaging     ConfigStageJob `yaml:"deploy-staging"`
	DeployStable      ConfigStageJob `yaml:"deploy-stable"`

	AlwaysCommon      ConfigStageJob `yaml:"always-common"`
	AlwaysDevelopment ConfigStageJob `yaml:"always-development"`
	AlwaysPrereleases ConfigStageJob `yaml:"always-prereleases"`
	AlwaysStaging     ConfigStageJob `yaml:"always-staging"`
	AlwaysStable      ConfigStageJob `yaml:"always-stable"`

	FailureCommon      ConfigStageJob `yaml:"failure-common"`
	FailureDevelopment ConfigStageJob `yaml:"failure-development"`
	FailurePrereleases ConfigStageJob `yaml:"failure-prereleases"`
	FailureStaging     ConfigStageJob `yaml:"failure-staging"`
	FailureStable      ConfigStageJob `yaml:"failure-stable"`

	SuccessCommon      ConfigStageJob `yaml:"success-common"`
	SuccessDevelopment ConfigStageJob `yaml:"success-development"`
	SuccessPrereleases ConfigStageJob `yaml:"success-prereleases"`
	SuccessStaging     ConfigStageJob `yaml:"success-staging"`
	SuccessStable      ConfigStageJob `yaml:"success-stable"`
}

// Load configuration from file
func (c *Config) Load(filepath string) error {
	path, err := path.New(filepath)
	if err != nil {
		return err
	}
	if !path.Exists() {
		return errors.Newf("configuration file missing %s", filepath)
	}

	f, err := os.Open(path.Abs())
	if err != nil {
		return err
	}
	defer f.Close()

	inputBytes, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(inputBytes, &c); err != nil {
		return err
	}
	return nil
}
