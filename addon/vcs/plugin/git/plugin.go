// Copyright 2005-2017 Marko Kungla. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// license that can be found in the LICENSE file.

package git

import (
	"os/exec"
	"strings"

	"github.com/howi-ce/howi/addon/filesystem"
	"github.com/howi-ce/howi/addon/filesystem/plugin/path"
	"github.com/howi-ce/howi/std/herrors"
	"github.com/howi-ce/howi/std/hos/hexec"
	"github.com/howi-ce/howi/std/hvars"
)

var (
	// Git executable
	executable string
	// Git version
	version hvars.Value
	// ErrNoGit is error used when no git is installed
	ErrNoGit = herrors.New("git binary not found")
	// ErrDestinationNotEmpty is error used when trying to clone into non empty directory
	ErrDestinationNotEmpty = herrors.New("destination path already exists and is not an empty directory")
	// ErrNotAGitRepository is used when git command is executed outside of repository
	ErrNotAGitRepository = herrors.New("not a git repository or any parent up to mount point /")
)

// NewPlugin returns new Git Plugin
func NewPlugin(wd string) (*Plugin, error) {
	p, err := path.NewPlugin(wd)
	if err != nil {
		return nil, err
	}
	if p.IsGitRepository() {
		fs, err2 := filesystem.NewAddon(p.Abs())
		return &Plugin{fs: fs}, err2
	}

	for !p.IsGitRepository() {
		p, err = path.NewPlugin(p.Join("../"))
		if err != nil {
			return nil, err
		}
		if p.Abs() == "/" {
			return nil, herrors.New("Not a git repository (or any parent up to mount point /)")
		}
		if p.IsGitRepository() {
			fs, err := filesystem.NewAddon(p.Abs())
			return &Plugin{fs: fs}, err
		}
	}
	return nil, herrors.New("failed to load git repository")
}

// Plugin for Git source contral managment
type Plugin struct {
	fs *filesystem.Addon
}

// Apply a patch to files and/or to the index. Reads the supplied diff output
// (i.e. "a patch") and applies it to files. When running from a subdirectory in
// a repository, patched paths outside the directory are ignored. It apply the
// patch to files, and does not require them to be in a Git repository.
//   git-apply - https://git-scm.com/docs/git-apply
func (p *Plugin) Apply() herrors.ErrNotImplemented {
	return newErrNotImplemented("apply")
}

// CheckRefFormat ensures that a reference name is well formed
//   check-ref-format - https://git-scm.com/docs/git-check-ref-format
func (p *Plugin) CheckRefFormat() herrors.ErrNotImplemented {
	return newErrNotImplemented("check-ref-format")
}

// Clone a repository into a newly created directory, creates remote-tracking
// branches for each branch in the cloned repository
//   (visible using git branch -r),
// and creates and checks out an initial branch that is forked from the cloned
// repository’s currently active branch.
//   git-clone - https://git-scm.com/docs/git-clone
func (p *Plugin) Clone() herrors.ErrNotImplemented {
	return newErrNotImplemented("clone")
}

// Column displays data in columns
//   git-column - https://git-scm.com/docs/git-column
func (p *Plugin) Column() herrors.ErrNotImplemented {
	return newErrNotImplemented("column")
}

// Add file contents to the index using the current content found in the working
// tree, to prepare the content staged for the next commit.
//   git-add - https://git-scm.com/docs/git-add
func (p *Plugin) Add() herrors.ErrNotImplemented {
	return newErrNotImplemented("add")
}

// Am splits mail messages in a mailbox into commit log message, authorship
// information and patches, and applies them to the current branch.
//   git-am - https://git-scm.com/docs/git-am
func (p *Plugin) Am() herrors.ErrNotImplemented {
	return newErrNotImplemented("am")
}

// Annotate (deprecated)
//   git-annotate - https://git-scm.com/docs/git-annotate
func (p *Plugin) Annotate() herrors.ErrDeprecated {
	return newErrDeprecated("annotate", "blame", "https://git-scm.com/docs/git-annotate")
}

// Archimport imports a project from one or more Arch repositories. It will
// follow branches and repositories within the namespaces defined by the
// <archive/branch> parameters supplied.
//   git-archimport - https://git-scm.com/docs/git-archimport
func (p *Plugin) Archimport() herrors.ErrNotImplemented {
	return newErrNotImplemented("archimport")
}

// Archive creates an archive of the specified format containing the tree
// structure for the named tree, and writes it out to the standard output.
// If <prefix> is specified it is prepended to the filenames in the archive.
//   git-archive - https://git-scm.com/docs/git-archive
func (p *Plugin) Archive() herrors.ErrNotImplemented {
	return newErrNotImplemented("am")
}

// Bisect binary search to find the commit that introduced a bug.
//   git-bisect - https://git-scm.com/docs/git-bisect
func (p *Plugin) Bisect() herrors.ErrNotImplemented {
	return newErrNotImplemented("am")
}

// Blame annotates each line in the given file with information from the
// revision which last modified the line.
//   git-blame - https://git-scm.com/docs/git-blame
func (p *Plugin) Blame() herrors.ErrNotImplemented {
	return newErrNotImplemented("am")
}

// Branch list, create, or delete branches.
//   git-branch - https://git-scm.com/docs/git-branch
func (p *Plugin) Branch() herrors.ErrNotImplemented {
	return newErrNotImplemented("branch")
}

// Bundle moves objects and refs by archive
//   git-bundle - https://git-scm.com/docs/git-bundle
func (p *Plugin) Bundle() herrors.ErrNotImplemented {
	return newErrNotImplemented("bundle")
}

// CatFile provides content or type and size information for repository objects
//   git-cat-file - https://git-scm.com/docs/git-cat-file
func (p *Plugin) CatFile() herrors.ErrNotImplemented {
	return newErrNotImplemented("cat-file")
}

// CheckAttr displays gitattributes information
//   git-check-attr - https://git-scm.com/docs/git-check-attr
func (p *Plugin) CheckAttr() herrors.ErrNotImplemented {
	return newErrNotImplemented("check-attr")
}

// CheckIgnore enables you to debug gitignore / exclude files.
//   git-check-ignore - https://git-scm.com/docs/git-check-ignore
func (p *Plugin) CheckIgnore() herrors.ErrNotImplemented {
	return newErrNotImplemented("check-ignore")
}

// CheckMailmap show canonical names and email addresses of contacts
//   git-check-mailmap - https://git-scm.com/docs/git-check-mailmap
func (p *Plugin) CheckMailmap() herrors.ErrNotImplemented {
	return newErrNotImplemented("check-mailmap")
}

// Checkout switch branches or restore working tree files. Updates files in
// the working tree to match the version in the index or the specified tree.
// If no paths are given, git checkout will also update HEAD to set
// the specified branch as the current branch.
//   git-checkout - https://git-scm.com/docs/git-checkout
func (p *Plugin) Checkout() herrors.ErrNotImplemented {
	return newErrNotImplemented("checkout")
}

// CheckoutIndex will copy all files listed from the index to the working
// directory (not overwriting existing files).
//   git-checkout-index - https://git-scm.com/docs/git-checkout-index
func (p *Plugin) CheckoutIndex() herrors.ErrNotImplemented {
	return newErrNotImplemented("checkout-index")
}

// Cherry finds commits yet to be applied to upstream. Determine whether there
// are commits in
//   <head>..<upstream>
// that are equivalent to those in the range
//   <limit>..<head>
//
//  git-cherry - https://git-scm.com/docs/git-cherry .
func (p *Plugin) Cherry() herrors.ErrNotImplemented {
	return newErrNotImplemented("cherry")
}

// CherryPick applies the changes introduced by some existing commits. Given one
// or more existing commits, apply the change each one introduces,recording a
// new commit for each. This requires your working tree to be clean
// (no modifications from the HEAD commit).
//   git-cherry-pick - https://git-scm.com/docs/git-cherry-pick
func (p *Plugin) CherryPick() herrors.ErrNotImplemented {
	return newErrNotImplemented("cherry-pick")
}

// Clean the working tree by recursively removing files that are not under
// version control, starting from the current directory. Normally, only files
// unknown to Git are removed, but if the -x option is specified, ignored files
// are also removed. This can, for example,
// be useful to remove all build products. If any optional <path>... arguments
// are given, only those paths are affected.
//   git-clean https://git-scm.com/docs/git-clean
func (p *Plugin) Clean() herrors.ErrNotImplemented {
	return newErrNotImplemented("clean")
}

// Citool Graphical alternative to git-commit is removed from this library.
//   git-citool - https://git-scm.com/docs/git-citool
func (p *Plugin) Citool() herrors.ErrDeprecated {
	return newErrDeprecated("citool", "none", "https://git-scm.com/docs/git-citool")
}

// Commit records changes to the repository. Stores the current contents of the
// index in a new commit along with a log message from the user describing
// the changes.
//   git-commit - https://git-scm.com/docs/git-commit
func (p *Plugin) Commit() herrors.ErrNotImplemented {
	return newErrNotImplemented("commit")
}

// CommitTree creates a new commit object. This is usually not what an end
// user wants to run directly. See .Commmit instead.
//   git-commit-tree - https://git-scm.com/docs/git-commit-tree
func (p *Plugin) CommitTree() herrors.ErrNotImplemented {
	return newErrNotImplemented("commit-tree")
}

// Config get and set repository or global options. You can query/set/replace/unset
// options with this command. The name is actually the section and the key separated
// by a dot, and the value will be escaped.
//   git-config - https://git-scm.com/docs/git-config
func (p *Plugin) Config(flags ...string) ([]string, error) {

	flags = append([]string{"config"}, flags...)
	gitconfig, err := exec.Command("git", flags...).Output()
	return strings.Split(string(gitconfig), "\n"), err
}

// CountObjects counts unpacked number of objects and their disk consumption.
// This counts the number of unpacked object files and disk space consumed by them,
// to help you decide when it is a good time to repack.
//   git-count-objects - https://git-scm.com/docs/git-count-objects
func (p *Plugin) CountObjects() herrors.ErrNotImplemented {
	return newErrNotImplemented("count-objects")
}

// Credential retrieve and store user credentials. Git has an internal interface
// for storing and retrieving credentials from system-specific helpers, as well
// as prompting the user for usernames and passwords. The git-credential command
// exposes this interface to scripts which may want to retrieve, store, or prompt
// for credentials in the same manner as Git. The design of this scriptable
// interface models the internal C API; see the Git credential API for more
// background on the concepts.
//   git-credential - https://git-scm.com/docs/git-credential
func (p *Plugin) Credential() herrors.ErrNotImplemented {
	return newErrNotImplemented("credential")
}

// CredentialCache caches credentials in memory for use by future Git programs.
// The stored credentials never touch the disk, and are forgotten after a
// configurable timeout. The cache is accessible over a Unix domain socket,
// restricted to the current user by filesystem permissions.
//   git-credential-cache - https://git-scm.com/docs/git-credential-cache
func (p *Plugin) CredentialCache() herrors.ErrNotImplemented {
	return newErrNotImplemented("credential")
}

// CredentialCacheDaemon This command listens on the Unix domain socket specified
// by <socket> for git-credential-cache clients. Clients may store and retrieve credentials.
// You probably don’t want to invoke this command yourself;
// it is started automatically when you use git-credential-cache[1].
//   git-credential-cache--daemon - https://git-scm.com/docs/git-credential-cache--daemon
func (p *Plugin) CredentialCacheDaemon() herrors.ErrDeprecated {
	return newErrDeprecated("credential-cache--daemon", "credential-cache",
		"https://git-scm.com/docs/git-credential-cache--daemon")
}

// CredentialStore Using this helper will store your passwords unencrypted on disk,
// protected only by filesystem permissions. If this is not an acceptable
// security tradeoff, try git-credential-cache[1], or find a helper that
// integrates with secure storage provided by your operating system.
// This command stores credentials indefinitely on disk for use by future Git programs.
//   git-credential-store - https://git-scm.com/docs/git-credential-store
func (p *Plugin) CredentialStore() herrors.ErrNotImplemented {
	return newErrNotImplemented("credential-store")
}

// CredentialLibsecret helper that talks via libsecret with
// implementations of XDG Secret Service API
//   git-credential-libsecret - #
func (p *Plugin) CredentialLibsecret() herrors.ErrNotImplemented {
	return newErrNotImplemented("credential-libsecret")
}

// CredentialNetrc helper credential helper
//   git-credential-netrc - #
func (p *Plugin) CredentialNetrc() herrors.ErrNotImplemented {
	return newErrNotImplemented("credential-netrc")
}

// Cvsexportcommit Exports a commit from Git to a CVS checkout, making it easier
// to merge patches from a Git repository into a CVS repository.
//   git-cvsexportcommit - https://git-scm.com/docs/git-cvsexportcommit
func (p *Plugin) Cvsexportcommit() herrors.ErrNotImplemented {
	return newErrNotImplemented("cvsexportcommit")
}

// Cvsimport  Salvage your data out of another SCM people love to hate.
// cvsps version 2 is deprecated.
//   git-cvsimport - https://git-scm.com/docs/git-cvsimport
func (p *Plugin) Cvsimport() herrors.ErrDeprecated {
	return newErrDeprecated("cvsimport", "", "https://git-scm.com/docs/git-cvsimport")
}

// Cvsserver A CVS server emulator for Git
//   git-cvsserver - https://git-scm.com/docs/git-cvsserver
func (p *Plugin) Cvsserver() herrors.ErrNotImplemented {
	return newErrNotImplemented("cvsserver")
}

// Daemon is a really simple TCP Git daemon that normally listens on port
// "DEFAULT_GIT_PORT" aka 9418. It waits for a connection asking for a service,
// and will serve that service if it is enabled.
//   git-daemon - https://git-scm.com/docs/git-daemon
func (p *Plugin) Daemon() herrors.ErrNotImplemented {
	return newErrNotImplemented("daemon")
}

// Describe command finds the most recent tag that is reachable from a commit.
// If the tag points to the commit, then only the tag is shown. Otherwise,
// it suffixes the tag name with the number of additional commits on top of the
// tagged object and the abbreviated object name of the most recent commit.
//   git-describe - https://git-scm.com/docs/git-describe
func (p *Plugin) Describe(s ...string) (hexec.Output, error) {
	s = append([]string{"describe"}, s...)
	return cmdgitInPath(p.fs.RealAbs(), s...)
}

// DiffFiles Compares the files in the working tree and the index. When paths
// are specified, compares only those named paths. Otherwise all entries in the
// index are compared. The output format is the same as for git diff-index and
// git diff-tree.
//   git-diff-files - https://git-scm.com/docs/git-diff-files
func (p *Plugin) DiffFiles() herrors.ErrNotImplemented {
	return newErrNotImplemented("diff-files")
}

// DiffIndex Compares the content and mode of the blobs found in a tree object with the
// corresponding tracked files in the working tree, or with the corresponding paths
// in the index. When <path> arguments are present, compares only paths matching those
// patterns. Otherwise all tracked files are compared.
//   git-diff-index - https://git-scm.com/docs/git-diff-index
func (p *Plugin) DiffIndex() herrors.ErrNotImplemented {
	return newErrNotImplemented("diff-index")
}

// DiffTree compares the content and mode of the blobs found via two tree objects.
// If there is only one <tree-ish> given, the commit is compared with its parents
// (see --stdin below). Note that git diff-tree can use the tree encapsulated in
// a commit object.
//   git-diff-tree - https://git-scm.com/docs/git-diff-tree
func (p *Plugin) DiffTree() herrors.ErrNotImplemented {
	return newErrNotImplemented("diff-tree")
}

// Diff shows changes between the working tree and the index or a tree, changes
// between the index and a tree, changes between two trees, changes between two
// blob objects, or changes between two files on disk.
//   git-diff - https://git-scm.com/docs/git-diff
func (p *Plugin) Diff() herrors.ErrNotImplemented {
	return newErrNotImplemented("diff")
}

// Difftool it difftool is a Git command that allows you to compare and edit
// files between revisions using common diff tools. git difftool is a frontend
// to git diff and accepts the same options and arguments.
//   git-difftool -
func (p *Plugin) Difftool() herrors.ErrNotImplemented {
	return newErrNotImplemented("difftool")
}

// FastExport This program dumps the given revisions in a form suitable to be
// piped into git fast-import. You can use it as a human-readable bundle
// replacement (see git-bundle[1]), or as a kind of an interactive git filter-branch.
//   git-fast-export - https://git-scm.com/docs/git-fast-export
func (p *Plugin) FastExport() herrors.ErrNotImplemented {
	return newErrNotImplemented("fast-export")
}

// FastImport This program is usually not what the end user wants to run directly.
// Most end users want to use one of the existing frontend programs, which parses
// a specific type of foreign source and feeds the contents stored there to git fast-import.
//   git-fast-import - https://git-scm.com/docs/git-fast-import
func (p *Plugin) FastImport() herrors.ErrNotImplemented {
	return newErrNotImplemented("fast-import")
}

// FetchPack Usually you would want to use git fetch, which is a higher level
// wrapper of this command, instead. Invokes git-upload-pack on a possibly
// remote repository and asks it to send objects missing from this repository,
// to update the named heads. The list of commits available locally is found out
// by scanning the local refs/ hierarchy and sent to git-upload-pack running on the other en
//   git-fetch-pack - https://git-scm.com/docs/git-fetch-pack
func (p *Plugin) FetchPack() herrors.ErrNotImplemented {
	return newErrNotImplemented("fetch-pack")
}

// Fetch branches and/or tags (collectively, "refs") from one or more other
// repositories, along with the objects necessary to complete their histories.
// Remote-tracking branches are updated (see the description of <refspec>
// below for ways to control this behavior).
//   git-fetch - https://git-scm.com/docs/git-fetch
func (p *Plugin) Fetch() herrors.ErrNotImplemented {
	return newErrNotImplemented("fetch")
}

// FilterBranch Lets you rewrite Git revision history by rewriting the branches
// mentioned in the <rev-list options>, applying custom filters on each revision.
// Those filters can modify each tree (e.g. removing a file or running a perl
// rewrite on all files) or information about each commit. Otherwise, all
// information (including original commit times or merge information) will be preserved.
//   git-filter-branch - https://git-scm.com/docs/git-filter-branch
func (p *Plugin) FilterBranch() herrors.ErrNotImplemented {
	return newErrNotImplemented("filter-branch")
}

// FmtMergeMsg Takes the list of merged objects on stdin and produces a suitable
// commit message to be used for the merge commit, usually to be passed as the
// <merge-message> argument of git merge.
//   git-fmt-merge-msg - https://git-scm.com/docs/git-fmt-merge-msg
func (p *Plugin) FmtMergeMsg() herrors.ErrNotImplemented {
	return newErrNotImplemented("fmt-merge-msg")
}

// ForEachRef Iterate over all refs that match <pattern> and show them according
// to the given <format>, after sorting them according to the given set of <key>.
// If <count> is given, stop after showing that many refs. The interpolated
// values in <format> can optionally be quoted as string literals in the
// specified host language allowing their direct evaluation in that language.
//   git-for-each-ref - https://git-scm.com/docs/git-for-each-ref
func (p *Plugin) ForEachRef() herrors.ErrNotImplemented {
	return newErrNotImplemented("for-each-ref")
}

// FormatPatch prepare each commit with its patch in one file per commit,
// formatted to resemble UNIX mailbox format. The output of this command is
// convenient for e-mail submission or for use with git am.
//   git-format-patch - https://git-scm.com/docs/git-format-patch
func (p *Plugin) FormatPatch() herrors.ErrNotImplemented {
	return newErrNotImplemented("format-patch")
}

// FsckObjects  is a synonym for git-fsck[1].
// Please refer to the documentation of that command.
//   git-fsck-objects - https://git-scm.com/docs/git-fsck-objects
func (p *Plugin) FsckObjects() herrors.ErrDeprecated {
	return newErrDeprecated("fsck-objects", "fsck",
		"https://git-scm.com/docs/git-fsck")
}

// Fsck verifies the connectivity and validity of the objects in the database.
//   git-fsck - https://git-scm.com/docs/git-fsck
func (p *Plugin) Fsck() herrors.ErrNotImplemented {
	return newErrNotImplemented("fsck")
}

// GC Runs a number of housekeeping tasks within the current repository,
// such as compressing file revisions (to reduce disk space and increase
// performance) and removing unreachable objects which may have been created
// from prior invocations of git add.
//   git-gc - https://git-scm.com/docs/git-gc
func (p *Plugin) GC() herrors.ErrNotImplemented {
	return newErrNotImplemented("gc")
}

// GetTarCommitID Read a tar archive created by git archive from the standard
// input and extract the commit ID stored in it. It reads only the first 1024
// bytes of input, thus its runtime is not influenced by the size of the tar
// archive very much.
//   git-get-tar-commit-id - https://git-scm.com/docs/git-get-tar-commit-id
func (p *Plugin) GetTarCommitID() herrors.ErrNotImplemented {
	return newErrNotImplemented("get-tar-commit-id")
}

// Grep Look for specified patterns in the tracked files in the work tree,
// blobs registered in the index file, or blobs in given tree objects.
//  Patterns are lists of one or more search expressions separated by newline
// characters. An empty string as search expression matches all lines
//   git-grep - https://git-scm.com/docs/git-grep
func (p *Plugin) Grep() herrors.ErrNotImplemented {
	return newErrNotImplemented("grep")
}

// GUI A Tcl/Tk based graphical user interface to Git. git gui focuses on
// allowing users to make changes to their repository by making new commits,
// amending existing ones, creating branches, performing local merges, and
// fetching/pushing to remote repositories.
//   git-gui - https://git-scm.com/docs/git-gui
func (p *Plugin) GUI() herrors.ErrNotImplemented {
	return newErrNotImplemented("gui")
}

// HashObject computes the object ID value for an object with specified type
// with the contents of the named file (which can be outside of the work tree),
// and optionally writes the resulting object into the object database. Reports
// its object ID to its standard output. This is used by git cvsimport to update
// the index without modifying files in the work tree.
// When <type> is not specified, it defaults to "blob".
//   git-hash-object - https://git-scm.com/docs/git-hash-object
func (p *Plugin) HashObject() herrors.ErrNotImplemented {
	return newErrNotImplemented("hash-object")
}

// Help display help information about Git
//   git-help - https://git-scm.com/docs/git-help
func (p *Plugin) Help() herrors.ErrNotImplemented {
	return newErrNotImplemented("help")
}

// HTTPBackend A simple CGI program to serve the contents of a Git repository
// to Git clients accessing the repository over http:// and https:// protocols.
// The program supports clients fetching using both the smart HTTP protocol and
// the backwards-compatible dumb HTTP protocol, as well as clients pushing using
// the smart HTTP protocol.
//   git-http-backend - https://git-scm.com/docs/git-http-backend
func (p *Plugin) HTTPBackend() herrors.ErrNotImplemented {
	return newErrNotImplemented("http-backend")
}

// HTTPFetch Download from a remote Git repository via HTTP
//   git-http-fetch - https://git-scm.com/docs/git-http-fetch
func (p *Plugin) HTTPFetch() herrors.ErrNotImplemented {
	return newErrNotImplemented("http-fetch")
}

// HTTPPush push objects over HTTP/DAV to another repository
//   git-http-push - https://git-scm.com/docs/git-http-push
func (p *Plugin) HTTPPush() herrors.ErrNotImplemented {
	return newErrNotImplemented("http-push")
}

// ImapSend command uploads a mailbox generated with git format-patch into an
// IMAP drafts folder. This allows patches to be sent as other email is when
// using mail clients that cannot read mailbox files directly. The command
// also works with any general mailbox in which emails have the fields "From",
// "Date", and "Subject" in that order.
//   git format-patch --signoff --stdout --attach origin | git imap-send
//
//   git-imap-send - https://git-scm.com/docs/git-imap-send
func (p *Plugin) ImapSend() herrors.ErrNotImplemented {
	return newErrNotImplemented("imap-send")
}

// IndexPack Reads a packed archive (.pack) from the specified file, and builds
// a pack index file (.idx) for it. The packed archive together with the pack
// index can then be placed in the objects/pack/ directory of a Git repository.
//   git-index-pack - https://git-scm.com/docs/git-index-pack
func (p *Plugin) IndexPack() herrors.ErrNotImplemented {
	return newErrNotImplemented("index-pack")
}

// InitDB This is a synonym for git-init[1].
// Please refer to the documentation of that command.
//   git-init-db - https://git-scm.com/docs/git-init-db
func (p *Plugin) InitDB() herrors.ErrDeprecated {
	return newErrDeprecated("init-db", "init", "https://git-scm.com/docs/git-init-db")
}

// Init creates an empty Git repository or reinitialize an existing one
//   git-init - https://git-scm.com/docs/git-init
func (p *Plugin) Init() herrors.ErrNotImplemented {
	return newErrNotImplemented("init")
}

// Instaweb Instantly browse your working repository in gitweb
//   git-instaweb - https://git-scm.com/docs/git-instaweb
func (p *Plugin) Instaweb() herrors.ErrNotImplemented {
	return newErrNotImplemented("instaweb")
}

// InterpretTrailers help add structured information into commit messages.
// Help adding trailers lines, that look similar to RFC 822 e-mail headers,
// at the end of the otherwise free-form part of a commit message.
//   git-interpret-trailers - https://git-scm.com/docs/git-interpret-trailers
func (p *Plugin) InterpretTrailers() herrors.ErrNotImplemented {
	return newErrNotImplemented("interpret-trailers")
}

// Gitk Displays changes in a repository or a selected set of commits.
// This includes visualizing the commit graph, showing information related to
// each commit, and the files in the trees of each revision.
//   git-gitk - https://git-scm.com/docs/gitk
func (p *Plugin) Gitk() herrors.ErrNotImplemented {
	return newErrNotImplemented("gitk")
}

// Log shows the commit logs.
//   git-log - https://git-scm.com/docs/git-log
func (p *Plugin) Log(s ...string) (hexec.Output, error) {
	s = append([]string{"log"}, s...)
	return cmdgitInPath(p.fs.RealAbs(), s...)
}

// LsFiles show information about files in the index and the working tree.
// This merges the file listing in the directory cache index with the actual
// working directory list, and shows different combinations of the two.
//   git-ls-files - https://git-scm.com/docs/git-ls-files
func (p *Plugin) LsFiles() herrors.ErrNotImplemented {
	return newErrNotImplemented("ls-files")
}

// LsRemote displays references available in a remote repository
// along with the associated commit IDs.
//   git-ls-remote - https://git-scm.com/docs/git-ls-remote
func (p *Plugin) LsRemote() herrors.ErrNotImplemented {
	return newErrNotImplemented("ls-remote")
}

// LsTree lists the contents of a given tree object, like what "/bin/ls -a"
// does in the current working directory.
//   git-ls-tree - https://git-scm.com/docs/git-ls-tree
func (p *Plugin) LsTree() herrors.ErrNotImplemented {
	return newErrNotImplemented("ls-tree")
}

// Mailinfo extracts patch and authorship from a single e-mail message.
// reads a single e-mail message from the standard input, and writes the commit
// log message in <msg> file, and the patches in <patch> file.
//   git-mailinfo - https://git-scm.com/docs/git-mailinfo
func (p *Plugin) Mailinfo() herrors.ErrNotImplemented {
	return newErrNotImplemented("mailinfo")
}

// Mailsplit Simple UNIX mbox splitter program. Splits a mbox file or a Maildir
// into a list of files: "0001" "0002" .. in the specified directory so you can
// process them further from there.
//   git-mailsplit - https://git-scm.com/docs/git-mailsplit
func (p *Plugin) Mailsplit() herrors.ErrNotImplemented {
	return newErrNotImplemented("mailsplit")
}

// MergeBase finds best common ancestor(s) between two commits to use in a
// three-way merge. One common ancestor is better than another common ancestor
// if the latter is an ancestor of the former. A common ancestor that does not
// have any better common ancestor is a best common ancestor, i.e. a merge base.
// Note that there can be more than one merge base for a pair of commits.
//   git-merge-base -
func (p *Plugin) MergeBase() herrors.ErrNotImplemented {
	return newErrNotImplemented("merge-base")
}

// MergeFile runs a three-way file merge.
//   git-merge-file - https://git-scm.com/docs/git-merge-file
func (p *Plugin) MergeFile() herrors.ErrNotImplemented {
	return newErrNotImplemented("merge-file")
}

// MergeIndex run a merge for files needing merging.
//   git-merge-index - https://git-scm.com/docs/git-merge-index
func (p *Plugin) MergeIndex() herrors.ErrNotImplemented {
	return newErrNotImplemented("merge-index")
}

// MergeOctopus common merge strategy, resolves cases with more than two heads,
// but refuses to do a complex merge that needs manual resolution
//   git-merge-octopus - https://git-scm.com/docs/git-merge
func (p *Plugin) MergeOctopus() herrors.ErrNotImplemented {
	return newErrNotImplemented("merge-octopus")
}

// MergeOneFile is the standard helper program to use with git merge-index to
// resolve a merge after the trivial merge done with git read-tree -m.
//   git-merge-one-file - https://git-scm.com/docs/git-merge-one-file
func (p *Plugin) MergeOneFile() herrors.ErrNotImplemented {
	return newErrNotImplemented("merge-one-file")
}

// MergeOurs resolves any number of heads, but the resulting tree of the merge
// is always that of the current branch head, effectively ignoring all changes
// from all other branches
//   git-merge-ours - https://git-scm.com/docs/git-merge
func (p *Plugin) MergeOurs() herrors.ErrNotImplemented {
	return newErrNotImplemented("merge-ours")
}

// MergeRecursive can only resolve two heads using a 3-way merge algorithm.
//   git-merge-recursive - https://git-scm.com/docs/git-merge
func (p *Plugin) MergeRecursive() herrors.ErrNotImplemented {
	return newErrNotImplemented("merge-recursive")
}

// MergeResolve This can only resolve two heads (i.e. the current branch and
// another branch you pulled from) using a 3-way merge algorithm.
//   git-merge-resolve - https://git-scm.com/docs/git-merge
func (p *Plugin) MergeResolve() herrors.ErrNotImplemented {
	return newErrNotImplemented("merge-resolve")
}

// MergeSubtree This is a modified recursive strategy. When merging trees A and B
//   git-merge-subtree - https://git-scm.com/docs/git-merge
func (p *Plugin) MergeSubtree() herrors.ErrNotImplemented {
	return newErrNotImplemented("merge-subtree")
}

// MergeTree Reads three tree-ish, and output trivial merge results and conflicting
// stages to the standard output. This is similar to what three-way git read-tree
// -m does, but instead of storing the results in the index, the command outputs
// the entries to the standard output.
//   git-merge-tree -
func (p *Plugin) MergeTree() herrors.ErrNotImplemented {
	return newErrNotImplemented("merge-tree")
}

// Merge Incorporates changes from the named commits (since the time their histories
// diverged from the current branch) into the current branch. This command is used
// by git pull to incorporate changes from another repository and can be used by
// hand to merge changes from one branch into another.
//   git-merge - https://git-scm.com/docs/git-merge
func (p *Plugin) Merge() herrors.ErrNotImplemented {
	return newErrNotImplemented("merge")
}

// Mergetool Run merge conflict resolution tools to resolve merge conflicts.
//   git-mergetool - https://git-scm.com/docs/git-mergetool
func (p *Plugin) Mergetool() herrors.ErrNotImplemented {
	return newErrNotImplemented("mergetool")
}

// Mktag reads a tag contents on standard input and creates a tag object that
// can also be used to sign other objects.
//   git-mktag - https://git-scm.com/docs/git-mktag
func (p *Plugin) Mktag() herrors.ErrNotImplemented {
	return newErrNotImplemented("mktag")
}

// Mktree Build a tree-object from ls-tree formatted text. Reads standard input
// in non-recursive ls-tree output format, and creates a tree object. The order
// of the tree entries is normalised by mktree so pre-sorting the input is not
// required. The object name of the tree object built is written to the standard output.
//   git-mktree - https://git-scm.com/docs/git-mktree
func (p *Plugin) Mktree() herrors.ErrNotImplemented {
	return newErrNotImplemented("mktree")
}

// Mv move or rename a file, a directory, or a symlink.
//   git-mv - https://git-scm.com/docs/git-mv
func (p *Plugin) Mv() herrors.ErrNotImplemented {
	return newErrNotImplemented("mv")
}

// NameRev finds symbolic names suitable for human digestion for revisions given
// in any format parsable by git rev-parse.
//   git-name-rev - https://git-scm.com/docs/git-name-rev
func (p *Plugin) NameRev() herrors.ErrNotImplemented {
	return newErrNotImplemented("name-rev")
}

// Notes adds, removes, or reads notes attached to objects,
// without touching the objects themselves.
//   git-notes - https://git-scm.com/docs/git-notes
func (p *Plugin) Notes() herrors.ErrNotImplemented {
	return newErrNotImplemented("notes")
}

// P4 Import from and submit to Perforce repositories
//   git-p4 - https://git-scm.com/docs/git-p4
func (p *Plugin) P4() herrors.ErrNotImplemented {
	return newErrNotImplemented("p4")
}

// PackObjects reads list of objects from the standard input, and writes a packed
// archive with specified base-name, or to the standard output.
//   git-pack-objects - https://git-scm.com/docs/git-pack-objects
func (p *Plugin) PackObjects() herrors.ErrNotImplemented {
	return newErrNotImplemented("pack-objects")
}

// PackRedundant computes which packs in your repository are redundant.
// The output is suitable for piping to xargs rm if you are in the root of the repositor
//   git-pack-redundant - https://git-scm.com/docs/git-pack-redundant
func (p *Plugin) PackRedundant() herrors.ErrNotImplemented {
	return newErrNotImplemented("pack-redundant")
}

// PackRefs Traditionally, tips of branches and tags (collectively known as refs)
// were stored one file per ref in a (sub)directory under $GIT_DIR/refs directory.
//   git-pack-refs - https://git-scm.com/docs/git-pack-refs
func (p *Plugin) PackRefs() herrors.ErrNotImplemented {
	return newErrNotImplemented("pack-refs")
}

// ParseRemote routines to help parsing remote repository access parameters.
//   git-parse-remote - https://git-scm.com/docs/git-parse-remote
func (p *Plugin) ParseRemote() herrors.ErrNotImplemented {
	return newErrNotImplemented("parse-remote")
}

// PatchID read a patch from the standard input and compute the patch ID for it.
//   git-patch-id - https://git-scm.com/docs/git-patch-id
func (p *Plugin) PatchID() herrors.ErrNotImplemented {
	return newErrNotImplemented("patch-id")
}

// Prune all unreachable objects from the object database
//   git-prune - https://git-scm.com/docs/git-prune
func (p *Plugin) Prune() herrors.ErrNotImplemented {
	return newErrNotImplemented("prune")
}

// PrunePacked removes extra objects that are already in pack files
//   git-prune-packed - https://git-scm.com/docs/git-prune-packed
func (p *Plugin) PrunePacked() herrors.ErrNotImplemented {
	return newErrNotImplemented("prune-packed")
}

// Pull fetch from and integrate with another repository or a local branch
//   git-pull - https://git-scm.com/docs/git-pull
func (p *Plugin) Pull() herrors.ErrNotImplemented {
	return newErrNotImplemented("pull")
}

// Push updates remote refs using local refs, while sending objects necessary
// to complete the given refs.
//   git-push - https://git-scm.com/docs/git-push
func (p *Plugin) Push() herrors.ErrNotImplemented {
	return newErrNotImplemented("push")
}

// Quiltimport Applies a quilt patchset onto the current Git branch, preserving
// the patch boundaries, patch order, and patch descriptions present in the quilt patchset.
//   git-quiltimport - https://git-scm.com/docs/git-quiltimport
func (p *Plugin) Quiltimport() herrors.ErrNotImplemented {
	return newErrNotImplemented("quiltimport")
}

// ReadTree Reads the tree information given by <tree-ish> into the index,
// but does not actually update any of the files it "caches". (see: git-checkout-index[1])
//   git-read-tree - https://git-scm.com/docs/git-read-tree
func (p *Plugin) ReadTree() herrors.ErrNotImplemented {
	return newErrNotImplemented("read-tree")
}

// Rebase reapply commits on top of another base tip
//   git-rebase - https://git-scm.com/docs/git-rebase
func (p *Plugin) Rebase() herrors.ErrNotImplemented {
	return newErrNotImplemented("rebase")
}

// ReceivePack invoked by git send-pack and updates the repository with the
// information fed from the remote end.
//   git-receive-pack - https://git-scm.com/docs/git-receive-pack
func (p *Plugin) ReceivePack() herrors.ErrNotImplemented {
	return newErrNotImplemented("receive-pack")
}

// Reflog the command takes various subcommands, and different options
// depending on the subcommand:
//   git-reflog - https://git-scm.com/docs/git-reflog
func (p *Plugin) Reflog() herrors.ErrNotImplemented {
	return newErrNotImplemented("reflog")
}

// RemoteExt Bridge smart transport to external command.Data written to stdin of
// the specified <command> is assumed to be sent to a git:// server, git-upload-pack,
// git-receive-pack or git-upload-archive (depending on situation), and data read
// from stdout of <command> is assumed to be received from the same service.
//   git-remote-ext - https://git-scm.com/docs/git-remote-ext
func (p *Plugin) RemoteExt() herrors.ErrNotImplemented {
	return newErrNotImplemented("remote-ext")
}

// RemoteFd This helper uses specified file descriptors to connect to a remote
// Git server. This is not meant for end users but for programs and scripts
// calling git fetch, push or archive.
//   git-remote-fd - https://git-scm.com/docs/git-remote-fd
func (p *Plugin) RemoteFd() herrors.ErrNotImplemented {
	return newErrNotImplemented("remote-fd")
}

// RemoteFtp ftp
//   git-remote-ftp -
func (p *Plugin) RemoteFtp() herrors.ErrNotImplemented {
	return newErrNotImplemented("remote-ftp")
}

// RemoteFtps ftps
//   git-remote-ftps -
func (p *Plugin) RemoteFtps() herrors.ErrNotImplemented {
	return newErrNotImplemented("remote-ftps")
}

// RemoteHTTP http
//   git-remote-http -
func (p *Plugin) RemoteHTTP() herrors.ErrNotImplemented {
	return newErrNotImplemented("remote-http")
}

// RemoteHTTPS https
//   git-remote-https -
func (p *Plugin) RemoteHTTPS() herrors.ErrNotImplemented {
	return newErrNotImplemented("remote-https")
}

// RemoteTestSvn svn
//   git-remote-testsvn -
func (p *Plugin) RemoteTestSvn() herrors.ErrNotImplemented {
	return newErrNotImplemented("remote-testsvn")
}

// RemoteTestGit git
//   git-remote-testgit -
func (p *Plugin) RemoteTestGit() herrors.ErrNotImplemented {
	return newErrNotImplemented("remote-testgit")
}

// Remote manage the set of repositories ("remotes") whose branches you track.
//   git-remote - https://git-scm.com/docs/git-remote
func (p *Plugin) Remote() herrors.ErrNotImplemented {
	return newErrNotImplemented("remote")
}

// Repack This command is used to combine all objects that do not currently
// reside in a "pack", into a pack. It can also be used to re-organize existing
// packs into a single, more efficient pack.
//   git-repack - https://git-scm.com/docs/git-repack
func (p *Plugin) Repack() herrors.ErrNotImplemented {
	return newErrNotImplemented("repack")
}

// Replace create, list, delete refs to replace objects
//   git-replace - https://git-scm.com/docs/git-replace
func (p *Plugin) Replace() herrors.ErrNotImplemented {
	return newErrNotImplemented("replace")
}

// RequestPull generates a summary of pending changes
//   git-request-pull - Generate a request asking your upstream project to pull
// changes into their tree. The request, printed to the standard output, begins
// with the branch description, summarizes the changes and indicates from where
// they can be pulled.
//   git-request-pull - https://git-scm.com/docs/git-request-pull
func (p *Plugin) RequestPull() herrors.ErrNotImplemented {
	return newErrNotImplemented("request-pull")
}

// Rerere Reuse recorded resolution of conflicted merges
//   git-rerere -
func (p *Plugin) Rerere() herrors.ErrNotImplemented {
	return newErrNotImplemented("rerere")
}

// Reset current HEAD to the specified state
//   git-reset - https://git-scm.com/docs/git-reset
func (p *Plugin) Reset() herrors.ErrNotImplemented {
	return newErrNotImplemented("reset")
}

// RevList Lists commit objects in reverse chronological order
//   git-rev-list - https://git-scm.com/docs/git-rev-list
func (p *Plugin) RevList() herrors.ErrNotImplemented {
	return newErrNotImplemented("rev-list")
}

// RevParse Many Git porcelainish commands take mixture of flags (i.e.
// parameters that begin with a dash -) and parameters meant for the underlying
// git rev-list command they use internally and flags and parameters for the other
// commands they use downstream of git rev-list.
// This command is used to distinguish between them.
//   git-rev-parse - https://git-scm.com/docs/git-rev-parse
func (p *Plugin) RevParse() herrors.ErrNotImplemented {
	return newErrNotImplemented("rev-parse")
}

// Revert Given one or more existing commits, revert the changes that the related
// patches introduce, and record some new commits that record them. This requires
// your working tree to be clean (no modifications from the HEAD commit).
//   git-revert -
func (p *Plugin) Revert() herrors.ErrNotImplemented {
	return newErrNotImplemented("revert")
}

// Rm Remove files from the index, or from the working tree and the index. git
// rm will not remove a file from just your working directory.
//   git-rm - https://git-scm.com/docs/git-rm
func (p *Plugin) Rm() herrors.ErrNotImplemented {
	return newErrNotImplemented("rm")
}

// SendEmail Send a collection of patches as emails
//   git-send-email - https://git-scm.com/docs/git-send-email
func (p *Plugin) SendEmail() herrors.ErrNotImplemented {
	return newErrNotImplemented("send-email")
}

// SendPack push objects over Git protocol to another repository. Usually you
// would want to use git push, which is a higher-level wrapper of this command,
//instead. See git-push[1].
//   git-send-pack - https://git-scm.com/docs/git-send-pack
func (p *Plugin) SendPack() herrors.ErrNotImplemented {
	return newErrNotImplemented("send-pack")
}

// Shell Restricted login shell for Git-only SSH access
//   git-shell - https://git-scm.com/docs/git-shell
func (p *Plugin) Shell() herrors.ErrNotImplemented {
	return newErrNotImplemented("shell")
}

// Shortlog summarizes git log output in a format suitable for inclusion in
// release announcements. Each commit will be grouped by author and title.
//   git-shortlog - https://git-scm.com/docs/git-shortlog
func (p *Plugin) Shortlog() herrors.ErrNotImplemented {
	return newErrNotImplemented("shortlog")
}

// ShowBranch Show branches and their commits. Shows the commit ancestry graph
// starting from the commits named with <rev>s or <globs>s (or all refs under
// refs/heads and/or refs/tags) semi-visually.
//   git-show-branch - https://git-scm.com/docs/git-show-branch
func (p *Plugin) ShowBranch() herrors.ErrNotImplemented {
	return newErrNotImplemented("show-branch")
}

// ShowIndex Show packed archive index
//   git-show-index - https://git-scm.com/docs/git-show-index
func (p *Plugin) ShowIndex() herrors.ErrNotImplemented {
	return newErrNotImplemented("show-index")
}

// ShowRef List references in a local repository
//   git-show-ref - https://git-scm.com/docs/git-show-ref
func (p *Plugin) ShowRef() herrors.ErrNotImplemented {
	return newErrNotImplemented("show-ref")
}

// Show  one or more objects (blobs, trees, tags and commits).
//   git-show - https://git-scm.com/docs/git-show
func (p *Plugin) Show() herrors.ErrNotImplemented {
	return newErrNotImplemented("show")
}

// Stage Add file contents to the staging area
//   git-stage - https://git-scm.com/docs/git-stage
func (p *Plugin) Stage() herrors.ErrNotImplemented {
	return newErrNotImplemented("stage")
}

// Stash Stash the changes in a dirty working directory away
//   git-stash - https://git-scm.com/docs/git-stash
func (p *Plugin) Stash() herrors.ErrNotImplemented {
	return newErrNotImplemented("stash")
}

// Status Displays paths that have differences between the index file and the
// current HEAD commit, paths that have differences between the working tree and
// the index file, and paths in the working tree that are not tracked by Git
// (and are not ignored by gitignore[5]).
//   git-status - https://git-scm.com/docs/git-status
func (p *Plugin) Status() herrors.ErrNotImplemented {
	return newErrNotImplemented("status")
}

// Stripspace Read text, such as commit messages, notes, tags and branch descriptions,
// from the standard input and clean it in the manner used by Git.
//   git-stripspace -
func (p *Plugin) Stripspace() herrors.ErrNotImplemented {
	return newErrNotImplemented("stripspace")
}

// Submodule Inspects, updates and manages submodules.
// A submodule allows you to keep another Git repository in a subdirectory of
// your repository. The other repository has its own history, which does not
// interfere with the history of the current repository. This can be used to have
// external dependencies such as third party libraries for example.
//   git-submodule -
func (p *Plugin) Submodule() herrors.ErrNotImplemented {
	return newErrNotImplemented("submodule")
}

// Svn Bidirectional operation between a Subversion repository and Git.
// git svn is a simple conduit for changesets between Subversion and Git.
// It provides a bidirectional flow of changes between a Subversion and a Git repository.
//   git-svn -
func (p *Plugin) Svn() herrors.ErrNotImplemented {
	return newErrNotImplemented("svn")
}

// SymbolicRef Read, modify and delete symbolic refs.
// Given one argument, reads which branch head the given symbolic ref refers to
// and outputs its path, relative to the .git/ directory. Typically you would give
// HEAD as the <name> argument to see which branch your working tree is on.
//   git-symbolic-ref - https://git-scm.com/docs/git-symbolic-ref
func (p *Plugin) SymbolicRef() herrors.ErrNotImplemented {
	return newErrNotImplemented("symbolic-ref")
}

// Tag Create, list, delete or verify a tag object signed with GPG.
//   git-tag - https://git-scm.com/docs/git-tag
func (p *Plugin) Tag() herrors.ErrNotImplemented {
	return newErrNotImplemented("tag")
}

// UnpackFile Creates a temporary file with a blob’s contents. Creates a file
// holding the contents of the blob specified by sha1. It returns the name of
// the temporary file in the following format: .merge_file_XXXXX
//   git-unpack-file - https://git-scm.com/docs/git-unpack-file
func (p *Plugin) UnpackFile() herrors.ErrNotImplemented {
	return newErrNotImplemented("unpack-file")
}

// UnpackObjects Unpack objects from a packed archive. Read a packed archive
// (.pack) from the standard input, expanding the objects contained within and
// writing them into the repository in "loose" (one object per file) format.
//   git-unpack-objects - https://git-scm.com/docs/git-unpack-objects
func (p *Plugin) UnpackObjects() herrors.ErrNotImplemented {
	return newErrNotImplemented("unpack-objects")
}

// UpdateIndex Register file contents in the working tree to the index.
// Modifies the index or directory cache. Each file mentioned is updated into
// the index and any unmerged or needs updating state is cleared.
//   git-update-index - https://git-scm.com/docs/git-update-index
func (p *Plugin) UpdateIndex() herrors.ErrNotImplemented {
	return newErrNotImplemented("update-index")
}

// UpdateRef Update the object name stored in a ref safely
//   git-update-ref - https://git-scm.com/docs/git-update-ref
func (p *Plugin) UpdateRef() herrors.ErrNotImplemented {
	return newErrNotImplemented("update-ref")
}

// UpdateServerInfo Update auxiliary info file to help dumb servers.
// A dumb server that does not do on-the-fly pack generations must have some
// auxiliary information files in $GIT_DIR/info and $GIT_OBJECT_DIRECTORY/info
// directories to help clients discover what references and packs the server has.
// This command generates such auxiliary files
//   git-update-server-info -
func (p *Plugin) UpdateServerInfo() herrors.ErrNotImplemented {
	return newErrNotImplemented("update-server-info")
}

// UploadArchive Send archive back to git-archive.
// Invoked by git archive --remote and sends a generated archive to the other
// end over the Git protocol.
//   git-upload-archive - https://git-scm.com/docs/git-upload-archive
func (p *Plugin) UploadArchive() herrors.ErrNotImplemented {
	return newErrNotImplemented("upload-archive")
}

// UploadPack Send objects packed back to git-fetch-pack.
// This command is usually not invoked directly by the end user.
// The UI for the protocol is on the git fetch-pack side, and the program pair
// is meant to be used to pull updates from a remote repository.
// For push operations, see git send-pack.
//   git-upload-pack - https://git-scm.com/docs/git-upload-pack
func (p *Plugin) UploadPack() herrors.ErrNotImplemented {
	return newErrNotImplemented("upload-pack")
}

// Var Show a Git logical variable. Cause the logical variables to be listed.
// In addition, all the variables of the Git configuration file .git/config are
// listed as well. (However, the configuration variables listing functionality
// is deprecated in favor of git config -l.)
//   git-var -
func (p *Plugin) Var() herrors.ErrNotImplemented {
	return newErrNotImplemented("var")
}

// VerifyCommit Check the GPG signature of commits. Validates the GPG signature
// created by git commit -S.
//   git-verify-commit - https://git-scm.com/docs/git-verify-commit
func (p *Plugin) VerifyCommit() herrors.ErrNotImplemented {
	return newErrNotImplemented("verify-commit")
}

// VerifyPack Reads given idx file for packed Git archive created with the git
// pack-objects command and verifies idx file and the corresponding pack file.
//   git-verify-pack - https://git-scm.com/docs/git-verify-pack
func (p *Plugin) VerifyPack() herrors.ErrNotImplemented {
	return newErrNotImplemented("verify-pack")
}

// VerifyTag Check the GPG signature of tags. Validates the gpg signature created by git tag.
//   git-verify-tag - https://git-scm.com/docs/git-verify-tag
func (p *Plugin) VerifyTag() herrors.ErrNotImplemented {
	return newErrNotImplemented("verify-tag")
}

// Whatchanged Show logs with difference each commit introduces
//   git-whatchanged - https://git-scm.com/docs/git-whatchanged
func (p *Plugin) Whatchanged() herrors.ErrNotImplemented {
	return newErrNotImplemented("whatchanged")
}

// Worktree Manage multiple working trees attached to the same repository.
//   git-worktree - https://git-scm.com/docs/git-worktree
func (p *Plugin) Worktree() herrors.ErrNotImplemented {
	return newErrNotImplemented("worktree")
}

// WriteTree Creates a tree object using the current index.
//  The name of the new tree object is printed to standard output.
//   git-write-tree - https://git-scm.com/docs/git-write-tree
func (p *Plugin) WriteTree() herrors.ErrNotImplemented {
	return newErrNotImplemented("write-tree")
}
