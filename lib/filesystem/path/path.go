// Copyright 2005-2017 Marko Kungla.
// Use of this source code is governed by a The MIT License
// license that can be found in the LICENSE file.

package path

import (
	"go/build"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/howi-ce/howi/std/errors"
)

// New returns Obj for given string,
// It tries to set absolute representation of path,
// but sets provided string if that fails.
func New(path string) (Obj, error) {
	abs, err := filepath.Abs(filepath.FromSlash(path))
	p := Obj{}
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

// Obj is a path object with methods to work with path
type Obj struct {
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
func (p *Obj) Abs() string {
	return p.abs
}

// Base returns the last element of path.
// result of filepath.Base when object was loaded
func (p *Obj) Base() string {
	return p.base
}

// Clean returns the shortest path name equivalent to path
// by purely lexical processing.
// result of filepath.Clean when object was loaded
func (p *Obj) Clean() string {
	return p.clean
}

// Dir returns all but the last element of path, typically the path's directory.
// After dropping the final element,
// result of filepath.Dir when object was loaded
func (p *Obj) Dir() string {
	return p.dir
}

// Ext returns the file name extension used by path.
// The extension is the suffix beginning at the final dot
// in the final element of path; it is empty if there is
// no dot.
// result of filepath.Ext when object was loaded
func (p *Obj) Ext() string {
	return p.ext
}

// IsAbs reports whether the path is absolute.
// result of filepath.IsAbs when object was loaded
func (p *Obj) IsAbs() bool {
	return p.isAbs
}

// VolumeName is abbreviation for filepath.VolumeName()
// with current path as argumentwhich returns leading volume name.
func (p *Obj) VolumeName() string {
	return filepath.VolumeName(p.abs)
}

// Split is abbreviation for filepath.Split() with current path as argument
// which splits path immediately following the final Separator,
func (p *Obj) Split() (dir, file string) {
	return filepath.Split(p.abs)
}

// ToSlash is abbreviation for filepath.ToSlash() with current path as argument
// which returns the result of replacing each separator character in path
// with a slash ('/') character.
func (p *Obj) ToSlash() string {
	return filepath.ToSlash(p.abs)
}

// Join joins any number of path elements into a single path
// and appends these to current path
func (p *Obj) Join(elem ...string) string {
	return filepath.Join(append([]string{p.abs}, elem...)...)
}

// Match is abbreviation for filepath.Match() with current paths basename as name argument
func (p *Obj) Match(pattern string) (matched bool, err error) {
	return filepath.Match(pattern, p.base)
}

// Rel is abbreviation for filepath.Rel() with current path as basepath argument
func (p *Obj) Rel(targpath string) (string, error) {
	return filepath.Rel(p.abs, targpath)
}

// Walk is abbreviation for filepath.Walk() with current path as root argument
func (p *Obj) Walk(walkFn func(path string, info os.FileInfo, err error) error) error {
	return filepath.Walk(p.abs, walkFn)
}

// Exists checks if a path exists.
func (p *Obj) Exists() bool {
	if _, err := p.Stat(); err != nil {
		return false
	}
	return true
}

// IsDir checks if a given path is a directory.
// func (os.FileInfo).IsDir() bool
func (p *Obj) IsDir() bool {
	p.Stat()
	return p.fileInfo.IsDir()
}

// IsGitRepository checks if a given path is a git repository directory.
func (p *Obj) IsGitRepository() bool {
	p.Stat()
	if !p.fileInfo.IsDir() {
		return false
	}
	gitRepo, _ := New(p.Join(".git"))
	return gitRepo.Exists()
}

// IsRegular reports whether opject describes a regular file.
func (p *Obj) IsRegular() bool {
	p.Stat()
	return p.fileInfo.Mode().IsRegular()
}

// InGOPATH reports whether path is in GOPATH.
func (p *Obj) InGOPATH() bool {
	for _, gopath := range filepath.SplitList(build.Default.GOPATH) {
		if strings.HasPrefix(p.abs, gopath) {
			return true
		}
	}
	return false
}

// Mode returns os.FileMode
func (p *Obj) Mode() os.FileMode {
	p.Stat()
	return p.fileInfo.Mode()
}

// Perm returns os.FileInfo.Mode().Perm()
func (p *Obj) Perm() os.FileMode {
	p.Stat()
	return p.fileInfo.Mode().Perm()
}

// ModTime returns modification time
// func (os.FileInfo).ModTime() time.Time
func (p *Obj) ModTime() time.Time {
	p.Stat()
	return p.fileInfo.ModTime()
}

// Size returns length in bytes for regular files; system-dependent for others
func (p *Obj) Size() int64 {
	p.Stat()
	return p.fileInfo.Size()
}

// Stat returns a FileInfo describing the named file.
// If there is an error, it will be of type *PluginError.
func (p *Obj) Stat() (os.FileInfo, error) {
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

func (p *Obj) error(msg ...string) error {
	return errors.Newf("%s (%s)", p.abs, msg)
}
