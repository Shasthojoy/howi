package git

import "strings"

// Output of git commands
type Output struct {
	b []byte
}

// Lines as string slice
func (o *Output) Lines() []string {
	return strings.Split(string(o.b), "\n")
}

// String retruns string representation of output
func (o *Output) String() string {
	return string(o.b)
}
