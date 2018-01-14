// Copyright 2012 Marko Kungla. All rights reserved.
// Use of this source code is governed by a The Apache-style
// license that can be found in the LICENSE file.

package vars

import (
	"regexp"
	"strings"
)

// ParseFromString parses variable from single "key=val" pair and
// returns (key string, val Value)
func ParseFromString(kv string) (key string, val Value) {
	if len(kv) == 0 {
		return
	}
	reg := regexp.MustCompile(`"([^"]*)"`)

	kv = reg.ReplaceAllString(kv, "${1}")
	l := len(kv)
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

// ParseFromBytes parses []bytes to string, creates []string by new line
// and calls ParseFromStrings.
func ParseFromBytes(b []byte) Collection {
	slice := strings.Split(string(b[0:]), "\n")
	return ParseFromStrings(slice)
}
