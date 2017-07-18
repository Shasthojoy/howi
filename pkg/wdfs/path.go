package wdfs

import (
	"go/build"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/howi-ce/howi/pkg/std/herrors"
)

// New returns PathInfo for given path, It tries to set absolute representation
// of path, but sets provided string if that fails
func New(path string) (*Path, error) {
	abs, err := filepath.Abs(filepath.FromSlash(path))
	p := &Path{}
	if err != nil {
		p.abs = path
	} else {
		p.abs = abs
	}
	// first load
	p.Stat()
	p.base = filepath.Base(abs)
	p.clean = filepath.Clean(abs)
	p.dir = filepath.Dir(abs)
	p.ext = filepath.Ext(abs)
	p.isAbs = filepath.IsAbs(abs)
	return p, err
}

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
func (o *Path) Abs() string {
	return o.abs
}

// Base returns the last element of path.
// result of filepath.Base when object was loaded
func (o *Path) Base() string {
	return o.base
}

// Clean returns the shortest path name equivalent to path
// by purely lexical processing.
// result of filepath.Clean when object was loaded
func (o *Path) Clean() string {
	return o.clean
}

// Dir returns all but the last element of path, typically the path's directory.
// After dropping the final element,
// result of filepath.Dir when object was loaded
func (o *Path) Dir() string {
	return o.dir
}

// Ext returns the file name extension used by path.
// The extension is the suffix beginning at the final dot
// in the final element of path; it is empty if there is
// no dot.
// result of filepath.Ext when object was loaded
func (o *Path) Ext() string {
	return o.ext
}

// IsAbs reports whether the path is absolute.
// result of filepath.IsAbs when object was loaded
func (o *Path) IsAbs() bool {
	return o.isAbs
}

// VolumeName is abbreviation for filepath.VolumeName()
// with current path as argumentwhich returns leading volume name.
func (o *Path) VolumeName() string {
	return filepath.VolumeName(o.abs)
}

// Split is abbreviation for filepath.Split() with current path as argument
// which splits path immediately following the final Separator,
func (o *Path) Split() (dir, file string) {
	return filepath.Split(o.abs)
}

// ToSlash is abbreviation for filepath.ToSlash() with current path as argument
// which returns the result of replacing each separator character in path
// with a slash ('/') character.
func (o *Path) ToSlash() string {
	return filepath.ToSlash(o.abs)
}

// Join joins any number of path elements into a single path
// and appends these to current path
func (o *Path) Join(elem ...string) string {
	return filepath.Join(append([]string{o.abs}, elem...)...)
}

// Match is abbreviation for filepath.Match() with current paths basename as name argument
func (o *Path) Match(pattern string) (matched bool, err error) {
	return filepath.Match(pattern, o.base)
}

// Rel is abbreviation for filepath.Rel() with current path as basepath argument
func (o *Path) Rel(targpath string) (string, error) {
	return filepath.Rel(o.abs, targpath)
}

// Walk is abbreviation for filepath.Walk() with current path as root argument
func (o *Path) Walk(walkFn func(path string, info os.FileInfo, err error) error) error {
	return filepath.Walk(o.abs, walkFn)
}

// Exists checks if a path exists.
func (o *Path) Exists() bool {
	if _, err := o.Stat(); err != nil {
		return false
	}
	return true
}

// IsDir checks if a given path is a directory.
// func (os.FileInfo).IsDir() bool
func (o *Path) IsDir() bool {
	o.Stat()
	return o.fileInfo.IsDir()
}

// IsRegular reports whether opject describes a regular file.
func (o *Path) IsRegular() bool {
	o.Stat()
	return o.fileInfo.Mode().IsRegular()
}

// InGOPATH reports whether path is in GOPATH.
func (o *Path) InGOPATH() bool {
	for _, gopath := range filepath.SplitList(build.Default.GOPATH) {
		if strings.HasPrefix(o.abs, gopath) {
			return true
		}
	}
	return false
}

// Mode returns os.FileMode
func (o *Path) Mode() os.FileMode {
	o.Stat()
	return o.fileInfo.Mode()
}

// Perm returns os.FileInfo.Mode().Perm()
func (o *Path) Perm() os.FileMode {
	o.Stat()
	return o.fileInfo.Mode().Perm()
}

// ModTime returns modification time
// func (os.FileInfo).ModTime() time.Time
func (o *Path) ModTime() time.Time {
	o.Stat()
	return o.fileInfo.ModTime()
}

// Size returns length in bytes for regular files; system-dependent for others
func (o *Path) Size() int64 {
	o.Stat()
	return o.fileInfo.Size()
}

// Stat returns a FileInfo describing the named file.
// If there is an error, it will be of type *PathError.
func (o *Path) Stat() (os.FileInfo, error) {
	fileInfo, err := os.Stat(o.abs)
	if os.IsNotExist(err) {
		if o.fileInfo != nil {
			o.deleted = true
			return nil, o.error("has been deleted")
		}
		return nil, o.error(err.Error())
	}
	o.deleted = false
	o.fileInfo = fileInfo
	return o.fileInfo, nil
}

func (o *Path) error(msg ...string) error {
	return herrors.Newf("%s (%s)", o.abs, msg)
}
