package hcli

import (
	"github.com/howi-ce/howi/pkg/hcli/flags"
	"github.com/howi-ce/howi/pkg/std/herrors"
	"github.com/howi-ce/howi/pkg/vars"
)

// Command for application
type Command struct {
	name           string
	hidden         bool
	category       string
	usage          string
	shortDesc      string
	longDesc       string
	errs           herrors.MultiError // internal errors
	beforeFn       func(w *Worker)
	doFn           func(w *Worker)
	afterFailureFn func(w *Worker)
	afterSuccessFn func(w *Worker)
	afterAlwaysFn  func(w *Worker)
	subCommands    map[string]Command      // subcommands
	flags          map[int]flags.Interface // command flags
	flagAliases    map[string]int          // command flag aliases
	acceptArgs     int
	args           []vars.Value
	subCmd         *Command // if subcommand was called
	parents        []string
}
