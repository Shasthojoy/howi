// Copyright 2005-2017 Marko Kungla. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// license that can be found in the LICENSE file.

package hlog

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"golang.org/x/crypto/ssh/terminal"
)

// A Logger represents an active logging object that generates lines of
// output to an io.Writer. Each logging operation makes a single call to
// the Writer's Write method. A Logger can be used simultaneously from
// multiple goroutines; it guarantees to serialize access to the Writer.
type Logger struct {
	mu           sync.Mutex // ensures atomic writes; protects all Logger fields
	w            io.Writer
	wt           byte
	level        int
	aligned      bool
	colors       bool
	inProgress   bool
	exit         func(int)
	msgBuf       []byte // for accumulating text to write out
	started      timestamp
	ts           *timestamp
	primaryColor []byte
	prfx         []byte
	levelLocked  bool
}

// Colors colirzes output
func (l *Logger) Colors() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.colors = true
}

// ColorsDisable calls std.ColorsDisable
func (l *Logger) ColorsDisable() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.colors = false
	padDef = 2
}

// Exit calls method set with SetExitFunc defaults to os.Exit
func (l *Logger) Exit(code int) {

	if t != nil {
		t.sch <- struct{}{}
		for t.monitoring { /* wait terminal to be restored */
		}
	}
	if l.exit != nil {
		l.exit(code)
	}
}

// InitTerm puts current terminal into raw mode and enables logger to use
// current terminal
func (l *Logger) InitTerm() {
	if t != nil {
		return
	}
	t = &term{}
	t.fd = int(os.Stdout.Fd())
	if terminal.IsTerminal(t.fd) {
		t.size.w, t.size.h, _ = terminal.GetSize(t.fd)
		if (t.size.w + t.size.h) > 0 {
			l.aligned = true
			t.evch = make(chan tsize, 1)
			t.sch = make(chan struct{})
			t.winch = make(chan os.Signal, 1)
			go t.monitor()
		}
	}
}

// isValid checks whether Logger has valid Writer.
// If not, it returns an bool false
func (l *Logger) isValid() bool {
	return l.w != nil
}

// New creates a new Logger. The w variable sets the
// destination to which log data will be written.
// The level argument sets log level
func New(w io.Writer, level int) *Logger {
	l := &Logger{
		w:       w,
		level:   level,
		exit:    os.Exit,
		wt:      t1,
		ts:      &timestamp{},
		started: timestamp{},
		prfx:    []byte(" "),
	}
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	go func() {
		<-sigc
		close(sigc)
		l.Exit(0)
	}()
	l.started.now(t0)
	return l
}

// NewStdout creates a new Logger writin to os.Stdout.
// The level argument sets log level
func NewStdout(level int) *Logger {
	return New(os.Stdout, level)
}

// NewProgress creates and returns new progress bar
func (l *Logger) NewProgress(name string, steps int) *Progress {
	return &Progress{name: name, steps: steps, log: l, started: time.Now()}
}

// SetExitFunc sets exit function defaults to os.Exit
func (l *Logger) SetExitFunc(exit func(code int)) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.exit = exit
}

// GetCurrentLevel returns current log level
func (l *Logger) GetCurrentLevel() int {
	return l.level
}

// SetLogLevel sets log level if loglevel is not locked by previous call
// to .LockLevel
func (l *Logger) SetLogLevel(level int) {
	if (level <= DEBUG || level >= QUIET) && !l.levelLocked {
		l.mu.Lock()
		defer l.mu.Unlock()
		if level == DEBUG {
			debug = true
		}
		l.level = level
	}
}

// LockLevel locks log level so it can not be modified by SetLogLevel again
func (l *Logger) LockLevel() {
	l.levelLocked = true
}

// SetOutput sets the output destination for the logger.
func (l *Logger) SetOutput(w io.Writer) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.w = w
}

// TsDisabled disables timestamping log messages
func (l *Logger) TsDisabled() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.wt = t0
}

// TsStandard enables timestamping of log messages with standard layout
func (l *Logger) TsStandard() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.wt = t1
}

// TsTime calls std.TsTime
func (l *Logger) TsTime() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.wt = t2
}

// PANIC (1)

// Panic tries to write string to writer followed by a call to panic()
// Arguments are handled in the manner of fmt.Println.
func (l *Logger) Panic(v ...interface{}) {
	s := fmt.Sprint(v...)
	if l.level >= PANIC && l.isValid() {
		l.write(s, l.prfx, sfxPanic[:], red)
	}
	panic(s)
}

// Panicf tries to write string to writer followed by a call to panic().
// Arguments are handled in the manner of fmt.Printf followed by \n.
func (l *Logger) Panicf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	if l.level >= PANIC && l.isValid() {
		l.write(s, l.prfx, sfxPanic[:], red)
	}
	panic(s)
}

// FATAL (1)

// Fatal is equivalent to l.Print() followed by a call to os.Exit(1).
// Arguments are handled in the manner of fmt.Println.
func (l *Logger) Fatal(v ...interface{}) {
	if l.level >= FATAL && l.isValid() {
		l.write(fmt.Sprint(v...), l.prfx, sfxFatal[:], red)
	}
	l.Exit(1)
}

// Fatalf is equivalent to l.Printf() followed by a call to os.Exit(1).
// Arguments are handled in the manner of fmt.Printf followed by \n.
func (l *Logger) Fatalf(format string, v ...interface{}) {
	if l.level >= FATAL && l.isValid() {
		l.write(fmt.Sprintf(format, v...), l.prfx, sfxFatal[:], red)
	}
	l.Exit(1)
}

// EMERGENCY (2)

// Emergency performs write to the loggers attached io.Writer.
// Arguments are handled in the manner of fmt.Println.
func (l *Logger) Emergency(v ...interface{}) {
	if l.level >= EMERGENCY && l.isValid() {
		l.write(fmt.Sprint(v...), l.prfx, sfxEmergency[:], red)
	}
}

// Emergencyf performs write to the loggers attached io.Writer.
// Arguments are handled in the manner of fmt.Printf followed by \n.
func (l *Logger) Emergencyf(format string, v ...interface{}) {
	if l.level >= EMERGENCY && l.isValid() {
		l.write(fmt.Sprintf(format, v...), l.prfx, sfxEmergency[:], red)
	}
}

// Deprecated performs write to the loggers attached io.Writer.
// Arguments are handled in the manner of fmt.Println.
// enables you to log and notice package users if any method is deprecated
func (l *Logger) Deprecated(v ...interface{}) {
	if l.level >= EMERGENCY && l.isValid() {
		l.write(fmt.Sprint(v...), l.prfx, sfxDeprecated[:], red)
	}
}

// Deprecatedf performs write to the loggers attached io.Writer.
// Arguments are handled in the manner of fmt.Printf followed by \n.
// enables you to log and notice package users if any method is deprecated
func (l *Logger) Deprecatedf(format string, v ...interface{}) {
	if l.level >= EMERGENCY && l.isValid() {
		l.write(fmt.Sprintf(format, v...), l.prfx, sfxDeprecated[:], red)
	}
}

// ALERT(3)

// Alert performs write to the loggers attached io.Writer.
// Arguments are handled in the manner of fmt.Println.
func (l *Logger) Alert(v ...interface{}) {
	if l.level >= ALERT && l.isValid() {
		l.write(fmt.Sprint(v...), l.prfx, sfxAlert[:], red)

	}
}

// Alertf performs write to the loggers attached io.Writer.
// Arguments are handled in the manner of fmt.Printf followed by \n.
func (l *Logger) Alertf(format string, v ...interface{}) {
	if l.level >= ALERT && l.isValid() {
		l.write(fmt.Sprintf(format, v...), l.prfx, sfxAlert[:], red)
	}
}

// CRITICAL(4)

// Critical performs write to the loggers attached io.Writer.
// Arguments are handled in the manner of fmt.Println.
func (l *Logger) Critical(v ...interface{}) {
	if l.level >= CRITICAL && l.isValid() {
		l.write(fmt.Sprint(v...), l.prfx, sfxCritical[:], red)
	}
}

// Criticalf performs write to the loggers attached io.Writer.
// Arguments are handled in the manner of fmt.Printf followed by \n.
func (l *Logger) Criticalf(format string, v ...interface{}) {
	if l.level >= CRITICAL && l.isValid() {
		l.write(fmt.Sprintf(format, v...), l.prfx, sfxCritical[:], red)
	}
}

// ERROR (5)

// Error performs write to the loggers attached io.Writer.
// Arguments are handled in the manner of fmt.Println.
func (l *Logger) Error(v ...interface{}) {
	if l.level >= ERROR && l.isValid() {
		l.write(fmt.Sprint(v...), l.prfx, sfxError[:], red)
	}
}

// Errorf performs write to the loggers attached io.Writer.
// Arguments are handled in the manner of fmt.Printf followed by \n.
func (l *Logger) Errorf(format string, v ...interface{}) {
	if l.level >= ERROR && l.isValid() {
		l.write(fmt.Sprintf(format, v...), l.prfx, sfxError[:], red)
	}
}

// WARNING (6)

// Warning performs write to the loggers attached io.Writer.
// Arguments are handled in the manner of fmt.Println.
func (l *Logger) Warning(v ...interface{}) {
	if l.level >= WARNING && l.isValid() {
		l.write(fmt.Sprint(v...), l.prfx, sfxWarning[:], yellow)
	}
}

// Warningf performs write to the loggers attached io.Writer.
// Arguments are handled in the manner of fmt.Printf followed by \n.
func (l *Logger) Warningf(format string, v ...interface{}) {
	if l.level >= WARNING && l.isValid() {
		l.write(fmt.Sprintf(format, v...), l.prfx, sfxWarning[:], yellow)
	}
}

// NOTICE (7)

// Notice performs write to the loggers attached io.Writer.
// Arguments are handled in the manner of fmt.Println.
func (l *Logger) Notice(v ...interface{}) {
	if l.level >= NOTICE && l.isValid() {
		l.write(fmt.Sprint(v...), l.prfx, sfxNotice[:], cyan)
	}
}

// Noticef performs write to the loggers attached io.Writer.
// Arguments are handled in the manner of fmt.Printf followed by \n.
func (l *Logger) Noticef(format string, v ...interface{}) {
	if l.level >= NOTICE && l.isValid() {
		l.write(fmt.Sprintf(format, v...), l.prfx, sfxNotice[:], cyan)
	}
}

// Line performs write to the loggers attached io.Writer.
// Arguments are handled in the manner of fmt.Println.
func (l *Logger) Line(v ...interface{}) {
	if l.level >= LINE && l.isValid() {
		l.write(fmt.Sprint(v...), nil, nil, nil)
	}
}

// Linef performs write to the loggers attached io.Writer.
// Arguments are handled in the manner of fmt.Printf followed by \n.
func (l *Logger) Linef(format string, v ...interface{}) {
	if l.level >= LINE && l.isValid() {
		l.write(fmt.Sprintf(format, v...), nil, nil, nil)
	}
}

// INFO (8)

// Info performs write to the loggers attached io.Writer.
// Arguments are handled in the manner of fmt.Println.
func (l *Logger) Info(v ...interface{}) {
	if l.level >= INFO && l.isValid() {
		l.write(fmt.Sprint(v...), l.prfx, sfxInfo[:], cyan)
	}
}

// Infof performs write to the loggers attached io.Writer.
// Arguments are handled in the manner of fmt.Printf followed by \n.
func (l *Logger) Infof(format string, v ...interface{}) {
	if l.level >= INFO && l.isValid() {
		l.write(fmt.Sprintf(format, v...), l.prfx, sfxInfo[:], cyan)
	}
}

// OK(8)

// Ok performs write to the loggers attached io.Writer.
// Arguments are handled in the manner of fmt.Println.
func (l *Logger) Ok(v ...interface{}) {
	if l.level >= OK && l.isValid() {
		l.write(fmt.Sprint(v...), l.prfx, sfxOk[:], green)
	}
}

// Okf performs write to the loggers attached io.Writer.
// Arguments are handled in the manner of fmt.Printf followed by \n.
func (l *Logger) Okf(format string, v ...interface{}) {
	if l.level >= OK && l.isValid() {
		l.write(fmt.Sprintf(format, v...), l.prfx, sfxOk[:], green)
	}
}

// DEBUG (9)

// Debug performs write to the loggers attached io.Writer.
// Arguments are handled in the manner of fmt.Println.
func (l *Logger) Debug(v ...interface{}) {
	if !debug {
		return
	}
	if l.isValid() {
		l.write(fmt.Sprint(v...), l.prfx, sfxDebug[:], white)
	}
}

// Debugf performs write to the loggers attached io.Writer.
// Arguments are handled in the manner of fmt.Printf followed by \n.
func (l *Logger) Debugf(format string, v ...interface{}) {
	if !debug {
		return
	}
	if l.isValid() {
		l.write(fmt.Sprintf(format, v...), l.prfx, sfxDebug[:], white)
	}
}

// SetPrimaryColor sets color for ColoredLine and ColoredLinef
func (l *Logger) SetPrimaryColor(color string) {
	switch color {
	case "black":
		l.primaryColor = black
	case "red":
		l.primaryColor = red
	case "green":
		l.primaryColor = green
	case "yellow":
		l.primaryColor = yellow
	case "blue":
		l.primaryColor = blue
	case "magenta":
		l.primaryColor = magenta
	case "cyan":
		l.primaryColor = cyan
	default:
		l.primaryColor = white
	}
}

// ColoredLine performs write to the loggers attached io.Writer.
// Arguments are handled in the manner of fmt.Println.
// line is colored with color set by SetPrimaryColor
func (l *Logger) ColoredLine(v ...interface{}) {
	if l.isValid() && l.level >= LINE {
		msg := fmt.Sprint(v...)
		if l.colors {
			msg = string(l.primaryColor) + msg + string(reset)
		}
		l.write(msg, nil, nil, nil)
	}
}

// ColoredLinef performs write to the loggers attached io.Writer.
// Arguments are handled in the manner of fmt.Printf followed by \n.
// line is colored with color set by SetPrimaryColor
func (l *Logger) ColoredLinef(format string, v ...interface{}) {
	if l.isValid() && l.level >= LINE {
		msg := fmt.Sprintf(format, v...)
		if l.colors {
			msg = string(l.primaryColor) + msg + string(reset)
		}
		l.write(msg, nil, nil, nil)
	}
}

// write writes the output for a logging event. The string s contains
func (l *Logger) write(s string, prfx []byte, suffix []byte, color []byte) error {

	if l.colors && color != nil && suffix != nil {
		suffix = append(color, suffix...)
		suffix = append(suffix, reset...)
		padDef = 11
	}
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.wt > t0 && prfx != nil {
		ts := l.ts.now(l.wt)
		prfx = append(ts, prfx...)
	}
	l.msgBuf = l.msgBuf[:0]
	if l.inProgress {
		// Delete current progressbar
		l.msgBuf = append(l.msgBuf, _cr)
	}
	if !l.aligned && suffix != nil {
		prfx = append(prfx, suffix...)
	}
	l.msgBuf = append(l.msgBuf, prfx...)
	l.msgBuf = append(l.msgBuf, 32)
	l.msgBuf = append(l.msgBuf, s...)

	if l.aligned && suffix != nil {
		bodyLen := len(l.msgBuf)
		sfxLen := len(suffix)
		termW := Width()
		padLen := termW - bodyLen - sfxLen
		r := padLen + padDef
		if r > 0 {
			pad := bytes.Repeat([]byte{32}, r)
			l.msgBuf = append(l.msgBuf, pad...)
		}

		l.msgBuf = append(l.msgBuf, suffix...)
	}
	l.msgBuf = append(l.msgBuf, _lf)
	_, err := l.w.Write(l.msgBuf)
	return err
}

func (l *Logger) printProgress(name string, pct float32, started time.Time) {
	isDone := int(pct) == 100
	if isDone {
		elapsed := time.Now().Sub(started)
		l.write(fmt.Sprintf("%s [100%% elapsed %s]", name, elapsed.String()), nil, sfxOk[:], green)
		return
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	l.inProgress = true
	l.msgBuf = l.msgBuf[:0]
	l.msgBuf = append(l.msgBuf, _cr)
	termW := Width()

	l.msgBuf = append(l.msgBuf, name...)
	l.msgBuf = append(l.msgBuf, 32)
	bodyLen := len(l.msgBuf)

	suffix := fmt.Sprintf(" %.2f%%", pct)
	sfxLen := len(suffix)
	barLen := float32(termW-bodyLen-sfxLen) / 100 * pct
	if barLen > 0 {
		bar := bytes.Repeat([]byte{'#'}, int(barLen))
		l.msgBuf = append(l.msgBuf, bar...)
	}
	padLen := termW - len(l.msgBuf) - sfxLen
	if padLen > 0 {
		pad := bytes.Repeat([]byte{' '}, padLen)
		l.msgBuf = append(l.msgBuf, pad...)
	}
	l.msgBuf = append(l.msgBuf, suffix...)
	l.w.Write(l.msgBuf)
}

// Cheap integer to fixed-width decimal ASCII.
// Give a negative width to avoid zero-padding.
// modified  go/src/log/log.go itoa
func itoa(buf *[]byte, i int, wid int) {
	// Assemble decimal in reverse order.
	var b [10]byte
	bp := len(b) - 1
	for i >= 10 || wid > 1 {
		wid--
		q := i / 10
		b[bp] = byte('0' + i - q*10)
		bp--
		i = q
	}
	b[bp] = byte('0' + i)
	*buf = append(*buf, b[bp:]...)

}
