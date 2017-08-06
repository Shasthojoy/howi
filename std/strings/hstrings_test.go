// Copyright 2005-2017 Marko Kungla. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// license that can be found in the LICENSE file.

package strings

import "testing"

func TestToCamelCaseAlnum(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"some-str", "SomeStr"},
		{"some str", "SomeStr"},
		{"SoMe STr", "SomeStr"},
		{"@SoMe!STr", "SomeStr"},
	}
	for _, tt := range tests {
		t.Run("TestToCamelCaseAlnum", func(t *testing.T) {
			if got := ToCamelCaseAlnum(tt.in); got != tt.want {
				t.Errorf("ToCamelCaseAlnum(%q) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}

func TestIsNamespace(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want bool
	}{
		{"valid", "NameSpace", true},
		{"valid", "NameSpace2", true},
		{"valid", "name-space", true},
		{"valid", "name_space", true},
		{"invalid", "2NameSpace", false},
		{"invalid", "name space", false},
		{"invalid", "name_space ", false},
		{"invalid", " name_space", false},
		{"invalid", "name_space_", false},
		{"invalid", "_name_space", false},
		{"invalid", "name_space-", false},
		{"invalid", "-name_space", false},
		{"invalid", "CamelCase ", false},
		{"invalid", "~abc", false},
		{"invalid", "a@bc", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNamespace(tt.arg); got != tt.want {
				t.Errorf("IsNamespace(%q) = %v, want %v", tt.arg, got, tt.want)
			}
		})
	}
}

func TestPadRight(t *testing.T) {
	type args struct {
		str    string
		length int
		pad    string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		0: {"pad alpha", args{"on", 3, "e"}, "one"},
		1: {"pad alpha", args{"one", 3, "e"}, "one"},
		2: {"pad alpha", args{"one two", 3, "a"}, "one two"},
		3: {"pad alpha", args{"", 3, "a"}, "aaa"},
		4: {"pad space", args{"", 3, " "}, "   "},
		5: {"pad num", args{"12", 3, "3"}, "123"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PadRight(tt.args.str, tt.args.length, tt.args.pad); got != tt.want {
				t.Errorf("PadRight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPadLeft(t *testing.T) {
	type args struct {
		str    string
		length int
		pad    string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		0: {"pad alpha", args{"ne", 3, "o"}, "one"},
		1: {"pad alpha", args{"one", 3, "o"}, "one"},
		2: {"pad alpha", args{"one two", 3, "a"}, "one two"},
		3: {"pad alpha", args{"", 3, "a"}, "aaa"},
		4: {"pad space", args{"", 3, " "}, "   "},
		5: {"pad num", args{"23", 3, "1"}, "123"}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PadLeft(tt.args.str, tt.args.length, tt.args.pad); got != tt.want {
				t.Errorf("PadLeft() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPadLeftUTF8(t *testing.T) {
	type args struct {
		str string
		len int
		pad string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		0: {"pad alpha", args{"ne", 3, "o"}, "one"},
		1: {"pad alpha", args{"one", 3, "o"}, "one"},
		2: {"pad alpha", args{"one two", 3, "a"}, "one two"},
		3: {"pad alpha", args{"", 3, "a"}, "aaa"},
		4: {"pad space", args{"", 3, " "}, "   "},
		5: {"pad num", args{"23", 3, "1"}, "123"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PadLeftUTF8(tt.args.str, tt.args.len, tt.args.pad); got != tt.want {
				t.Errorf("PadLeftUTF8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPadRightUTF8(t *testing.T) {
	type args struct {
		str string
		len int
		pad string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		0: {"pad alpha", args{"on", 3, "e"}, "one"},
		1: {"pad alpha", args{"one", 3, "e"}, "one"},
		2: {"pad alpha", args{"one two", 3, "a"}, "one two"},
		3: {"pad alpha", args{"", 3, "a"}, "aaa"},
		4: {"pad space", args{"", 3, " "}, "   "},
		5: {"pad num", args{"12", 3, "3"}, "123"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PadRightUTF8(tt.args.str, tt.args.len, tt.args.pad); got != tt.want {
				t.Errorf("PadRightUTF8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimpleRepeater(t *testing.T) {
	type args struct {
		str string
		n   int
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
	}{
		{"repeat x", args{"x", 10}, "xxxxxxxxxx"},
		{"repeat 1", args{"1", 10}, "1111111111"},
		{"repeat ' '", args{" ", 10}, "          "},
		{"repeat '\v'", args{string('\n'), 10}, "\n\n\n\n\n\n\n\n\n\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := simpleRepeater(tt.args.str, tt.args.n); gotOut != tt.wantOut {
				t.Errorf("simpleRepeater() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
