// Copyright 2005-2017 Marko Kungla. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// license that can be found in the LICENSE file.

package log_test

import (
	"bytes"
	"log"
	"testing"

	hlog "github.com/howi-ce/howi/std/log"
)

const testString = "Dummy log message to have something to log, length of this string is avarage message length."

func BenchmarkHOWIloggerStd(b *testing.B) {
	var buf bytes.Buffer
	hlog.SetOutput(&buf)
	for i := 0; i < b.N; i++ {
		buf.Reset()
		hlog.Line(testString)
	}
}
func BenchmarkGoLoggerStd(b *testing.B) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	for i := 0; i < b.N; i++ {
		buf.Reset()
		log.Println(testString)
	}
}

func BenchmarkHOWIlogger(b *testing.B) {
	var buf bytes.Buffer
	l := hlog.New(&buf, hlog.DEBUG)
	for i := 0; i < b.N; i++ {
		buf.Reset()
		l.Line(testString)
	}
}

func BenchmarkGoLogger(b *testing.B) {
	var buf bytes.Buffer
	l := log.New(&buf, "", log.LstdFlags)
	for i := 0; i < b.N; i++ {
		buf.Reset()
		l.Println(testString)
	}
}

func BenchmarkHOWIloggerRecreate(b *testing.B) {
	var buf bytes.Buffer
	for i := 0; i < b.N; i++ {
		buf.Reset()
		l := hlog.New(&buf, hlog.DEBUG)
		l.Line(testString)
	}
}

func BenchmarkGoLoggerRecreate(b *testing.B) {
	var buf bytes.Buffer
	for i := 0; i < b.N; i++ {
		buf.Reset()
		l := log.New(&buf, "", log.LstdFlags)
		l.Println(testString)
	}
}
func BenchmarkHOWIloggerNoTime(b *testing.B) {
	var buf bytes.Buffer
	hlog.SetOutput(&buf)
	hlog.TsDisabled()
	for i := 0; i < b.N; i++ {
		buf.Reset()
		hlog.Line(testString)
	}
}

func BenchmarkHOWIWithTimeAndResize(b *testing.B) {
	var buf bytes.Buffer
	hlog.SetOutput(&buf)
	hlog.TsStandard()
	hlog.InitTerm()
	for i := 0; i < b.N; i++ {
		buf.Reset()
		hlog.Line(testString)
	}
}
