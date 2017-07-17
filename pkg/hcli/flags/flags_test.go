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
