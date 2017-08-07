package cli

import (
	"strings"

	"github.com/howi-ce/howi/addon/application/plugin/cli/clitmpl"
	"github.com/howi-ce/howi/addon/application/plugin/cli/flags"
	"github.com/howi-ce/howi/std/log"
)

var (
	helpGlobalTmpl = `{{if .Info.ShortDescription}}{{ .Info.ShortDescription }}{{end}}

 Usage:
  {{ .Info.Name }} command
  {{ .Info.Name }} command [command-flags] [arguments]
  {{ .Info.Name }} [global-flags] command [command-flags] [arguments]
  {{ .Info.Name }} [global-flags] command ...subcommand [command-flags] [arguments]
  
 The commands are:{{ if .PrimaryCommands }}{{ range $cmdObj := .PrimaryCommands }}
  {{ $cmdObj.Name | funcCmdName }}{{ $cmdObj.ShortDesc }}{{ end }}{{ end }}
{{ if .CommandsCategorized }}{{ range $cat, $cmds := .CommandsCategorized }}
 {{ $cat | funcCmdCategory }}
 {{ range $cmdObj := $cmds }}
 {{$cmdObj.Name | funcCmdName }}{{ $cmdObj.ShortDesc }}{{ end }}
 {{ end }}{{ end }}
 
 The global flags are:{{ if .Flags }}{{ range $flag := .Flags }}{{ if not .IsHidden }}
  {{$flag.HelpName | funcFlagName }}{{ $flag.Usage }}{{ if $flag.HelpAliases }}
   {{$flag.HelpAliases}}
{{ end }}{{ end }}{{ end }}{{ end }}
{{if .Info.LongDescription}}{{ .Info.LongDescription }}{{end}}`

	helpCommandTmpl = `{{ if .Command.LongDesc }}{{.Command.LongDesc}}{{ else }}{{ .Command.ShortDesc }}{{ end }}

 Usage:
   {{ .Usage | funcTextBold }}{{ if .Command.Usage }}
   {{ .Command.Usage }}{{ end }}
{{ if .Command.HasSubcommands }}
 {{ print "Subcommands" | funcCmdCategory }}
{{ range $cmdObj := .Command.GetSubcommands }}
{{ $cmdObj.Name | funcCmdName }}{{ $cmdObj.ShortDesc }}{{ end }}
{{ end }}
{{ if .Command.AcceptsFlags }} Accepts following flags:{{ range $flag := .Flags }}{{ if not .IsHidden }}
 {{$flag.HelpName | funcFlagName }}{{ $flag.Usage }}{{ if $flag.HelpAliases }}
	{{$flag.HelpAliases}}
{{ end }}{{ end }}{{ end }}{{ end }}`
)

// HelpGlobal used to show help for application
type HelpGlobal struct {
	clitmpl.TmplParser
	Info                ApplicationInfo
	Commands            map[string]Command
	Flags               map[int]flags.Interface
	PrimaryCommands     []Command
	CommandsCategorized map[string][]Command
}

// Print application help
func (h *HelpGlobal) Print(log *log.Logger) {
	h.SetTemplate(helpGlobalTmpl)
	for _, cmdObj := range h.Commands {
		if cmdObj.category == "" {
			h.PrimaryCommands = append(h.PrimaryCommands, cmdObj)
		} else {
			if h.CommandsCategorized == nil {
				h.CommandsCategorized = make(map[string][]Command)
			}
			h.CommandsCategorized[cmdObj.category] = append(h.CommandsCategorized[cmdObj.category],
				cmdObj)
		}
	}
	err := h.ParseTmpl("help-global-tmpl", h, 0)
	if err != nil {
		log.Fatal(err)
	}
	log.Line(h.String())
}

// HelpCommand is used to display help for command
type HelpCommand struct {
	clitmpl.TmplParser
	Info    ApplicationInfo
	Command Command
	Usage   string
	Flags   []flags.Interface
}

// Print command help
func (h *HelpCommand) Print(log *log.Logger) {
	h.SetTemplate(helpCommandTmpl)
	usage := []string{h.Info.Name}
	for _, parent := range h.Command.getParents() {
		usage = append(usage, parent)
	}
	usage = append(usage, h.Command.Name())
	if h.Command.AcceptsFlags() {
		usage = append(usage, "[flags]")
	}
	if h.Command.HasSubcommands() {
		usage = append(usage, "[subcommands]")
	}
	if h.Command.AcceptsArgs() {
		usage = append(usage, "[args]")
	}
	h.Usage = strings.Join(usage, " ")

	for _, flag := range h.Command.getFlags() {
		h.Flags = append(h.Flags, flag)
	}
	err := h.ParseTmpl("help-global-tmpl", h, 0)
	if err != nil {
		log.Fatal(err)
	}
	log.Line(h.String())
}
