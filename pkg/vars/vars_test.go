// Copyright 2012 Marko Kungla. All rights reserved.
// Use of this source code is governed by a The Apache-style
// license that can be found in the LICENSE file.

package vars

import (
	"strconv"
	"strings"
	"testing"
)

///// TEST API
func TestParseFromStrings(t *testing.T) {
	slice := strings.Split(string(genStringTestBytes()), "\n")
	collection := ParseKeyValSlice(slice)
	for _, test := range stringTests {
		if actual := collection.Getvar(test.key); actual.String() != test.want {
			t.Errorf("Collection.Getvar(%q) = %q, want %q", test.key, actual.String(), test.want)
		}
	}

	collection2 := ParseKeyValSlice([]string{"X"})
	if actual := collection2.Getvar("x"); actual.String() != "" {
		t.Errorf("Collection.Getvar(\"X\") = %q, want \"\"", actual.String())
	}
}

func TestParseFromBytes(t *testing.T) {
	collection := ParseFromBytes(genStringTestBytes())
	for _, test := range stringTests {
		if actual := collection.Getvar(test.key); actual.String() != test.want {
			t.Errorf("Collection.Getvar(%q) = %q, want %q", test.key, actual.String(), test.want)
		}
	}
}

func TestValueFromString(t *testing.T) {
	tests := []struct {
		name string
		val  string
		want string
	}{
		{"STRING", "some-string", "some-string"},
		{"STRING", "some-string with space ", "some-string with space"},
		{"STRING", " some-string with space", "some-string with space"},
		{"STRING", "1234567", "1234567"},
	}
	for _, tt := range tests {
		if got := NewValue(tt.val); got.String() != tt.want {
			t.Errorf("ValueFromString() = %q, want %q", got.String(), tt.want)
		}
		if rv := NewValue(tt.val); string(rv.Rune()) != tt.want {
			t.Errorf("Value.Rune() = %q, want %q", string(rv.Rune()), tt.want)
		}
	}
}

///// TEST Collection
func TestCollection_GetvarOrDefaultTo(t *testing.T) {
	collection := ParseFromBytes([]byte{})
	tests := []struct {
		k      string
		defVal string
		want   string
	}{
		{"STRING", "some-string", "some-string"},
		{"STRING", "some-string with space ", "some-string with space"},
		{"STRING", " some-string with space", "some-string with space"},
		{"STRING", "1234567", "1234567"},
		{"", "1234567", "1234567"},
	}
	for _, tt := range tests {
		if actual := collection.GetVarOrDefaultTo(tt.k, tt.defVal); actual.String() != tt.want {
			t.Errorf("Collection.GetvarOrDefaultTo(%q, %q) = %q, want %q", tt.k, tt.defVal, actual, tt.want)
		}
	}
}

func TestCollection_GetvarsWithPrefix(t *testing.T) {
	collection := ParseFromBytes(genStringTestBytes())
	p := collection.GetVarsWithPrefix("CGO")
	if len(p) != 6 {
		t.Errorf("Collection.GetvarsWithPrefix(\"CGO\") = %d, want (6)", len(p))
	}
}

///// TEST Value
func TestValue_ParseBool(t *testing.T) {
	collection := ParseFromBytes(genAtobTestBytes())
	for _, test := range atobTests {
		val := collection.Getvar(test.key)
		b, err := val.Bool()
		if test.wantErr != nil {
			if err == nil {
				t.Errorf("Value(%s).ParseBool(): expected %s but got nil", test.key, test.wantErr)
			} else {
				if test.wantErr != err.(*strconv.NumError).Err {
					t.Errorf("Value(%s).ParseBool(): expected %s but got %s", test.key, test.wantErr, err)
				}
			}
		} else {
			if err != nil {
				t.Errorf("Value(%s).ParseBool(): expected no error but got %s", test.key, err)
			}
			if b != test.want {
				t.Errorf("Value(%s).ParseBool(): = %t, want %t", test.key, b, test.want)
			}
		}
	}
}

func TestValue_ParseFloat(t *testing.T) {
	collection := ParseFromBytes(genAtofTestBytes())
	for _, test := range atofTests {
		val := collection.Getvar(test.key)
		out, err := val.Float(64)
		outs := strconv.FormatFloat(out, 'g', -1, 64)
		if test.wantErr != nil {
			if err == nil {
				t.Errorf("Value(%s).ParseFloat(64) = %v, err(%s) want %v, err(%s)",
					test.key, out, err, test.want, test.wantErr)
			} else {
				if test.wantErr != err.(*strconv.NumError).Err {
					t.Errorf("Value(%s).ParseFloat(64) = %v, err(%s) want %v, err(%s)",
						test.key, out, err, test.want, test.wantErr)
				}
			}
		}
		if outs != test.want {
			t.Errorf("Value(%s).ParseFloat(64) = %v, err(%s) want %v, err(%s)",
				test.key, out, err, test.want, test.wantErr)
		}

		if float64(float32(out)) == out {
			out, err := val.Float(32)
			out32 := float32(out)
			if float64(out32) != out {
				t.Errorf("Value(%s).ParseFloat(32) = %v, not a float32 (closest is %v)", test.key, out, float64(out32))
				continue
			}
			outs := strconv.FormatFloat(float64(out32), 'g', -1, 32)
			if outs != test.want {
				t.Errorf("Value(%s).ParseFloat(32) = %v, %s want %v, %s  # %v",
					test.key, out32, err, test.want, test.wantErr, out)
			}
		}
	}
}

func TestValue_ParseFloat32(t *testing.T) {
	collection := ParseFromBytes(genAtof32TestBytes())
	for _, test := range atof32Tests {
		val := collection.Getvar(test.key)
		out, err := val.Float(32)
		out32 := float32(out)
		if float64(out32) != out {
			t.Errorf("Value(%s).ParseFloat(32) = %v, not a float32 (closest is %v)",
				test.key, out, float64(out32))
			continue
		}
		outs := strconv.FormatFloat(float64(out32), 'g', -1, 32)
		if test.wantErr != nil {
			if err == nil {
				t.Errorf("Value(%s).ParseFloat(32) = %v, err(%s) want %v, err(%s)",
					test.key, out, err, test.want, test.wantErr)
			} else {
				if test.wantErr != err.(*strconv.NumError).Err {
					t.Errorf("Value(%s).ParseFloat(32) = %v, err(%s) want %v, err(%s)",
						test.key, out, err, test.want, test.wantErr)
				}
			}
		}
		if outs != test.want {
			t.Errorf("Value(%s).ParseFloat(32) = %v, err(%s) want %v, err(%s)",
				test.key, out, err, test.want, test.wantErr)
		}
	}
}

func TestValue_ParseUint64(t *testing.T) {
	collection := ParseFromBytes(genAtoui64TestBytes())
	for _, test := range atoui64Tests {
		val := collection.Getvar(test.key)
		out, err := val.Uint(10, 64)
		if test.wantErr != nil {
			if err == nil {
				t.Errorf("Value(%s).ParseUint(10, 64) = %v, err(%s) want %v, err(%s)",
					test.key, out, err, test.want, test.wantErr)
			} else {
				if test.wantErr != err.(*strconv.NumError).Err {
					t.Errorf("Value(%s).ParseUint(10, 64) = %v, err(%s) want %v, err(%s)",
						test.key, out, err, test.want, test.wantErr)
				}
			}
		}
		if out != test.want {
			t.Errorf("Value(%s).ParseUint(10, 64) = %v, err(%s) want %v, err(%s)",
				test.key, out, err, test.want, test.wantErr)
		}
	}
}

func TestValue_ParseUint64Base(t *testing.T) {
	collection := ParseFromBytes(genBtoui64TestBytes())
	for _, test := range btoui64Tests {
		val := collection.Getvar(test.key)
		out, err := val.Uint(0, 64)
		if test.wantErr != nil {
			if err == nil {
				t.Errorf("Value(%s).ParseUint(0, 64) = %v, err(%s) want %v, err(%s)",
					test.key, out, err, test.want, test.wantErr)
			} else {
				if test.wantErr != err.(*strconv.NumError).Err {
					t.Errorf("Value(%s).ParseUint(0, 64) = %v, err(%s) want %v, err(%s)",
						test.key, out, err, test.want, test.wantErr)
				}
			}
		}
		if out != test.want {
			t.Errorf("Value(%s).ParseUint(0, 64) = %v, err(%s) want %v, err(%s)",
				test.key, out, err, test.want, test.wantErr)
		}
	}
}

func TestValue_ParseInt64(t *testing.T) {
	val := Value("200")
	iout, erri1 := val.AsInt()
	if iout != 200 {
		t.Errorf("Value(11).AsInt() = %d, err(%v) want 200", iout, erri1)
	}

	val2 := Value("x")
	iout2, erri2 := val2.AsInt()
	if iout2 != 0 || erri2 == nil {
		t.Errorf("Value(11).AsInt() = %d, err(%v) want 0 and err", iout2, erri2)
	}

	collection := ParseFromBytes(genAtoi64TestBytes())
	for _, test := range atoi64Tests {
		val := collection.Getvar(test.key)
		out, err := val.Int(10, 64)
		if test.wantErr != nil {
			if err == nil {
				t.Errorf("Value(%s).ParseInt(10, 64) = %v, err(%s) want %v, err(%s)",
					test.key, out, err, test.want, test.wantErr)
			} else {
				if test.wantErr != err.(*strconv.NumError).Err {
					t.Errorf("Value(%s).ParseInt(10, 64) = %v, err(%s) want %v, err(%s)",
						test.key, out, err, test.want, test.wantErr)
				}
			}
		}
		if out != test.want {
			t.Errorf("Value(%s).ParseInt(10, 64) = %v, err(%s) want %v, err(%s)",
				test.key, out, err, test.want, test.wantErr)
		}
	}
}

func TestValue_ParseInt64Base(t *testing.T) {
	collection := ParseFromBytes(genBtoi64TestBytes())
	for _, test := range btoi64Tests {
		val := collection.Getvar(test.key)
		out, err := val.Int(test.base, 64)
		if test.wantErr != nil {
			if err == nil {
				t.Errorf("Value(%s).ParseInt(%d, 64) = %v, err(%v) want %v, err(%v)",
					test.key, test.base, out, err, test.want, test.wantErr)
			} else {
				if test.wantErr != err.(*strconv.NumError).Err {
					t.Errorf("Value(%s).ParseInt(%d, 64) = %v, err(%v) want %v, err(%v)",
						test.key, test.base, out, err, test.want, test.wantErr)
				}
			}
		}

		if out != test.want {
			t.Errorf("Value(%s).ParseInt(%d, 64) = %v, err(%v) want %v, err(%v)",
				test.key, test.base, out, err, test.want, test.wantErr)
		}
	}
}

func TestValue_TestParseUint(t *testing.T) {
	switch strconv.IntSize {
	case 32:
		collection := ParseFromBytes(genAtoui32TestBytes())
		for _, test := range atoui32Tests {
			val := collection.Getvar(test.key)
			out, err := val.Uint(10, 0)
			if test.wantErr != nil {
				if err == nil {
					t.Errorf("Value(%s).ParseUint(10, 0) = %v, err(%s) want %v, err(%s)",
						test.key, out, err, test.want, test.wantErr)
				} else {
					if test.wantErr != err.(*strconv.NumError).Err {
						t.Errorf("Value(%s).ParseUint(10, 0) = %v, err(%s) want %v, err(%s)",
							test.key, out, err, test.want, test.wantErr)
					}
				}
			}
			if uint32(out) != test.want {
				t.Errorf("Value(%s).ParseUint(10, 0) = %v, err(%s) want %v, err(%s)",
					test.key, out, err, test.want, test.wantErr)
			}
		}
	case 64:
		collection := ParseFromBytes(genAtoui64TestBytes())
		for _, test := range atoui64Tests {
			val := collection.Getvar(test.key)
			out, err := val.Uint(10, 0)
			if test.wantErr != nil {
				if err == nil {
					t.Errorf("Value(%s).ParseUint(10, 0) = %v, err(%s) want %v, err(%s)",
						test.key, out, err, test.want, test.wantErr)
				} else {
					if test.wantErr != err.(*strconv.NumError).Err {
						t.Errorf("Value(%s).ParseUint(10, 0) = %v, err(%s) want %v, err(%s)",
							test.key, out, err, test.want, test.wantErr)
					}
				}
			}
			if uint64(out) != test.want {
				t.Errorf("Value(%s).ParseUint(10, 0) = %v, err(%s) want %v, err(%s)",
					test.key, out, err, test.want, test.wantErr)
			}
		}
	}
}

func TestValue_TestParseInt(t *testing.T) {
	switch strconv.IntSize {
	case 32:
		collection := ParseFromBytes(genAtoi32TestBytes())
		for _, test := range atoi32tests {
			val := collection.Getvar(test.key)
			out, err := val.Int(10, 0)
			if test.wantErr != nil {
				if err == nil {
					t.Errorf("Value(%s).ParseInt(10, 0) = %v, err(%s) want %v, err(%s)",
						test.key, out, err, test.want, test.wantErr)
				} else {
					if test.wantErr != err.(*strconv.NumError).Err {
						t.Errorf("Value(%s).ParseInt(10, 0)= %v, err(%s) want %v, err(%s)",
							test.key, out, err, test.want, test.wantErr)
					}
				}
			}
			if int32(out) != test.want {
				t.Errorf("Value(%s).ParseInt(10, 64) = %v, err(%s) want %v, err(%s)",
					test.key, out, err, test.want, test.wantErr)
			}
		}
	case 64:
		collection := ParseFromBytes(genAtoi64TestBytes())
		for _, test := range atoi64Tests {
			val := collection.Getvar(test.key)
			out, err := val.Int(10, 64)
			if test.wantErr != nil {
				if err == nil {
					t.Errorf("Value(%s).ParseInt(10, 64) = %v, err(%s) want %v, err(%s)",
						test.key, out, err, test.want, test.wantErr)
				} else {
					if test.wantErr != err.(*strconv.NumError).Err {
						t.Errorf("Value(%s).ParseInt(10, 64) = %v, err(%s) want %v, err(%s)",
							test.key, out, err, test.want, test.wantErr)
					}
				}
			}
			if int64(out) != test.want {
				t.Errorf("Value(%s).ParseInt(10, 64) = %v, err(%s) want %v, err(%s)",
					test.key, out, err, test.want, test.wantErr)
			}
		}
	}
}

func TestValue_ParseFields(t *testing.T) {
	collection := ParseFromBytes([]byte{})
	tests := []struct {
		k       string
		defVal  string
		wantLen int
	}{
		{"STRING", "one two", 2},
		{"STRING", "one two three four ", 4},
		{"STRING", " one two three four ", 4},
		{"STRING", "1 2 3 4 5 6 7 8.1", 8},
	}
	for _, tt := range tests {
		val := collection.GetVarOrDefaultTo(tt.k, tt.defVal)
		actual := len(val.ParseFields())
		if actual != tt.wantLen {
			t.Errorf("Value.(%q).ParseFields() len = %d, want %d", tt.k, actual, tt.wantLen)
		}
	}
}

func TestValue_ParseComplex64(t *testing.T) {
	collection := ParseFromBytes(genComplex64TestBytes())
	for _, test := range complex64Tests {
		val := collection.Getvar(test.key)
		out, err := val.Complex64()
		if test.wantErr != nil {
			if err == nil {
				t.Errorf("Value(%s).ParseComplex64() = %v, err(%s) want %v, err(%s)",
					test.key, out, err, test.want, test.wantErr)
			}
		}
		if out != test.want {
			t.Errorf("Value(%s).ParseComplex64() = %v, err(%s) want %v, err(%s)",
				test.key, out, err, test.want, test.wantErr)
		}
	}
}

func TestValue_ParseComplex128(t *testing.T) {
	collection := ParseFromBytes(genComplex128TestBytes())
	for _, test := range complex128Tests {
		val := collection.Getvar(test.key)
		out, err := val.Complex128()
		if test.wantErr != nil {
			if err == nil {
				t.Errorf("Value(%s).ParseComplex128() = %v, err(%s) want %v, err(%s)",
					test.key, out, err, test.want, test.wantErr)
			}
		}

		if out != test.want {
			t.Errorf("Value(%s).ParseComplex128() = %v, err(%s) want %v, err(%s)",
				test.key, out, err, test.want, test.wantErr)
		}
	}
}

func TestValue_Len(t *testing.T) {
	collection := ParseKeyValSlice([]string{})
	tests := []struct {
		k       string
		defVal  string
		wantLen int
	}{
		{"STRING", "one two", 2},
		{"STRING", "one two three four ", 4},
		{"STRING", " one two three four ", 4},
		{"STRING", "1 2 3 4 5 6 7 8.1", 8},
		{"STRING", "", 0},
	}
	for _, tt := range tests {
		val := collection.GetVarOrDefaultTo(tt.k, tt.defVal)
		actual := len(val.String())
		if actual != val.Len() {
			t.Errorf("Value.(%q).Len() len = %d, want %d", tt.k, actual, tt.wantLen)
		}
		if tt.defVal == "" && !val.Empty() {
			t.Errorf("Value.(%q).Empty() = %t for value(%q), want true", tt.k, val.Empty(), val.String())
		}
		if tt.defVal != "" && val.Empty() {
			t.Errorf("Value.(%q).Empty() = %t for value(%q), want true", tt.k, val.Empty(), val.String())
		}
	}
}

func TestParseFromString(t *testing.T) {
	key, val := ParseKeyVal("X=1")
	if key != "X" {
		t.Errorf("Key should be X got %q", key)
	}
	if val.Empty() {
		t.Error("Val should be 1")
	}
	if i, err := val.Int(0, 10); i != 1 || err != nil {
		t.Error("ParseInt should be 1")
	}
}

func TestParseKeyValEmpty(t *testing.T) {
	ek, ev := ParseKeyVal("")
	if ek != "" || ev != "" {
		t.Errorf("TestParseKeyValEmpty(\"\") = %q=%q, want ", ek, ev)
	}
}

func TestParseKeyValEmptyVal(t *testing.T) {
	key, val := ParseKeyVal("X")
	if key != "X" {
		t.Errorf("Key should be X got %q", key)
	}
	if !val.Empty() {
		t.Error("Val should be empty")
	}
}
