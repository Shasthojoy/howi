package hcli

const (
	// PhasePending marks that phase is pending (800)
	PhasePending = iota + 800
	// PhaseSkipped marks that phase was skipped (801)
	PhaseSkipped
	// PhaseSuccess marks that phase was successful (802)
	PhaseSuccess
	// PhaseRunning marks that phase is running (803)
	PhaseRunning
	// PhaseFailed marks that phase failed (804)
	PhaseFailed
	// FmtErrFlagShadowing formats shadowed flag error.
	FmtErrFlagShadowing = "flag(%s) alias %q shadows existing flag for %q"
	// FmtErrCommandFlagShadowing formats command shaddowed flag error.
	FmtErrCommandFlagShadowing = "command (%s) flag %q alias %q shadows existing flag"
	// FmtErrCommandNameInvalid formats invalid command name error.
	FmtErrCommandNameInvalid = "command name %q is invalid - must match following regex %v"
	// FmtErrCommandNameInUse formats command name in use error.
	FmtErrCommandNameInUse = "command %q is already in use"
	// FmtErrCommandMissingDoFn formats command missing Do function error.
	FmtErrCommandMissingDoFn = "command (%s) must have DoFn"
	// FmtErrPhaseFailed formats phase failure error.
	FmtErrPhaseFailed = "phase: %q failed (%s)"
	// FmtErrAppWithNoCommandsOrFlags formats error when application is started
	// without any commands or flags configured.
	FmtErrAppWithNoCommandsOrFlags = "application has no flags or commands"
	// FmtErrAppUnnamed formats error for unnamed application.
	FmtErrAppUnnamed = "application must have a name"
	// FmtErrUnknownGlobalFlag formats invalid global flag error.
	FmtErrUnknownGlobalFlag = "unknown global flag %q"
	// FmtErrUnknownCommand formats unknown command error.
	FmtErrUnknownCommand = "unknown command %q"
	// FmtErrUnknownFlag formats error for any request looking non existing flag.
	FmtErrUnknownFlag = "unknown flag %q for command %q"
	// FmtErrUnknownSubcommand formats error for unknown subcommand request.
	FmtErrUnknownSubcommand = "unknown subcommand %q for command %q"
	// FmtErrTooManyArgs formats error when too many arguments are passed.
	FmtErrTooManyArgs = "too many arguments for command %q which accepts max (%d) args"
	// FmtErrInvalidCommandArgs is returned when invalid args are recieved by
	// command parser.
	FmtErrInvalidCommandArgs = "invalid arguments passed for (%s).Parse"
)

// NewApplication constructs new application and returns it's instance for
// configuration. Returned Application has basic initialization done and
// all defaults set.
func NewApplication() *Application {
	return &Application{}
}

// NewCommand returns new command constructor.
func NewCommand(name string) Command {
	return Command{name: name}
}
