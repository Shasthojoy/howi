// Copyright 2016 Marko Kungla. All rights reserved.
// Use of this source code is governed by a The Apache-style
// license that can be found in the LICENSE file.

package cli

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
	"time"

	"github.com/digaverse/howi/pkg/log"
)

// TmplParser enables to parse templates for cli apps
type TmplParser struct {
	tmpl   string
	buffer bytes.Buffer
	t      *template.Template
}

// SetTemplate sets template to be parsed
func (t *TmplParser) SetTemplate(tmpl string) {
	t.tmpl = tmpl
}

// ParseTmpl parses template for cli application
// arg name is template name, arg info is data passed to template
// and elapsed is time duration used by specific type of templates and can usually set to "0"
func (t *TmplParser) ParseTmpl(name string, info interface{}, elapsed time.Duration) error {
	t.t = template.New(name)
	t.t.Funcs(template.FuncMap{
		"funcTextBold":    t.textBold,
		"funcCmdCategory": t.cmdCategory,
		"funcCmdName":     t.cmdName,
		"funcFlagName":    t.flagName,
		"funcDate":        t.dateOnly,
		"funcElapsed":     func() string { return elapsed.String() },
	})
	tmpl := template.Must(t.t.Parse(t.tmpl))
	return tmpl.Execute(&t.buffer, info)
}

// String returns parsed template as string
func (t *TmplParser) String() string {
	return t.buffer.String()
}

func (t *TmplParser) cmdCategory(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(s)
}

func (t *TmplParser) cmdName(s string) string {
	if s == "" {
		return s
	}
	return fmt.Sprintf("\033[1m %-20s\033[0m", s)
}

func (t *TmplParser) flagName(s string) string {
	if s == "" {
		return s
	}
	return fmt.Sprintf("%-25s", s)
}

func (t *TmplParser) textBold(s string) string {
	if s == "" {
		return s
	}
	return fmt.Sprintf("\033[1m%s\033[0m", s)
}

func (t *TmplParser) dateOnly(ts time.Time) string {
	y, m, d := ts.Date()
	return fmt.Sprintf("%.2d-%.2d-%d", d, m, y)
}

// Header can be shown right after application is loaded.
type Header struct {
	TmplParser
}

// Defaults for header
func (h *Header) Defaults() {
	h.SetTemplate(`
################################################################################
# {{ .Title }}{{ if .Copyright.Since }}
#  Copyright © {{ .Copyright.Since }} {{ .Copyright.By }}. All rights reserved.{{end}}
# {{if .Version}}
#   Version:    {{ .Version }}{{end}}{{if .BuildDate}}
#   Build date: {{ .BuildDate | funcDate }}{{end}}
################################################################################`)
}

// Print application header
func (h *Header) Print(log *log.Logger, project interface{}, elapsed time.Duration) {
	err := h.ParseTmpl("header-tmpl", project, elapsed)
	if err != nil {
		log.Fatal(err)
	}
	log.ColoredLine(h.buffer.String())
}

// Footer can be shown right before application exits
type Footer struct {
	TmplParser
}

// Defaults for footer
func (f *Footer) Defaults() {
	f.SetTemplate(`
################################################################################
# elapsed: {{ funcElapsed }}
################################################################################`)
}

// Print application footer
func (f *Footer) Print(log *log.Logger, project interface{}, elapsed time.Duration) {
	err := f.ParseTmpl("footer-tmpl", project, elapsed)
	if err != nil {
		log.Fatal(err)
	}
	log.ColoredLine(f.buffer.String())
}
