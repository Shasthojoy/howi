// Copyright 2016 Marko Kungla. All rights reserved.
// Use of this source code is governed by a The Apache-style
// license that can be found in the LICENSE file.

package log

import "time"

// Progress provides progress bar
type Progress struct {
	name    string
	steps   int
	step    int
	pct     float32
	log     *Logger
	started time.Time
}

// Done reports has progres made enough steps based on provided step count
func (p *Progress) Done() bool {
	done := p.steps == p.step
	if done {
		p.log = nil
	}
	return done
}

// Next increments current step counter
func (p *Progress) Next() {
	p.step++
	p.pct = 100 / float32(p.steps) * float32(p.step)
	if p.log != nil {
		p.log.printProgress(p.name, p.pct, p.started)
	}
}
