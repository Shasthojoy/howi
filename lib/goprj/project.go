// Copyright 2005-2017 Marko Kungla.
// Use of this source code is governed by a The MIT License
// license that can be found in the LICENSE file.

package goprj

import (
	"github.com/blang/semver"
	"github.com/howi-ce/howi/lib/alm/pipeline"
	"github.com/howi-ce/howi/lib/filesystem/path"
	"github.com/howi-ce/howi/lib/goprj/contributors"
	"github.com/howi-ce/howi/lib/vcs/git"
	"github.com/howi-ce/howi/std/errors"
	"github.com/howi-ce/howi/std/vars"
)

// Project holds information about the golang project
type Project struct {
	errs        errors.MultiError // internal errors
	exists      bool
	pipeline    *pipeline.Library
	Contributor *contributors.Contributor // Current Contributor
	Git         *git.Git                  // Git instance for project
	Config      Config
	Vars        vars.Collection
	Path        path.Obj
	Version     semver.Version
}

// Exists check whether or not .howi/project.yaml exists at projects path
func (p *Project) Exists() bool {
	conf, _ := p.Git.FS.LoadPath(".howi", "project.yaml")
	p.exists = conf.Exists()
	p.Config.filepath = conf
	p.Path, _ = p.Git.FS.LoadPath("/")
	return p.exists
}

// Pipeline of the project
func (p *Project) Pipeline() (*pipeline.Library, error) {
	if p.pipeline == nil && p.exists {
		p.pipeline = pipeline.New()
		err := p.pipeline.SetConfigDir(p.Path.Join(".howi", "pipeline"))
		if err != nil {
			return nil, err
		}
		p.pipeline.Vars = p.Vars
	}
	if p.pipeline == nil {
		return nil, errors.New("failed to load project pipeline")
	}
	return p.pipeline, nil
}
