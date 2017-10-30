// Copyright 2005-2017 Marko Kungla.
// Use of this source code is governed by a The MIT License
// license that can be found in the LICENSE file.

package cli

import (
	"sync"
	"time"

	"github.com/howi-ce/howi/addon/application/plugin/cli/flags"
	"github.com/howi-ce/howi/lib/appinfo"
	"github.com/howi-ce/howi/std/errors"
	"github.com/howi-ce/howi/std/log"
	"github.com/howi-ce/howi/std/strings"
	"github.com/howi-ce/howi/std/vars"
)

// NewWorker constructs new worker
func newWorker(logger *log.Logger) *Worker {
	w := &Worker{
		started:      time.Now(),
		phases:       make(map[string]*Phase),
		taskPayloads: make(map[string]chan []byte),
		Log:          logger,
		Config: WorkerConfig{
			ShowHeader: true,
			ShowFooter: true,
		},
	}
	w.phases["before"] = newPhase("before")
	w.phases["do"] = newPhase("do")
	w.phases["after-failure"] = newPhase("after-failure")
	w.phases["after-success"] = newPhase("after-success")
	w.phases["after-always"] = newPhase("after-always")
	return w
}

// Worker is instance shared between command phases
type Worker struct {
	mu           sync.Mutex // ensures atomic writes; protects worker fields
	wg           sync.WaitGroup
	started      time.Time
	phase        string
	phases       map[string]*Phase
	taskPayloads map[string]chan []byte
	args         []vars.Value
	flags        map[int]flags.Interface // global flags
	flagAliases  map[string]int          // global flag aliases
	Log          *log.Logger
	Config       WorkerConfig
	Info         appinfo.Info
}

// Fail marks phase as failed
func (w *Worker) Fail(msg string) {
	w.Phase().msg = msg
	w.Phase().status = StatusFailed
}

// Failed returns true if tasks in current phase have failed
func (w *Worker) Failed() bool {
	return w.Phase().status == StatusFailed
}

// Task for worker
func (w *Worker) Task(name string, wt func(task *Task)) {
	if w.Phase().status == StatusFailed {
		w.Log.Warningf("skipping task %q since previous taks failed.", name)
		return
	}
	// Check task name and exit on failure
	if !strings.IsNamespace(name) {
		w.Log.Fatalf("task name %q is invalid - must match following regex %v",
			name, strings.NamespaceMustCompile)
	}
	if _, exists := w.taskPayloads[name]; exists {
		w.Log.Fatalf("task name %q is already in use", name)
	}
	w.mu.Lock()
	w.taskPayloads[name] = make(chan []byte)
	w.mu.Unlock()

	w.wg.Add(1)
	t := &Task{}
	go func() {
		defer func() {
			w.wg.Done()
			w.taskPayloads[name] <- t.payload
			close(w.taskPayloads[name])
		}()
		wt(t)
		// Mark phase as failed if task failed without AllowFailure
		if t.status == StatusFailed {
			w.Fail(t.msg)
		}
	}()

}

// Wait for all previous tasks to complete before scheduling next task
func (w *Worker) Wait() {
	w.Log.Debug("waiting running tasks to complete before command can proceed")
	w.wg.Wait()

}

// Phase information
func (w *Worker) Phase() *Phase {
	return w.phases[w.phase]
}

// Args returns arguments passed to command.
func (w *Worker) Args() []vars.Value {
	return w.args
}

// Flag looks up flag by name or alias and returns flags.Interface.
// If no flag was found error (nil, ErrUnknownFlag) will be returned
func (w *Worker) Flag(alias string) (flags.Interface, error) {
	if w.flagAliases != nil {
		if id, exists := w.flagAliases[alias]; exists {
			return w.flags[id], nil
		}
	}
	return nil, errors.Newf(FmtErrUnknownFlag, alias)
}

// WaitTaskPayloadFrom enables you to wait payload from specific tasks.
// Only one task can consume the payload, Having more consumers expecting same payload
// then first consumer should pass it as it's payload.
func (w *Worker) WaitTaskPayloadFrom(name string) ([]byte, error) {
	w.mu.Lock()
	_, isChan := w.taskPayloads[name]
	w.mu.Unlock()
	if isChan {
		payload := <-w.taskPayloads[name]
		return payload, nil
	}

	return nil, errors.Newf("no such task registered %q", name)
}

// wait for the phase to return
func (w *Worker) phasewait() {
	w.Log.Debugf("phase: %s status: %s, started: %s",
		w.Phase().Name(), w.Phase().Status(), w.Phase().started.String())

	w.wg.Wait()
	w.Phase().finish()
	w.Log.Debugf("phase: %s status: %s, elapsed: %s", w.Phase().Name(),
		w.Phase().Status(), w.Phase().Elapsed())
}

func (w *Worker) attachFlag(f flags.Interface) {
	if w.flags == nil {
		w.flags = make(map[int]flags.Interface)
		w.flagAliases = make(map[string]int)
	}
	nextFlagID := len(w.flags) + 1
	w.flags[nextFlagID] = f
	// create flag aliases
	for _, alias := range f.GetAliases() {
		w.flagAliases[alias] = nextFlagID
	}
}

// WorkerConfig enables some application runtime configuration options
// which can be changed within command phases and task functions.
type WorkerConfig struct {
	ShowHeader bool
	ShowFooter bool
}

func newPhase(name string) *Phase {
	return &Phase{
		name:   name,
		status: StatusPending,
	}
}

// Phase tracks execution of specific phase
type Phase struct {
	started    time.Time
	finished   time.Time
	status     uint
	msg        string
	name       string
	totalTasks int
}

// Name returns name of the phase
func (p *Phase) Name() string {
	return p.name
}

// Elapsed returns how long phase has been running
func (p *Phase) Elapsed() string {
	if p.status == StatusRunning {
		return time.Now().Sub(p.started).String()
	}
	return p.finished.Sub(p.started).String()
}

// Status returns string representation of current phase status
func (p *Phase) Status() (status string) {
	switch p.status {
	case StatusPending:
		status = "pending"
	case StatusRunning:
		status = "running"
	case StatusSuccess:
		status = "success"
	case StatusSkipped:
		status = "skipped"
	default: // PhaseFailed
		status = "failed"
	}
	return
}

func (p *Phase) start() {
	p.started = time.Now()
	p.status = StatusRunning
}

func (p *Phase) finish() {
	p.finished = time.Now()
	if p.status == StatusRunning {
		p.status = StatusSuccess
	}
}

// Task is single task which will be executed in it's own go routine
// within the execution phase it was attached to.
type Task struct {
	name         string
	payload      []byte
	status       uint
	msg          string
	allowFailure bool
}

// Name returns the name of the task
func (t *Task) Name() string {
	return t.name
}

// SetPayload sets payload which can be retrieved by next tasks or phases.
func (t *Task) SetPayload(p []byte) {
	t.payload = p
}

// AllowFailure marks this task to be allowed to fail
func (t *Task) AllowFailure() {
	t.allowFailure = true
}

// Fail marks tasks as failed it updates status only if AllowFailure was not called
func (t *Task) Fail(msg string) {
	t.msg = msg
	if !t.allowFailure {
		t.status = StatusFailed
	}
}

// Failed returns true if task has failed
func (t *Task) Failed() bool {
	return t.status == StatusFailed
}
