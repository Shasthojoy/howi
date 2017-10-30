// Copyright 2005-2017 Marko Kungla.
// Use of this source code is governed by a The MIT License
// license that can be found in the LICENSE file.

package appinfo

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/blang/semver"
	"github.com/howi-ce/howi/std/net/mail"
)

// Metadata contains general information about the app
type Metadata struct {
	name         string
	title        string
	desc         Description
	contributors []mail.Address
	copyright    Copyright
	url          string
	version      semver.Version
	buildDate    time.Time
}

// SetName of tshe application
func (m *Metadata) SetName(name string) error {
	if len(name) > 36 {
		return errors.New("name is too long max char allowed 36")
	}
	m.name = name
	return nil
}

// Name returns application name
func (m *Metadata) Name() string {
	return m.name
}

// SetTitle of the application used as terminal title or application title
func (m *Metadata) SetTitle(title string) error {
	if len(title) > 72 {
		return errors.New("title is too long max char allowed 72")
	}
	m.title = title
	return nil
}

// SetShortDesc set short description of the application max 160char
func (m *Metadata) SetShortDesc(sdesc string) error {
	if len(sdesc) > 160 {
		return errors.New("description is too long max char allowed 160")
	}
	m.desc.short = sdesc
	return nil
}

// SetLongDesc sets long description of the application used in man, about or help pages.
func (m *Metadata) SetLongDesc(ldesc string) {
	m.desc.long = ldesc
}

// AddContributor to application. argument should be valid RFC 5322 address,
// e.g. "John Doe <john.doe@example.com>"
func (m *Metadata) AddContributor(addr string) error {
	contributor, err := mail.ParseAddress(addr)
	if err == nil {
		m.contributors = append(m.contributors, *contributor)
	}
	return err
}

// SetCopyRightInfo of the application
func (m *Metadata) SetCopyRightInfo(year int, message string) {
	m.copyright.since = year
	m.copyright.msg = message
}

// SetURL for application main site or company site
func (m *Metadata) SetURL(url string) {
	m.url = url
}

// SetVersion of application
func (m *Metadata) SetVersion(version string) error {
	v, err := semver.Make(version)
	if err != nil {
		return err
	}
	m.version = v
	return nil
}

// SetBuildDate of application
func (m *Metadata) SetBuildDate(buildDate time.Time) {
	m.buildDate = buildDate
}

// GetCopyMessage returns copyright message
func (m *Metadata) GetCopyMessage() string {
	ty := time.Now()
	y := strconv.Itoa(m.copyright.since)
	if m.copyright.since < ty.Year() {
		y = fmt.Sprintf("%d - %d", m.copyright.since, ty.Year())
	}
	return fmt.Sprintf("Copyright © %s %s", m.copyright.msg, y)
}

// GetInfo returns application info which can be consumed by templates or output as json
func (m *Metadata) GetInfo() Info {
	var contributors []string
	for _, contributor := range m.contributors {
		contributors = append(contributors, contributor.String())
	}
	info := Info{
		Name:             m.name,
		Title:            m.title,
		ShortDescription: m.desc.short,
		LongDescription:  m.desc.long,
		CopyRight:        m.GetCopyMessage(),
		CopySince:        m.copyright.since,
		CopyMsg:          m.copyright.msg,
		URL:              m.url,
		Version:          m.version.String(),
		BuildDate:        m.buildDate.Format(time.RFC3339),
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
