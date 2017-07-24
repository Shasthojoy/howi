package path

import (
	"go/build"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/howi-ce/howi/std/herrors"
)

// Path retruns a path object with methods to work with path
type Path struct {
	abs      string
	base     string
	clean    string
	dir      string
	ext      string
	isAbs    bool
	deleted  bool
	fileInfo os.FileInfo
}

// Abs returns an absolute representation of path.
// result of filepath.Abs when object was loaded
func (p *Path) Abs() string {
	return p.abs
}

// Base returns the last element of path.
// result of filepath.Base when object was loaded
func (p *Path) Base() string {
	return p.base
}

// Clean returns the shortest path name equivalent to path
// by purely lexical processing.
// result of filepath.Clean when object was loaded
func (p *Path) Clean() string {
	return p.clean
}

// Dir returns all but the last element of path, typically the path's directory.
// After dropping the final element,
// result of filepath.Dir when object was loaded
func (p *Path) Dir() string {
	return p.dir
}

// Ext returns the file name extension used by path.
// The extension is the suffix beginning at the final dot
// in the final element of path; it is empty if there is
// no dot.
// result of filepath.Ext when object was loaded
func (p *Path) Ext() string {
	return p.ext
}

// IsAbs reports whether the path is absolute.
// result of filepath.IsAbs when object was loaded
func (p *Path) IsAbs() bool {
	return p.isAbs
}

// VolumeName is abbreviation for filepath.VolumeName()
// with current path as argumentwhich returns leading volume name.
func (p *Path) VolumeName() string {
	return filepath.VolumeName(p.abs)
}

// Split is abbreviation for filepath.Split() with current path as argument
// which splits path immediately following the final Separator,
func (p *Path) Split() (dir, file string) {
	return filepath.Split(p.abs)
}

// ToSlash is abbreviation for filepath.ToSlash() with current path as argument
// which returns the result of replacing each separator character in path
// with a slash ('/') character.
func (p *Path) ToSlash() string {
	return filepath.ToSlash(p.abs)
}

// Join joins any number of path elements into a single path
// and appends these to current path
func (p *Path) Join(elem ...string) string {
	return filepath.Join(append([]string{p.abs}, elem...)...)
}

// Match is abbreviation for filepath.Match() with current paths basename as name argument
func (p *Path) Match(pattern string) (matched bool, err error) {
	return filepath.Match(pattern, p.base)
}

// Rel is abbreviation for filepath.Rel() with current path as basepath argument
func (p *Path) Rel(targpath string) (string, error) {
	return filepath.Rel(p.abs, targpath)
}

// Walk is abbreviation for filepath.Walk() with current path as root argument
func (p *Path) Walk(walkFn func(path string, info os.FileInfo, err error) error) error {
	return filepath.Walk(p.abs, walkFn)
}

// Exists checks if a path exists.
func (p *Path) Exists() bool {
	if _, err := p.Stat(); err != nil {
		return false
	}
	return true
}

// IsDir checks if a given path is a directory.
// func (os.FileInfo).IsDir() bool
func (p *Path) IsDir() bool {
	p.Stat()
	return p.fileInfo.IsDir()
}

// IsGitRepository checks if a given path is a git repository directory.
func (p *Path) IsGitRepository() bool {
	p.Stat()
	if !p.fileInfo.IsDir() {
		return false
	}
	gitRepo, _ := Plugin(p.Join(".git"))
	return gitRepo.Exists()
}

// IsRegular reports whether opject describes a regular file.
func (p *Path) IsRegular() bool {
	p.Stat()
	return p.fileInfo.Mode().IsRegular()
}

// InGOPATH reports whether path is in GOPATH.
func (p *Path) InGOPATH() bool {
	for _, gopath := range filepath.SplitList(build.Default.GOPATH) {
		if strings.HasPrefix(p.abs, gopath) {
			return true
		}
	}
	return false
}

// Mode returns os.FileMode
func (p *Path) Mode() os.FileMode {
	p.Stat()
	return p.fileInfo.Mode()
}

// Perm returns os.FileInfo.Mode().Perm()
func (p *Path) Perm() os.FileMode {
	p.Stat()
	return p.fileInfo.Mode().Perm()
}

// ModTime returns modification time
// func (os.FileInfo).ModTime() time.Time
func (p *Path) ModTime() time.Time {
	p.Stat()
	return p.fileInfo.ModTime()
}

// Size returns length in bytes for regular files; system-dependent for others
func (p *Path) Size() int64 {
	p.Stat()
	return p.fileInfo.Size()
}

// Stat returns a FileInfo describing the named file.
// If there is an error, it will be of type *PathError.
func (p *Path) Stat() (os.FileInfo, error) {
	fileInfo, err := os.Stat(p.abs)
	if os.IsNotExist(err) {
		if p.fileInfo != nil {
			p.deleted = true
			return nil, p.error("has been deleted")
		}
		return nil, p.error(err.Error())
	}
	p.deleted = false
	p.fileInfo = fileInfo
	return p.fileInfo, nil
}

func (p *Path) error(msg ...string) error {
	return herrors.Newf("%s (%s)", p.abs, msg)
}
