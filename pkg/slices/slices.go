package slices

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

var pfx = fmt.Sprintf("sl:::%d:::", time.Now().UTC().UnixNano())

// Slice interface
type Slice interface {
	Add(string)
	AddFromInt(int)
	SetFromSerialized(string)
	String() string
	Serialize() string
	Raw() interface{}
	Delete(string) bool
	JSON() string
}

// Slice interface
type slice struct {
	raw      interface{}
	json     string
	notEmpty bool
}

// Add appends the string value to the list of values
func (s *slice) Add(value string) error { return nil }

// Add appends or updates the string value to the list of values
func (s *slice) Set(value string) error { return s.Add(value) }

// AddFromInt directly adds an integer to the list of values
func (s *slice) AddFromInt(value int) {
	s.Add(string(value))
}

// AddFromInt directly adds an integer to the list of values
func (s *slice) AddFromFloat32(value float32) {
	s.Add(strconv.FormatFloat(float64(value), 'f', -1, 32))
}

// AddFromInt directly adds an integer to the list of values
func (s *slice) AddFromFloat64(value float64) {
	s.Add(strconv.FormatFloat(value, 'f', -1, 64))
}

// String returns a readable representation of this value (for usage defaults)
func (s *slice) String() string {
	return fmt.Sprintf("%#v", s.raw)
}

// JSON return json string of the slice
func (s *slice) JSON() string {
	bytes, _ := json.Marshal(s.raw)
	s.json = string(bytes)
	return s.json
}

// Serialize slice
func (s *slice) Serialize() string {
	return fmt.Sprintf("%s%s", pfx, s.json)
}

// SetFromSerialized set value from serialized string returns bool
// whether provided string was used to update values
func (s *slice) SetFromSerialized(str string) bool {
	if strings.HasPrefix(str, pfx) {
		_ = json.Unmarshal([]byte(strings.Replace(str, pfx, "", 1)), &s.raw)
		s.notEmpty = true
		return true
	}
	return false
}
