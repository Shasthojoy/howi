package hlog

import (
	"io"
	"os"
)

const (
	// QUIET (0) Does not log anything except Fatal and Panic
	QUIET = 0
	// PANIC (1) Will call panic right after writing the message
	PANIC = 1
	// FATAL (1) will handle emergency exit with os.Exit(1) after writing the message
	FATAL = PANIC
	// EMERGENCY (2) Service is unusable
	EMERGENCY = iota - 1
	// ALERT (3) Condition or event occurred which requires that
	// immediate action will be taken.
	ALERT
	// CRITICAL (4) Condition or event occurred which may be followed
	// by ALERT if action is not taken
	CRITICAL
	// ERROR (5) Condition or event occurred which does not require
	// immediate action, but should typically be logged and monitored.
	ERROR
	// WARNING (6) Exceptional occurrences which are not errors and current
	// instance can resolve that exception and which does not require any action.
	WARNING
	// NOTICE (7) Normal but significant events.
	NOTICE
	// INFO (8) Interesting events for monitoring.
	INFO
	// OK (8) Presents another type of interesting event upon success level = INFO
	OK = INFO
	// DEBUG (9) is detailed debug information
	DEBUG = INFO + 1
	// LINE (7) Normal non verbose output line
	LINE = NOTICE
	// CR
	_cr uint8 = 13
	// newline
	_lf uint8 = 10
	t0  uint8 = 0 // Disable timestamp
	t1  uint8 = 1 // Prefix standard timestamp
	t2  uint8 = 2 // Prefix time only
)

var (
	std          = New(os.Stderr, INFO)
	sfxPanic     = [17]byte{ /*\u2718*/ 91, 32, 226, 156, 152, 32, 112, 97, 110, 105, 99, 32, 32, 32, 32, 32, 93}
	sfxFatal     = [17]byte{ /*\u2718*/ 91, 32, 226, 156, 152, 32, 102, 97, 116, 97, 108, 32, 32, 32, 32, 32, 93}
	sfxEmergency = [17]byte{ /*\u2717*/ 91, 32, 226, 156, 151, 32, 101, 109, 101, 114, 103, 101, 110, 99, 121, 32, 93}
	sfxAlert     = [17]byte{ /*\u2717*/ 91, 32, 226, 156, 151, 32, 97, 108, 101, 114, 116, 32, 32, 32, 32, 32, 93}
	sfxCritical  = [17]byte{ /*\u2717*/ 91, 32, 226, 156, 151, 32, 99, 114, 105, 116, 105, 99, 97, 108, 32, 32, 93}
	sfxError     = [17]byte{ /*\u2717*/ 91, 32, 226, 156, 151, 32, 101, 114, 114, 111, 114, 32, 32, 32, 32, 32, 93}
	sfxWarning   = [17]byte{ /*\u26A0*/ 91, 32, 226, 154, 160, 32, 119, 97, 114, 110, 105, 110, 103, 32, 32, 32, 93}
	sfxNotice    = [17]byte{ /*\u26A0*/ 91, 32, 226, 154, 160, 32, 110, 111, 116, 105, 99, 101, 32, 32, 32, 32, 93}
	sfxInfo      = [17]byte{ /*\u26A0*/ 91, 32, 226, 154, 160, 32, 105, 110, 102, 111, 32, 32, 32, 32, 32, 32, 93}
	sfxOk        = [17]byte{ /*\u2714*/ 91, 32, 226, 156, 148, 32, 111, 107, 32, 32, 32, 32, 32, 32, 32, 32, 93}
	sfxDebug     = [17]byte{ /*\u2699*/ 91, 32, 226, 154, 153, 32, 100, 101, 98, 117, 103, 32, 32, 32, 32, 32, 93}
	esc          = byte(27)
	black        = []byte{esc, '[', '3', '0', 'm'}
	red          = []byte{esc, '[', '3', '1', 'm'}
	green        = []byte{esc, '[', '3', '2', 'm'}
	yellow       = []byte{esc, '[', '3', '3', 'm'}
	blue         = []byte{esc, '[', '3', '4', 'm'}
	magenta      = []byte{esc, '[', '3', '5', 'm'}
	cyan         = []byte{esc, '[', '3', '6', 'm'}
	white        = []byte{esc, '[', '3', '7', 'm'}
	reset        = []byte{esc, '[', '0', 'm'}
	padDef       = 2
)

// Colors calls std.Colors
func Colors() {
	std.Colors()
}

// ColorsDisable calls std.ColorsDisable
func ColorsDisable() {
	std.ColorsDisable()
}

// Exit calls std.Exit
func Exit(code int) {
	std.Exit(code)
}

// InitTerm sets terminal to raw mode and aligns output based on terminal width
func InitTerm() {
	std.InitTerm()
}

// NewProgress calls std.NewProgress
func NewProgress(name string, steps int) *Progress {
	return std.NewProgress(name, steps)
}

// SetExitFunc calls std.SetExitFunc
func SetExitFunc(exit func(code int)) {
	std.SetExitFunc(exit)
}

// SetLogLevel calls std.SetLogLevel
func SetLogLevel(level int) {
	std.SetLogLevel(level)
}

// SetOutput calls std.SetOutput
func SetOutput(w io.Writer) {
	std.SetOutput(w)
}

// TsDisabled calls std.TsDisabled
func TsDisabled() {
	std.TsDisabled()
}

// TsStandard calls std.TsStandard
func TsStandard() {
	std.TsStandard()
}

// TsTime calls std.TsTime
func TsTime() {
	std.TsTime()
}

// Panic calls std.Panic
func Panic(v ...interface{}) {
	std.Panic(v...)
}

// Panicf calls std.Panicf
func Panicf(format string, v ...interface{}) {
	std.Panicf(format, v...)
}

// Fatal calls std.Fatal
func Fatal(v ...interface{}) {
	std.Fatal(v...)
}

// Fatalf calls std.Fatalf
func Fatalf(format string, v ...interface{}) {
	std.Fatalf(format, v...)
}

// Emergency calls std.Emergency
func Emergency(v ...interface{}) {
	std.Emergency(v...)
}

// Emergencyf calls std.Emergencyf
func Emergencyf(format string, v ...interface{}) {
	std.Emergencyf(format, v...)
}

// Alert calls std.Alert
func Alert(v ...interface{}) {
	std.Alert(v...)
}

// Alertf calls std.Alertf
func Alertf(format string, v ...interface{}) {
	std.Alertf(format, v...)
}

// Critical calls std.Critical
func Critical(v ...interface{}) {
	std.Critical(v...)
}

// Criticalf calls std.Criticalf
func Criticalf(format string, v ...interface{}) {
	std.Criticalf(format, v...)
}

// Error calls std.Error
func Error(v ...interface{}) {
	std.Error(v...)
}

// Errorf calls std.Errorf
func Errorf(format string, v ...interface{}) {
	std.Errorf(format, v...)
}

// Warning calls std.Warning
func Warning(v ...interface{}) {
	std.Warning(v...)
}

// Warningf calls std.Warningf
func Warningf(format string, v ...interface{}) {
	std.Warningf(format, v...)
}

// Notice calls std.Notice
func Notice(v ...interface{}) {
	std.Notice(v...)
}

// Noticef calls std.Noticef
func Noticef(format string, v ...interface{}) {
	std.Noticef(format, v...)
}

// Line calls std.Line
func Line(v ...interface{}) {
	std.Line(v...)
}

// Linef calls std.Linef
func Linef(format string, v ...interface{}) {
	std.Linef(format, v...)
}

// Info calls std.Infof
func Info(v ...interface{}) {
	std.Info(v...)
}

// Infof calls std.Infof
func Infof(format string, v ...interface{}) {
	std.Infof(format, v...)
}

// Ok calls std.Ok
func Ok(v ...interface{}) {
	std.Ok(v...)
}

// Okf calls std.Okf
func Okf(format string, v ...interface{}) {
	std.Okf(format, v...)
}

// Debug calls std.Debug
func Debug(v ...interface{}) {
	std.Debug(v...)
}

// Debugf calls std.Debugf
func Debugf(format string, v ...interface{}) {
	std.Debugf(format, v...)
}

// SetPrimaryColor calls std.SetPrimaryColor
func SetPrimaryColor(color string) {
	std.SetPrimaryColor(color)
}

// ColoredLine calls std.ColoredLine
func ColoredLine(v ...interface{}) {
	std.ColoredLine(v...)
}

// ColoredLinef calls std.ColoredLinef
func ColoredLinef(format string, v ...interface{}) {
	std.ColoredLinef(format, v...)
}
