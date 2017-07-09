package git

import (
	"errors"
	"fmt"
)

var (
	// ErrNoGit is error used when no git is installed
	ErrNoGit = errors.New("git binary not found")
	// ErrDestinationNotEmpty is error used when trying to clone into non empty directory
	ErrDestinationNotEmpty = errors.New("destination path already exists and is not an empty directory")
	// ErrNotAGitRepository is used when git command is executed outside of repository
	ErrNotAGitRepository = errors.New("not a git repository or any parent up to mount point /")
)

// ErrDeprecated is error returned when git command with matching name is deprecated.
type ErrDeprecated error

func errDeprecated(method string, alternative string, docs string) ErrDeprecated {
	return fmt.Errorf("method %s is deprecated - alternatives (%s) - docs (%s)",
		method, alternative, docs)
}

// ErrNotImplemented is error returned when git command with matching name is
// not implemented at this point and may be removed in next release.
type ErrNotImplemented error

func errNotImplemented(method string, reason ...string) ErrNotImplemented {
	return fmt.Errorf("method %s is not implemented - it may be removed in next version or is in development (%s)",
		method, reason)
}
