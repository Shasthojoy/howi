// Copyright 2018 DIGAVERSE. All rights reserved.
// Use of this source code is governed by a The Apache-style
// license that can be found in the LICENSE file.

package project

import (
	"encoding/json"
	"time"

	"github.com/blang/semver"
	"github.com/digaverse/howi/pkg/emailaddr"
	"github.com/digaverse/howi/pkg/errors"
	"github.com/digaverse/howi/pkg/namespace"
)

// New metadata
func New(config []byte) (*Project, error) {
	prj := &Project{}
	prj.errors = errors.NewMultiError()
	prj.load(config)
	if !prj.errors.Nil() {
		return prj, prj.errors.AsError()
	}
	return prj, nil
}

// Project data
type Project struct {
	errors errors.MultiError
	// data
	ID              int                 `json:"id,omitempty"`
	Name            string              `json:"name,omitempty"`
	Version         semver.Version      `json:"version,omitempty"`
	Namespace       string              `json:"namespace,omitempty"`
	Title           string              `json:"title,omitempty"`
	License         string              `json:"license,omitempty"`
	Author          emailaddr.Address   `json:"author,omitempty"`
	Copyright       Copyright           `json:"copyright,omitempty"`
	BuildDate       time.Time           `json:"builddate,omitempty"`
	Description     string              `json:"description,omitempty"`
	Keywords        []string            `json:"keywords,omitempty"`
	Homepage        string              `json:"homepage,omitempty"`
	Repository      string              `json:"repository,omitempty"`
	Bugs            Bugs                `json:"bugs,omitempty"`
	Contributors    []emailaddr.Address `json:"contributors,omitempty"`
	Scripts         map[string]string   `json:"scripts,omitempty"`
	Config          Config              `json:"config,omitempty"`
	Dependencies    map[string]string   `json:"dependencies,omitempty"`
	DevDependencies map[string]string   `json:"devDependencies,omitempty"`
}

// Errors within project configuration
func (prj *Project) Errors() errors.MultiError {
	return prj.errors
}

func (prj *Project) load(config []byte) {
	if err := json.Unmarshal(config, &prj); err != nil {
		prj.errors.Add(err)
		return
	}

	if len(prj.Name) > 72 {
		prj.errors.Append("name is too long max char allowed 72")
	}
	if !namespace.IsValid(prj.Name) {
		prj.errors.Appendf("invalid name %q, name must only consist a-zA-Z0-9_-", prj.Name)
	}
	if len(prj.Namespace) > 72 {
		prj.errors.Append("namespace is too long max char allowed 72")
	}
	if !namespace.IsValid(prj.Namespace) {
		prj.errors.Appendf("invalid name %q, name must only consist a-zA-Z0-9_-", prj.Namespace)
	}
}
