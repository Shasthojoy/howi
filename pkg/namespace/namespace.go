// Copyright 2017 Marko Kungla. All rights reserved.
// Use of this source code is governed by a The Apache-style
// license that can be found in the LICENSE file.

package namespace

import (
	"bytes"
	"regexp"
	"unicode"
	"unicode/utf8"
)

const (
	// NamespaceMustCompile against following expression
	NamespaceMustCompile = "^[a-zA-Z][a-zA-Z0-9_-]*[a-zA-Z0-9]$"
)

var (
	alnum = &unicode.RangeTable{
		R16: []unicode.Range16{
			{'0', '9', 1},
			{'A', 'Z', 1},
			{'a', 'z', 1},
		},
	}
)

// ToCamelCaseAlnum returns a camel case representation of the string all
// non alpha numeric characters removed. Uppercase characters are mapped
// first alnum in string and after each non alnum character is removed.
func ToCamelCaseAlnum(s string) string {
	var b bytes.Buffer
	tu := true
	for _, c := range s {
		isAlnum := unicode.Is(alnum, c)
		isSpace := unicode.IsSpace(c)
		isLower := unicode.IsLower(c)
		if isSpace || !isAlnum {
			tu = true
			continue
		}
		if tu {
			if isLower {
				b.WriteRune(unicode.ToUpper(c))
			} else {
				b.WriteRune(c)
			}
			tu = false
			continue
		} else {
			if !isLower {
				c = unicode.ToLower(c)
			}
			b.WriteRune(c)
		}
	}
	return b.String()
}

// IsValid returns true if s is string which is valid name space
func IsValid(s string) bool {
	re := regexp.MustCompile(NamespaceMustCompile)
	return re.MatchString(s)
}

// PadRight string
func PadRight(str string, length int, pad string) string {
	return str + simpleRepeater(pad, length-len(str))
}

// PadLeft string
func PadLeft(str string, length int, pad string) string {
	return simpleRepeater(pad, length-len(str)) + str
}

// PadLeftUTF8 left-pads the string with pad up to len runes
// len may be exceeded if
func PadLeftUTF8(str string, len int, pad string) string {
	return simpleRepeater(pad, len-utf8.RuneCountInString(str)) + str
}

// PadRightUTF8 right-pads the string with pad up to len runes
func PadRightUTF8(str string, len int, pad string) string {
	return str + simpleRepeater(pad, len-utf8.RuneCountInString(str))
}

func simpleRepeater(str string, n int) (out string) {
	for i := 0; i < n; i++ {
		out += str
	}
	return
}
