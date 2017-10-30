// Copyright 2005-2017 Marko Kungla.
// Use of this source code is governed by a The MIT License
// license that can be found in the LICENSE file.

package vars

import (
	"regexp"
	"strconv"
	"strings"
)

// Value describes the variable value
type Value string

func (v Value) String() string {
	return string(v)
}

// ParseBool calls strconv.ParseBool on Value
func (v Value) ParseBool() (bool, error) {
	return strconv.ParseBool(v.String())
}

// ParseFloat calls strconv.ParseFloat on Value
func (v Value) ParseFloat(bitSize int) (float64, error) {
	return strconv.ParseFloat(v.String(), bitSize)
}

// ParseInt calls strconv.ParseInt on Value
func (v Value) ParseInt(base int, bitSize int) (int64, error) {
	return strconv.ParseInt(v.String(), base, bitSize)
}

// ParseUint calls strconv.ParseUint on Value
func (v Value) ParseUint(base int, bitSize int) (uint64, error) {
	return strconv.ParseUint(v.String(), base, bitSize)
}

// ParseFields calls strings.Fields on Value string
func (v Value) ParseFields() []string {
	return strings.Fields(v.String())
}

// ParseComplex64 tries to split Value to strings.Fields and
// use 2 first fields to return complex64
func (v Value) ParseComplex64() (complex64, error) {
	var err error
	fields := v.ParseFields()
	if len(fields) != 2 {
		return complex64(0), strconv.ErrSyntax
	}

	var f1 float64
	var f2 float64
	if f1, err = strconv.ParseFloat(fields[0], 32); err != nil {
		return complex64(0), err
	}
	if f2, err = strconv.ParseFloat(fields[1], 32); err != nil {
		return complex64(0), err
	}
	return complex64(complex(f1, f2)), nil
}

// ParseComplex128 tries to split Value to strings.Fields and
// use 2 first fields to return complex128
func (v Value) ParseComplex128() (complex128, error) {
	var err error
	fields := v.ParseFields()
	if len(fields) != 2 {
		return complex128(0), strconv.ErrSyntax
	}
	var f1 float64
	var f2 float64
	if f1, err = strconv.ParseFloat(fields[0], 64); err != nil {
		return complex128(0), err
	}
	if f2, err = strconv.ParseFloat(fields[1], 64); err != nil {
		return complex128(0), err
	}
	return complex128(complex(f1, f2)), nil
}

// Len returns the lenght of the string representation of the Value
func (v Value) Len() int {
	return len(v.String())
}

// Empty returns true if this Value is empty
func (v Value) Empty() bool {
	return v.Len() == 0
}

// ValueFromString trims spaces and returns Value
func ValueFromString(val string) Value {
	return Value(strings.TrimSpace(val))
}

// Collection holds collection of variables
type Collection map[string]Value

// ParseFromStrings parses variables from any []"key=val" slice and
// returns map[string]string (map[key]val)
func ParseFromStrings(kv []string) Collection {
	vars := make(Collection)
	if len(kv) == 0 {
		return vars
	}
	reg := regexp.MustCompile(`"([^"]*)"`)

NextVar:
	for _, v := range kv {

		v = reg.ReplaceAllString(v, "${1}")
		l := len(v)
		if l == 0 {
			continue
		}
		for i := 0; i < l; i++ {
			if v[i] == '=' {
				vars[v[:i]] = ValueFromString(v[i+1:])
				if i < l {
					continue NextVar
				}
			}
		}
		// VAR did not have any value
		vars[strings.TrimRight(v[:l], "=")] = ""
	}
	return vars
}

// ParseFromString parses variable from single "key=val" pair and
// returns (key string, val Value)
func ParseFromString(kv string) (key string, val Value) {
	if len(kv) == 0 {
		return
	}
	reg := regexp.MustCompile(`"([^"]*)"`)

	kv = reg.ReplaceAllString(kv, "${1}")
	l := len(kv)
	if l == 0 {
		return
	}
	for i := 0; i < l; i++ {
		if kv[i] == '=' {
			key = kv[:i]
			val = ValueFromString(kv[i+1:])
			if i < l {
				return
			}
		}
	}
	// VAR did not have any value
	key = kv[:l]
	val = ""
	return
}

// ParseFromBytes parses []bytes to string, creates []string by new line
// and calls ParseFromStrings.
func ParseFromBytes(b []byte) Collection {
	slice := strings.Split(string(b[0:len(b)]), "\n")
	return ParseFromStrings(slice)
}

// Getvar retrieves the value of the variable named by the key.
// It returns the value, which will be empty string if the variable is not set
// or value was empty.
func (c Collection) Getvar(k string) (v Value) {
	if len(k) == 0 {
		return ""
	}
	v, _ = c[k]
	return
}

// GetvarOrDefaultTo is same as Getvar but returns default value if
// value of variable [key] is empty or does not exist.
// It only retuns this case default it neither sets or exports that default
func (c Collection) GetvarOrDefaultTo(k string, defVal string) (v Value) {
	v = c.Getvar(k)
	if v == "" {
		v = ValueFromString(defVal)
	}
	return
}

// GetvarsWithPrefix return all variables with prefix if any as map[]
func (c Collection) GetvarsWithPrefix(prfx string) (vars Collection) {
	vars = make(Collection)
	for k, v := range c {
		if len(k) >= len(prfx) && k[0:len(prfx)] == prfx {
			vars[k] = v
		}
	}
	return
}
