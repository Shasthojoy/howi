// Copyright 2012 Marko Kungla. All rights reserved.
// Use of this source code is governed by a The Apache-style
// license that can be found in the LICENSE file.

package errors

import (
	"fmt"
	"runtime"
	"strings"
)

// ErrorWithStackTrace implements error and includes stack trace
type ErrorWithStackTrace struct {
	// custom message
	msg    string
	frames []Frame
}

// GetStackTrace returns all stack frames in the current stack trace.
func (stc *ErrorWithStackTrace) GetStackTrace() []Frame {
	return stc.frames
}

// Error implements error.Error method
func (stc *ErrorWithStackTrace) Error() string {
	return stc.msg
}

func (stc *ErrorWithStackTrace) trace() {
	// default max depth
	const d = 32
	var cntr [d]uintptr
	n := runtime.Callers(3, cntr[:])
	stc.frames = make([]Frame, n)
	for i := 0; i < len(stc.frames); i++ {
		stc.frames[i] = newFrame(cntr[i])
	}
}

// Frame of the stack
type Frame struct {
	pc uintptr
	// file returns the full path to the file that contains the
	// function for this Frame's pc.
	file string
	// line returns the line number of source code of the
	// function for this Frame's pc.
	line int
	// name of the package
	pkgName string
	// function name
	fnName string
}

// Line returns file line nr of that frame
func (f Frame) Line() int {
	return f.line
}

// File returns file name reative to $GOPATH/src
func (f Frame) File() string {
	return f.file
}

// Package returns name of the package
func (f Frame) Package() string {
	return f.pkgName
}

// Func returns name of the function
func (f Frame) Func() string {
	return f.fnName
}

// Error implements error.Error method
func (f Frame) String() string {
	return fmt.Sprintf("%s:%d %s.%s", f.file, f.line, f.pkgName, f.fnName)
}

func newFrame(pc uintptr) Frame {
	frame := Frame{}
	frame.pc = uintptr(pc) - 1

	fn := runtime.FuncForPC(frame.pc)
	if fn != nil {
		frame.file, frame.line = fn.FileLine(frame.pc)
		frame.pkgName, frame.fnName = packageInfo(fn.Name())
		frame.file = fixPath(frame.file)
	}
	return frame
}

// packageInfo removes the path prefix from string reported by runtime.FuncForPC(uintptr).Name()
// and returns name of the package and function
func packageInfo(str string) (pkgName, fnName string) {
	if len(str) == 0 {
		return
	}
	pos := strings.LastIndex(str, "/")
	str = str[pos+1:]
	dotpos := strings.Index(str, ".")
	// treat string as fn name if there is no "."
	if dotpos == -1 {
		fnName = str
		return
	}
	pkgName = str[:dotpos]
	fnName = str[dotpos+1:]
	return
}

// Clean the file path
func fixPath(path string) string {
	srci := strings.Index(path, "/src/")
	if srci != -1 {
		path = path[srci+5:]
	}
	return path
}
