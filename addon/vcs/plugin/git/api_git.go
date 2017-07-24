package git

import (
	"os/exec"
	"strings"

	"github.com/howi-ce/howi/pkg/std/herrors"
	"github.com/howi-ce/howi/pkg/wdfs"
)

// OpenRepositoryPath returns git instance
func OpenRepositoryPath(wd string) (*Git, error) {
	p, err := wdfs.LoadPath(wd)
	if err != nil {
		return nil, err
	}
	if p.IsGitRepository() {
		return &Git{wd: p}, err
	}

	for !p.IsGitRepository() {
		p, err = wdfs.LoadPath(p.Join("../"))
		if err != nil {
			return nil, err
		}
		if p.Abs() == "/" {
			return nil, herrors.New("Not a git repository (or any parent up to mount point /)")
		}
		if p.IsGitRepository() {
			return &Git{wd: p}, err
		}
	}
	return nil, herrors.New("failed to load git repository")
}

// Git global
type Git struct {
	wd wdfs.Path
}

// Apply a patch to files and/or to the index. Reads the supplied diff output
// (i.e. "a patch") and applies it to files. When running from a subdirectory in
// a repository, patched paths outside the directory are ignored. It apply the
// patch to files, and does not require them to be in a Git repository.
//   git-apply - https://git-scm.com/docs/git-apply
func (g *Git) Apply() ErrNotImplemented {
	return errNotImplemented("apply")
}

// CheckRefFormat ensures that a reference name is well formed
//   check-ref-format - https://git-scm.com/docs/git-check-ref-format
func (g *Git) CheckRefFormat() ErrNotImplemented {
	return errNotImplemented("check-ref-format")
}

// Clone a repository into a newly created directory, creates remote-tracking
// branches for each branch in the cloned repository
//   (visible using git branch -r),
// and creates and checks out an initial branch that is forked from the cloned
// repository’s currently active branch.
//   git-clone - https://git-scm.com/docs/git-clone
func (g *Git) Clone() ErrNotImplemented {
	return errNotImplemented("clone")
}

// Column displays data in columns
//   git-column - https://git-scm.com/docs/git-column
func (g *Git) Column() ErrNotImplemented {
	return errNotImplemented("column")
}

// Add file contents to the index using the current content found in the working
// tree, to prepare the content staged for the next commit.
//   git-add - https://git-scm.com/docs/git-add
func (g *Git) Add() ErrNotImplemented {
	return errNotImplemented("add")
}

// Am splits mail messages in a mailbox into commit log message, authorship
// information and patches, and applies them to the current branch.
//   git-am - https://git-scm.com/docs/git-am
func (g *Git) Am() ErrNotImplemented {
	return errNotImplemented("am")
}

// Annotate (deprecated)
//   git-annotate - https://git-scm.com/docs/git-annotate
func (g *Git) Annotate() ErrDeprecated {
	return errDeprecated("annotate", "blame", "https://git-scm.com/docs/git-annotate")
}

// Archimport imports a project from one or more Arch repositories. It will
// follow branches and repositories within the namespaces defined by the
// <archive/branch> parameters supplied.
//   git-archimport - https://git-scm.com/docs/git-archimport
func (g *Git) Archimport() ErrNotImplemented {
	return errNotImplemented("archimport")
}

// Archive creates an archive of the specified format containing the tree
// structure for the named tree, and writes it out to the standard output.
// If <prefix> is specified it is prepended to the filenames in the archive.
//   git-archive - https://git-scm.com/docs/git-archive
func (g *Git) Archive() ErrNotImplemented {
	return errNotImplemented("am")
}

// Bisect binary search to find the commit that introduced a bug.
//   git-bisect - https://git-scm.com/docs/git-bisect
func (g *Git) Bisect() ErrNotImplemented {
	return errNotImplemented("am")
}

// Blame annotates each line in the given file with information from the
// revision which last modified the line.
//   git-blame - https://git-scm.com/docs/git-blame
func (g *Git) Blame() ErrNotImplemented {
	return errNotImplemented("am")
}

// Branch list, create, or delete branches.
//   git-branch - https://git-scm.com/docs/git-branch
func (g *Git) Branch() ErrNotImplemented {
	return errNotImplemented("branch")
}

// Bundle moves objects and refs by archive
//   git-bundle - https://git-scm.com/docs/git-bundle
func (g *Git) Bundle() ErrNotImplemented {
	return errNotImplemented("bundle")
}

// CatFile provides content or type and size information for repository objects
//   git-cat-file - https://git-scm.com/docs/git-cat-file
func (g *Git) CatFile() ErrNotImplemented {
	return errNotImplemented("cat-file")
}

// CheckAttr displays gitattributes information
//   git-check-attr - https://git-scm.com/docs/git-check-attr
func (g *Git) CheckAttr() ErrNotImplemented {
	return errNotImplemented("check-attr")
}

// CheckIgnore enables you to debug gitignore / exclude files.
//   git-check-ignore - https://git-scm.com/docs/git-check-ignore
func (g *Git) CheckIgnore() ErrNotImplemented {
	return errNotImplemented("check-ignore")
}

// CheckMailmap show canonical names and email addresses of contacts
//   git-check-mailmap - https://git-scm.com/docs/git-check-mailmap
func (g *Git) CheckMailmap() ErrNotImplemented {
	return errNotImplemented("check-mailmap")
}

// Checkout switch branches or restore working tree files. Updates files in
// the working tree to match the version in the index or the specified tree.
// If no paths are given, git checkout will also update HEAD to set
// the specified branch as the current branch.
//   git-checkout - https://git-scm.com/docs/git-checkout
func (g *Git) Checkout() ErrNotImplemented {
	return errNotImplemented("checkout")
}

// CheckoutIndex will copy all files listed from the index to the working
// directory (not overwriting existing files).
//   git-checkout-index - https://git-scm.com/docs/git-checkout-index
func (g *Git) CheckoutIndex() ErrNotImplemented {
	return errNotImplemented("checkout-index")
}

// Cherry finds commits yet to be applied to upstream. Determine whether there
// are commits in
//   <head>..<upstream>
// that are equivalent to those in the range
//   <limit>..<head>
//
//  git-cherry - https://git-scm.com/docs/git-cherry .
func (g *Git) Cherry() ErrNotImplemented {
	return errNotImplemented("cherry")
}

// CherryPick applies the changes introduced by some existing commits. Given one
// or more existing commits, apply the change each one introduces,recording a
// new commit for each. This requires your working tree to be clean
// (no modifications from the HEAD commit).
//   git-cherry-pick - https://git-scm.com/docs/git-cherry-pick
func (g *Git) CherryPick() ErrNotImplemented {
	return errNotImplemented("cherry-pick")
}

// Clean the working tree by recursively removing files that are not under
// version control, starting from the current directory. Normally, only files
// unknown to Git are removed, but if the -x option is specified, ignored files
// are also removed. This can, for example,
// be useful to remove all build products. If any optional <path>... arguments
// are given, only those paths are affected.
//   git-clean https://git-scm.com/docs/git-clean
func (g *Git) Clean() ErrNotImplemented {
	return errNotImplemented("clean")
}

// Citool Graphical alternative to git-commit is removed from this library.
//   git-citool - https://git-scm.com/docs/git-citool
func (g *Git) Citool() ErrDeprecated {
	return errDeprecated("citool", "none", "https://git-scm.com/docs/git-citool")
}

// Commit records changes to the repository. Stores the current contents of the
// index in a new commit along with a log message from the user describing
// the changes.
//   git-commit - https://git-scm.com/docs/git-commit
func (g *Git) Commit() ErrNotImplemented {
	return errNotImplemented("commit")
}

// CommitTree creates a new commit object. This is usually not what an end
// user wants to run directly. See .Commmit instead.
//   git-commit-tree - https://git-scm.com/docs/git-commit-tree
func (g *Git) CommitTree() ErrNotImplemented {
	return errNotImplemented("commit-tree")
}

// Config get and set repository or global options. You can query/set/replace/unset
// options with this command. The name is actually the section and the key separated
// by a dot, and the value will be escaped.
//   git-config - https://git-scm.com/docs/git-config
func (g *Git) Config(flags ...string) ([]string, error) {

	flags = append([]string{"config"}, flags...)
	gitconfig, err := exec.Command("git", flags...).Output()
	return strings.Split(string(gitconfig), "\n"), err
}

// CountObjects counts unpacked number of objects and their disk consumption.
// This counts the number of unpacked object files and disk space consumed by them,
// to help you decide when it is a good time to repack.
//   git-count-objects - https://git-scm.com/docs/git-count-objects
func (g *Git) CountObjects() ErrNotImplemented {
	return errNotImplemented("count-objects")
}

// Credential retrieve and store user credentials. Git has an internal interface
// for storing and retrieving credentials from system-specific helpers, as well
// as prompting the user for usernames and passwords. The git-credential command
// exposes this interface to scripts which may want to retrieve, store, or prompt
// for credentials in the same manner as Git. The design of this scriptable
// interface models the internal C API; see the Git credential API for more
// background on the concepts.
//   git-credential - https://git-scm.com/docs/git-credential
func (g *Git) Credential() ErrNotImplemented {
	return errNotImplemented("credential")
}

// CredentialCache caches credentials in memory for use by future Git programs.
// The stored credentials never touch the disk, and are forgotten after a
// configurable timeout. The cache is accessible over a Unix domain socket,
// restricted to the current user by filesystem permissions.
//   git-credential-cache - https://git-scm.com/docs/git-credential-cache
func (g *Git) CredentialCache() ErrNotImplemented {
	return errNotImplemented("credential")
}

// CredentialCacheDaemon This command listens on the Unix domain socket specified
// by <socket> for git-credential-cache clients. Clients may store and retrieve credentials.
// You probably don’t want to invoke this command yourself;
// it is started automatically when you use git-credential-cache[1].
//   git-credential-cache--daemon - https://git-scm.com/docs/git-credential-cache--daemon
func (g *Git) CredentialCacheDaemon() ErrDeprecated {
	return errDeprecated("credential-cache--daemon", "credential-cache",
		"https://git-scm.com/docs/git-credential-cache--daemon")
}

// CredentialStore Using this helper will store your passwords unencrypted on disk,
// protected only by filesystem permissions. If this is not an acceptable
// security tradeoff, try git-credential-cache[1], or find a helper that
// integrates with secure storage provided by your operating system.
// This command stores credentials indefinitely on disk for use by future Git programs.
//   git-credential-store - https://git-scm.com/docs/git-credential-store
func (g *Git) CredentialStore() ErrNotImplemented {
	return errNotImplemented("credential-store")
}

// CredentialLibsecret helper that talks via libsecret with
// implementations of XDG Secret Service API
//   git-credential-libsecret - #
func (g *Git) CredentialLibsecret() ErrNotImplemented {
	return errNotImplemented("credential-libsecret")
}

// CredentialNetrc helper credential helper
//   git-credential-netrc - #
func (g *Git) CredentialNetrc() ErrNotImplemented {
	return errNotImplemented("credential-netrc")
}

// Cvsexportcommit Exports a commit from Git to a CVS checkout, making it easier
// to merge patches from a Git repository into a CVS repository.
//   git-cvsexportcommit - https://git-scm.com/docs/git-cvsexportcommit
func (g *Git) Cvsexportcommit() ErrNotImplemented {
	return errNotImplemented("cvsexportcommit")
}

// Cvsimport  Salvage your data out of another SCM people love to hate.
// cvsps version 2 is deprecated.
//   git-cvsimport - https://git-scm.com/docs/git-cvsimport
func (g *Git) Cvsimport() ErrDeprecated {
	return errDeprecated("cvsimport", "", "https://git-scm.com/docs/git-cvsimport")
}

// Cvsserver A CVS server emulator for Git
//   git-cvsserver - https://git-scm.com/docs/git-cvsserver
func (g *Git) Cvsserver() ErrNotImplemented {
	return errNotImplemented("cvsserver")
}

// Daemon is a really simple TCP Git daemon that normally listens on port
// "DEFAULT_GIT_PORT" aka 9418. It waits for a connection asking for a service,
// and will serve that service if it is enabled.
//   git-daemon - https://git-scm.com/docs/git-daemon
func (g *Git) Daemon() ErrNotImplemented {
	return errNotImplemented("daemon")
}

// Describe command finds the most recent tag that is reachable from a commit.
// If the tag points to the commit, then only the tag is shown. Otherwise,
// it suffixes the tag name with the number of additional commits on top of the
// tagged object and the abbreviated object name of the most recent commit.
//   git-describe - https://git-scm.com/docs/git-describe
func (g *Git) Describe(s ...string) (*Output, error) {
	s = append([]string{"describe"}, s...)
	return cmdgitInPath(g.wd.Abs(), s...)
}

// DiffFiles Compares the files in the working tree and the index. When paths
// are specified, compares only those named paths. Otherwise all entries in the
// index are compared. The output format is the same as for git diff-index and
// git diff-tree.
//   git-diff-files - https://git-scm.com/docs/git-diff-files
func (g *Git) DiffFiles() ErrNotImplemented {
	return errNotImplemented("diff-files")
}

// DiffIndex Compares the content and mode of the blobs found in a tree object with the
// corresponding tracked files in the working tree, or with the corresponding paths
// in the index. When <path> arguments are present, compares only paths matching those
// patterns. Otherwise all tracked files are compared.
//   git-diff-index - https://git-scm.com/docs/git-diff-index
func (g *Git) DiffIndex() ErrNotImplemented {
	return errNotImplemented("diff-index")
}

// DiffTree compares the content and mode of the blobs found via two tree objects.
// If there is only one <tree-ish> given, the commit is compared with its parents
// (see --stdin below). Note that git diff-tree can use the tree encapsulated in
// a commit object.
//   git-diff-tree - https://git-scm.com/docs/git-diff-tree
func (g *Git) DiffTree() ErrNotImplemented {
	return errNotImplemented("diff-tree")
}

// Diff shows changes between the working tree and the index or a tree, changes
// between the index and a tree, changes between two trees, changes between two
// blob objects, or changes between two files on disk.
//   git-diff - https://git-scm.com/docs/git-diff
func (g *Git) Diff() ErrNotImplemented {
	return errNotImplemented("diff")
}

// Difftool it difftool is a Git command that allows you to compare and edit
// files between revisions using common diff tools. git difftool is a frontend
// to git diff and accepts the same options and arguments.
//   git-difftool -
func (g *Git) Difftool() ErrNotImplemented {
	return errNotImplemented("difftool")
}

// FastExport This program dumps the given revisions in a form suitable to be
// piped into git fast-import. You can use it as a human-readable bundle
// replacement (see git-bundle[1]), or as a kind of an interactive git filter-branch.
//   git-fast-export - https://git-scm.com/docs/git-fast-export
func (g *Git) FastExport() ErrNotImplemented {
	return errNotImplemented("fast-export")
}

// FastImport This program is usually not what the end user wants to run directly.
// Most end users want to use one of the existing frontend programs, which parses
// a specific type of foreign source and feeds the contents stored there to git fast-import.
//   git-fast-import - https://git-scm.com/docs/git-fast-import
func (g *Git) FastImport() ErrNotImplemented {
	return errNotImplemented("fast-import")
}

// FetchPack Usually you would want to use git fetch, which is a higher level
// wrapper of this command, instead. Invokes git-upload-pack on a possibly
// remote repository and asks it to send objects missing from this repository,
// to update the named heads. The list of commits available locally is found out
// by scanning the local refs/ hierarchy and sent to git-upload-pack running on the other en
//   git-fetch-pack - https://git-scm.com/docs/git-fetch-pack
func (g *Git) FetchPack() ErrNotImplemented {
	return errNotImplemented("fetch-pack")
}

// Fetch branches and/or tags (collectively, "refs") from one or more other
// repositories, along with the objects necessary to complete their histories.
// Remote-tracking branches are updated (see the description of <refspec>
// below for ways to control this behavior).
//   git-fetch - https://git-scm.com/docs/git-fetch
func (g *Git) Fetch() ErrNotImplemented {
	return errNotImplemented("fetch")
}

// FilterBranch Lets you rewrite Git revision history by rewriting the branches
// mentioned in the <rev-list options>, applying custom filters on each revision.
// Those filters can modify each tree (e.g. removing a file or running a perl
// rewrite on all files) or information about each commit. Otherwise, all
// information (including original commit times or merge information) will be preserved.
//   git-filter-branch - https://git-scm.com/docs/git-filter-branch
func (g *Git) FilterBranch() ErrNotImplemented {
	return errNotImplemented("filter-branch")
}

// FmtMergeMsg Takes the list of merged objects on stdin and produces a suitable
// commit message to be used for the merge commit, usually to be passed as the
// <merge-message> argument of git merge.
//   git-fmt-merge-msg - https://git-scm.com/docs/git-fmt-merge-msg
func (g *Git) FmtMergeMsg() ErrNotImplemented {
	return errNotImplemented("fmt-merge-msg")
}

// ForEachRef Iterate over all refs that match <pattern> and show them according
// to the given <format>, after sorting them according to the given set of <key>.
// If <count> is given, stop after showing that many refs. The interpolated
// values in <format> can optionally be quoted as string literals in the
// specified host language allowing their direct evaluation in that language.
//   git-for-each-ref - https://git-scm.com/docs/git-for-each-ref
func (g *Git) ForEachRef() ErrNotImplemented {
	return errNotImplemented("for-each-ref")
}

// FormatPatch prepare each commit with its patch in one file per commit,
// formatted to resemble UNIX mailbox format. The output of this command is
// convenient for e-mail submission or for use with git am.
//   git-format-patch - https://git-scm.com/docs/git-format-patch
func (g *Git) FormatPatch() ErrNotImplemented {
	return errNotImplemented("format-patch")
}

// FsckObjects  is a synonym for git-fsck[1].
// Please refer to the documentation of that command.
//   git-fsck-objects - https://git-scm.com/docs/git-fsck-objects
func (g *Git) FsckObjects() ErrDeprecated {
	return errDeprecated("fsck-objects", "fsck",
		"https://git-scm.com/docs/git-fsck")
}

// Fsck verifies the connectivity and validity of the objects in the database.
//   git-fsck - https://git-scm.com/docs/git-fsck
func (g *Git) Fsck() ErrNotImplemented {
	return errNotImplemented("fsck")
}

// GC Runs a number of housekeeping tasks within the current repository,
// such as compressing file revisions (to reduce disk space and increase
// performance) and removing unreachable objects which may have been created
// from prior invocations of git add.
//   git-gc - https://git-scm.com/docs/git-gc
func (g *Git) GC() ErrNotImplemented {
	return errNotImplemented("gc")
}

// GetTarCommitID Read a tar archive created by git archive from the standard
// input and extract the commit ID stored in it. It reads only the first 1024
// bytes of input, thus its runtime is not influenced by the size of the tar
// archive very much.
//   git-get-tar-commit-id - https://git-scm.com/docs/git-get-tar-commit-id
func (g *Git) GetTarCommitID() ErrNotImplemented {
	return errNotImplemented("get-tar-commit-id")
}

// Grep Look for specified patterns in the tracked files in the work tree,
// blobs registered in the index file, or blobs in given tree objects.
//  Patterns are lists of one or more search expressions separated by newline
// characters. An empty string as search expression matches all lines
//   git-grep - https://git-scm.com/docs/git-grep
func (g *Git) Grep() ErrNotImplemented {
	return errNotImplemented("grep")
}

// GUI A Tcl/Tk based graphical user interface to Git. git gui focuses on
// allowing users to make changes to their repository by making new commits,
// amending existing ones, creating branches, performing local merges, and
// fetching/pushing to remote repositories.
//   git-gui - https://git-scm.com/docs/git-gui
func (g *Git) GUI() ErrNotImplemented {
	return errNotImplemented("gui")
}

// HashObject computes the object ID value for an object with specified type
// with the contents of the named file (which can be outside of the work tree),
// and optionally writes the resulting object into the object database. Reports
// its object ID to its standard output. This is used by git cvsimport to update
// the index without modifying files in the work tree.
// When <type> is not specified, it defaults to "blob".
//   git-hash-object - https://git-scm.com/docs/git-hash-object
func (g *Git) HashObject() ErrNotImplemented {
	return errNotImplemented("hash-object")
}

// Help display help information about Git
//   git-help - https://git-scm.com/docs/git-help
func (g *Git) Help() ErrNotImplemented {
	return errNotImplemented("help")
}

// HTTPBackend A simple CGI program to serve the contents of a Git repository
// to Git clients accessing the repository over http:// and https:// protocols.
// The program supports clients fetching using both the smart HTTP protocol and
// the backwards-compatible dumb HTTP protocol, as well as clients pushing using
// the smart HTTP protocol.
//   git-http-backend - https://git-scm.com/docs/git-http-backend
func (g *Git) HTTPBackend() ErrNotImplemented {
	return errNotImplemented("http-backend")
}

// HTTPFetch Download from a remote Git repository via HTTP
//   git-http-fetch - https://git-scm.com/docs/git-http-fetch
func (g *Git) HTTPFetch() ErrNotImplemented {
	return errNotImplemented("http-fetch")
}

// HTTPPush push objects over HTTP/DAV to another repository
//   git-http-push - https://git-scm.com/docs/git-http-push
func (g *Git) HTTPPush() ErrNotImplemented {
	return errNotImplemented("http-push")
}

// ImapSend command uploads a mailbox generated with git format-patch into an
// IMAP drafts folder. This allows patches to be sent as other email is when
// using mail clients that cannot read mailbox files directly. The command
// also works with any general mailbox in which emails have the fields "From",
// "Date", and "Subject" in that order.
//   git format-patch --signoff --stdout --attach origin | git imap-send
//
//   git-imap-send - https://git-scm.com/docs/git-imap-send
func (g *Git) ImapSend() ErrNotImplemented {
	return errNotImplemented("imap-send")
}

// IndexPack Reads a packed archive (.pack) from the specified file, and builds
// a pack index file (.idx) for it. The packed archive together with the pack
// index can then be placed in the objects/pack/ directory of a Git repository.
//   git-index-pack - https://git-scm.com/docs/git-index-pack
func (g *Git) IndexPack() ErrNotImplemented {
	return errNotImplemented("index-pack")
}

// InitDB This is a synonym for git-init[1].
// Please refer to the documentation of that command.
//   git-init-db - https://git-scm.com/docs/git-init-db
func (g *Git) InitDB() ErrDeprecated {
	return errDeprecated("init-db", "init", "https://git-scm.com/docs/git-init-db")
}

// Init creates an empty Git repository or reinitialize an existing one
//   git-init - https://git-scm.com/docs/git-init
func (g *Git) Init() ErrNotImplemented {
	return errNotImplemented("init")
}

// Instaweb Instantly browse your working repository in gitweb
//   git-instaweb - https://git-scm.com/docs/git-instaweb
func (g *Git) Instaweb() ErrNotImplemented {
	return errNotImplemented("instaweb")
}

// InterpretTrailers help add structured information into commit messages.
// Help adding trailers lines, that look similar to RFC 822 e-mail headers,
// at the end of the otherwise free-form part of a commit message.
//   git-interpret-trailers - https://git-scm.com/docs/git-interpret-trailers
func (g *Git) InterpretTrailers() ErrNotImplemented {
	return errNotImplemented("interpret-trailers")
}

// Gitk Displays changes in a repository or a selected set of commits.
// This includes visualizing the commit graph, showing information related to
// each commit, and the files in the trees of each revision.
//   git-gitk - https://git-scm.com/docs/gitk
func (g *Git) Gitk() ErrNotImplemented {
	return errNotImplemented("gitk")
}

// Log shows the commit logs.
//   git-log - https://git-scm.com/docs/git-log
func (g *Git) Log(s ...string) (*Output, error) {
	s = append([]string{"log"}, s...)
	return cmdgitInPath(g.wd.Abs(), s...)
}

// LsFiles show information about files in the index and the working tree.
// This merges the file listing in the directory cache index with the actual
// working directory list, and shows different combinations of the two.
//   git-ls-files - https://git-scm.com/docs/git-ls-files
func (g *Git) LsFiles() ErrNotImplemented {
	return errNotImplemented("ls-files")
}

// LsRemote displays references available in a remote repository
// along with the associated commit IDs.
//   git-ls-remote - https://git-scm.com/docs/git-ls-remote
func (g *Git) LsRemote() ErrNotImplemented {
	return errNotImplemented("ls-remote")
}

// LsTree lists the contents of a given tree object, like what "/bin/ls -a"
// does in the current working directory.
//   git-ls-tree - https://git-scm.com/docs/git-ls-tree
func (g *Git) LsTree() ErrNotImplemented {
	return errNotImplemented("ls-tree")
}

// Mailinfo extracts patch and authorship from a single e-mail message.
// reads a single e-mail message from the standard input, and writes the commit
// log message in <msg> file, and the patches in <patch> file.
//   git-mailinfo - https://git-scm.com/docs/git-mailinfo
func (g *Git) Mailinfo() ErrNotImplemented {
	return errNotImplemented("mailinfo")
}

// Mailsplit Simple UNIX mbox splitter program. Splits a mbox file or a Maildir
// into a list of files: "0001" "0002" .. in the specified directory so you can
// process them further from there.
//   git-mailsplit - https://git-scm.com/docs/git-mailsplit
func (g *Git) Mailsplit() ErrNotImplemented {
	return errNotImplemented("mailsplit")
}

// MergeBase finds best common ancestor(s) between two commits to use in a
// three-way merge. One common ancestor is better than another common ancestor
// if the latter is an ancestor of the former. A common ancestor that does not
// have any better common ancestor is a best common ancestor, i.e. a merge base.
// Note that there can be more than one merge base for a pair of commits.
//   git-merge-base -
func (g *Git) MergeBase() ErrNotImplemented {
	return errNotImplemented("merge-base")
}

// MergeFile runs a three-way file merge.
//   git-merge-file - https://git-scm.com/docs/git-merge-file
func (g *Git) MergeFile() ErrNotImplemented {
	return errNotImplemented("merge-file")
}

// MergeIndex run a merge for files needing merging.
//   git-merge-index - https://git-scm.com/docs/git-merge-index
func (g *Git) MergeIndex() ErrNotImplemented {
	return errNotImplemented("merge-index")
}

// MergeOctopus common merge strategy, resolves cases with more than two heads,
// but refuses to do a complex merge that needs manual resolution
//   git-merge-octopus - https://git-scm.com/docs/git-merge
func (g *Git) MergeOctopus() ErrNotImplemented {
	return errNotImplemented("merge-octopus")
}

// MergeOneFile is the standard helper program to use with git merge-index to
// resolve a merge after the trivial merge done with git read-tree -m.
//   git-merge-one-file - https://git-scm.com/docs/git-merge-one-file
func (g *Git) MergeOneFile() ErrNotImplemented {
	return errNotImplemented("merge-one-file")
}

// MergeOurs resolves any number of heads, but the resulting tree of the merge
// is always that of the current branch head, effectively ignoring all changes
// from all other branches
//   git-merge-ours - https://git-scm.com/docs/git-merge
func (g *Git) MergeOurs() ErrNotImplemented {
	return errNotImplemented("merge-ours")
}

// MergeRecursive can only resolve two heads using a 3-way merge algorithm.
//   git-merge-recursive - https://git-scm.com/docs/git-merge
func (g *Git) MergeRecursive() ErrNotImplemented {
	return errNotImplemented("merge-recursive")
}

// MergeResolve This can only resolve two heads (i.e. the current branch and
// another branch you pulled from) using a 3-way merge algorithm.
//   git-merge-resolve - https://git-scm.com/docs/git-merge
func (g *Git) MergeResolve() ErrNotImplemented {
	return errNotImplemented("merge-resolve")
}

// MergeSubtree This is a modified recursive strategy. When merging trees A and B
//   git-merge-subtree - https://git-scm.com/docs/git-merge
func (g *Git) MergeSubtree() ErrNotImplemented {
	return errNotImplemented("merge-subtree")
}

// MergeTree Reads three tree-ish, and output trivial merge results and conflicting
// stages to the standard output. This is similar to what three-way git read-tree
// -m does, but instead of storing the results in the index, the command outputs
// the entries to the standard output.
//   git-merge-tree -
func (g *Git) MergeTree() ErrNotImplemented {
	return errNotImplemented("merge-tree")
}

// Merge Incorporates changes from the named commits (since the time their histories
// diverged from the current branch) into the current branch. This command is used
// by git pull to incorporate changes from another repository and can be used by
// hand to merge changes from one branch into another.
//   git-merge - https://git-scm.com/docs/git-merge
func (g *Git) Merge() ErrNotImplemented {
	return errNotImplemented("merge")
}

// Mergetool Run merge conflict resolution tools to resolve merge conflicts.
//   git-mergetool - https://git-scm.com/docs/git-mergetool
func (g *Git) Mergetool() ErrNotImplemented {
	return errNotImplemented("mergetool")
}

// Mktag reads a tag contents on standard input and creates a tag object that
// can also be used to sign other objects.
//   git-mktag - https://git-scm.com/docs/git-mktag
func (g *Git) Mktag() ErrNotImplemented {
	return errNotImplemented("mktag")
}

// Mktree Build a tree-object from ls-tree formatted text. Reads standard input
// in non-recursive ls-tree output format, and creates a tree object. The order
// of the tree entries is normalised by mktree so pre-sorting the input is not
// required. The object name of the tree object built is written to the standard output.
//   git-mktree - https://git-scm.com/docs/git-mktree
func (g *Git) Mktree() ErrNotImplemented {
	return errNotImplemented("mktree")
}

// Mv move or rename a file, a directory, or a symlink.
//   git-mv - https://git-scm.com/docs/git-mv
func (g *Git) Mv() ErrNotImplemented {
	return errNotImplemented("mv")
}

// NameRev finds symbolic names suitable for human digestion for revisions given
// in any format parsable by git rev-parse.
//   git-name-rev - https://git-scm.com/docs/git-name-rev
func (g *Git) NameRev() ErrNotImplemented {
	return errNotImplemented("name-rev")
}

// Notes adds, removes, or reads notes attached to objects,
// without touching the objects themselves.
//   git-notes - https://git-scm.com/docs/git-notes
func (g *Git) Notes() ErrNotImplemented {
	return errNotImplemented("notes")
}

// P4 Import from and submit to Perforce repositories
//   git-p4 - https://git-scm.com/docs/git-p4
func (g *Git) P4() ErrNotImplemented {
	return errNotImplemented("p4")
}

// PackObjects reads list of objects from the standard input, and writes a packed
// archive with specified base-name, or to the standard output.
//   git-pack-objects - https://git-scm.com/docs/git-pack-objects
func (g *Git) PackObjects() ErrNotImplemented {
	return errNotImplemented("pack-objects")
}

// PackRedundant computes which packs in your repository are redundant.
// The output is suitable for piping to xargs rm if you are in the root of the repositor
//   git-pack-redundant - https://git-scm.com/docs/git-pack-redundant
func (g *Git) PackRedundant() ErrNotImplemented {
	return errNotImplemented("pack-redundant")
}

// PackRefs Traditionally, tips of branches and tags (collectively known as refs)
// were stored one file per ref in a (sub)directory under $GIT_DIR/refs directory.
//   git-pack-refs - https://git-scm.com/docs/git-pack-refs
func (g *Git) PackRefs() ErrNotImplemented {
	return errNotImplemented("pack-refs")
}

// ParseRemote routines to help parsing remote repository access parameters.
//   git-parse-remote - https://git-scm.com/docs/git-parse-remote
func (g *Git) ParseRemote() ErrNotImplemented {
	return errNotImplemented("parse-remote")
}

// PatchID read a patch from the standard input and compute the patch ID for it.
//   git-patch-id - https://git-scm.com/docs/git-patch-id
func (g *Git) PatchID() ErrNotImplemented {
	return errNotImplemented("patch-id")
}

// Prune all unreachable objects from the object database
//   git-prune - https://git-scm.com/docs/git-prune
func (g *Git) Prune() ErrNotImplemented {
	return errNotImplemented("prune")
}

// PrunePacked removes extra objects that are already in pack files
//   git-prune-packed - https://git-scm.com/docs/git-prune-packed
func (g *Git) PrunePacked() ErrNotImplemented {
	return errNotImplemented("prune-packed")
}

// Pull fetch from and integrate with another repository or a local branch
//   git-pull - https://git-scm.com/docs/git-pull
func (g *Git) Pull() ErrNotImplemented {
	return errNotImplemented("pull")
}

// Push updates remote refs using local refs, while sending objects necessary
// to complete the given refs.
//   git-push - https://git-scm.com/docs/git-push
func (g *Git) Push() ErrNotImplemented {
	return errNotImplemented("push")
}

// Quiltimport Applies a quilt patchset onto the current Git branch, preserving
// the patch boundaries, patch order, and patch descriptions present in the quilt patchset.
//   git-quiltimport - https://git-scm.com/docs/git-quiltimport
func (g *Git) Quiltimport() ErrNotImplemented {
	return errNotImplemented("quiltimport")
}

// ReadTree Reads the tree information given by <tree-ish> into the index,
// but does not actually update any of the files it "caches". (see: git-checkout-index[1])
//   git-read-tree - https://git-scm.com/docs/git-read-tree
func (g *Git) ReadTree() ErrNotImplemented {
	return errNotImplemented("read-tree")
}

// Rebase reapply commits on top of another base tip
//   git-rebase - https://git-scm.com/docs/git-rebase
func (g *Git) Rebase() ErrNotImplemented {
	return errNotImplemented("rebase")
}

// ReceivePack invoked by git send-pack and updates the repository with the
// information fed from the remote end.
//   git-receive-pack - https://git-scm.com/docs/git-receive-pack
func (g *Git) ReceivePack() ErrNotImplemented {
	return errNotImplemented("receive-pack")
}

// Reflog the command takes various subcommands, and different options
// depending on the subcommand:
//   git-reflog - https://git-scm.com/docs/git-reflog
func (g *Git) Reflog() ErrNotImplemented {
	return errNotImplemented("reflog")
}

// RemoteExt Bridge smart transport to external command.Data written to stdin of
// the specified <command> is assumed to be sent to a git:// server, git-upload-pack,
// git-receive-pack or git-upload-archive (depending on situation), and data read
// from stdout of <command> is assumed to be received from the same service.
//   git-remote-ext - https://git-scm.com/docs/git-remote-ext
func (g *Git) RemoteExt() ErrNotImplemented {
	return errNotImplemented("remote-ext")
}

// RemoteFd This helper uses specified file descriptors to connect to a remote
// Git server. This is not meant for end users but for programs and scripts
// calling git fetch, push or archive.
//   git-remote-fd - https://git-scm.com/docs/git-remote-fd
func (g *Git) RemoteFd() ErrNotImplemented {
	return errNotImplemented("remote-fd")
}

// RemoteFtp ftp
//   git-remote-ftp -
func (g *Git) RemoteFtp() ErrNotImplemented {
	return errNotImplemented("remote-ftp")
}

// RemoteFtps ftps
//   git-remote-ftps -
func (g *Git) RemoteFtps() ErrNotImplemented {
	return errNotImplemented("remote-ftps")
}

// RemoteHTTP http
//   git-remote-http -
func (g *Git) RemoteHTTP() ErrNotImplemented {
	return errNotImplemented("remote-http")
}

// RemoteHTTPS https
//   git-remote-https -
func (g *Git) RemoteHTTPS() ErrNotImplemented {
	return errNotImplemented("remote-https")
}

// RemoteTestSvn svn
//   git-remote-testsvn -
func (g *Git) RemoteTestSvn() ErrNotImplemented {
	return errNotImplemented("remote-testsvn")
}

// RemoteTestGit git
//   git-remote-testgit -
func (g *Git) RemoteTestGit() ErrNotImplemented {
	return errNotImplemented("remote-testgit")
}

// Remote manage the set of repositories ("remotes") whose branches you track.
//   git-remote - https://git-scm.com/docs/git-remote
func (g *Git) Remote() ErrNotImplemented {
	return errNotImplemented("remote")
}

// Repack This command is used to combine all objects that do not currently
// reside in a "pack", into a pack. It can also be used to re-organize existing
// packs into a single, more efficient pack.
//   git-repack - https://git-scm.com/docs/git-repack
func (g *Git) Repack() ErrNotImplemented {
	return errNotImplemented("repack")
}

// Replace create, list, delete refs to replace objects
//   git-replace - https://git-scm.com/docs/git-replace
func (g *Git) Replace() ErrNotImplemented {
	return errNotImplemented("replace")
}

// RequestPull generates a summary of pending changes
//   git-request-pull - Generate a request asking your upstream project to pull
// changes into their tree. The request, printed to the standard output, begins
// with the branch description, summarizes the changes and indicates from where
// they can be pulled.
//   git-request-pull - https://git-scm.com/docs/git-request-pull
func (g *Git) RequestPull() ErrNotImplemented {
	return errNotImplemented("request-pull")
}

// Rerere Reuse recorded resolution of conflicted merges
//   git-rerere -
func (g *Git) Rerere() ErrNotImplemented {
	return errNotImplemented("rerere")
}

// Reset current HEAD to the specified state
//   git-reset - https://git-scm.com/docs/git-reset
func (g *Git) Reset() ErrNotImplemented {
	return errNotImplemented("reset")
}

// RevList Lists commit objects in reverse chronological order
//   git-rev-list - https://git-scm.com/docs/git-rev-list
func (g *Git) RevList() ErrNotImplemented {
	return errNotImplemented("rev-list")
}

// RevParse Many Git porcelainish commands take mixture of flags (i.e.
// parameters that begin with a dash -) and parameters meant for the underlying
// git rev-list command they use internally and flags and parameters for the other
// commands they use downstream of git rev-list.
// This command is used to distinguish between them.
//   git-rev-parse - https://git-scm.com/docs/git-rev-parse
func (g *Git) RevParse() ErrNotImplemented {
	return errNotImplemented("rev-parse")
}

// Revert Given one or more existing commits, revert the changes that the related
// patches introduce, and record some new commits that record them. This requires
// your working tree to be clean (no modifications from the HEAD commit).
//   git-revert -
func (g *Git) Revert() ErrNotImplemented {
	return errNotImplemented("revert")
}

// Rm Remove files from the index, or from the working tree and the index. git
// rm will not remove a file from just your working directory.
//   git-rm - https://git-scm.com/docs/git-rm
func (g *Git) Rm() ErrNotImplemented {
	return errNotImplemented("rm")
}

// SendEmail Send a collection of patches as emails
//   git-send-email - https://git-scm.com/docs/git-send-email
func (g *Git) SendEmail() ErrNotImplemented {
	return errNotImplemented("send-email")
}

// SendPack push objects over Git protocol to another repository. Usually you
// would want to use git push, which is a higher-level wrapper of this command,
//instead. See git-push[1].
//   git-send-pack - https://git-scm.com/docs/git-send-pack
func (g *Git) SendPack() ErrNotImplemented {
	return errNotImplemented("send-pack")
}

// Shell Restricted login shell for Git-only SSH access
//   git-shell - https://git-scm.com/docs/git-shell
func (g *Git) Shell() ErrNotImplemented {
	return errNotImplemented("shell")
}

// Shortlog summarizes git log output in a format suitable for inclusion in
// release announcements. Each commit will be grouped by author and title.
//   git-shortlog - https://git-scm.com/docs/git-shortlog
func (g *Git) Shortlog() ErrNotImplemented {
	return errNotImplemented("shortlog")
}

// ShowBranch Show branches and their commits. Shows the commit ancestry graph
// starting from the commits named with <rev>s or <globs>s (or all refs under
// refs/heads and/or refs/tags) semi-visually.
//   git-show-branch - https://git-scm.com/docs/git-show-branch
func (g *Git) ShowBranch() ErrNotImplemented {
	return errNotImplemented("show-branch")
}

// ShowIndex Show packed archive index
//   git-show-index - https://git-scm.com/docs/git-show-index
func (g *Git) ShowIndex() ErrNotImplemented {
	return errNotImplemented("show-index")
}

// ShowRef List references in a local repository
//   git-show-ref - https://git-scm.com/docs/git-show-ref
func (g *Git) ShowRef() ErrNotImplemented {
	return errNotImplemented("show-ref")
}

// Show  one or more objects (blobs, trees, tags and commits).
//   git-show - https://git-scm.com/docs/git-show
func (g *Git) Show() ErrNotImplemented {
	return errNotImplemented("show")
}

// Stage Add file contents to the staging area
//   git-stage - https://git-scm.com/docs/git-stage
func (g *Git) Stage() ErrNotImplemented {
	return errNotImplemented("stage")
}

// Stash Stash the changes in a dirty working directory away
//   git-stash - https://git-scm.com/docs/git-stash
func (g *Git) Stash() ErrNotImplemented {
	return errNotImplemented("stash")
}

// Status Displays paths that have differences between the index file and the
// current HEAD commit, paths that have differences between the working tree and
// the index file, and paths in the working tree that are not tracked by Git
// (and are not ignored by gitignore[5]).
//   git-status - https://git-scm.com/docs/git-status
func (g *Git) Status() ErrNotImplemented {
	return errNotImplemented("status")
}

// Stripspace Read text, such as commit messages, notes, tags and branch descriptions,
// from the standard input and clean it in the manner used by Git.
//   git-stripspace -
func (g *Git) Stripspace() ErrNotImplemented {
	return errNotImplemented("stripspace")
}

// Submodule Inspects, updates and manages submodules.
// A submodule allows you to keep another Git repository in a subdirectory of
// your repository. The other repository has its own history, which does not
// interfere with the history of the current repository. This can be used to have
// external dependencies such as third party libraries for example.
//   git-submodule -
func (g *Git) Submodule() ErrNotImplemented {
	return errNotImplemented("submodule")
}

// Svn Bidirectional operation between a Subversion repository and Git.
// git svn is a simple conduit for changesets between Subversion and Git.
// It provides a bidirectional flow of changes between a Subversion and a Git repository.
//   git-svn -
func (g *Git) Svn() ErrNotImplemented {
	return errNotImplemented("svn")
}

// SymbolicRef Read, modify and delete symbolic refs.
// Given one argument, reads which branch head the given symbolic ref refers to
// and outputs its path, relative to the .git/ directory. Typically you would give
// HEAD as the <name> argument to see which branch your working tree is on.
//   git-symbolic-ref - https://git-scm.com/docs/git-symbolic-ref
func (g *Git) SymbolicRef() ErrNotImplemented {
	return errNotImplemented("symbolic-ref")
}

// Tag Create, list, delete or verify a tag object signed with GPG.
//   git-tag - https://git-scm.com/docs/git-tag
func (g *Git) Tag() ErrNotImplemented {
	return errNotImplemented("tag")
}

// UnpackFile Creates a temporary file with a blob’s contents. Creates a file
// holding the contents of the blob specified by sha1. It returns the name of
// the temporary file in the following format: .merge_file_XXXXX
//   git-unpack-file - https://git-scm.com/docs/git-unpack-file
func (g *Git) UnpackFile() ErrNotImplemented {
	return errNotImplemented("unpack-file")
}

// UnpackObjects Unpack objects from a packed archive. Read a packed archive
// (.pack) from the standard input, expanding the objects contained within and
// writing them into the repository in "loose" (one object per file) format.
//   git-unpack-objects - https://git-scm.com/docs/git-unpack-objects
func (g *Git) UnpackObjects() ErrNotImplemented {
	return errNotImplemented("unpack-objects")
}

// UpdateIndex Register file contents in the working tree to the index.
// Modifies the index or directory cache. Each file mentioned is updated into
// the index and any unmerged or needs updating state is cleared.
//   git-update-index - https://git-scm.com/docs/git-update-index
func (g *Git) UpdateIndex() ErrNotImplemented {
	return errNotImplemented("update-index")
}

// UpdateRef Update the object name stored in a ref safely
//   git-update-ref - https://git-scm.com/docs/git-update-ref
func (g *Git) UpdateRef() ErrNotImplemented {
	return errNotImplemented("update-ref")
}

// UpdateServerInfo Update auxiliary info file to help dumb servers.
// A dumb server that does not do on-the-fly pack generations must have some
// auxiliary information files in $GIT_DIR/info and $GIT_OBJECT_DIRECTORY/info
// directories to help clients discover what references and packs the server has.
// This command generates such auxiliary files
//   git-update-server-info -
func (g *Git) UpdateServerInfo() ErrNotImplemented {
	return errNotImplemented("update-server-info")
}

// UploadArchive Send archive back to git-archive.
// Invoked by git archive --remote and sends a generated archive to the other
// end over the Git protocol.
//   git-upload-archive - https://git-scm.com/docs/git-upload-archive
func (g *Git) UploadArchive() ErrNotImplemented {
	return errNotImplemented("upload-archive")
}

// UploadPack Send objects packed back to git-fetch-pack.
// This command is usually not invoked directly by the end user.
// The UI for the protocol is on the git fetch-pack side, and the program pair
// is meant to be used to pull updates from a remote repository.
// For push operations, see git send-pack.
//   git-upload-pack - https://git-scm.com/docs/git-upload-pack
func (g *Git) UploadPack() ErrNotImplemented {
	return errNotImplemented("upload-pack")
}

// Var Show a Git logical variable. Cause the logical variables to be listed.
// In addition, all the variables of the Git configuration file .git/config are
// listed as well. (However, the configuration variables listing functionality
// is deprecated in favor of git config -l.)
//   git-var -
func (g *Git) Var() ErrNotImplemented {
	return errNotImplemented("var")
}

// VerifyCommit Check the GPG signature of commits. Validates the GPG signature
// created by git commit -S.
//   git-verify-commit - https://git-scm.com/docs/git-verify-commit
func (g *Git) VerifyCommit() ErrNotImplemented {
	return errNotImplemented("verify-commit")
}

// VerifyPack Reads given idx file for packed Git archive created with the git
// pack-objects command and verifies idx file and the corresponding pack file.
//   git-verify-pack - https://git-scm.com/docs/git-verify-pack
func (g *Git) VerifyPack() ErrNotImplemented {
	return errNotImplemented("verify-pack")
}

// VerifyTag Check the GPG signature of tags. Validates the gpg signature created by git tag.
//   git-verify-tag - https://git-scm.com/docs/git-verify-tag
func (g *Git) VerifyTag() ErrNotImplemented {
	return errNotImplemented("verify-tag")
}

// Whatchanged Show logs with difference each commit introduces
//   git-whatchanged - https://git-scm.com/docs/git-whatchanged
func (g *Git) Whatchanged() ErrNotImplemented {
	return errNotImplemented("whatchanged")
}

// Worktree Manage multiple working trees attached to the same repository.
//   git-worktree - https://git-scm.com/docs/git-worktree
func (g *Git) Worktree() ErrNotImplemented {
	return errNotImplemented("worktree")
}

// WriteTree Creates a tree object using the current index.
//  The name of the new tree object is printed to standard output.
//   git-write-tree - https://git-scm.com/docs/git-write-tree
func (g *Git) WriteTree() ErrNotImplemented {
	return errNotImplemented("write-tree")
}
