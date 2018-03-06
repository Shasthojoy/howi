// Copyright 2016 Marko Kungla. All rights reserved.
// Use of this source code is governed by a The Apache-style
// license that can be found in the LICENSE file.

package flags

import "testing"

func TestName(t *testing.T) {
	flag := NewBoolFlag("some-flag", "sf")
	actual := flag.HelpName()
	if flag.Name() != "some-flag" {
		t.Errorf(".Name want = some-flag, got = %q", flag.Name())
	}
	if actual != "--some-flag" {
		t.Errorf(".HelpName want = --some-flag, got = %q", actual)
	}

	flag2 := NewBoolFlag("s")
	actual2 := flag2.HelpName()
	if flag2.Name() != "s" {
		t.Errorf(".Name want = some-flag, got = %q", flag.Name())
	}
	if actual2 != "-s" {
		t.Errorf(".HelpName want = -s, got = %q", actual2)
	}
}

func TestUsage(t *testing.T) {
	flag := NewBoolFlag("some-flag")
	flag.SetUsage("description")
	if flag.Usage() != "description" {
		t.Errorf("Usage() want 'description' got %q", flag.Usage())
	}
	flag.SetUsagef("description %d", 2)
	if flag.Usage() != "description 2" {
		t.Errorf("Usage() want 'description 2' got %q", flag.Usage())
	}
}

func TestHelpAliases(t *testing.T) {
	flag := NewBoolFlag("some-flag")
	if flag.HelpAliases() != "" {
		t.Errorf("HelpAliases should be empty got %q", flag.HelpAliases())
	}

	flag2 := NewBoolFlag("some-flag", "a", "bbbb")
	if flag2.HelpAliases() != "-a,--bbbb" {
		t.Errorf("HelpAliases want -a,--bbbb, got %q", flag2.HelpAliases())
	}
}

func TestHidden(t *testing.T) {
	flag := NewBoolFlag("some-flag")
	if flag.IsHidden() {
		t.Error("by default flag must not be hidden")
	}
	flag.Hide()
	if !flag.IsHidden() {
		t.Error("flag should be hidden now")
	}
}

func TestGlobal(t *testing.T) {
	flag := NewBoolFlag("some-flag")
	if flag.IsGlobal() {
		t.Error("flag should not be gloabal by default")
	}
}

func TestPos(t *testing.T) {
	flag := NewBoolFlag("some-flag")
	if flag.Pos() != 0 {
		t.Errorf("flag.Pos want 0 got %d", flag.Pos())
	}
}
