// Copyright 2005-2017 Marko Kungla. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// license that can be found in the LICENSE file.

package log

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"
)

const testString = "Dummy log message to have something to log, length of this string is avarage message length."

func TestSetLogLevel(t *testing.T) {
	type args struct {
		level int
	}
	tests := []struct {
		name string
		args args
	}{
		{"QUIET", args{QUIET}},
		{"PANIC", args{PANIC}},
		{"FATAL", args{FATAL}},
		{"EMERGENCY", args{EMERGENCY}},
		{"ALERT", args{ALERT}},
		{"CRITICAL", args{CRITICAL}},
		{"ERROR", args{ERROR}},
		{"WARNING", args{WARNING}},
		{"NOTICE", args{NOTICE}},
		{"INFO", args{INFO}},
		{"OK", args{OK}},
		{"DEBUG", args{DEBUG}},
		{"LINE", args{LINE}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetLogLevel(tt.args.level)
		})
	}
}

func TestNewStdout(t *testing.T) {
	l := NewStdout(INFO)
	if l.w != os.Stdout {
		t.Fatal("should have os.Stdout")
	}
}

func TestLogger_SetOutput(t *testing.T) {
	var buf bytes.Buffer
	NewStdout(DEBUG)
	SetOutput(&buf)
	Line("TestLogger_SetOutput")
	if !strings.Contains(buf.String(), "TestLogger_SetOutput") {
		t.Error("buffer should contain log message")
	}
}

func TestLogger_TsDisabled(t *testing.T) {
	now := time.Now()
	date := now.Format("2006-01-02")
	var buf bytes.Buffer
	SetOutput(&buf)
	TsStandard()
	Notice("TestLogger_TsDisabled1")
	if !strings.Contains(buf.String(), date) {
		t.Errorf("buffer should contain date got %q", buf.String())
	}
	buf.Reset()
	TsDisabled()
	Notice("TestLogger_TsDisabled1")
	if strings.Contains(buf.String(), date) {
		t.Error("buffer should not contain date")
	}
}

func TestLogger_TsStandard(t *testing.T) {
	now := time.Now()
	date := now.Format("2006-01-02 15:")
	var buf bytes.Buffer
	SetOutput(&buf)
	TsStandard()
	Notice("TestLogger_TsStandard")
	if !strings.Contains(buf.String(), date) {
		t.Errorf("buffer should contain date got %q", buf.String())
	}
}

func TestLogger_TsTime(t *testing.T) {
	now := time.Now()
	date := now.Format("2006-01-02")
	time := now.Format("15:")
	var buf bytes.Buffer
	SetOutput(&buf)
	TsTime()
	Notice("TestLogger_TsTime")
	if strings.Contains(buf.String(), date) {
		t.Error("buffer should not contain date")
	}
	if !strings.Contains(buf.String(), time) {
		t.Errorf("buffer should contain time got %q", buf.String())
	}
}

func TestSetExitFunc(t *testing.T) {
	e := func(d int) { fmt.Print("") }
	std.exit = nil
	SetExitFunc(e)
	if std.exit == nil {
		t.Error("exit function should not be nil")
	}
}

func TestPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			if _, ok := r.(error); ok {
				t.Errorf("should panic with message ok(%t)", ok)
			}
		}

	}()
	Panic("the msg")
}

func TestPanicf(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			if _, ok := r.(error); ok {
				t.Errorf("should panic with message ok(%t)", ok)
			}
		}

	}()
	Panicf("the %s", "msg")
}

func TestFatal(t *testing.T) {
	t.Run("exits", func(t *testing.T) {
		SetExitFunc(func(code int) {})
		Fatal("some message")
	})
}

func TestFatalf(t *testing.T) {
	t.Run("exits", func(t *testing.T) {
		SetExitFunc(func(code int) {})
		Fatalf("some message")
	})
}

func TestEmergencyf(t *testing.T) {
	var buf bytes.Buffer
	SetOutput(&buf)
	Colors()
	Emergencyf("some msg")
	if !strings.Contains(buf.String(), string(red)) {
		t.Error("Emergencyf want color red")
	}
	if !strings.Contains(buf.String(), "emergency") {
		t.Error("Emergencyf want substr emergency")
	}
}

func TestEmergency(t *testing.T) {
	var buf bytes.Buffer
	SetOutput(&buf)
	Colors()
	Emergency("some msg")
	if !strings.Contains(buf.String(), string(red)) {
		t.Error("Emergencyf want color red")
	}
	if !strings.Contains(buf.String(), "emergency") {
		t.Error("Emergencyf want substr emergency")
	}
}
func TestAlertf(t *testing.T) {
	var buf bytes.Buffer
	SetOutput(&buf)
	Colors()
	Alertf("some msg")
	if !strings.Contains(buf.String(), string(red)) {
		t.Error("Alertf want color red")
	}
	if !strings.Contains(buf.String(), "alert") {
		t.Error("Alertf want substr alert")
	}
}
func TestAlert(t *testing.T) {
	var buf bytes.Buffer
	SetOutput(&buf)
	Colors()
	Alert("some msg")
	if !strings.Contains(buf.String(), string(red)) {
		t.Error("Alertf want color red")
	}
	if !strings.Contains(buf.String(), "alert") {
		t.Error("Alertf want substr alert")
	}
}

func TestCritical(t *testing.T) {
	var buf bytes.Buffer
	SetOutput(&buf)
	Colors()
	Critical("some msg")
	if !strings.Contains(buf.String(), string(red)) {
		t.Error("Critical want color red")
	}
	if !strings.Contains(buf.String(), "critical") {
		t.Error("Critical want substr critical")
	}
}

func TestCriticalf(t *testing.T) {
	var buf bytes.Buffer
	SetOutput(&buf)
	Colors()
	Criticalf("some msg")
	if !strings.Contains(buf.String(), string(red)) {
		t.Error("Criticalf want color red")
	}
	if !strings.Contains(buf.String(), "critical") {
		t.Error("Criticalf want substr critical")
	}
}

func TestError(t *testing.T) {
	var buf bytes.Buffer
	SetOutput(&buf)
	Colors()
	Error("some msg")
	if !strings.Contains(buf.String(), string(red)) {
		t.Error("Error want color red")
	}
	if !strings.Contains(buf.String(), "error") {
		t.Error("Error want substr error")
	}
}

func TestErrorf(t *testing.T) {
	var buf bytes.Buffer
	SetOutput(&buf)
	Colors()
	Errorf("some msg")
	if !strings.Contains(buf.String(), string(red)) {
		t.Error("Errorf want color red")
	}
	if !strings.Contains(buf.String(), "error") {
		t.Error("Errorf want substr error")
	}
}

func TestWarning(t *testing.T) {
	var buf bytes.Buffer
	SetOutput(&buf)
	Colors()
	Warning("some msg")
	if !strings.Contains(buf.String(), string(yellow)) {
		t.Error("Warning want color yellow")
	}
	if !strings.Contains(buf.String(), "warning") {
		t.Error("Warning want substr warning")
	}
}

func TestWarningf(t *testing.T) {
	var buf bytes.Buffer
	SetOutput(&buf)
	Colors()
	Warningf("some msg")
	if !strings.Contains(buf.String(), string(yellow)) {
		t.Error("Warningf want color yellow")
	}
	if !strings.Contains(buf.String(), "warning") {
		t.Error("Warningf want substr warning")
	}
}

func TestNotice(t *testing.T) {
	var buf bytes.Buffer
	SetOutput(&buf)
	Colors()
	Notice("some msg")
	if !strings.Contains(buf.String(), string(cyan)) {
		t.Error("Notice want color cyan")
	}
	if !strings.Contains(buf.String(), "notice") {
		t.Error("Notice want substr notice")
	}
}
func TestNoticef(t *testing.T) {
	var buf bytes.Buffer
	SetOutput(&buf)
	Colors()
	Noticef("some msg")
	if !strings.Contains(buf.String(), string(cyan)) {
		t.Error("Noticef want color cyan")
	}
	if !strings.Contains(buf.String(), "notice") {
		t.Error("Noticef want substr notice")
	}
}
func TestLine(t *testing.T) {
	var buf bytes.Buffer
	SetOutput(&buf)
	Colors()
	Line(testString)
	if !strings.Contains(buf.String(), testString) {
		t.Errorf("Line wants (%q)", testString)
	}
}
func TestLinef(t *testing.T) {
	var buf bytes.Buffer
	SetOutput(&buf)
	Colors()
	Linef("some %d %t %s", 1, true, "a")
	if !strings.Contains(buf.String(), "some 1 true a") {
		t.Error("Linef want substr some 1 true a")
	}
}

func TestInfo(t *testing.T) {
	var buf bytes.Buffer
	SetOutput(&buf)
	SetLogLevel(INFO)
	Info(testString)
	if !strings.Contains(buf.String(), "info") {
		t.Errorf("Info want substr info")
	}
	if !strings.Contains(buf.String(), string(cyan)) {
		t.Error("Info want color cyan")
	}
}

func TestInfof(t *testing.T) {
	var buf bytes.Buffer
	SetOutput(&buf)
	SetLogLevel(INFO)
	Infof(testString)
	if !strings.Contains(buf.String(), testString) {
		t.Errorf("Infof want substr (%q)", testString)
	}
	if !strings.Contains(buf.String(), string(cyan)) {
		t.Error("Infof want color cyan")
	}
}

func TestOk(t *testing.T) {
	var buf bytes.Buffer
	SetOutput(&buf)
	SetLogLevel(INFO)
	Ok(testString)
	if !strings.Contains(buf.String(), "ok") {
		t.Errorf("Ok want substr (%q)", testString)
	}
	if !strings.Contains(buf.String(), string(green)) {
		t.Error("Ok want color green")
	}
}

func TestOkf(t *testing.T) {
	var buf bytes.Buffer
	SetOutput(&buf)
	SetLogLevel(INFO)
	Okf("some")
	if !strings.Contains(buf.String(), "ok") {
		t.Error("Okf want substr ok")
	}
	if !strings.Contains(buf.String(), string(green)) {
		t.Error("Okf want color green")
	}
}

func TestDebug(t *testing.T) {
	var buf bytes.Buffer
	SetOutput(&buf)
	SetLogLevel(DEBUG)
	Debug(testString)
	if !strings.Contains(buf.String(), "debug") {
		t.Errorf("Debug want substr (%q)", testString)
	}
}

func TestDebugf(t *testing.T) {
	var buf bytes.Buffer
	SetOutput(&buf)
	SetLogLevel(DEBUG)
	Debugf("some")
	if !strings.Contains(buf.String(), "debug") {
		t.Errorf("Debugf want substr debug got %q", buf.String())
	}
}

func TestItoa(t *testing.T) {
	dst := make([]byte, 0, 64)

	tests := []struct {
		want string
		i    int
		wid  int
		err  string
	}{
		{"2017", 2017, 4, "Wrong year"},
		{"01", 1, 2, "Wrong month"},
		{"30", 30, 2, "Wrong day"},
		{"23", 23, 2, "Wrong hour"},
		{"59", 59, 2, "Wrong minute"},
		{"00", 0, 2, "Wrong second"},
		{"987654", 987654, 6, "Wrong microsecond"},
		{"1234567890", 1234567890, 10, "Wrong result"},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			dst = dst[0:0]
			itoa(&dst, tt.i, tt.wid)
			if tt.want != string(dst) {
				t.Errorf("%s (%d, %d) want (%q) got = %s", tt.err, tt.i, tt.wid, tt.want, string(dst))
			}
		})
	}
}
