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
	"github.com/okramlabs/howi/pkg/emailaddr"
	"github.com/okramlabs/howi/pkg/namespace"
)

// Basic application metadata
type Basic struct {
	title        string
	name         string
	namespace    string
	desc         Description
	contributors []emailaddr.Address
	keywords     []string
	copyright    Copyright
	url          string
	version      semver.Version
	buildDate    time.Time
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
	if !namespace.IsValid(name) {
		return errors.New("name must only consist a-zA-Z0-9_-")
	}
	b.name = name
	return nil
}

// Name returns application name
func (b *Basic) Name() string {
	return b.name
}

// SetCopyRightInfo of the application
func (b *Basic) SetCopyRightInfo(year int, message string) {
	b.copyright.since = year
	b.copyright.msg = message
}

// SetURL for application main site or company site
func (b *Basic) SetURL(url string) {
	b.url = url
}

// SetNamespace of the application
func (b *Basic) SetNamespace(ns string) error {
	if len(ns) > 72 {
		return errors.New("namespace is too long max char allowed 72")
	}
	if !namespace.IsValid(ns) {
		return errors.New("namespace must only consist a-zA-Z0-9_-")
	}
	b.namespace = ns
	return nil
}

// AddContributor to application. argument should be valid RFC 5322 address,
// e.g. "John Doe <john.doe@example.com>"
func (b *Basic) AddContributor(addr string) error {
	contributor, err := emailaddr.ParseAddress(addr)
	if err == nil {
		b.contributors = append(b.contributors, *contributor)
	}
	return err
}

// Namespace returns application namespace
func (b *Basic) Namespace() string {
	return b.namespace
}

// SetShortDesc set short description of the application max 160char
func (b *Basic) SetShortDesc(sdesc string) error {
	if len(sdesc) > 160 {
		return errors.New("description is too long max char allowed 160")
	}
	b.desc.short = sdesc
	return nil
}

// SetLongDesc sets long description of the application used in man, about or help pages.
func (b *Basic) SetLongDesc(ldesc string) {
	b.desc.long = ldesc
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

// JSON returns application info which can be consumed by templates or output as json
func (b *Basic) JSON() JSON {
	var contributors []string
	for _, contributor := range b.contributors {
		contributors = append(contributors, contributor.String())
	}
	info := JSON{
		Name:             b.name,
		Title:            b.title,
		ShortDescription: b.desc.short,
		LongDescription:  b.desc.long,
		CopyRight:        b.GetCopyMessage(),
		CopySince:        b.copyright.since,
		CopyMsg:          b.copyright.msg,
		URL:              b.url,
		Version:          b.version.String(),
		BuildDate:        b.buildDate.Format(time.RFC3339),
		Contributors:     contributors,
	}
	return info
}

// Description of the application
type Description struct {
	short string
	long  string
}

// Copyright info of the app
type Copyright struct {
	since int
	msg   string
}
