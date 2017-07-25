package cli

import (
	"fmt"
	"strconv"
	"time"

	"github.com/howi-ce/howi/std/herrors"
	"github.com/howi-ce/howi/std/hnet/hmail"
)

// MetaData contains application metadata
type MetaData struct {
	name         string
	title        string
	sdesc        string
	ldesc        string
	contributors []hmail.Address
	copySince    int // copy year
	copyMsg      string
	url          string
	version      string
	buildDate    string
}

// SetName of the application
func (m *MetaData) SetName(name string) {
	m.name = name
}

// SetTitle of the application used as terminal title or application title
func (m *MetaData) SetTitle(title string) {
	m.title = title
}

// SetShortDesc set short description of the application max 160char
func (m *MetaData) SetShortDesc(sdesc string) (errDescToLong error) {
	if len(sdesc) > 160 {
		errDescToLong = herrors.New("description is to long max char allowed 160")
		return
	}
	m.sdesc = sdesc
	return
}

// SetLongDesc sets long description of the application used in man, about or help pages.
func (m *MetaData) SetLongDesc(ldesc string) {
	m.ldesc = ldesc
}

// AddContributor to application. argument should be valid RFC 5322 address,
// e.g. "John Doe <john.doe@example.com>"
func (m *MetaData) AddContributor(addr string) error {
	contributor, err := hmail.ParseAddress(addr)
	if err == nil {
		m.contributors = append(m.contributors, *contributor)
	}
	return err
}

// SetCopyRightInfo of the application
func (m *MetaData) SetCopyRightInfo(year int, message string) {
	m.copySince = year
	m.copyMsg = message
}

// SetURL for application main site or company site
func (m *MetaData) SetURL(url string) {
	m.url = url
}

// SetVersion of application
func (m *MetaData) SetVersion(version string) {
	m.version = version
}

// SetBuildDate of application
func (m *MetaData) SetBuildDate(buildDate string) {
	m.buildDate = buildDate
}

// GetInfo returns application info which can be consumed by templates or output as json
func (m *MetaData) GetInfo() ApplicationInfo {
	var contributors []string
	for _, contributor := range m.contributors {
		contributors = append(contributors, contributor.String())
	}
	info := ApplicationInfo{
		Name:             m.name,
		Title:            m.title,
		ShortDescription: m.sdesc,
		LongDescription:  m.ldesc,
		CopyRight:        m.GetCopyMessage(),
		CopySince:        m.copySince,
		CopyMsg:          m.copyMsg,
		URL:              m.url,
		Version:          m.version,
		BuildDate:        m.buildDate,
		Contributors:     contributors,
	}
	return info
}

// GetCopyMessage returns copyright message
func (m *MetaData) GetCopyMessage() string {
	thisYear := time.Now()
	msg := strconv.Itoa(m.copySince)
	if m.copySince < thisYear.Year() {
		msg = fmt.Sprintf("%d - %d", m.copySince, thisYear.Year())
	}
	return fmt.Sprintf("Copyright © %s %s", m.copyMsg, msg)
}
