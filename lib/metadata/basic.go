// Copyright 2012 Marko Kungla. All rights reserved.
// Use of this source code is governed by a The Apache-style
// license that can be found in the LICENSE file.

package metadata

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/blang/semver"
	"github.com/okramlabs/howi/pkg/strings"
)

// Basic application metadata
type Basic struct {
	title     string
	name      string
	namespace string
	desc      string
	keywords  []string
	copyright Copyright
	version   semver.Version
	buildDate time.Time
}

// Copyright info of the app
type Copyright struct {
	since int
	msg   string
}

// SetTitle of the application used as terminal title or application title
func (b *Basic) SetTitle(title string) error {
	if len(title) > 72 {
		return errors.New("title is too long max char allowed 72")
	}
	b.title = title
	return nil
}

// Title returns application title
func (b *Basic) Title() string {
	return b.title
}

// SetName of the application used as command name
func (b *Basic) SetName(name string) error {
	if len(name) > 72 {
		return errors.New("name is too long max char allowed 72")
	}
	if !strings.IsNamespace(name) {
		return errors.New("name must only consist a-zA-Z0-9_-")
	}
	b.name = name
	return nil
}

// Name returns application name
func (b *Basic) Name() string {
	return b.name
}

// SetNamespace of the application
func (b *Basic) SetNamespace(namespace string) error {
	if len(namespace) > 72 {
		return errors.New("namespace is too long max char allowed 72")
	}
	if !strings.IsNamespace(namespace) {
		return errors.New("namespace must only consist a-zA-Z0-9_-")
	}
	b.namespace = namespace
	return nil
}

// Namespace returns application namespace
func (b *Basic) Namespace() string {
	return b.namespace
}

// SetDesc set short description of the application max 160char
func (b *Basic) SetDesc(desc string) error {
	if len(desc) > 160 {
		return errors.New("description is too long max char allowed 160")
	}
	b.desc = desc
	return nil
}

// Desc returns application description
func (b *Basic) Desc() string {
	return b.desc
}

// SetKeywords for application
func (b *Basic) SetKeywords(keywords ...string) {
	b.keywords = append([]string{}, keywords...)
}

// Keywords returns application keywords
func (b *Basic) Keywords() []string {
	return b.keywords
}

// SetCopyRight of the application
func (b *Basic) SetCopyRight(year int, message string) {
	b.copyright.since = year
	b.copyright.msg = message
}

// GetCopyMessage returns copyright message
func (b *Basic) GetCopyMessage() string {
	ty := time.Now()
	y := strconv.Itoa(b.copyright.since)
	if b.copyright.since < ty.Year() {
		y = fmt.Sprintf("%d - %d", b.copyright.since, ty.Year())
	}
	return fmt.Sprintf("Copyright © %s %s", b.copyright.msg, y)
}

// SetVersion of application
func (b *Basic) SetVersion(version string) error {
	v, err := semver.Make(version)
	if err != nil {
		return err
	}
	b.version = v
	return nil
}

// GetVersion returns application version
func (b *Basic) GetVersion() semver.Version {
	return b.version
}

// SetBuildDate of application
func (b *Basic) SetBuildDate(buildDate time.Time) {
	b.buildDate = buildDate
}

// GetBuildDate returns raw build date
func (b *Basic) GetBuildDate() time.Time {
	return b.buildDate
}

// GetBuildDateRFC3339 return build date as RFC3339 formated string
func (b *Basic) GetBuildDateRFC3339() string {
	return b.buildDate.Format(time.RFC3339)
}
