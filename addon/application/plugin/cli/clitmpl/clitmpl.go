package clitmpl

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
	"time"

	"github.com/howi-ce/howi/std/log"
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
		"funcTextBold":    textBold,
		"funcCmdCategory": cmdCategory,
		"funcCmdName":     cmdName,
		"funcFlagName":    flagName,
		"funcDate":        dateOnly,
		"funcElapsed":     func() string { return elapsed.String() },
	})
	tmpl := template.Must(t.t.Parse(t.tmpl))
	return tmpl.Execute(&t.buffer, info)
}

// String returns parsed template as string
func (t *TmplParser) String() string {
	return t.buffer.String()
}

func cmdCategory(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(s)
}

func cmdName(s string) string {
	if s == "" {
		return s
	}
	return fmt.Sprintf("\033[1m %-20s\033[0m", s)
}

func flagName(s string) string {
	if s == "" {
		return s
	}
	return fmt.Sprintf("%-25s", s)
}

func textBold(s string) string {
	if s == "" {
		return s
	}
	return fmt.Sprintf("\033[1m%s\033[0m", s)
}

func dateOnly(s string) string {
	date, _ := time.Parse(time.RFC3339, s)
	y, m, d := date.Date()
	return fmt.Sprintf("%.2d-%.2d-%d", d, m, y)
}

// Header can be shown right after application is loaded.
type Header struct {
	TmplParser
}

// Print application header
func (h *Header) Print(log *log.Logger, info interface{}, elapsed time.Duration) {
	err := h.ParseTmpl("header-tmpl", info, elapsed)
	if err != nil {
		log.Fatal(err)
	}
	log.ColoredLine(h.buffer.String())
}

// Footer can be shown right before application exits
type Footer struct {
	TmplParser
}

// Print application footer
func (f *Footer) Print(log *log.Logger, info interface{}, elapsed time.Duration) {
	err := f.ParseTmpl("footer-tmpl", info, elapsed)
	if err != nil {
		log.Fatal(err)
	}
	log.ColoredLine(f.buffer.String())
}
