package git

import (
	"os/exec"

	"github.com/howi-ce/howi/pkg/vars"
)

var (
	// Git executable
	executable string

	// Git version
	version vars.Value
)

// LookPath searches for an executable git binary
// in the directories named by the PATH environment variable.
// If file contains a slash, it is tried directly and the PATH is not consulted.
// The result may be an absolute path or a path relative to the current directory.
func lookPath() (string, error) {
	var err error
	if executable == "" {
		executable, err = exec.LookPath("git")
	}
	return executable, err
}
