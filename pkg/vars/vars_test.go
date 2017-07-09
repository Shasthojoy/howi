package vars

import (
	"strconv"
	"strings"
	"testing"
)

///// TEST API
func TestParseFromStrings(t *testing.T) {
	slice := strings.Split(string(genStringTestBytes()), "\n")
	collection := ParseFromStrings(slice)
	for _, test := range stringTests {
		if actual := collection.Getvar(test.key); actual.String() != test.want {
			t.Errorf("Collection.Getvar(%q) = %q, want %q", test.key, actual.String(), test.want)
		}
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
		if got := ValueFromString(tt.val); got.String() != tt.want {
			t.Errorf("ValueFromString() = %q, want %q", got.String(), tt.want)
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
		if actual := collection.GetvarOrDefaultTo(tt.k, tt.defVal); actual.String() != tt.want {
			t.Errorf("Collection.GetvarOrDefaultTo(%q, %q) = %q, want %q", tt.k, tt.defVal, actual, tt.want)
		}
	}
}

func TestCollection_GetvarsWithPrefix(t *testing.T) {
	collection := ParseFromBytes(genStringTestBytes())
	p := collection.GetvarsWithPrefix("CGO")
	if len(p) != 6 {
		t.Errorf("Collection.GetvarsWithPrefix(\"CGO\") = %d, want (6)", len(p))
	}
}

///// TEST Value
func TestValue_ParseBool(t *testing.T) {
	collection := ParseFromBytes(genAtobTestBytes())
	for _, test := range atobTests {
		val := collection.Getvar(test.key)
		b, err := val.ParseBool()
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
		out, err := val.ParseFloat(64)
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
			out, err := val.ParseFloat(32)
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
		out, err := val.ParseFloat(32)
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
		out, err := val.ParseUint(10, 64)
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
		out, err := val.ParseUint(0, 64)
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
	collection := ParseFromBytes(genAtoi64TestBytes())
	for _, test := range atoi64Tests {
		val := collection.Getvar(test.key)
		out, err := val.ParseInt(10, 64)
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
		out, err := val.ParseInt(test.base, 64)
		if test.wantErr != nil {
			if err == nil {
				t.Errorf("Value(%s).ParseInt(%d, 64) = %v, err(%s) want %v, err(%s)",
					test.key, test.base, out, err, test.want, test.wantErr)
			} else {
				if test.wantErr != err.(*strconv.NumError).Err {
					t.Errorf("Value(%s).ParseInt(%d, 64) = %v, err(%s) want %v, err(%s)",
						test.key, test.base, out, err, test.want, test.wantErr)
				}
			}
		}
		if out != test.want {
			t.Errorf("Value(%s).ParseInt(%d, 64) = %v, err(%s) want %v, err(%s)",
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
			out, err := val.ParseUint(10, 0)
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
			if uint32(out) != test.want {
				t.Errorf("Value(%s).ParseUint(10, 64) = %v, err(%s) want %v, err(%s)",
					test.key, out, err, test.want, test.wantErr)
			}
		}
	case 64:
		collection := ParseFromBytes(genAtoui64TestBytes())
		for _, test := range atoui64Tests {
			val := collection.Getvar(test.key)
			out, err := val.ParseUint(10, 0)
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
			if uint64(out) != test.want {
				t.Errorf("Value(%s).ParseUint(10, 64) = %v, err(%s) want %v, err(%s)",
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
			out, err := val.ParseInt(10, 0)
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
			if int32(out) != test.want {
				t.Errorf("Value(%s).ParseInt(10, 64) = %v, err(%s) want %v, err(%s)",
					test.key, out, err, test.want, test.wantErr)
			}
		}
	case 64:
		collection := ParseFromBytes(genAtoi64TestBytes())
		for _, test := range atoi64Tests {
			val := collection.Getvar(test.key)
			out, err := val.ParseInt(10, 64)
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
		val := collection.GetvarOrDefaultTo(tt.k, tt.defVal)
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
		out, err := val.ParseComplex64()
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
		out, err := val.ParseComplex128()
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
	collection := ParseFromStrings([]string{})
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
		val := collection.GetvarOrDefaultTo(tt.k, tt.defVal)
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
