// Copyright 2005-2017 Marko Kungla.
// Use of this source code is governed by a The MIT License
// license that can be found in the LICENSE file.

package goprj

import (
	"errors"
	"io/ioutil"
	"os"

	"github.com/howi-ce/howi/lib/filesystem/path"
	"gopkg.in/yaml.v2"
)

// Config of project
type Config struct {
	filepath path.Obj
	Info     struct {
		Name string `yaml:"name"`
	} `yaml:"info"`
}

// Reload configuration file
func (c *Config) Reload() error {
	if c.filepath.Base() != "project.yaml" {
		return errors.New("project config filepath missing")
	}
	f, err := os.Open(c.filepath.Abs())
	if err != nil {
		return err
	}
	defer f.Close()

	inputBytes, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(inputBytes, &c)
}

// Save configuration to file at repo root `.howi/project.yaml`
func (c *Config) Save() error {
	contents, err := yaml.Marshal(&c)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(c.filepath.Abs(), []byte(contents), 0644)
}
