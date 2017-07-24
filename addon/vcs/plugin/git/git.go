package git

import (
	"os"
	"os/exec"

	"github.com/howi-ce/howi/pkg/vars"
)

var (
	// Git executable
	executable string

	// Git version
	version vars.Value
)

// GlobalConfig returns global git config
func GlobalConfig() ([]string, error) {
	gitconfig, err := cmdgit("config", "--global", "--list")
	return gitconfig.Lines(), err
}

// LookPath searches for an executable git binary
// in the directories named by the PATH environment variable.
// If file contains a slash, it is tried directly and the PATH is not consulted.
// The result may be an absolute path or a path relative to the current directory.
func LookPath() (string, error) {
	var err error
	if executable == "" {
		executable, err = exec.LookPath("git")
	}
	return executable, err
}

func cmdgit(v ...string) (*Output, error) {
	b, err := exec.Command("git", v...).Output()
	return &Output{b: b}, err
}
func cmdgitInPath(p string, v ...string) (*Output, error) {
	cur, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	if err := os.Chdir(p); err != nil {
		return nil, err
	}
	resp, rderr := exec.Command("git", v...).Output()
	if err := os.Chdir(cur); err != nil {
		return nil, err
	}
	return &Output{b: resp}, rderr
}
