package hcli

// NewApplication constructs new application and returns it's instance for
// configuration. Returned Application has basic initialization done and
// all defaults set.
func NewApplication() *Application {
	return &Application{}
}

// NewCommand returns new command constructor.
func NewCommand() Command {
	return Command{}
}
