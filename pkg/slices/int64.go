package slices

import "strconv"

// MakeInt64Slice returns new *Int64Slice with optionally default values
func MakeInt64Slice(defaults ...int64) *Int64Slice {
	return &Int64Slice{raw: append([]int64{}, defaults...)}
}

// Int64Slice wraps a []int64
type Int64Slice struct {
	slice
	raw []int64
}

// Add parses the value into an integer and appends it to the list of values
func (s *Int64Slice) Add(value string) error {
	if !s.notEmpty {
		s.raw = []int64{}
		s.notEmpty = true
	}

	if s.SetFromSerialized(value) {
		return nil
	}

	tmp, err := strconv.ParseInt(value, 0, 64)
	if err != nil {
		return err
	}

	s.raw = append(s.raw, tmp)
	return nil
}

// Raw returns []int64 slice
func (s *Int64Slice) Raw() []int64 {
	return s.raw
}
