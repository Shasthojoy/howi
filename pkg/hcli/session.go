package hcli

// Session is instance shared between command phases
type Session struct{}

// SessionConfig enables some application runtime configuration options
// which can be changed within command phases and task functions.
type SessionConfig struct{}

// Phase tracks execution of specific phase
type Phase struct{}

// Task is single task which will be executed in it's own go routine
// within the execution phase it was attached to.
type Task struct{}
