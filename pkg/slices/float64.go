package slices

import "strconv"

// MakeFloat64Slice returns new *Float64Slice with optionally default values
func MakeFloat64Slice(defaults ...float64) *Float64Slice {
	return &Float64Slice{raw: append([]float64{}, defaults...)}
}

// Float64Slice wraps a []float64
type Float64Slice struct {
	slice
	raw []float64
}

// Add parses the value into a float64 and appends it to the list of values
func (s *Float64Slice) Add(value string) error {
	if !s.notEmpty {
		s.raw = []float64{}
		s.notEmpty = true
	}

	if s.SetFromSerialized(value) {
		return nil
	}

	tmp, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return err
	}

	s.raw = append(s.raw, tmp)
	return nil
}

// Raw returns []float64 slice
func (s *Float64Slice) Raw() []float64 {
	return s.raw
}
