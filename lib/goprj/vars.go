// Copyright 2005-2017 Marko Kungla.
// Use of this source code is governed by a The MIT License
// license that can be found in the LICENSE file.

package goprj

import (
	"fmt"
	"strconv"

	"github.com/howi-ce/howi/std/vars"
)

func (p *Project) createVarsPrj() {
	p.Vars["HOWI_PRJ_NAME"] = vars.Value(p.Config.Info.Name)
	p.Vars["HOWI_PRJ_PATH"] = vars.Value(p.Config.filepath.Dir())
	p.Vars["HOWI_PRJ_CONTRIBUTOR_NAME"] = vars.Value(p.Contributor.Name())
	p.Vars["HOWI_PRJ_CONTRIBUTOR_EMAIL"] = vars.Value(p.Contributor.Email())
	p.Vars["HOWI_PRJ_VERSION"] = vars.Value(p.Version.String())
	p.Vars["HOWI_PRJ_VERSION_MAJOR"] = vars.Value(strconv.FormatUint(p.Version.Major, 10))
	p.Vars["HOWI_PRJ_VERSION_MINOR"] = vars.Value(strconv.FormatUint(p.Version.Minor, 10))
	p.Vars["HOWI_PRJ_VERSION_PATCH"] = vars.Value(strconv.FormatUint(p.Version.Patch, 10))
	p.Vars["HOWI_PRJ_VERSION_PRE"] = vars.Value(fmt.Sprint(p.Version.Pre))
	p.Vars["HOWI_PRJ_VERSION_BUILD"] = vars.Value(fmt.Sprint(p.Version.Build))
}

func (p *Project) createVarsGit() {

	// Latest commit hash
	p.Vars["HOWI_GIT_COMMIT_SHA"] = vars.Value(func() string {
		o, _ := p.Git.RevParse("--verify", "HEAD")
		return o.String()
	}())
	// latest shortened commit hash
	// (effective value of the core.abbrev configuration variable)
	p.Vars["HOWI_GIT_COMMIT_SHA_ABBREV"] = vars.Value(func() string {
		o, _ := p.Git.RevParse("--verify", "--short", "HEAD")
		return o.String()
	}())
	// Latest tag commit
	p.Vars["HOWI_GIT_LAST_TAG_COMMIT"] = vars.Value(func() string {
		o, _ := p.Git.RevList("--tags", "--max-count=1")
		return o.String()
	}())
	// Latest tag
	p.Vars["HOWI_GIT_LAST_TAG"] = vars.Value(func() string {
		o, _ := p.Git.Describe("--tags", p.Vars.Getvar("HOWI_GIT_LAST_TAG_COMMIT").String())
		return o.String()
	}())
	p.Vars["HOWI_GIT_NUM_COMMITS_SINCE_LAST_TAG"] = vars.Value(func() string {
		o, _ := p.Git.RevList(fmt.Sprintf("%s..", p.Vars.Getvar("HOWI_GIT_LAST_TAG_COMMIT").String()), "--count")
		return o.String()
	}())
	p.Vars["HOWI_GIT_BRANCH"] = vars.Value(func() string {
		o, _ := p.Git.RevParse("--abbrev-ref", "HEAD")
		return o.String()
	}())
	// The branch or tag name for which project is on
	isTag, _ := p.Git.Describe("--exact-match", "HEAD")
	if len(isTag.String()) == 0 {
		p.Vars["HOWI_GIT_REF_NAME"] = p.Vars["HOWI_GIT_BRANCH"]
		p.Vars["HOWI_GIT_IS_TAG"] = vars.Value("false")
	} else {
		p.Vars["HOWI_GIT_REF_NAME"] = p.Vars["HOWI_GIT_LAST_TAG"]
		p.Vars["HOWI_GIT_IS_TAG"] = vars.Value("true")
	}

	// $HOWI_GIT_REF_NAME lowercased, shortened to 63 bytes,
	// and with everything except 0-9 and a-z replaced with -.
	// No leading / trailing -. Use in URLs, host names and domain names.
	// https://github.com/howi-ce/howi/issues/34
	p.Vars["HOWI_GIT_REF_SLUG"] = vars.Value("")
}
