package git

import (
	"reflect"
	"strings"
	"testing"

	"github.com/howi-ce/howi/pkg/std/hstrings"
)

func TestGitAPI(t *testing.T) {
	tests := []struct {
		name string
		desc string
	}{
		{"add", "Add file contents to the index"},
		{"am", "Apply a series of patches from a mailbox"},
		{"annotate", "Annotate file lines with commit information (deprecated)"},
		{"apply", "Apply a patch to files and/or to the index"},
		{"archimport", "Import an Arch repository into Git"},
		{"archive", "Create an archive of files from a named tree"},
		{"bisect", "Use binary search to find the commit that introduced a bug"},
		{"blame", "Show what revision and author last modified each line of a file"},
		{"branch", "List, create, or delete branches"},
		{"bundle", "Move objects and refs by archive"},
		{"cat-file", "Provide content or type and size information for repository objects"},
		{"check-attr", "Display gitattributes information"},
		{"check-ignore", "Debug gitignore / exclude files"},
		{"check-mailmap", "Show canonical names and email addresses of contacts"},
		{"check-ref-format", "Ensures that a reference name is well formed"},
		{"checkout-index", "Copy files from the index to the working tree"},
		{"checkout", "Switch branches or restore working tree files"},
		{"cherry-pick", "Apply the changes introduced by some existing commits"},
		{"cherry", "Find commits yet to be applied to upstream"},
		{"citool", "Graphical alternative to git-commit"},
		{"clean", " Remove untracked files from the working tree"},
		{"clone", "Clone a repository into a new directory"},
		{"column", "Display data in columns"},
		{"commit-tree", "Create a new commit object"},
		{"commit", "Record changes to the repository"},
		{"config", "Get and set repository or global options"},
		{"count-objects", "Count unpacked number of objects and their disk consumption"},
		{"credential-cache--daemon", "Temporarily store user credentials in memory"},
		{"credential-cache", "Helper to temporarily store passwords in memory"},
		{"credential-store", "Helper to store credentials on disk"},
		{"credential", "Retrieve and store user credentials"},
		{"credential-libsecret", "credential helper that talks via libsecret with implementations of XDG Secret Service API"},
		{"credential-netrc", "credential helper"},
		{"cvsexportcommit", "Export a single commit to a CVS checkout"},
		{"cvsimport", "Salvage your data out of another SCM people love to hate"},
		{"cvsserver", "A CVS server emulator for Git"},
		{"daemon", "A really simple server for Git repositories"},
		{"describe", "Describe a commit using the most recent tag reachable from it"},
		{"diff-files", "Compares files in the working tree and the index"},
		{"diff-index", "Compare a tree to the working tree or index"},
		{"diff-tree", "Compares the content and mode of blobs found via two tree objects"},
		{"diff", "Show changes between commits, commit and working tree, etc"},
		{"difftool", "Show changes using common diff tools"},
		{"fast-export", "Git data exporter"},
		{"fast-import", "Backend for fast Git data importers"},
		{"fetch-pack", "Receive missing objects from another repository"},
		{"fetch", "Download objects and refs from another repository"},
		{"filter-branch", "Rewrite branches"},
		{"fmt-merge-msg", "Produce a merge commit message"},
		{"for-each-ref", "Output information on each ref"},
		{"format-patch", "Prepare patches for e-mail submission"},
		{"fsck-objects", "Verifies the connectivity and validity of the objects in the database"},
		{"fsck", "Verifies the connectivity and validity of the objects in the database"},
		{"gc", "Cleanup unnecessary files and optimize the local repository"},
		{"get-tar-commit-id", "Extract commit ID from an archive created using git-archive"},
		{"grep", "Print lines matching a pattern"},
		{"gui", "A portable graphical interface to Git"},
		{"hash-object", "Compute object ID and optionally creates a blob from a file"},
		{"help", "Display help information about Git"},
		{"http-backend", "Server side implementation of Git over HTTP"},
		{"http-fetch", "Download from a remote Git repository via HTTP"},
		{"http-push", "Push objects over HTTP/DAV to another repository"},
		{"imap-send", "Send a collection of patches from stdin to an IMAP folder"},
		{"index-pack", "Build pack index file for an existing packed archive"},
		{"init-db", "Creates an empty Git repository"},
		{"init", "Create an empty Git repository or reinitialize an existing one"},
		{"instaweb", "Instantly browse your working repository in gitweb"},
		{"interpret-trailers", "help add structured information into commit messages"},
		{"gitk", "The Git repository browser"},
		{"log", "Show commit logs"},
		{"ls-files", "Show information about files in the index and the working tree"},
		{"ls-remote", "List references in a remote repository"},
		{"ls-tree", "List the contents of a tree object"},
		{"mailinfo", "Extracts patch and authorship from a single e-mail message"},
		{"mailsplit", "imple UNIX mbox splitter program"},
		{"merge-base", "Find as good common ancestors as possible for a merge"},
		{"merge-file", "Run a three-way file merge"},
		{"merge-index", "Run a merge for files needing merging"},
		{"merge-octopus", "Resolves cases with more than two heads, but refuses to do a complex merge that needs manual resolution"},
		{"merge-one-file", "The standard helper program to use with git-merge-index"},
		{"merge-ours", "Resolves any number of heads, but the resulting tree of the merge is always that of the current branch head, effectively ignoring all changes from all other branches"},
		{"merge-recursive", "This can only resolve two heads using a 3-way merge algorithm"},
		{"merge-resolve", "This can only resolve two heads (i.e. the current branch and another branch you pulled from) using a 3-way merge algorithm."},
		{"merge-subtree", "This is a modified recursive strategy. When merging trees A and B"},
		{"merge-tree", "Show three-way merge without touching index"},
		{"merge", "Join two or more development histories together"},
		{"mergetool", "Run merge conflict resolution tools to resolve merge conflicts"},
		{"mktag", "Creates a tag object"},
		{"mktree", "Build a tree-object from ls-tree formatted text"},
		{"mv", "Move or rename a file, a directory, or a symlink"},
		{"name-rev", "Find symbolic names for given revs"},
		{"notes", "Add or inspect object notes"},
		{"p4", "Import from and submit to Perforce repositories"},
		{"pack-objects", "Create a packed archive of objects"},
		{"pack-redundant", "Find redundant pack files"},
		{"pack-refs", "Pack heads and tags for efficient repository access"},
		{"parse-remote", "Routines to help parsing remote repository access parameters"},
		{"patch-id", "Compute unique ID for a patch"},
		{"prune", "Prune all unreachable objects from the object database"},
		{"prune-packed", "Remove extra objects that are already in pack files"},
		{"pull", "Fetch from and integrate with another repository or a local branch"},
		{"push", "Update remote refs along with associated objects"},
		{"quiltimport", "Applies a quilt patchset onto the current branch"},
		{"read-tree", "Reads tree information into the index"},
		{"rebase", "Reapply commits on top of another base tip"},
		{"receive-pack", "Receive what is pushed into the repository"},
		{"reflog", "Manage reflog information"},
		{"remote-ext", "Bridge smart transport to external command."},
		{"remote-fd", "Reflect smart transport stream back to caller"},
		{"remote-ftp", ""},
		{"remote-ftps", ""},
		{"remote-http", ""},
		{"remote-https", ""},
		{"remote-testsvn", ""},
		{"remote-testgit", "Example remote-helper"},
		{"remote", "Manage set of tracked repositories"},
		{"repack", "Pack unpacked objects in a repository"},
		{"replace", "reate, list, delete refs to replace objects"},
		{"request-pull", "Generates a summary of pending changes"},
		{"rerere", "Reuse recorded resolution of conflicted merges"},
		{"reset", "Reset current HEAD to the specified state"},
		{"rev-list", "Lists commit objects in reverse chronological order"},
		{"rev-parse", "Pick out and massage parameters"},
		{"revert", "Revert some existing commits"},
		{"rm", "Remove files from the working tree and from the index"},
		{"send-email", "Send a collection of patches as emails"},
		{"send-pack", "Push objects over Git protocol to another repository"},
		{"shell", "Restricted login shell for Git-only SSH access"},
		{"shortlog", "Summarize 'git log' output"},
		{"show-branch", "Show branches and their commits"},
		{"show-index", "Show packed archive index"},
		{"show-ref", "List references in a local repository"},
		{"show", "Show various types of objects"},
		{"stage", "Add file contents to the staging area"},
		{"stash", "Stash the changes in a dirty working directory away"},
		{"status", "Show the working tree status"},
		{"stripspace", "Remove unnecessary whitespace"},
		{"submodule", "Initialize, update or inspect submodules"},
		{"svn", "Bidirectional operation between a Subversion repository and Git"},
		{"symbolic-ref", "Read, modify and delete symbolic refs"},
		{"tag", " Create, list, delete or verify a tag object signed with GPG"},
		{"unpack-file", "Creates a temporary file with a blob's contents"},
		{"unpack-objects", "Unpack objects from a packed archive"},
		{"update-index", "Register file contents in the working tree to the index"},
		{"update-ref", "Update the object name stored in a ref safely"},
		{"update-server-info", "Update auxiliary info file to help dumb servers"},
		{"upload-archive", "Send archive back to git-archive"},
		{"upload-pack", "Send objects packed back to git-fetch-pack"},
		{"var", "Show a Git logical variable"},
		{"verify-commit", "Check the GPG signature of commits"},
		{"verify-pack", "Validate packed Git archive files"},
		{"verify-tag", "Check the GPG signature of tags"},
		{"whatchanged", "Show logs with difference each commit introduces"},
		{"worktree", "Manage multiple working trees"},
		{"write-tree", "Create a tree object from the current index"},
	}

	gitType := reflect.TypeOf(&Git{})
	for _, tt := range tests {
		expected := hstrings.ToCamelCaseAlnum(tt.name)
		if expected == "Gc" {
			expected = "GC"
		}
		if expected == "Gui" {
			expected = "GUI"
		}
		if expected == "RemoteTestsvn" {
			expected = "RemoteTestSvn"
		}
		if expected == "RemoteTestgit" {
			expected = "RemoteTestGit"
		}
		if strings.HasSuffix(expected, "Id") {
			expected = string(expected[:len(expected)-2]) + "ID"
		}
		if strings.HasSuffix(expected, "Db") {
			expected = string(expected[:len(expected)-2]) + "DB"
		}
		if strings.HasSuffix(expected, "Http") {
			expected = string(expected[:len(expected)-4]) + "HTTP"
		}
		if strings.HasSuffix(expected, "Https") {
			expected = string(expected[:len(expected)-5]) + "HTTPS"
		}
		if strings.HasPrefix(expected, "Http") {
			expected = "HTTP" + string(expected[4:])
		}

		if _, exits := gitType.MethodByName(expected); !exits {
			t.Errorf("Git.%q (git-%s) is not implemented: %s", expected, tt.name, tt.desc)
		}
	}
}
