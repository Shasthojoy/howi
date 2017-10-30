// Copyright 2005-2017 Marko Kungla.
// Use of this source code is governed by a The MIT License
// license that can be found in the LICENSE file.

package flags

import "testing"

func TestIsHidden(t *testing.T) {
	flag := NewBoolFlag("some-flag", "sf")
	if flag.IsHidden() {
		t.Error("by default flag must not mbe hidden")
	}
}
func TestHelpName(t *testing.T) {
	flag := NewBoolFlag("some-flag", "sf")
	actual := flag.HelpName()
	if actual != "--some-flag" {
		t.Errorf(".HelpName want = --some-flag, got = %q", actual)
	}
}
