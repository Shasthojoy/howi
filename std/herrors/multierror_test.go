package herrors

import (
	"fmt"
	"testing"
)

func TestNewMultiError(t *testing.T) {
	merr := NewMultiError()
	tests := []struct {
		errstr string
	}{
		{"new error"},
		{"another error"},
		{""},
	}
	for i := 0; i < len(tests); i++ {
		err := merr.AppendString(tests[i].errstr)
		actual := merr.AppendError(err)
		if actual != err {
			t.Errorf("multierror.AppendError() = %v, want %v", actual, err)
		}

		total := (i + 1) * 2
		if tests[i].errstr != "" && merr.Len() != total {
			t.Errorf("multierror.Len() = %v, want %v", merr.Len(), total)
		}
		expected := fmt.Sprintf("%s (total errors: %d)", tests[0].errstr, total)
		if tests[i].errstr != "" && expected != merr.Error() {
			t.Errorf("multierror.Error() = %q, want %q", merr.Error(), expected)
		}
	}

}
