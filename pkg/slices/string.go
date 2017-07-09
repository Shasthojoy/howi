package slices

// MakeStringSlice returns new *StringSlice with optionally default values
func MakeStringSlice(defaults ...string) *StringSlice {
	return &StringSlice{raw: append([]string{}, defaults...)}
}

// StringSlice wraps a []string
type StringSlice struct {
	slice
	raw []string
}

// Add appends the string value to the list of values
func (s *StringSlice) Add(value string) error {
	if !s.notEmpty {
		s.raw = []string{}
		s.notEmpty = true
	}

	if s.SetFromSerialized(value) {
		return nil
	}

	s.raw = append(s.raw, value)
	return nil
}

// Raw returns the slice of strings
func (s *StringSlice) Raw() []string {
	return s.raw
}
